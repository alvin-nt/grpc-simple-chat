// Code generated by protoc-gen-go.
// source: ChatService.proto
// DO NOT EDIT!

/*
Package simplechat is a generated protocol buffer package.

It is generated from these files:
	ChatService.proto

It has these top-level messages:
	Channel
	Message
	MessageResponse
	UserMessage
*/
package simplechat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Channel struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}

type Message struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	Channel string `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

type MessageResponse struct {
	Status  uint32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}

type UserMessage struct {
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Nickname  string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ChatService service

type ChatServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error)
	GetChannelMessages(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_GetChannelMessagesClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := grpc.Invoke(ctx, "/simplechat.ChatService/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChannelMessages(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_GetChannelMessagesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatService_serviceDesc.Streams[0], c.cc, "/simplechat.ChatService/GetChannelMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetChannelMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetChannelMessagesClient interface {
	Recv() (*UserMessage, error)
	grpc.ClientStream
}

type chatServiceGetChannelMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetChannelMessagesClient) Recv() (*UserMessage, error) {
	m := new(UserMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChatService service

type ChatServiceServer interface {
	SendMessage(context.Context, *Message) (*MessageResponse, error)
	GetChannelMessages(*Channel, ChatService_GetChannelMessagesServer) error
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Message)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ChatServiceServer).SendMessage(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ChatService_GetChannelMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetChannelMessages(m, &chatServiceGetChannelMessagesServer{stream})
}

type ChatService_GetChannelMessagesServer interface {
	Send(*UserMessage) error
	grpc.ServerStream
}

type chatServiceGetChannelMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetChannelMessagesServer) Send(m *UserMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simplechat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChannelMessages",
			Handler:       _ChatService_GetChannelMessages_Handler,
			ServerStreams: true,
		},
	},
}
