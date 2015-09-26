package main

import pb "github.com/alvin-nt/grpc-simple-chat/proto"

import (
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Channel struct {
	// store a duplicate of the member names and ids
	members map[uint32]string

	last      time.Time
	connected uint32

	mu sync.RWMutex
}

type chatServer struct {
	name     map[uint32]string
	channels map[string]*Channel

	buf map[uint32]chan *pb.Event

	// variables to keep track of number of invocations received/events sent
	in  uint32
	out uint32

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

func (cs *chatServer) unsafeExpire(sid uint32) {
	if buf, ok := cs.buf[sid]; ok {
		close(buf)
	}
	delete(cs.name, sid)
	delete(cs.buf, sid)

	// remove the member from the channels
	for k, ch := range cs.channels {
		delete(ch.members, sid)
	}
}

func (cs *chatServer) generateSessionID() uint32 {
	return rand.Uint32()
}

func (cs *chatServer) Authorize(ctx context.Context, req *pb.RequestAuthorize) (*pb.ResponseAuthorize, error) {
	cs.in++

	// assign the required values
	if len(req.Name) == 0 {
		req.Name = RandString(16)
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
	defer cs.withReadLock(func() {
		log.Printf("User with name %s disconnected.", name)

		// broadcast this message to all of the chat rooms
		// treat as a 'leave' event
		// only send once per user
		users := make(map[uint32]bool)
		for _, ch := range cs.channels {
			for sid := range ch.members {
				if user, sent := users[sid]; !sent {
					cs.buf[sid] <- &pb.Event{
						Leave: &pb.EventLeave{
							Name: name,
						},
					}
				}
			}
		}
	})

	// in case of error/disconnect, execute this
	defer cs.withWriteLock(func() { cs.unsafeExpire(sid) })

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
			event := &pb.Event{None: &pb.EventNone{}}
			if err := stream.Send(event); err != nil {
				return err
			}
			cs.out++
		}
	}
}

func (cs *chatServer) SendCommand(ctx context.Context, req *pb.Command) (*None, error) {
	var (
		sid  = req.SessionId
		name string
		err  error
	)

	// check for the existence of the user
	cs.withReadLock(func() {
		var ok bool
		name, ok = cs.name[sid]
		if !ok {
			err = errors.New("Not authorized")
		}
	})
	if err != nil {
		return nil, err
	}

	switch cmd := req.GetCommand(); cmd {
	case condition:

	}
}

func (cs *chatServer) Say(sid uint32, req *pb.CommandSay) error {
	cs.in++

	var (
		channel = req.ChannelName
		name    string
		err     error
	)

	cs.withReadLock(func() {
		var ok bool
		name, ok = cs.name[sid]
		if !ok {
			err = errors.New("Not authorized")
		}

		// check whether the user is the member of the channel
		if len(channel) > 0 {
			if _, ok := cs.channels[channel]; !ok {
				err = fmt.Errorf("Channel %s not found", channel)
			}
		}
	})
	if err != nil {
		return err
	}

}

func (cs *chatServer) ChangeNick(cmd *pb.CommandChangeNick) error {

}

func (cs *chatServer) JoinChannel(cmd *pb.CommandJoinChannel) error {

}

func (cs *chatServer) LeaveChannel(cmd *pb.CommandLeaveChannel) error {

}

func (cs *chatServer) Exit(cmd *pb.CommandExit) error {

}

func main() {

}
