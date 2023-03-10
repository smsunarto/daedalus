// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/prover.proto

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

// ProverClient is the client API for Prover service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProverClient interface {
	GenerateProof(ctx context.Context, in *GenerateProofRequest, opts ...grpc.CallOption) (*GenerateProofResponse, error)
}

type proverClient struct {
	cc grpc.ClientConnInterface
}

func NewProverClient(cc grpc.ClientConnInterface) ProverClient {
	return &proverClient{cc}
}

func (c *proverClient) GenerateProof(ctx context.Context, in *GenerateProofRequest, opts ...grpc.CallOption) (*GenerateProofResponse, error) {
	out := new(GenerateProofResponse)
	err := c.cc.Invoke(ctx, "/Prover/GenerateProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProverServer is the server API for Prover service.
// All implementations must embed UnimplementedProverServer
// for forward compatibility
type ProverServer interface {
	GenerateProof(context.Context, *GenerateProofRequest) (*GenerateProofResponse, error)
	mustEmbedUnimplementedProverServer()
}

// UnimplementedProverServer must be embedded to have forward compatible implementations.
type UnimplementedProverServer struct {
}

func (UnimplementedProverServer) GenerateProof(context.Context, *GenerateProofRequest) (*GenerateProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateProof not implemented")
}
func (UnimplementedProverServer) mustEmbedUnimplementedProverServer() {}

// UnsafeProverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProverServer will
// result in compilation errors.
type UnsafeProverServer interface {
	mustEmbedUnimplementedProverServer()
}

func RegisterProverServer(s grpc.ServiceRegistrar, srv ProverServer) {
	s.RegisterService(&Prover_ServiceDesc, srv)
}

func _Prover_GenerateProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProverServer).GenerateProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Prover/GenerateProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProverServer).GenerateProof(ctx, req.(*GenerateProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Prover_ServiceDesc is the grpc.ServiceDesc for Prover service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Prover_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Prover",
	HandlerType: (*ProverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateProof",
			Handler:    _Prover_GenerateProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/prover.proto",
}
