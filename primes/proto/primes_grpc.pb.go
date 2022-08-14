// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: primes.proto

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

// PrimesServiceClient is the client API for PrimesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimesServiceClient interface {
	Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (PrimesService_PrimesClient, error)
}

type primesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimesServiceClient(cc grpc.ClientConnInterface) PrimesServiceClient {
	return &primesServiceClient{cc}
}

func (c *primesServiceClient) Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (PrimesService_PrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimesService_ServiceDesc.Streams[0], "/primes.PrimesService/Primes", opts...)
	if err != nil {
		return nil, err
	}
	x := &primesServicePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimesService_PrimesClient interface {
	Recv() (*PrimesResponse, error)
	grpc.ClientStream
}

type primesServicePrimesClient struct {
	grpc.ClientStream
}

func (x *primesServicePrimesClient) Recv() (*PrimesResponse, error) {
	m := new(PrimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimesServiceServer is the server API for PrimesService service.
// All implementations must embed UnimplementedPrimesServiceServer
// for forward compatibility
type PrimesServiceServer interface {
	Primes(*PrimesRequest, PrimesService_PrimesServer) error
	mustEmbedUnimplementedPrimesServiceServer()
}

// UnimplementedPrimesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimesServiceServer struct {
}

func (UnimplementedPrimesServiceServer) Primes(*PrimesRequest, PrimesService_PrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedPrimesServiceServer) mustEmbedUnimplementedPrimesServiceServer() {}

// UnsafePrimesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimesServiceServer will
// result in compilation errors.
type UnsafePrimesServiceServer interface {
	mustEmbedUnimplementedPrimesServiceServer()
}

func RegisterPrimesServiceServer(s grpc.ServiceRegistrar, srv PrimesServiceServer) {
	s.RegisterService(&PrimesService_ServiceDesc, srv)
}

func _PrimesService_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimesServiceServer).Primes(m, &primesServicePrimesServer{stream})
}

type PrimesService_PrimesServer interface {
	Send(*PrimesResponse) error
	grpc.ServerStream
}

type primesServicePrimesServer struct {
	grpc.ServerStream
}

func (x *primesServicePrimesServer) Send(m *PrimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PrimesService_ServiceDesc is the grpc.ServiceDesc for PrimesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "primes.PrimesService",
	HandlerType: (*PrimesServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _PrimesService_Primes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "primes.proto",
}
