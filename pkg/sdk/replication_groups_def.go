package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var replicationGroupObjectTypes = g.NewQueryStruct("ReplicationGroupObjectTypes").
	OptionalSQL("ACCOUNT PARAMETERS").
	OptionalSQL("DATABASES").
	OptionalSQL("INTEGRATIONS").
	OptionalSQL("NETWORK POLICIES").
	OptionalSQL("RESOURCE MONITORS").
	OptionalSQL("ROLES").
	OptionalSQL("SHARES").
	OptionalSQL("USERS").
	OptionalSQL("WAREHOUSES").
	WithValidation(g.AtLeastOneValueSet, "AccountParameters", "Databases", "Integrations", "NetworkPolicies", "ResourceMonitors", "Roles", "Shares", "Users", "Warehouses")

var (
	replicationGroupDatabase        = g.NewQueryStruct("ReplicationGroupDatabase").Text("Database", g.KeywordOptions())
	replicationGroupShare           = g.NewQueryStruct("ReplicationGroupShare").Text("Share", g.KeywordOptions())
	replicationGroupAccount         = g.NewQueryStruct("ReplicationGroupAccount").Text("Account", g.KeywordOptions())
	replicationGroupIntegrationType = g.NewQueryStruct("ReplicationGroupIntegrationType").Text("IntegrationType", g.KeywordOptions())
)

var ReplicationGroupsDef = g.NewInterface(
	"ReplicationGroups",
	"ReplicationGroup",
	g.KindOfT[AccountObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-replication-group",
	g.NewQueryStruct("CreateReplicationGroup").
		Create().
		SQL("REPLICATION GROUP").
		IfNotExists().
		Name().
		QueryStructField(
			"ObjectTypes",
			replicationGroupObjectTypes,
			g.ListOptions().NoParentheses().SQL("OBJECT_TYPES ="),
		).
		ListQueryStructField(
			"Databases",
			replicationGroupDatabase,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_DATABASES"),
		).
		ListQueryStructField(
			"Shares",
			replicationGroupShare,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_SHARES"),
		).
		ListQueryStructField(
			"IntegrationTypes",
			replicationGroupIntegrationType,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_INTEGRATION_TYPES"),
		).
		ListQueryStructField(
			"Accounts",
			replicationGroupAccount,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_ACCOUNTS"),
		).
		OptionalSQL("IGNORE EDITION CHECK").
		WithValidation(g.ValidIdentifier, "name"),
)
