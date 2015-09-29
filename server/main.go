package main

import pb "github.com/alvin-nt/grpc-simple-chat/proto"

import (
	"errors"
	"fmt"
	"github.com/alvin-nt/grpc-simple-chat/util"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

// Channel represents a 'chat room' that can be joined by a user.
type Channel struct {
	name string

	members []uint32

	last      time.Time
	connected uint32

	mu sync.RWMutex
}

type chatServer struct {
	name     map[uint32]string
	channels map[string]*Channel

	buf map[uint32]chan *pb.Event

	// variables to keep track of number of invocations received/events sent per second
	in   uint32
	out  uint32
	last time.Time

	mu sync.RWMutex
}

func (cs *chatServer) withReadLock(f func()) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	f()
}

func (cs *chatServer) withWriteLock(f func()) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	f()
}

func (ch *Channel) withReadLock(f func()) {
	ch.mu.RLock()
	defer ch.mu.RUnlock()
	f()
}

func (ch *Channel) withWriteLock(f func()) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	f()
}

func (ch *Channel) findMemberIdx(sid uint32) int {
	for idx, memberSid := range ch.members {
		if memberSid == sid {
			return idx
		}
	}

	return -1
}

func (ch *Channel) hasMember(sid uint32) bool {
	for _, memberSid := range ch.members {
		if memberSid == sid {
			return true
		}
	}

	return false
}

func (ch *Channel) addMember(sid uint32) error {
	if ch.hasMember(sid) {
		return fmt.Errorf("Member with id %d is already in channel %s", sid, ch.name)
	}

	// add it
	ch.members = append(ch.members, sid)
	ch.connected++

	return nil
}

func (ch *Channel) removeMember(sid uint32) error {
	idx := ch.findMemberIdx(sid)

	if idx == -1 {
		return fmt.Errorf("Member with id %d is not in channel %s", sid, ch.name)
	}

	// remove the member
	ch.members[idx] = ch.members[len(ch.members)-1]
	ch.members = ch.members[:len(ch.members)-1]
	ch.connected--

	return nil
}

func (cs *chatServer) unsafeExpire(sid uint32) {
	if buf, ok := cs.buf[sid]; ok {
		close(buf)
	}
	delete(cs.name, sid)
	delete(cs.buf, sid)

	// remove the member from the channels
	for _, ch := range cs.channels {
		if ch.hasMember(sid) {
			ch.connected--
			err := ch.removeMember(sid)
			if err != nil {
				log.Printf("Caught error: %s\n", err.Error())
			}
		}
	}
}

func (cs *chatServer) removeEmptyChannels() {
	var channelsToDelete []string

	cs.withReadLock(func() {
		for _, ch := range cs.channels {
			ch.withReadLock(func() {
				if ch.connected == 0 {
					channelsToDelete = append(channelsToDelete, ch.name)
				}
			})
		}
	})

	cs.withWriteLock(func() {
		for _, ch := range channelsToDelete {
			delete(cs.channels, ch)
		}
	})
}

func (cs *chatServer) generateSessionID() uint32 {
	return rand.Uint32()
}

func (cs *chatServer) Authorize(ctx context.Context, req *pb.RequestAuthorize) (*pb.ResponseAuthorize, error) {
	cs.in++

	// assign the required values
	if len(req.Name) == 0 {
		req.Name = util.RandString(16)
	}
	sid := cs.generateSessionID()

	cs.withWriteLock(func() {
		cs.name[sid] = req.Name
	})

	// if the user is not connected within 5 seconds, then disconnect
	go func() {
		time.Sleep(5 * time.Second)
		cs.withWriteLock(func() {
			if _, ok := cs.buf[sid]; ok {
				return
			}
			cs.unsafeExpire(sid)
		})
	}()

	// return a response
	res := pb.ResponseAuthorize{
		SessionId: sid,
		Name:      req.Name,
	}

	log.Printf("User with name %s authorized with id %d", req.Name, sid)

	return &res, nil
}

func (cs *chatServer) Connect(req *pb.RequestConnect, stream pb.SimpleChat_ConnectServer) error {
	cs.in++

	var (
		sid  = req.SessionId
		buf  = make(chan *pb.Event, 1000)
		err  error
		name string
	)

	// check for the existence of the user
	cs.withWriteLock(func() {
		var ok bool
		name, ok = cs.name[sid]
		if !ok {
			err = errors.New("Unauthorized")
			return
		}

		// check whether the user has already connected
		if _, ok := cs.buf[sid]; ok {
			err = errors.New("Already connected")
			return
		}
	})
	if err != nil {
		return err
	}

	// log
	log.Printf("User with name %s connected to server", name)

	// in case of error/disconnect, execute this
	defer cs.removeEmptyChannels()
	defer log.Printf("User with name %s disconnected.", name)
	defer cs.withWriteLock(func() { cs.unsafeExpire(sid) })
	defer cs.BroadcastDisconnectedMessage(sid)

	// the main loop
	tick := time.Tick(time.Second)
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-buf:
			if err := stream.Send(event); err != nil {
				return err
			}
			cs.out++
		case <-tick:
			event := &pb.Event{Event: &pb.Event_None{None: &pb.EventNone{}}}
			if err := stream.Send(event); err != nil {
				return err
			}
			cs.out++
		}
	}
}

func (cs *chatServer) SendCommand(ctx context.Context, req *pb.Command) (*pb.None, error) {
	cs.in++

	var (
		sid  = req.SessionId
		name string
		err  error
	)

	// check for the existence of the user
	cs.withReadLock(func() {
		if val, ok := cs.name[sid]; !ok {
			err = errors.New("Not authorized")
		} else {
			name = val
		}
	})
	if err != nil {
		return nil, err
	}

	// execute the apropriate command
	switch cmdType := req.GetCommand().(type) {
	case *pb.Command_Say:
		cmd := req.GetSay()
		err = cs.Say(sid, cmd)
	case *pb.Command_Nick:
		cmd := req.GetNick()
		err = cs.ChangeNick(sid, cmd)
	case *pb.Command_Join:
		cmd := req.GetJoin()
		err = cs.JoinChannel(sid, cmd)
	case *pb.Command_Leave:
		cmd := req.GetLeave()
		err = cs.LeaveChannel(sid, cmd)
	case *pb.Command_Exit:
		cmd := req.GetExit()
		err = cs.Exit(sid, cmd)
	default:
		err = fmt.Errorf("Unknown command type: %s", cmdType)
	}

	if err != nil {
		return nil, err
	}

	return &pb.None{}, nil
}

func (cs *chatServer) Say(sid uint32, req *pb.CommandSay) error {
	var (
		channel     = req.ChannelName
		message     = req.Message
		name        string
		err         error
		currentTime = time.Now()
	)

	if len(channel) > 0 {
		// publish message to that channel only
		cs.withReadLock(func() {
			var ch = cs.channels[channel]
			if ch == nil {
				err = fmt.Errorf("Channel %s not found", channel)
				return
			}

			// check whether the user is the member of the channel
			if !ch.hasMember(sid) {
				err = fmt.Errorf("You are not a member of channel %s", channel)
				return
			}

			// send the message to all of the members of the channels
			ch.withReadLock(func() {
				log.Printf("Message from %s to channel %s: %s", name, channel, message)
				for _, memberSid := range ch.members {
					cs.buf[memberSid] <- &pb.Event{
						Event: &pb.Event_Log{
							Log: &pb.EventLog{
								Timestamp: currentTime.String(),
								Name:      name,
								Message:   message,
								Channel:   channel,
							},
						},
					}
				}
			})

			// update the last sent message from the channel
			ch.withWriteLock(func() {
				ch.last = currentTime
			})
		})
	} else {
		// publish message to all channels
		go cs.BroadcastMessage(sid, message, currentTime)
	}
	if err != nil {
		return err
	}

	return nil
}

func (cs *chatServer) ChangeNick(sid uint32, cmd *pb.CommandChangeNick) error {
	var (
		newNick = cmd.Name
		oldNick string
	)

	if len(newNick) < 3 {
		return errors.New("New nickname cannot be less than three chars.")
	}
	cs.withReadLock(func() {
		oldNick = cs.name[sid]
	})

	cs.withWriteLock(func() {
		cs.name[sid] = newNick
	})

	// broadcast message
	go cs.BroadcastMessage(sid, fmt.Sprintf("User %s changed name to %s", oldNick, newNick), time.Now())

	return nil
}

func (cs *chatServer) BroadcastMessage(sid uint32, message string, timestamp time.Time) error {
	name := cs.name[sid]

	// publish message to all channels joined by sid
	cs.withReadLock(func() {
		log.Printf("Message from %s: %s", name, message)
		for _, ch := range cs.channels {
			userFound := false
			ch.withReadLock(func() {
				// send to that channel if the user is there
				if ch.hasMember(sid) {
					for _, memberSid := range ch.members {
						cs.buf[memberSid] <- &pb.Event{
							Event: &pb.Event_Log{
								Log: &pb.EventLog{
									Timestamp: timestamp.String(),
									Name:      name,
									Message:   message,
									Channel:   ch.name,
								},
							},
						}
					}
				}
			})

			// update the last time message is sent if user is found
			if userFound {
				ch.withWriteLock(func() {
					ch.last = timestamp
				})
			}
		}
	})

	return nil
}

func (cs *chatServer) BroadcastDisconnectedMessage(sid uint32) error {
	// publish message to all channels joined by sid
	cs.withReadLock(func() {
		for _, ch := range cs.channels {
			ch.withReadLock(func() {
				// send to that channel if the user is there
				if ch.hasMember(sid) {
					for _, memberSid := range ch.members {
						if sid != memberSid {
							cs.buf[memberSid] <- &pb.Event{
								Event: &pb.Event_Leave{
									Leave: &pb.EventLeave{Name: cs.name[sid]},
								},
							}
						}
					}
				}
			})
		}
	})

	return nil
}

func (cs *chatServer) JoinChannel(sid uint32, cmd *pb.CommandJoinChannel) error {
	var (
		chName  = cmd.Name
		err     error
		chFound bool
	)

	cs.withReadLock(func() {
		if ch, chFound := cs.channels[chName]; chFound {
			if ch.hasMember(sid) {
				err = fmt.Errorf("Already registered at channel %s", chName)
				return
			}

			// add the user to the existing channel
			log.Printf("User %s joined to channel %s", cs.name[sid], chName)
			ch.withWriteLock(func() {
				ch.addMember(sid)
			})

			// broadcast the join message to the channel
			ch.withReadLock(func() {
				for _, memberSid := range ch.members {
					cs.buf[memberSid] <- &pb.Event{
						Event: &pb.Event_Join{
							Join: &pb.EventJoin{
								Name: cs.name[sid],
								Channel: ch.name,
							},
						},
					}
				}
			})
		}
	})
	if err != nil {
		return err
	}

	// create new channel
	if !chFound {
		cs.withWriteLock(func() {
			ch := &Channel{name: chName}
			ch.addMember(sid)

			cs.channels[chName] = ch
		})
		log.Printf("Added new channel %s, by user %s\n", chName, cs.name[sid])
	}

	return nil
}

func (cs *chatServer) LeaveChannel(sid uint32, cmd *pb.CommandLeaveChannel) error {
	var (
		chName   = cmd.Name
		err      error
		removeCh bool
		ch       *Channel
	)

	cs.withReadLock(func() {
		ch = cs.channels[chName]
		if ch == nil {
			err = fmt.Errorf("Channel with name %s not found", chName)
			return
		}

		ch.withWriteLock(func() {
			err = ch.removeMember(sid)
		})
		if err != nil {
			return
		}

		// broadcast leave message to the rest of the members of the channel
		ch.withReadLock(func() {
			if ch.connected > 0 {
				removeCh = false
				for _, memberSid := range ch.members {
					cs.buf[memberSid] <- &pb.Event{
						Event: &pb.Event_Leave{
							Leave: &pb.EventLeave{
								Name: cs.name[sid],
								Channel: ch.name,
							},
						},
					}
				}
			} else {
				removeCh = true
			}
		})
	})
	if err != nil {
		return err
	}

	go func() {
		if removeCh {
			cs.withWriteLock(func() {
				delete(cs.channels, chName)
			})
		}
	}()

	return nil
}

func (cs *chatServer) Exit(sid uint32, cmd *pb.CommandExit) error {
	return nil
}

func newChatServer() *chatServer {
	return &chatServer{
		name:     make(map[uint32]string),
		buf:      make(map[uint32]chan *pb.Event),
		channels: make(map[string]*Channel),
	}
}

const (
	port = ":50001"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("net.Listen:", err)
	}

	cs := newChatServer()

	// start the server monitor
	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				now := time.Now()
				duration := now.Sub(cs.last)

				log.Printf("IN=%d OUT=%d IPS=%.0f OPS=%.0f connected=%d\n",
					cs.in,
					cs.out,
					float64(cs.in)/float64(duration)*float64(time.Second),
					float64(cs.out)/float64(duration)*float64(time.Second),
					len(cs.buf),
				)

				// reset the params
				cs.last = now
				cs.in = 0
				cs.out = 0
			}
		}
	}()

	// start the server
	server := grpc.NewServer()
	pb.RegisterSimpleChatServer(server, cs)
	log.Println("Server running at ", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatalln("Serve:", err)
	}
}
