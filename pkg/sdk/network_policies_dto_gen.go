package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateNetworkPolicyOptions]   = new(CreateNetworkPolicyRequest)
	_ optionsProvider[AlterNetworkPolicyOptions]    = new(AlterNetworkPolicyRequest)
	_ optionsProvider[DropNetworkPolicyOptions]     = new(DropNetworkPolicyRequest)
	_ optionsProvider[ShowNetworkPolicyOptions]     = new(ShowNetworkPolicyRequest)
	_ optionsProvider[DescribeNetworkPolicyOptions] = new(DescribeNetworkPolicyRequest)
)

type CreateNetworkPolicyRequest struct {
	OrReplace              *bool
	name                   AccountObjectIdentifier // required
	AllowedNetworkRuleList []SchemaObjectIdentifier
	BlockedNetworkRuleList []SchemaObjectIdentifier
	AllowedIpList          []IPRequest
	BlockedIpList          []IPRequest
	Comment                *string
}

type IPRequest struct {
	IP string // required
}

type AlterNetworkPolicyRequest struct {
	IfExists     *bool
	name         AccountObjectIdentifier // required
	Set          *NetworkPolicySetRequest
	Add          *AddNetworkRuleRequest
	Remove       *RemoveNetworkRuleRequest
	UnsetComment *bool
	RenameTo     *AccountObjectIdentifier
}

type NetworkPolicySetRequest struct {
	AllowedNetworkRuleList []SchemaObjectIdentifier
	BlockedNetworkRuleList []SchemaObjectIdentifier
	AllowedIpList          []IPRequest
	BlockedIpList          []IPRequest
	Comment                *string
}

type AddNetworkRuleRequest struct {
	AddToAllowedNetworkRuleList *SchemaObjectIdentifier
	AddToBlockedNetworkRuleList *SchemaObjectIdentifier
}

type RemoveNetworkRuleRequest struct {
	RemoveFromAllowedNetworkRuleList *SchemaObjectIdentifier
	RemoveFromBlockedNetworkRuleList *SchemaObjectIdentifier
}

type DropNetworkPolicyRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
}

type ShowNetworkPolicyRequest struct{}

type DescribeNetworkPolicyRequest struct {
	name AccountObjectIdentifier // required
}
