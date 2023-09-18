// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: metadata/metadata.proto

package metadata

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
	Metadata_Watch_FullMethodName = "/metadata.Metadata/Watch"
)

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	// Returns a stream of events for the resources that match the selector.
	Watch(ctx context.Context, in *Selector, opts ...grpc.CallOption) (Metadata_WatchClient, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) Watch(ctx context.Context, in *Selector, opts ...grpc.CallOption) (Metadata_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metadata_ServiceDesc.Streams[0], Metadata_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &metadataWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metadata_WatchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type metadataWatchClient struct {
	grpc.ClientStream
}

func (x *metadataWatchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetadataServer is the server API for Metadata service.
// All implementations must embed UnimplementedMetadataServer
// for forward compatibility
type MetadataServer interface {
	// Returns a stream of events for the resources that match the selector.
	Watch(*Selector, Metadata_WatchServer) error
	mustEmbedUnimplementedMetadataServer()
}

// UnimplementedMetadataServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (UnimplementedMetadataServer) Watch(*Selector, Metadata_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedMetadataServer) mustEmbedUnimplementedMetadataServer() {}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Selector)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetadataServer).Watch(m, &metadataWatchServer{stream})
}

type Metadata_WatchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type metadataWatchServer struct {
	grpc.ServerStream
}

func (x *metadataWatchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metadata.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Metadata_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "metadata/metadata.proto",
}
