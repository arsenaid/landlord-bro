// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/lease_service.proto

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

const (
	LeaseService_CreateLease_FullMethodName = "/api.v1.LeaseService/CreateLease"
)

// LeaseServiceClient is the client API for LeaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeaseServiceClient interface {
	CreateLease(ctx context.Context, in *CreateLeaseRequest, opts ...grpc.CallOption) (*Lease, error)
}

type leaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaseServiceClient(cc grpc.ClientConnInterface) LeaseServiceClient {
	return &leaseServiceClient{cc}
}

func (c *leaseServiceClient) CreateLease(ctx context.Context, in *CreateLeaseRequest, opts ...grpc.CallOption) (*Lease, error) {
	out := new(Lease)
	err := c.cc.Invoke(ctx, LeaseService_CreateLease_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaseServiceServer is the server API for LeaseService service.
// All implementations should embed UnimplementedLeaseServiceServer
// for forward compatibility
type LeaseServiceServer interface {
	CreateLease(context.Context, *CreateLeaseRequest) (*Lease, error)
}

// UnimplementedLeaseServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLeaseServiceServer struct {
}

func (UnimplementedLeaseServiceServer) CreateLease(context.Context, *CreateLeaseRequest) (*Lease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLease not implemented")
}

// UnsafeLeaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaseServiceServer will
// result in compilation errors.
type UnsafeLeaseServiceServer interface {
	mustEmbedUnimplementedLeaseServiceServer()
}

func RegisterLeaseServiceServer(s grpc.ServiceRegistrar, srv LeaseServiceServer) {
	s.RegisterService(&LeaseService_ServiceDesc, srv)
}

func _LeaseService_CreateLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServiceServer).CreateLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaseService_CreateLease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServiceServer).CreateLease(ctx, req.(*CreateLeaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeaseService_ServiceDesc is the grpc.ServiceDesc for LeaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.LeaseService",
	HandlerType: (*LeaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLease",
			Handler:    _LeaseService_CreateLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/lease_service.proto",
}