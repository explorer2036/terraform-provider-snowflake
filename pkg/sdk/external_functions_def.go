package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var externalFunctionArgument = g.NewQueryStruct("ExternalFunctionArgument").
	Text("ArgName", g.KeywordOptions().NoQuotes().Required()).
	PredefinedQueryStructField("ArgDataType", "DataType", g.KeywordOptions().NoQuotes().Required())

var externalFunctionHeader = g.NewQueryStruct("ExternalFunctionHeader").
	Text("Name", g.KeywordOptions().SingleQuotes().Required()).
	PredefinedQueryStructField("Value", "string", g.ParameterOptions().SingleQuotes().Required())

var externalFunctionContextHeader = g.NewQueryStruct("ExternalFunctionContextHeader").Text("ContextFunction", g.KeywordOptions().NoQuotes())

var ExternalFunctionsDef = g.NewInterface(
	"ExternalFunctions",
	"ExternalFunction",
	g.KindOfT[SchemaObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-external-function",
	g.NewQueryStruct("CreateExternalFunction").
		Create().
		OrReplace().
		OptionalSQL("SECURE").
		SQL("EXTERNAL FUNCTION").
		Name().
		ListQueryStructField(
			"Arguments",
			externalFunctionArgument,
			g.ListOptions().MustParentheses()).
		PredefinedQueryStructField("ResultDataType", "DataType", g.ParameterOptions().NoEquals().SQL("RETURNS").Required()).
		PredefinedQueryStructField("ReturnNullValues", "*ReturnNullValues", g.KeywordOptions()).
		PredefinedQueryStructField("NullInputBehavior", "*NullInputBehavior", g.KeywordOptions()).
		PredefinedQueryStructField("ReturnResultsBehavior", "*ReturnResultsBehavior", g.KeywordOptions()).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		Identifier("ApiIntegration", g.KindOfTPointer[AccountObjectIdentifier](), g.IdentifierOptions().SQL("API_INTEGRATION =")).
		ListQueryStructField(
			"Headers",
			externalFunctionHeader,
			g.ParameterOptions().Parentheses().SQL("HEADERS"),
		).
		ListQueryStructField(
			"ContextHeaders",
			externalFunctionContextHeader,
			g.ParameterOptions().Parentheses().SQL("CONTEXT_HEADERS"),
		).
		OptionalNumberAssignment("MAX_BATCH_ROWS", g.ParameterOptions()).
		OptionalTextAssignment("COMPRESSION", g.ParameterOptions()).
		OptionalIdentifier("RequestTranslator", g.KindOfTPointer[SchemaObjectIdentifier](), g.IdentifierOptions().SQL("REQUEST_TRANSLATOR =")).
		OptionalIdentifier("ResponseTranslator", g.KindOfTPointer[SchemaObjectIdentifier](), g.IdentifierOptions().SQL("RESPONSE_TRANSLATOR =")).
		TextAssignment("AS", g.ParameterOptions().SingleQuotes().Required()).
		WithValidation(g.ValidIdentifier, "name"),
)
