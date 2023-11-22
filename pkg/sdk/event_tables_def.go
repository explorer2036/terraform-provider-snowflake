package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var EventTablesDef = g.NewInterface(
	"EventTables",
	"EventTable",
	g.KindOfT[SchemaObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-event-table",
	g.NewQueryStruct("CreateEventTable").
		Create().
		OrReplace().
		SQL("EVENT TABLE").
		IfNotExists().
		Name().
		PredefinedQueryStructField("ClusterBy", "[]string", g.ParameterOptions().Parentheses().NoEquals().SQL("CLUSTER BY")).
		WithValidation(g.ValidIdentifier, "name"),
)
