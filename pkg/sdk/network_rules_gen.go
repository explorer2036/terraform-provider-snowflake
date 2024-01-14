package sdk

import "context"

type NetworkRules interface {
	Create(ctx context.Context, request *CreateNetworkRuleRequest) error
}

// CreateNetworkRuleOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-network-rule.
type CreateNetworkRuleOptions struct {
	create                bool                   `ddl:"static" sql:"CREATE"`
	OrReplace             *bool                  `ddl:"keyword" sql:"OR REPLACE"`
	networkRule           bool                   `ddl:"static" sql:"NETWORK RULE"`
	name                  SchemaObjectIdentifier `ddl:"identifier"`
	NetworkIdentifierType *NetworkIdentifierType `ddl:"parameter" sql:"TYPE"`
	NetworkIdentifiers    []NetworkIdentifier    `ddl:"parameter,must_parentheses" sql:"VALUE_LIST"`
	NetworkRuleMode       *NetworkRuleMode       `ddl:"parameter" sql:"MODE"`
	Comment               *string                `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

type NetworkIdentifier struct {
	Value string `ddl:"keyword,single_quotes"`
}
