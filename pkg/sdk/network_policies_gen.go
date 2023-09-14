package sdk

import "context"

type NetworkPolicies interface {
	Create(ctx context.Context, request *CreateNetworkPolicyRequest) error
	Drop(ctx context.Context, request *DropNetworkPolicyRequest) error
	Show(ctx context.Context, request *ShowNetworkPolicyRequest) (any, error)

	Describe(ctx context.Context, request *DescribeNetworkPolicyRequest) error
}

// CreateNetworkPolicyOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-network-policy.
type CreateNetworkPolicyOptions struct {
	create        bool                    `ddl:"static" sql:"CREATE"`
	OrReplace     *bool                   `ddl:"keyword" sql:"OR REPLACE"`
	networkPolicy bool                    `ddl:"static" sql:"NETWORK POLICY"`
	name          AccountObjectIdentifier `ddl:"identifier"`
	AllowedIpList []string                `ddl:"parameter,parentheses" sql:"ALLOWED_IP_LIST"`
	Comment       *string                 `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

// DropNetworkPolicyOptions is based on https://docs.snowflake.com/en/sql-reference/sql/drop-network-policy.
type DropNetworkPolicyOptions struct {
	drop        bool                    `ddl:"static" sql:"DROP"`
	networkRule bool                    `ddl:"static" sql:"NETWORK RULE"`
	IfExists    *bool                   `ddl:"keyword" sql:"IF EXISTS"`
	name        AccountObjectIdentifier `ddl:"identifier"`
}

// ShowNetworkPolicyOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-network-policies.
type ShowNetworkPolicyOptions struct {
	show            bool `ddl:"static" sql:"SHOW"`
	networkPolicies bool `ddl:"static" sql:"NETWORK POLICIES"`
}

type showNetworkPolicyDBRow struct {
	CreatedOn              string `db:"created_on"`
	Name                   string `db:"name"`
	Comment                string `db:"comment"`
	EntriesInAllowedIpList int    `db:"entries_in_allowed_ip_list"`
	EntriesInBlockedIpList int    `db:"entries_in_blocked_ip_list"`
}

type NetworkPolicy struct {
	CreatedOn              string
	Name                   string
	Comment                string
	EntriesInAllowedIpList int
	EntriesInBlockedIpList int
}

// DescribeNetworkPolicyOptions is based on https://docs.snowflake.com/en/sql-reference/sql/desc-network-policy.
type DescribeNetworkPolicyOptions struct {
	describeNetworkPolicy bool                    `ddl:"static" sql:"DESCRIBE NETWORK POLICY"`
	name                  AccountObjectIdentifier `ddl:"identifier"`
}

type describeNetworkPolicyDBRow struct {
	Name  string `db:"name"`
	Value string `db:"value"`
}
