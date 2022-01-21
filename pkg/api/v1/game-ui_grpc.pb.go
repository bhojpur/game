// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GameUIClient is the client API for GameUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameUIClient interface {
	// ListEngineSpecs returns a list of Game Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (GameUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type gameUIClient struct {
	cc grpc.ClientConnInterface
}

func NewGameUIClient(cc grpc.ClientConnInterface) GameUIClient {
	return &gameUIClient{cc}
}

func (c *gameUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (GameUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameUI_ServiceDesc.Streams[0], "/v1.GameUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type gameUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *gameUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.GameUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameUIServer is the server API for GameUI service.
// All implementations must embed UnimplementedGameUIServer
// for forward compatibility
type GameUIServer interface {
	// ListEngineSpecs returns a list of Game Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, GameUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedGameUIServer()
}

// UnimplementedGameUIServer must be embedded to have forward compatible implementations.
type UnimplementedGameUIServer struct {
}

func (UnimplementedGameUIServer) ListEngineSpecs(*ListEngineSpecsRequest, GameUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedGameUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedGameUIServer) mustEmbedUnimplementedGameUIServer() {}

// UnsafeGameUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameUIServer will
// result in compilation errors.
type UnsafeGameUIServer interface {
	mustEmbedUnimplementedGameUIServer()
}

func RegisterGameUIServer(s grpc.ServiceRegistrar, srv GameUIServer) {
	s.RegisterService(&GameUI_ServiceDesc, srv)
}

func _GameUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameUIServer).ListEngineSpecs(m, &gameUIListEngineSpecsServer{stream})
}

type GameUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type gameUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *gameUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GameUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.GameUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameUI_ServiceDesc is the grpc.ServiceDesc for GameUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GameUI",
	HandlerType: (*GameUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _GameUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _GameUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "game-ui.proto",
}