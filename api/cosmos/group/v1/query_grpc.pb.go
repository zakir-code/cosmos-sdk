// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cosmos/group/v1/query.proto

package groupv1

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// GroupInfo queries group info based on group id.
	GroupInfo(ctx context.Context, in *QueryGroupInfoRequest, opts ...grpc.CallOption) (*QueryGroupInfoResponse, error)
	// GroupPolicyInfo queries group policy info based on account address of group policy.
	GroupPolicyInfo(ctx context.Context, in *QueryGroupPolicyInfoRequest, opts ...grpc.CallOption) (*QueryGroupPolicyInfoResponse, error)
	// GroupMembers queries members of a group by group id.
	GroupMembers(ctx context.Context, in *QueryGroupMembersRequest, opts ...grpc.CallOption) (*QueryGroupMembersResponse, error)
	// GroupsByAdmin queries groups by admin address.
	GroupsByAdmin(ctx context.Context, in *QueryGroupsByAdminRequest, opts ...grpc.CallOption) (*QueryGroupsByAdminResponse, error)
	// GroupPoliciesByGroup queries group policies by group id.
	GroupPoliciesByGroup(ctx context.Context, in *QueryGroupPoliciesByGroupRequest, opts ...grpc.CallOption) (*QueryGroupPoliciesByGroupResponse, error)
	// GroupPoliciesByAdmin queries group policies by admin address.
	GroupPoliciesByAdmin(ctx context.Context, in *QueryGroupPoliciesByAdminRequest, opts ...grpc.CallOption) (*QueryGroupPoliciesByAdminResponse, error)
	// Proposal queries a proposal based on proposal id.
	Proposal(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error)
	// ProposalsByGroupPolicy queries proposals based on account address of group policy.
	ProposalsByGroupPolicy(ctx context.Context, in *QueryProposalsByGroupPolicyRequest, opts ...grpc.CallOption) (*QueryProposalsByGroupPolicyResponse, error)
	// VoteByProposalVoter queries a vote by proposal id and voter.
	VoteByProposalVoter(ctx context.Context, in *QueryVoteByProposalVoterRequest, opts ...grpc.CallOption) (*QueryVoteByProposalVoterResponse, error)
	// VotesByProposal queries a vote by proposal id.
	VotesByProposal(ctx context.Context, in *QueryVotesByProposalRequest, opts ...grpc.CallOption) (*QueryVotesByProposalResponse, error)
	// VotesByVoter queries a vote by voter.
	VotesByVoter(ctx context.Context, in *QueryVotesByVoterRequest, opts ...grpc.CallOption) (*QueryVotesByVoterResponse, error)
	// GroupsByMember queries groups by member address.
	GroupsByMember(ctx context.Context, in *QueryGroupsByMemberRequest, opts ...grpc.CallOption) (*QueryGroupsByMemberResponse, error)
	// TallyResult returns the tally result of a proposal. If the proposal is
	// still in voting period, then this query computes the current tally state,
	// which might not be final. On the other hand, if the proposal is final,
	// then it simply returns the `final_tally_result` state stored in the
	// proposal itself.
	TallyResult(ctx context.Context, in *QueryTallyResultRequest, opts ...grpc.CallOption) (*QueryTallyResultResponse, error)
	// Groups queries all groups in state.
	Groups(ctx context.Context, in *QueryGroupsRequest, opts ...grpc.CallOption) (*QueryGroupsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GroupInfo(ctx context.Context, in *QueryGroupInfoRequest, opts ...grpc.CallOption) (*QueryGroupInfoResponse, error) {
	out := new(QueryGroupInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GroupPolicyInfo(ctx context.Context, in *QueryGroupPolicyInfoRequest, opts ...grpc.CallOption) (*QueryGroupPolicyInfoResponse, error) {
	out := new(QueryGroupPolicyInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupPolicyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GroupMembers(ctx context.Context, in *QueryGroupMembersRequest, opts ...grpc.CallOption) (*QueryGroupMembersResponse, error) {
	out := new(QueryGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GroupsByAdmin(ctx context.Context, in *QueryGroupsByAdminRequest, opts ...grpc.CallOption) (*QueryGroupsByAdminResponse, error) {
	out := new(QueryGroupsByAdminResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupsByAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GroupPoliciesByGroup(ctx context.Context, in *QueryGroupPoliciesByGroupRequest, opts ...grpc.CallOption) (*QueryGroupPoliciesByGroupResponse, error) {
	out := new(QueryGroupPoliciesByGroupResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupPoliciesByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GroupPoliciesByAdmin(ctx context.Context, in *QueryGroupPoliciesByAdminRequest, opts ...grpc.CallOption) (*QueryGroupPoliciesByAdminResponse, error) {
	out := new(QueryGroupPoliciesByAdminResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupPoliciesByAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Proposal(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error) {
	out := new(QueryProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/Proposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProposalsByGroupPolicy(ctx context.Context, in *QueryProposalsByGroupPolicyRequest, opts ...grpc.CallOption) (*QueryProposalsByGroupPolicyResponse, error) {
	out := new(QueryProposalsByGroupPolicyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/ProposalsByGroupPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VoteByProposalVoter(ctx context.Context, in *QueryVoteByProposalVoterRequest, opts ...grpc.CallOption) (*QueryVoteByProposalVoterResponse, error) {
	out := new(QueryVoteByProposalVoterResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/VoteByProposalVoter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VotesByProposal(ctx context.Context, in *QueryVotesByProposalRequest, opts ...grpc.CallOption) (*QueryVotesByProposalResponse, error) {
	out := new(QueryVotesByProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/VotesByProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VotesByVoter(ctx context.Context, in *QueryVotesByVoterRequest, opts ...grpc.CallOption) (*QueryVotesByVoterResponse, error) {
	out := new(QueryVotesByVoterResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/VotesByVoter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GroupsByMember(ctx context.Context, in *QueryGroupsByMemberRequest, opts ...grpc.CallOption) (*QueryGroupsByMemberResponse, error) {
	out := new(QueryGroupsByMemberResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/GroupsByMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TallyResult(ctx context.Context, in *QueryTallyResultRequest, opts ...grpc.CallOption) (*QueryTallyResultResponse, error) {
	out := new(QueryTallyResultResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/TallyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Groups(ctx context.Context, in *QueryGroupsRequest, opts ...grpc.CallOption) (*QueryGroupsResponse, error) {
	out := new(QueryGroupsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1.Query/Groups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// GroupInfo queries group info based on group id.
	GroupInfo(context.Context, *QueryGroupInfoRequest) (*QueryGroupInfoResponse, error)
	// GroupPolicyInfo queries group policy info based on account address of group policy.
	GroupPolicyInfo(context.Context, *QueryGroupPolicyInfoRequest) (*QueryGroupPolicyInfoResponse, error)
	// GroupMembers queries members of a group by group id.
	GroupMembers(context.Context, *QueryGroupMembersRequest) (*QueryGroupMembersResponse, error)
	// GroupsByAdmin queries groups by admin address.
	GroupsByAdmin(context.Context, *QueryGroupsByAdminRequest) (*QueryGroupsByAdminResponse, error)
	// GroupPoliciesByGroup queries group policies by group id.
	GroupPoliciesByGroup(context.Context, *QueryGroupPoliciesByGroupRequest) (*QueryGroupPoliciesByGroupResponse, error)
	// GroupPoliciesByAdmin queries group policies by admin address.
	GroupPoliciesByAdmin(context.Context, *QueryGroupPoliciesByAdminRequest) (*QueryGroupPoliciesByAdminResponse, error)
	// Proposal queries a proposal based on proposal id.
	Proposal(context.Context, *QueryProposalRequest) (*QueryProposalResponse, error)
	// ProposalsByGroupPolicy queries proposals based on account address of group policy.
	ProposalsByGroupPolicy(context.Context, *QueryProposalsByGroupPolicyRequest) (*QueryProposalsByGroupPolicyResponse, error)
	// VoteByProposalVoter queries a vote by proposal id and voter.
	VoteByProposalVoter(context.Context, *QueryVoteByProposalVoterRequest) (*QueryVoteByProposalVoterResponse, error)
	// VotesByProposal queries a vote by proposal id.
	VotesByProposal(context.Context, *QueryVotesByProposalRequest) (*QueryVotesByProposalResponse, error)
	// VotesByVoter queries a vote by voter.
	VotesByVoter(context.Context, *QueryVotesByVoterRequest) (*QueryVotesByVoterResponse, error)
	// GroupsByMember queries groups by member address.
	GroupsByMember(context.Context, *QueryGroupsByMemberRequest) (*QueryGroupsByMemberResponse, error)
	// TallyResult returns the tally result of a proposal. If the proposal is
	// still in voting period, then this query computes the current tally state,
	// which might not be final. On the other hand, if the proposal is final,
	// then it simply returns the `final_tally_result` state stored in the
	// proposal itself.
	TallyResult(context.Context, *QueryTallyResultRequest) (*QueryTallyResultResponse, error)
	// Groups queries all groups in state.
	Groups(context.Context, *QueryGroupsRequest) (*QueryGroupsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) GroupInfo(context.Context, *QueryGroupInfoRequest) (*QueryGroupInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupInfo not implemented")
}
func (UnimplementedQueryServer) GroupPolicyInfo(context.Context, *QueryGroupPolicyInfoRequest) (*QueryGroupPolicyInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPolicyInfo not implemented")
}
func (UnimplementedQueryServer) GroupMembers(context.Context, *QueryGroupMembersRequest) (*QueryGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupMembers not implemented")
}
func (UnimplementedQueryServer) GroupsByAdmin(context.Context, *QueryGroupsByAdminRequest) (*QueryGroupsByAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupsByAdmin not implemented")
}
func (UnimplementedQueryServer) GroupPoliciesByGroup(context.Context, *QueryGroupPoliciesByGroupRequest) (*QueryGroupPoliciesByGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPoliciesByGroup not implemented")
}
func (UnimplementedQueryServer) GroupPoliciesByAdmin(context.Context, *QueryGroupPoliciesByAdminRequest) (*QueryGroupPoliciesByAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPoliciesByAdmin not implemented")
}
func (UnimplementedQueryServer) Proposal(context.Context, *QueryProposalRequest) (*QueryProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposal not implemented")
}
func (UnimplementedQueryServer) ProposalsByGroupPolicy(context.Context, *QueryProposalsByGroupPolicyRequest) (*QueryProposalsByGroupPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposalsByGroupPolicy not implemented")
}
func (UnimplementedQueryServer) VoteByProposalVoter(context.Context, *QueryVoteByProposalVoterRequest) (*QueryVoteByProposalVoterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteByProposalVoter not implemented")
}
func (UnimplementedQueryServer) VotesByProposal(context.Context, *QueryVotesByProposalRequest) (*QueryVotesByProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotesByProposal not implemented")
}
func (UnimplementedQueryServer) VotesByVoter(context.Context, *QueryVotesByVoterRequest) (*QueryVotesByVoterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotesByVoter not implemented")
}
func (UnimplementedQueryServer) GroupsByMember(context.Context, *QueryGroupsByMemberRequest) (*QueryGroupsByMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupsByMember not implemented")
}
func (UnimplementedQueryServer) TallyResult(context.Context, *QueryTallyResultRequest) (*QueryTallyResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TallyResult not implemented")
}
func (UnimplementedQueryServer) Groups(context.Context, *QueryGroupsRequest) (*QueryGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Groups not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_GroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupInfo(ctx, req.(*QueryGroupInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GroupPolicyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupPolicyInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupPolicyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupPolicyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupPolicyInfo(ctx, req.(*QueryGroupPolicyInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupMembers(ctx, req.(*QueryGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GroupsByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupsByAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupsByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupsByAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupsByAdmin(ctx, req.(*QueryGroupsByAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GroupPoliciesByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupPoliciesByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupPoliciesByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupPoliciesByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupPoliciesByGroup(ctx, req.(*QueryGroupPoliciesByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GroupPoliciesByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupPoliciesByAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupPoliciesByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupPoliciesByAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupPoliciesByAdmin(ctx, req.(*QueryGroupPoliciesByAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Proposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/Proposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposal(ctx, req.(*QueryProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProposalsByGroupPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalsByGroupPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProposalsByGroupPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/ProposalsByGroupPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProposalsByGroupPolicy(ctx, req.(*QueryProposalsByGroupPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VoteByProposalVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVoteByProposalVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VoteByProposalVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/VoteByProposalVoter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VoteByProposalVoter(ctx, req.(*QueryVoteByProposalVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VotesByProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVotesByProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VotesByProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/VotesByProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VotesByProposal(ctx, req.(*QueryVotesByProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VotesByVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVotesByVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VotesByVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/VotesByVoter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VotesByVoter(ctx, req.(*QueryVotesByVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GroupsByMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupsByMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GroupsByMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/GroupsByMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GroupsByMember(ctx, req.(*QueryGroupsByMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TallyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTallyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TallyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/TallyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TallyResult(ctx, req.(*QueryTallyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Groups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Groups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1.Query/Groups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Groups(ctx, req.(*QueryGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.group.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupInfo",
			Handler:    _Query_GroupInfo_Handler,
		},
		{
			MethodName: "GroupPolicyInfo",
			Handler:    _Query_GroupPolicyInfo_Handler,
		},
		{
			MethodName: "GroupMembers",
			Handler:    _Query_GroupMembers_Handler,
		},
		{
			MethodName: "GroupsByAdmin",
			Handler:    _Query_GroupsByAdmin_Handler,
		},
		{
			MethodName: "GroupPoliciesByGroup",
			Handler:    _Query_GroupPoliciesByGroup_Handler,
		},
		{
			MethodName: "GroupPoliciesByAdmin",
			Handler:    _Query_GroupPoliciesByAdmin_Handler,
		},
		{
			MethodName: "Proposal",
			Handler:    _Query_Proposal_Handler,
		},
		{
			MethodName: "ProposalsByGroupPolicy",
			Handler:    _Query_ProposalsByGroupPolicy_Handler,
		},
		{
			MethodName: "VoteByProposalVoter",
			Handler:    _Query_VoteByProposalVoter_Handler,
		},
		{
			MethodName: "VotesByProposal",
			Handler:    _Query_VotesByProposal_Handler,
		},
		{
			MethodName: "VotesByVoter",
			Handler:    _Query_VotesByVoter_Handler,
		},
		{
			MethodName: "GroupsByMember",
			Handler:    _Query_GroupsByMember_Handler,
		},
		{
			MethodName: "TallyResult",
			Handler:    _Query_TallyResult_Handler,
		},
		{
			MethodName: "Groups",
			Handler:    _Query_Groups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/group/v1/query.proto",
}
