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

// [ REPLICATION_SCHEDULE = '{ <num> MINUTE | USING CRON <expr> <time_zone> }' ]
var replicationGroupSchedule = g.NewQueryStruct("ReplicationGroupSchedule").
	OptionalQueryStructField("IntervalMinutes",
		g.NewQueryStruct("ScheduleIntervalMinutes").
			Number("Minutes", g.KeywordOptions().Required()).
			SQL("MINUTES"),
		g.KeywordOptions(),
	).
	OptionalQueryStructField("CronExpression",
		g.NewQueryStruct("ScheduleCronExpression").
			SQL("USING CRON").
			Text("Expression", g.KeywordOptions().Required()).
			Text("TimeZone", g.KeywordOptions().Required()),
		g.KeywordOptions(),
	).
	WithValidation(g.ExactlyOneValueSet, "IntervalMinutes", "CronExpression")

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
			"AllowedDatabases",
			replicationGroupDatabase,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_DATABASES"),
		).
		ListQueryStructField(
			"AllowedShares",
			replicationGroupShare,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_SHARES"),
		).
		ListQueryStructField(
			"AllowedIntegrationTypes",
			replicationGroupIntegrationType,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_INTEGRATION_TYPES"),
		).
		ListQueryStructField(
			"AllowedAccounts",
			replicationGroupAccount,
			g.ParameterOptions().NoParentheses().SQL("ALLOWED_ACCOUNTS"),
		).
		OptionalSQL("IGNORE EDITION CHECK").
		OptionalQueryStructField(
			"ReplicationSchedule",
			replicationGroupSchedule,
			g.ParameterOptions().SingleQuotes().SQL("REPLICATION_SCHEDULE"),
		).
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ValidateValueSet, "ObjectTypes").
		WithValidation(g.ValidateValueSet, "AllowedAccounts"),
)
