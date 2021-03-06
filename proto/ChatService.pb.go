// Code generated by protoc-gen-go.
// source: ChatService.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ChatService.proto

It has these top-level messages:
	None
	RequestAuthorize
	ResponseAuthorize
	RequestConnect
	Command
	CommandSay
	CommandChangeNick
	CommandJoinChannel
	CommandLeaveChannel
	CommandExit
	Event
	EventNone
	EventJoin
	EventLeave
	EventLog
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type None struct {
}

func (m *None) Reset()         { *m = None{} }
func (m *None) String() string { return proto1.CompactTextString(m) }
func (*None) ProtoMessage()    {}

type RequestAuthorize struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RequestAuthorize) Reset()         { *m = RequestAuthorize{} }
func (m *RequestAuthorize) String() string { return proto1.CompactTextString(m) }
func (*RequestAuthorize) ProtoMessage()    {}

type ResponseAuthorize struct {
	SessionId uint32 `protobuf:"varint,1,opt,name=session_id" json:"session_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ResponseAuthorize) Reset()         { *m = ResponseAuthorize{} }
func (m *ResponseAuthorize) String() string { return proto1.CompactTextString(m) }
func (*ResponseAuthorize) ProtoMessage()    {}

type RequestConnect struct {
	SessionId uint32 `protobuf:"varint,1,opt,name=session_id" json:"session_id,omitempty"`
}

func (m *RequestConnect) Reset()         { *m = RequestConnect{} }
func (m *RequestConnect) String() string { return proto1.CompactTextString(m) }
func (*RequestConnect) ProtoMessage()    {}

type Command struct {
	SessionId uint32 `protobuf:"varint,1,opt,name=session_id" json:"session_id,omitempty"`
	// Types that are valid to be assigned to Command:
	//	*Command_Say
	//	*Command_Nick
	//	*Command_Join
	//	*Command_Leave
	//	*Command_Exit
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto1.CompactTextString(m) }
func (*Command) ProtoMessage()    {}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_Say struct {
	Say *CommandSay `protobuf:"bytes,2,opt,name=say,oneof"`
}
type Command_Nick struct {
	Nick *CommandChangeNick `protobuf:"bytes,3,opt,name=nick,oneof"`
}
type Command_Join struct {
	Join *CommandJoinChannel `protobuf:"bytes,4,opt,name=join,oneof"`
}
type Command_Leave struct {
	Leave *CommandLeaveChannel `protobuf:"bytes,5,opt,name=leave,oneof"`
}
type Command_Exit struct {
	Exit *CommandExit `protobuf:"bytes,6,opt,name=exit,oneof"`
}

func (*Command_Say) isCommand_Command()   {}
func (*Command_Nick) isCommand_Command()  {}
func (*Command_Join) isCommand_Command()  {}
func (*Command_Leave) isCommand_Command() {}
func (*Command_Exit) isCommand_Command()  {}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Command) GetSay() *CommandSay {
	if x, ok := m.GetCommand().(*Command_Say); ok {
		return x.Say
	}
	return nil
}

func (m *Command) GetNick() *CommandChangeNick {
	if x, ok := m.GetCommand().(*Command_Nick); ok {
		return x.Nick
	}
	return nil
}

func (m *Command) GetJoin() *CommandJoinChannel {
	if x, ok := m.GetCommand().(*Command_Join); ok {
		return x.Join
	}
	return nil
}

func (m *Command) GetLeave() *CommandLeaveChannel {
	if x, ok := m.GetCommand().(*Command_Leave); ok {
		return x.Leave
	}
	return nil
}

func (m *Command) GetExit() *CommandExit {
	if x, ok := m.GetCommand().(*Command_Exit); ok {
		return x.Exit
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Command) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), []interface{}) {
	return _Command_OneofMarshaler, _Command_OneofUnmarshaler, []interface{}{
		(*Command_Say)(nil),
		(*Command_Nick)(nil),
		(*Command_Join)(nil),
		(*Command_Leave)(nil),
		(*Command_Exit)(nil),
	}
}

func _Command_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*Command)
	// command
	switch x := m.Command.(type) {
	case *Command_Say:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Say); err != nil {
			return err
		}
	case *Command_Nick:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Nick); err != nil {
			return err
		}
	case *Command_Join:
		b.EncodeVarint(4<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Join); err != nil {
			return err
		}
	case *Command_Leave:
		b.EncodeVarint(5<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Leave); err != nil {
			return err
		}
	case *Command_Exit:
		b.EncodeVarint(6<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Exit); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Command.Command has unexpected type %T", x)
	}
	return nil
}

func _Command_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*Command)
	switch tag {
	case 2: // command.say
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(CommandSay)
		err := b.DecodeMessage(msg)
		m.Command = &Command_Say{msg}
		return true, err
	case 3: // command.nick
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(CommandChangeNick)
		err := b.DecodeMessage(msg)
		m.Command = &Command_Nick{msg}
		return true, err
	case 4: // command.join
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(CommandJoinChannel)
		err := b.DecodeMessage(msg)
		m.Command = &Command_Join{msg}
		return true, err
	case 5: // command.leave
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(CommandLeaveChannel)
		err := b.DecodeMessage(msg)
		m.Command = &Command_Leave{msg}
		return true, err
	case 6: // command.exit
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(CommandExit)
		err := b.DecodeMessage(msg)
		m.Command = &Command_Exit{msg}
		return true, err
	default:
		return false, nil
	}
}

type CommandSay struct {
	Message     string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name" json:"channel_name,omitempty"`
}

func (m *CommandSay) Reset()         { *m = CommandSay{} }
func (m *CommandSay) String() string { return proto1.CompactTextString(m) }
func (*CommandSay) ProtoMessage()    {}

type CommandChangeNick struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CommandChangeNick) Reset()         { *m = CommandChangeNick{} }
func (m *CommandChangeNick) String() string { return proto1.CompactTextString(m) }
func (*CommandChangeNick) ProtoMessage()    {}

type CommandJoinChannel struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CommandJoinChannel) Reset()         { *m = CommandJoinChannel{} }
func (m *CommandJoinChannel) String() string { return proto1.CompactTextString(m) }
func (*CommandJoinChannel) ProtoMessage()    {}

type CommandLeaveChannel struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CommandLeaveChannel) Reset()         { *m = CommandLeaveChannel{} }
func (m *CommandLeaveChannel) String() string { return proto1.CompactTextString(m) }
func (*CommandLeaveChannel) ProtoMessage()    {}

type CommandExit struct {
}

func (m *CommandExit) Reset()         { *m = CommandExit{} }
func (m *CommandExit) String() string { return proto1.CompactTextString(m) }
func (*CommandExit) ProtoMessage()    {}

type Event struct {
	// Types that are valid to be assigned to Event:
	//	*Event_None
	//	*Event_Join
	//	*Event_Leave
	//	*Event_Log
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto1.CompactTextString(m) }
func (*Event) ProtoMessage()    {}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_None struct {
	None *EventNone `protobuf:"bytes,1,opt,name=none,oneof"`
}
type Event_Join struct {
	Join *EventJoin `protobuf:"bytes,2,opt,name=join,oneof"`
}
type Event_Leave struct {
	Leave *EventLeave `protobuf:"bytes,3,opt,name=leave,oneof"`
}
type Event_Log struct {
	Log *EventLog `protobuf:"bytes,4,opt,name=log,oneof"`
}

func (*Event_None) isEvent_Event()  {}
func (*Event_Join) isEvent_Event()  {}
func (*Event_Leave) isEvent_Event() {}
func (*Event_Log) isEvent_Event()   {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetNone() *EventNone {
	if x, ok := m.GetEvent().(*Event_None); ok {
		return x.None
	}
	return nil
}

func (m *Event) GetJoin() *EventJoin {
	if x, ok := m.GetEvent().(*Event_Join); ok {
		return x.Join
	}
	return nil
}

func (m *Event) GetLeave() *EventLeave {
	if x, ok := m.GetEvent().(*Event_Leave); ok {
		return x.Leave
	}
	return nil
}

func (m *Event) GetLog() *EventLog {
	if x, ok := m.GetEvent().(*Event_Log); ok {
		return x.Log
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, []interface{}{
		(*Event_None)(nil),
		(*Event_Join)(nil),
		(*Event_Leave)(nil),
		(*Event_Log)(nil),
	}
}

func _Event_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*Event)
	// event
	switch x := m.Event.(type) {
	case *Event_None:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *Event_Join:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Join); err != nil {
			return err
		}
	case *Event_Leave:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Leave); err != nil {
			return err
		}
	case *Event_Log:
		b.EncodeVarint(4<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Event has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // event.none
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(EventNone)
		err := b.DecodeMessage(msg)
		m.Event = &Event_None{msg}
		return true, err
	case 2: // event.join
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(EventJoin)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Join{msg}
		return true, err
	case 3: // event.leave
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(EventLeave)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Leave{msg}
		return true, err
	case 4: // event.log
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(EventLog)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Log{msg}
		return true, err
	default:
		return false, nil
	}
}

type EventNone struct {
}

func (m *EventNone) Reset()         { *m = EventNone{} }
func (m *EventNone) String() string { return proto1.CompactTextString(m) }
func (*EventNone) ProtoMessage()    {}

type EventJoin struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
}

func (m *EventJoin) Reset()         { *m = EventJoin{} }
func (m *EventJoin) String() string { return proto1.CompactTextString(m) }
func (*EventJoin) ProtoMessage()    {}

type EventLeave struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
}

func (m *EventLeave) Reset()         { *m = EventLeave{} }
func (m *EventLeave) String() string { return proto1.CompactTextString(m) }
func (*EventLeave) ProtoMessage()    {}

type EventLog struct {
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Channel   string `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
}

func (m *EventLog) Reset()         { *m = EventLog{} }
func (m *EventLog) String() string { return proto1.CompactTextString(m) }
func (*EventLog) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for SimpleChat service

type SimpleChatClient interface {
	Authorize(ctx context.Context, in *RequestAuthorize, opts ...grpc.CallOption) (*ResponseAuthorize, error)
	Connect(ctx context.Context, in *RequestConnect, opts ...grpc.CallOption) (SimpleChat_ConnectClient, error)
	SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*None, error)
}

type simpleChatClient struct {
	cc *grpc.ClientConn
}

func NewSimpleChatClient(cc *grpc.ClientConn) SimpleChatClient {
	return &simpleChatClient{cc}
}

func (c *simpleChatClient) Authorize(ctx context.Context, in *RequestAuthorize, opts ...grpc.CallOption) (*ResponseAuthorize, error) {
	out := new(ResponseAuthorize)
	err := grpc.Invoke(ctx, "/proto.SimpleChat/Authorize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleChatClient) Connect(ctx context.Context, in *RequestConnect, opts ...grpc.CallOption) (SimpleChat_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SimpleChat_serviceDesc.Streams[0], c.cc, "/proto.SimpleChat/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleChatConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimpleChat_ConnectClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type simpleChatConnectClient struct {
	grpc.ClientStream
}

func (x *simpleChatConnectClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simpleChatClient) SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := grpc.Invoke(ctx, "/proto.SimpleChat/SendCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SimpleChat service

type SimpleChatServer interface {
	Authorize(context.Context, *RequestAuthorize) (*ResponseAuthorize, error)
	Connect(*RequestConnect, SimpleChat_ConnectServer) error
	SendCommand(context.Context, *Command) (*None, error)
}

func RegisterSimpleChatServer(s *grpc.Server, srv SimpleChatServer) {
	s.RegisterService(&_SimpleChat_serviceDesc, srv)
}

func _SimpleChat_Authorize_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(RequestAuthorize)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SimpleChatServer).Authorize(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SimpleChat_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestConnect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimpleChatServer).Connect(m, &simpleChatConnectServer{stream})
}

type SimpleChat_ConnectServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type simpleChatConnectServer struct {
	grpc.ServerStream
}

func (x *simpleChatConnectServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _SimpleChat_SendCommand_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Command)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SimpleChatServer).SendCommand(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _SimpleChat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SimpleChat",
	HandlerType: (*SimpleChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _SimpleChat_Authorize_Handler,
		},
		{
			MethodName: "SendCommand",
			Handler:    _SimpleChat_SendCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _SimpleChat_Connect_Handler,
			ServerStreams: true,
		},
	},
}
