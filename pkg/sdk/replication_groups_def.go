package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var dbReplicationGroupRow = g.DbStruct("replicationGroupRow").
	Field("snowflake_region", "string").
	Field("created_on", "string").
	Field("account_name", "string").
	Field("name", "string").
	Field("type", "string").
	Field("comment", "sql.NullString").
	Field("is_primary", "string").
	Field("primary", "string").
	Field("object_types", "string").
	Field("allowed_integration_types", "string").
	Field("allowed_accounts", "string").
	Field("organization_name", "string").
	Field("account_locator", "string").
	Field("replication_schedule", "string").
	Field("secondary_state", "sql.NullString").
	Field("next_scheduled_refresh", "sql.NullString").
	Field("owner", "string")

var plainReplicationGroup = g.PlainStruct("ReplicationGroup").
	Field("SnowflakeRegion", "string").
	Field("CreatedOn", "string").
	Field("AccountName", "string").
	Field("Name", "string").
	Field("Type", "string").
	Field("Comment", "string").
	Field("IsPrimary", "bool").
	Field("Primary", "string").
	Field("ObjectTypes", "string").
	Field("AllowedIntegrationTypes", "string").
	Field("AllowedAccounts", "string").
	Field("OrganizationName", "string").
	Field("AccountLocator", "string").
	Field("ReplicationSchedule", "string").
	Field("SecondaryState", "string").
	Field("NextScheduledRefresh", "string").
	Field("Owner", "string")

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

var replicationGroupSet = g.NewQueryStruct("ReplicationGroupSet").
	OptionalQueryStructField(
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
	OptionalQueryStructField(
		"ReplicationSchedule",
		replicationGroupSchedule,
		g.ParameterOptions().SingleQuotes().SQL("REPLICATION_SCHEDULE"),
	).
	OptionalBooleanAssignment("ENABLE_ETL_REPLICATION", g.ParameterOptions())

var replicationGroupSetIntegration = g.NewQueryStruct("ReplicationGroupSetIntegration").
	OptionalQueryStructField(
		"ObjectTypes",
		replicationGroupObjectTypes,
		g.ListOptions().NoParentheses().SQL("OBJECT_TYPES ="),
	).
	ListQueryStructField(
		"AllowedIntegrationTypes",
		replicationGroupIntegrationType,
		g.ParameterOptions().NoParentheses().SQL("ALLOWED_INTEGRATION_TYPES"),
	).
	OptionalQueryStructField(
		"ReplicationSchedule",
		replicationGroupSchedule,
		g.ParameterOptions().SingleQuotes().SQL("REPLICATION_SCHEDULE"),
	)

var replicationGroupAddDatabases = g.NewQueryStruct("ReplicationGroupAddDatabases").
	ListQueryStructField(
		"Databases",
		replicationGroupDatabase,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("ADD"),
	).
	SQL("TO ALLOWED_DATABASES")

var replicationGroupRemoveDatabases = g.NewQueryStruct("ReplicationGroupRemoveDatabases").
	ListQueryStructField(
		"Databases",
		replicationGroupDatabase,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("REMOVE"),
	).
	SQL("FROM ALLOWED_DATABASES")

var replicationGroupMoveDatabases = g.NewQueryStruct("ReplicationGroupMoveDatabases").
	ListQueryStructField(
		"Databases",
		replicationGroupDatabase,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("MOVE DATABASES"),
	).
	Identifier("MoveTo", g.KindOfTPointer[AccountObjectIdentifier](), g.IdentifierOptions().SQL("TO REPLICATION GROUP"))

var replicationGroupAddShares = g.NewQueryStruct("ReplicationGroupAddShares").
	ListQueryStructField(
		"Shares",
		replicationGroupShare,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("ADD"),
	).
	SQL("TO ALLOWED_SHARES")

var replicationGroupRemoveShares = g.NewQueryStruct("ReplicationGroupRemoveShares").
	ListQueryStructField(
		"Shares",
		replicationGroupShare,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("REMOVE"),
	).
	SQL("FROM ALLOWED_SHARES")

var replicationGroupMoveShares = g.NewQueryStruct("ReplicationGroupMoveShares").
	ListQueryStructField(
		"Shares",
		replicationGroupShare,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("MOVE SHARES"),
	).
	Identifier("MoveTo", g.KindOfTPointer[AccountObjectIdentifier](), g.IdentifierOptions().SQL("TO REPLICATION GROUP"))

var replicationGroupAddAccounts = g.NewQueryStruct("ReplicationGroupAddAccounts").
	ListQueryStructField(
		"Accounts",
		replicationGroupAccount,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("ADD"),
	).
	SQL("TO ALLOWED_ACCOUNTS").
	OptionalSQL("IGNORE EDITION CHECK")

var replicationGroupRemoveAccounts = g.NewQueryStruct("ReplicationGroupRemoveAccounts").
	ListQueryStructField(
		"Accounts",
		replicationGroupAccount,
		g.ParameterOptions().NoParentheses().NoEquals().SQL("REMOVE"),
	).
	SQL("FROM ALLOWED_ACCOUNTS")

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
).CustomOperation(
	"CreateSecondary",
	"https://docs.snowflake.com/en/sql-reference/sql/create-replication-group",
	g.NewQueryStruct("CreateSecondary").
		Create().
		SQL("REPLICATION GROUP").
		IfNotExists().
		Name().
		SQL("AS REPLICA OF").
		Identifier("Primary", g.KindOfTPointer[ExternalObjectIdentifier](), g.IdentifierOptions().Required()).
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ValidIdentifier, "Primary"),
).AlterOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/alter-replication-group",
	g.NewQueryStruct("AlterReplicationGroup").
		Alter().
		SQL("REPLICATION GROUP").
		IfExists().
		Name().
		Identifier("RenameTo", g.KindOfTPointer[SchemaObjectIdentifier](), g.IdentifierOptions().SQL("RENAME TO")).
		OptionalQueryStructField(
			"Set",
			replicationGroupSet,
			g.KeywordOptions().SQL("SET"),
		).
		OptionalQueryStructField(
			"SetIntegration",
			replicationGroupSetIntegration,
			g.KeywordOptions().SQL("SET"),
		).
		OptionalQueryStructField(
			"AddDatabases",
			replicationGroupAddDatabases,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"RemoveDatabases",
			replicationGroupRemoveDatabases,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"MoveDatabases",
			replicationGroupMoveDatabases,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"AddShares",
			replicationGroupAddShares,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"RemoveShares",
			replicationGroupRemoveShares,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"MoveShares",
			replicationGroupMoveShares,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"AddAccounts",
			replicationGroupAddAccounts,
			g.KeywordOptions(),
		).
		OptionalQueryStructField(
			"RemoveAccounts",
			replicationGroupRemoveAccounts,
			g.KeywordOptions(),
		).
		OptionalSQL("REFRESH").
		OptionalSQL("SUSPEND").
		OptionalSQL("RESUME").
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ValidIdentifierIfSet, "RenameTo").
		WithValidation(g.ExactlyOneValueSet, "Set", "SetIntegration", "AddDatabases", "RemoveDatabases", "MoveDatabases", "AddShares", "RemoveShares", "MoveShares", "AddAccounts", "RemoveAccounts", "Refresh", "Suspend", "Resume"),
).ShowOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/show-replication-groups",
	dbReplicationGroupRow,
	plainReplicationGroup,
	g.NewQueryStruct("ShowReplicationGroups").
		Show().
		SQL("REPLICATION GROUPS").
		OptionalIdentifier("InAccount", g.KindOfTPointer[AccountObjectIdentifier](), g.IdentifierOptions().SQL("IN ACCOUNT")),
).ShowByIdOperation().DropOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/drop-replication-group",
	g.NewQueryStruct("DropReplicationGroup").
		Drop().
		SQL("REPLICATION GROUP").
		IfExists().
		Name().
		WithValidation(g.ValidIdentifier, "name"),
)
