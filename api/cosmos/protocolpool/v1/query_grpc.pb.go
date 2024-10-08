// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/protocolpool/v1/query.proto

package protocolpoolv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Query_CommunityPool_FullMethodName   = "/cosmos.protocolpool.v1.Query/CommunityPool"
	Query_UnclaimedBudget_FullMethodName = "/cosmos.protocolpool.v1.Query/UnclaimedBudget"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service for community pool module.
type QueryClient interface {
	// CommunityPool queries the community pool coins.
	CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error)
	// UnclaimedBudget queries the remaining budget left to be claimed and it gives overall budget allocation view.
	UnclaimedBudget(ctx context.Context, in *QueryUnclaimedBudgetRequest, opts ...grpc.CallOption) (*QueryUnclaimedBudgetResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryCommunityPoolResponse)
	err := c.cc.Invoke(ctx, Query_CommunityPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UnclaimedBudget(ctx context.Context, in *QueryUnclaimedBudgetRequest, opts ...grpc.CallOption) (*QueryUnclaimedBudgetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryUnclaimedBudgetResponse)
	err := c.cc.Invoke(ctx, Query_UnclaimedBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query defines the gRPC querier service for community pool module.
type QueryServer interface {
	// CommunityPool queries the community pool coins.
	CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error)
	// UnclaimedBudget queries the remaining budget left to be claimed and it gives overall budget allocation view.
	UnclaimedBudget(context.Context, *QueryUnclaimedBudgetRequest) (*QueryUnclaimedBudgetResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPool not implemented")
}
func (UnimplementedQueryServer) UnclaimedBudget(context.Context, *QueryUnclaimedBudgetRequest) (*QueryUnclaimedBudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnclaimedBudget not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_CommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CommunityPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommunityPool(ctx, req.(*QueryCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UnclaimedBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnclaimedBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnclaimedBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_UnclaimedBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnclaimedBudget(ctx, req.(*QueryUnclaimedBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.protocolpool.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommunityPool",
			Handler:    _Query_CommunityPool_Handler,
		},
		{
			MethodName: "UnclaimedBudget",
			Handler:    _Query_UnclaimedBudget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/protocolpool/v1/query.proto",
}
