package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var versionAndPatch = g.NewQueryStruct("VersionAndPatch").
	TextAssignment("VERSION", g.ParameterOptions().NoEquals().NoQuotes().Required()).
	NumberAssignment("PATCH", g.ParameterOptions().NoEquals().Required())

var applicationVersion = g.NewQueryStruct("ApplicationVersion").
	OptionalText("VersionDirectory", g.KeywordOptions().NoQuotes()).
	OptionalQueryStructField("VersionAndPatch", versionAndPatch, g.KeywordOptions().NoQuotes()).
	WithValidation(g.ExactlyOneValueSet, "VersionDirectory", "VersionAndPatch")

var ApplicationsDef = g.NewInterface(
	"Applications",
	"Application",
	g.KindOfT[AccountObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-application",
	g.NewQueryStruct("CreateApplication").
		Create().
		SQL("APPLICATION").
		Name().
		SQL("FROM APPLICATION PACKAGE").
		Identifier("PackageName", g.KindOfT[AccountObjectIdentifier](), g.IdentifierOptions().Required()).
		OptionalQueryStructField(
			"ApplicationVersion",
			applicationVersion,
			g.KeywordOptions().SQL("USING"),
		).
		OptionalBooleanAssignment("DEBUG_MODE", g.ParameterOptions()).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		OptionalTags().
		WithValidation(g.ValidIdentifier, "name"),
)
