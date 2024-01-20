package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var streamlitsSet = g.NewQueryStruct("StreamlitsSet").
	OptionalTextAssignment("ROOT_LOCATION", g.ParameterOptions().SingleQuotes().Required()).
	OptionalIdentifier("Warehouse", g.KindOfT[AccountObjectIdentifier](), g.IdentifierOptions().Equals().SQL("QUERY_WAREHOUSE")).
	OptionalTextAssignment("MAIN_FILE", g.ParameterOptions().SingleQuotes().Required()).
	OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes())

var StreamlitsDef = g.NewInterface(
	"Streamlits",
	"Streamlit",
	g.KindOfT[SchemaObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-streamlit",
	g.NewQueryStruct("CreateStreamlit").
		OrReplace().
		SQL("STREAMLIT").
		IfNotExists().
		Name().
		TextAssignment("ROOT_LOCATION", g.ParameterOptions().SingleQuotes().Required()).
		TextAssignment("MAIN_FILE", g.ParameterOptions().SingleQuotes().Required()).
		OptionalIdentifier("Warehouse", g.KindOfT[AccountObjectIdentifier](), g.IdentifierOptions().Equals().SQL("QUERY_WAREHOUSE")).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ConflictingFields, "IfNotExists", "OrReplace"),
).AlterOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/alter-streamlit",
	g.NewQueryStruct("AlterStreamlit").
		Alter().
		SQL("STREAMLIT").
		IfExists().
		Name().
		OptionalQueryStructField(
			"Set",
			streamlitsSet,
			g.KeywordOptions().SQL("SET"),
		).
		Identifier("RenameTo", g.KindOfTPointer[SchemaObjectIdentifier](), g.IdentifierOptions().SQL("RENAME TO")).
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ExactlyOneValueSet, "RenameTo", "Set"),
)
