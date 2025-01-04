// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: service.proto

package proto

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

const (
	PersonalWordService_SaveWord_FullMethodName = "/proto.PersonalWordService/SaveWord"
)

// PersonalWordServiceClient is the client API for PersonalWordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonalWordServiceClient interface {
	SaveWord(ctx context.Context, in *SaveWordRequest, opts ...grpc.CallOption) (*SaveWordResponse, error)
}

type personalWordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonalWordServiceClient(cc grpc.ClientConnInterface) PersonalWordServiceClient {
	return &personalWordServiceClient{cc}
}

func (c *personalWordServiceClient) SaveWord(ctx context.Context, in *SaveWordRequest, opts ...grpc.CallOption) (*SaveWordResponse, error) {
	out := new(SaveWordResponse)
	err := c.cc.Invoke(ctx, PersonalWordService_SaveWord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonalWordServiceServer is the server API for PersonalWordService service.
// All implementations must embed UnimplementedPersonalWordServiceServer
// for forward compatibility
type PersonalWordServiceServer interface {
	SaveWord(context.Context, *SaveWordRequest) (*SaveWordResponse, error)
	mustEmbedUnimplementedPersonalWordServiceServer()
}

// UnimplementedPersonalWordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonalWordServiceServer struct {
}

func (UnimplementedPersonalWordServiceServer) SaveWord(context.Context, *SaveWordRequest) (*SaveWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveWord not implemented")
}
func (UnimplementedPersonalWordServiceServer) mustEmbedUnimplementedPersonalWordServiceServer() {}

// UnsafePersonalWordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonalWordServiceServer will
// result in compilation errors.
type UnsafePersonalWordServiceServer interface {
	mustEmbedUnimplementedPersonalWordServiceServer()
}

func RegisterPersonalWordServiceServer(s grpc.ServiceRegistrar, srv PersonalWordServiceServer) {
	s.RegisterService(&PersonalWordService_ServiceDesc, srv)
}

func _PersonalWordService_SaveWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonalWordServiceServer).SaveWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonalWordService_SaveWord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonalWordServiceServer).SaveWord(ctx, req.(*SaveWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonalWordService_ServiceDesc is the grpc.ServiceDesc for PersonalWordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonalWordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PersonalWordService",
	HandlerType: (*PersonalWordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveWord",
			Handler:    _PersonalWordService_SaveWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
