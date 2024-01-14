package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var networkIdentifier = g.NewQueryStruct("NetworkIdentifier").Text("Value", g.KeywordOptions().SingleQuotes())

var NetworkRulesDef = g.NewInterface(
	"NetworkRules",
	"NetworkRule",
	g.KindOfT[SchemaObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-network-rule",
	g.NewQueryStruct("CreateNetworkRule").
		Create().
		OrReplace().
		SQL("NETWORK RULE").
		Name().
		PredefinedQueryStructField("NetworkIdentifierType", "*NetworkIdentifierType", g.ParameterOptions().SQL("TYPE")).
		ListQueryStructField(
			"NetworkIdentifiers",
			networkIdentifier,
			g.ParameterOptions().MustParentheses().SQL("VALUE_LIST"),
		).
		PredefinedQueryStructField("NetworkRuleMode", "*NetworkRuleMode", g.ParameterOptions().SQL("MODE")).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		WithValidation(g.ValidIdentifier, "name"),
)
