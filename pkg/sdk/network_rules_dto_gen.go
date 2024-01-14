package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateNetworkRuleOptions] = new(CreateNetworkRuleRequest)
)

type CreateNetworkRuleRequest struct {
	OrReplace             *bool
	name                  SchemaObjectIdentifier // required
	NetworkIdentifierType *NetworkIdentifierType
	NetworkIdentifiers    []NetworkIdentifierRequest
	NetworkRuleMode       *NetworkRuleMode
	Comment               *string
}

type NetworkIdentifierRequest struct {
	Value string
}
