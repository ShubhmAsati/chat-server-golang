// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package message_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	Connect(ctx context.Context, in *UserById, opts ...grpc.CallOption) (Message_ConnectClient, error)
	BroadCastMessage(ctx context.Context, in *UserChatMessage, opts ...grpc.CallOption) (*MsgSent, error)
	SendMessage(ctx context.Context, in *UserChatMessage, opts ...grpc.CallOption) (*MsgSent, error)
	RemoveConnection(ctx context.Context, in *UserById, opts ...grpc.CallOption) (*StatusOkResponse, error)
	EmitWebRTCevents(ctx context.Context, in *EmitWebRTCeventsRequest, opts ...grpc.CallOption) (*StatusOkResponse, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Connect(ctx context.Context, in *UserById, opts ...grpc.CallOption) (Message_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Message_serviceDesc.Streams[0], "/message_grpc.Message/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Message_ConnectClient interface {
	Recv() (*UserChatMessage, error)
	grpc.ClientStream
}

type messageConnectClient struct {
	grpc.ClientStream
}

func (x *messageConnectClient) Recv() (*UserChatMessage, error) {
	m := new(UserChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageClient) BroadCastMessage(ctx context.Context, in *UserChatMessage, opts ...grpc.CallOption) (*MsgSent, error) {
	out := new(MsgSent)
	err := c.cc.Invoke(ctx, "/message_grpc.Message/BroadCastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendMessage(ctx context.Context, in *UserChatMessage, opts ...grpc.CallOption) (*MsgSent, error) {
	out := new(MsgSent)
	err := c.cc.Invoke(ctx, "/message_grpc.Message/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) RemoveConnection(ctx context.Context, in *UserById, opts ...grpc.CallOption) (*StatusOkResponse, error) {
	out := new(StatusOkResponse)
	err := c.cc.Invoke(ctx, "/message_grpc.Message/RemoveConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) EmitWebRTCevents(ctx context.Context, in *EmitWebRTCeventsRequest, opts ...grpc.CallOption) (*StatusOkResponse, error) {
	out := new(StatusOkResponse)
	err := c.cc.Invoke(ctx, "/message_grpc.Message/emitWebRTCevents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	Connect(*UserById, Message_ConnectServer) error
	BroadCastMessage(context.Context, *UserChatMessage) (*MsgSent, error)
	SendMessage(context.Context, *UserChatMessage) (*MsgSent, error)
	RemoveConnection(context.Context, *UserById) (*StatusOkResponse, error)
	EmitWebRTCevents(context.Context, *EmitWebRTCeventsRequest) (*StatusOkResponse, error)
	MustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (*UnimplementedMessageServer) Connect(*UserById, Message_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedMessageServer) BroadCastMessage(context.Context, *UserChatMessage) (*MsgSent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadCastMessage not implemented")
}
func (*UnimplementedMessageServer) SendMessage(context.Context, *UserChatMessage) (*MsgSent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedMessageServer) RemoveConnection(context.Context, *UserById) (*StatusOkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveConnection not implemented")
}
func (*UnimplementedMessageServer) EmitWebRTCevents(context.Context, *EmitWebRTCeventsRequest) (*StatusOkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitWebRTCevents not implemented")
}
func (*UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserById)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServer).Connect(m, &messageConnectServer{stream})
}

type Message_ConnectServer interface {
	Send(*UserChatMessage) error
	grpc.ServerStream
}

type messageConnectServer struct {
	grpc.ServerStream
}

func (x *messageConnectServer) Send(m *UserChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Message_BroadCastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).BroadCastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_grpc.Message/BroadCastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).BroadCastMessage(ctx, req.(*UserChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_grpc.Message/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMessage(ctx, req.(*UserChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_RemoveConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).RemoveConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_grpc.Message/RemoveConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).RemoveConnection(ctx, req.(*UserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_EmitWebRTCevents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitWebRTCeventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).EmitWebRTCevents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_grpc.Message/EmitWebRTCevents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).EmitWebRTCevents(ctx, req.(*EmitWebRTCeventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message_grpc.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadCastMessage",
			Handler:    _Message_BroadCastMessage_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Message_SendMessage_Handler,
		},
		{
			MethodName: "RemoveConnection",
			Handler:    _Message_RemoveConnection_Handler,
		},
		{
			MethodName: "emitWebRTCevents",
			Handler:    _Message_EmitWebRTCevents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Message_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message/message_service.proto",
}
