// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatBoxClient is the client API for ChatBox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatBoxClient interface {
	SendChat(ctx context.Context, in *ChatData, opts ...grpc.CallOption) (*Empty, error)
	GetChat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChatList, error)
}

type chatBoxClient struct {
	cc grpc.ClientConnInterface
}

func NewChatBoxClient(cc grpc.ClientConnInterface) ChatBoxClient {
	return &chatBoxClient{cc}
}

func (c *chatBoxClient) SendChat(ctx context.Context, in *ChatData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatbox.ChatBox/SendChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatBoxClient) GetChat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChatList, error) {
	out := new(ChatList)
	err := c.cc.Invoke(ctx, "/chatbox.ChatBox/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatBoxServer is the server API for ChatBox service.
// All implementations should embed UnimplementedChatBoxServer
// for forward compatibility
type ChatBoxServer interface {
	SendChat(context.Context, *ChatData) (*Empty, error)
	GetChat(context.Context, *Empty) (*ChatList, error)
}

// UnimplementedChatBoxServer should be embedded to have forward compatible implementations.
type UnimplementedChatBoxServer struct {
}

func (UnimplementedChatBoxServer) SendChat(context.Context, *ChatData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChat not implemented")
}
func (UnimplementedChatBoxServer) GetChat(context.Context, *Empty) (*ChatList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}

// UnsafeChatBoxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatBoxServer will
// result in compilation errors.
type UnsafeChatBoxServer interface {
	mustEmbedUnimplementedChatBoxServer()
}

func RegisterChatBoxServer(s grpc.ServiceRegistrar, srv ChatBoxServer) {
	s.RegisterService(&ChatBox_ServiceDesc, srv)
}

func _ChatBox_SendChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBoxServer).SendChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbox.ChatBox/SendChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBoxServer).SendChat(ctx, req.(*ChatData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatBox_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBoxServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbox.ChatBox/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBoxServer).GetChat(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatBox_ServiceDesc is the grpc.ServiceDesc for ChatBox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatBox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatbox.ChatBox",
	HandlerType: (*ChatBoxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChat",
			Handler:    _ChatBox_SendChat_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _ChatBox_GetChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chatbox.proto",
}
