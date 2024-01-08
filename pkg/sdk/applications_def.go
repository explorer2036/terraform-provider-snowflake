package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var versionAndPatch = g.NewQueryStruct("VersionAndPatch").
	TextAssignment("VERSION", g.ParameterOptions().NoEquals().NoQuotes().Required()).
	OptionalNumberAssignment("PATCH", g.ParameterOptions().NoEquals().Required())

var applicationVersion = g.NewQueryStruct("ApplicationVersion").
	OptionalText("VersionDirectory", g.KeywordOptions().SingleQuotes().Required()).
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
			"Version",
			applicationVersion,
			g.KeywordOptions().SQL("USING"),
		).
		OptionalBooleanAssignment("DEBUG_MODE", g.ParameterOptions()).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		OptionalTags().
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ValidIdentifier, "PackageName"),
).DropOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/drop-application",
	g.NewQueryStruct("DropApplication").
		Drop().
		SQL("APPLICATION").
		IfExists().
		Name().
		OptionalSQL("CASCADE").
		WithValidation(g.ValidIdentifier, "name"),
).ShowOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/show-applications",
	g.DbStruct("applicationRow").
		Field("created_on", "string").
		Field("name", "string").
		Field("is_default", "string").
		Field("is_current", "string").
		Field("source_type", "string").
		Field("source", "string").
		Field("owner", "string").
		Field("comment", "string").
		Field("version", "string").
		Field("label", "string").
		Field("patch", "int").
		Field("options", "string").
		Field("retention_time", "int"),
	g.PlainStruct("Application").
		Field("CreatedOn", "string").
		Field("Name", "string").
		Field("IsDefault", "bool").
		Field("IsCurrent", "bool").
		Field("SourceType", "string").
		Field("Source", "string").
		Field("Owner", "string").
		Field("Comment", "string").
		Field("Version", "string").
		Field("Label", "string").
		Field("Patch", "int").
		Field("Options", "string").
		Field("RetentionTime", "int"),
	g.NewQueryStruct("ShowApplications").
		Show().
		SQL("APPLICATIONS").
		OptionalLike().
		OptionalStartsWith().
		OptionalLimit(),
).ShowByIdOperation().DescribeOperation(
	g.DescriptionMappingKindSlice,
	"https://docs.snowflake.com/en/sql-reference/sql/desc-application",
	g.DbStruct("applicationDetailRow").
		Field("property", "string").
		Field("value", "string"),
	g.PlainStruct("ApplicationDetail").
		Field("Property", "string").
		Field("Value", "string"),
	g.NewQueryStruct("DescribeApplication").
		Describe().
		SQL("APPLICATION").
		Name().
		WithValidation(g.ValidIdentifier, "name"),
)
