// Code generated by dto builder generator; DO NOT EDIT.

package sdk

func NewCreateNetworkPolicyRequest(
	name AccountObjectIdentifier,
) *CreateNetworkPolicyRequest {
	s := CreateNetworkPolicyRequest{}
	s.name = name
	return &s
}

func (s *CreateNetworkPolicyRequest) WithOrReplace(OrReplace *bool) *CreateNetworkPolicyRequest {
	s.OrReplace = OrReplace
	return s
}

func (s *CreateNetworkPolicyRequest) WithAllowedNetworkRuleList(AllowedNetworkRuleList []SchemaObjectIdentifier) *CreateNetworkPolicyRequest {
	s.AllowedNetworkRuleList = AllowedNetworkRuleList
	return s
}

func (s *CreateNetworkPolicyRequest) WithBlockedNetworkRuleList(BlockedNetworkRuleList []SchemaObjectIdentifier) *CreateNetworkPolicyRequest {
	s.BlockedNetworkRuleList = BlockedNetworkRuleList
	return s
}

func (s *CreateNetworkPolicyRequest) WithAllowedIpList(AllowedIpList []IPRequest) *CreateNetworkPolicyRequest {
	s.AllowedIpList = AllowedIpList
	return s
}

func (s *CreateNetworkPolicyRequest) WithBlockedIpList(BlockedIpList []IPRequest) *CreateNetworkPolicyRequest {
	s.BlockedIpList = BlockedIpList
	return s
}

func (s *CreateNetworkPolicyRequest) WithComment(Comment *string) *CreateNetworkPolicyRequest {
	s.Comment = Comment
	return s
}

func NewIPRequest(
	IP string,
) *IPRequest {
	s := IPRequest{}
	s.IP = IP
	return &s
}

func NewAlterNetworkPolicyRequest(
	name AccountObjectIdentifier,
) *AlterNetworkPolicyRequest {
	s := AlterNetworkPolicyRequest{}
	s.name = name
	return &s
}

func (s *AlterNetworkPolicyRequest) WithIfExists(IfExists *bool) *AlterNetworkPolicyRequest {
	s.IfExists = IfExists
	return s
}

func (s *AlterNetworkPolicyRequest) WithSet(Set *NetworkPolicySetRequest) *AlterNetworkPolicyRequest {
	s.Set = Set
	return s
}

func (s *AlterNetworkPolicyRequest) WithUnsetComment(UnsetComment *bool) *AlterNetworkPolicyRequest {
	s.UnsetComment = UnsetComment
	return s
}

func (s *AlterNetworkPolicyRequest) WithRenameTo(RenameTo *AccountObjectIdentifier) *AlterNetworkPolicyRequest {
	s.RenameTo = RenameTo
	return s
}

func NewNetworkPolicySetRequest() *NetworkPolicySetRequest {
	return &NetworkPolicySetRequest{}
}

func (s *NetworkPolicySetRequest) WithAllowedIpList(AllowedIpList []IPRequest) *NetworkPolicySetRequest {
	s.AllowedIpList = AllowedIpList
	return s
}

func (s *NetworkPolicySetRequest) WithBlockedIpList(BlockedIpList []IPRequest) *NetworkPolicySetRequest {
	s.BlockedIpList = BlockedIpList
	return s
}

func (s *NetworkPolicySetRequest) WithComment(Comment *string) *NetworkPolicySetRequest {
	s.Comment = Comment
	return s
}

func NewDropNetworkPolicyRequest(
	name AccountObjectIdentifier,
) *DropNetworkPolicyRequest {
	s := DropNetworkPolicyRequest{}
	s.name = name
	return &s
}

func (s *DropNetworkPolicyRequest) WithIfExists(IfExists *bool) *DropNetworkPolicyRequest {
	s.IfExists = IfExists
	return s
}

func NewShowNetworkPolicyRequest() *ShowNetworkPolicyRequest {
	return &ShowNetworkPolicyRequest{}
}

func NewDescribeNetworkPolicyRequest(
	name AccountObjectIdentifier,
) *DescribeNetworkPolicyRequest {
	s := DescribeNetworkPolicyRequest{}
	s.name = name
	return &s
}
