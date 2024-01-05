package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

// https://medium.com/snowflake/hello-world-snowflake-native-apps-76e1ce82e1be

var applicationPackageModifyReleaseDirective = g.NewQueryStruct("ModifyReleaseDirective").
	Text("ReleaseDirective", g.KeywordOptions().NoQuotes().Required()).
	TextAssignment("VERSION", g.ParameterOptions().NoQuotes().Required()).
	NumberAssignment("PATCH", g.ParameterOptions().NoQuotes().Required())

var applicationPackageSetReleaseDirective = g.NewQueryStruct("SetReleaseDirective").
	Text("ReleaseDirective", g.KeywordOptions().NoQuotes().Required()).
	PredefinedQueryStructField("Accounts", "[]string", g.ParameterOptions().MustParentheses().NoQuotes().Required().SQL("ACCOUNTS")).
	TextAssignment("VERSION", g.ParameterOptions().NoQuotes().Required()).
	NumberAssignment("PATCH", g.ParameterOptions().NoQuotes().Required())

var applicationPackageSetDefaultReleaseDirective = g.NewQueryStruct("SetDefaultReleaseDirective").
	TextAssignment("VERSION", g.ParameterOptions().NoQuotes().Required()).
	NumberAssignment("PATCH", g.ParameterOptions().NoQuotes().Required())

var applicationPackageSet = g.NewQueryStruct("ApplicationPackageSet").
	OptionalNumberAssignment("DATA_RETENTION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
	OptionalNumberAssignment("MAX_DATA_EXTENSION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
	OptionalTextAssignment("DEFAULT_DDL_COLLATION", g.ParameterOptions().SingleQuotes()).
	OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
	PredefinedQueryStructField("Distribution", "*Distribution", g.ParameterOptions().SQL("DISTRIBUTION"))

var applicationPackageUnset = g.NewQueryStruct("ApplicationPackageUnset").
	OptionalSQL("DATA_RETENTION_TIME_IN_DAYS").
	OptionalSQL("MAX_DATA_EXTENSION_TIME_IN_DAYS").
	OptionalSQL("DEFAULT_DDL_COLLATION").
	OptionalSQL("COMMENT").
	OptionalSQL("DISTRIBUTION")

var ApplicationPackagesDef = g.NewInterface(
	"ApplicationPackages",
	"ApplicationPackage",
	g.KindOfT[AccountObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-application-package",
	g.NewQueryStruct("CreateApplicationPackage").
		Create().
		SQL("APPLICATION PACKAGE").
		IfNotExists().
		Name().
		OptionalNumberAssignment("DATA_RETENTION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
		OptionalNumberAssignment("MAX_DATA_EXTENSION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
		OptionalTextAssignment("DEFAULT_DDL_COLLATION", g.ParameterOptions().SingleQuotes()).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		PredefinedQueryStructField("Distribution", "*Distribution", g.ParameterOptions().SQL("DISTRIBUTION")).
		OptionalTags().
		WithValidation(g.ValidIdentifier, "name"),
).AlterOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/alter-application-package",
	g.NewQueryStruct("AlterApplicationPackage").
		Alter().
		SQL("APPLICATION PACKAGE").
		IfExists().
		Name().
		OptionalQueryStructField(
			"Set",
			applicationPackageSet,
			g.KeywordOptions().SQL("SET"),
		).
		OptionalQueryStructField(
			"Unset",
			applicationPackageUnset,
			g.KeywordOptions().SQL("UNSET"),
		).
		OptionalQueryStructField(
			"ModifyReleaseDirective",
			applicationPackageModifyReleaseDirective,
			g.KeywordOptions().SQL("MODIFY RELEASE DIRECTIVE"),
		).
		OptionalQueryStructField(
			"SetDefaultReleaseDirective",
			applicationPackageSetDefaultReleaseDirective,
			g.KeywordOptions().SQL("SET DEFAULT RELEASE DIRECTIVE"),
		).
		OptionalQueryStructField(
			"SetReleaseDirective",
			applicationPackageSetReleaseDirective,
			g.KeywordOptions().SQL("SET RELEASE DIRECTIVE"),
		).
		OptionalTextAssignment("UNSET RELEASE DIRECTIVE", g.ParameterOptions().NoEquals().NoQuotes()).
		OptionalSetTags().
		OptionalUnsetTags().
		WithValidation(g.ValidIdentifier, "name"),
).DropOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/drop-application-package",
	g.NewQueryStruct("DropApplicationPackage").
		Drop().
		SQL("APPLICATION PACKAGE").
		Name().
		WithValidation(g.ValidIdentifier, "name"),
).ShowOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/show-application-packages",
	g.DbStruct("applicationPackageRow").
		Field("created_on", "string").
		Field("name", "string").
		Field("is_default", "string").
		Field("is_current", "string").
		Field("distribution", "string").
		Field("owner", "string").
		Field("comment", "string").
		Field("retention_time", "int").
		Field("options", "string").
		Field("dropped_on", "sql.NullString").
		Field("application_class", "sql.NullString"),
	g.PlainStruct("ApplicationPackage").
		Field("CreatedOn", "string").
		Field("Name", "string").
		Field("IsDefault", "bool").
		Field("IsCurrent", "bool").
		Field("Distribution", "string").
		Field("Owner", "string").
		Field("Comment", "string").
		Field("RetentionTime", "int").
		Field("Options", "string").
		Field("DroppedOn", "string").
		Field("ApplicationClass", "string"),
	g.NewQueryStruct("ShowApplicationPackages").
		Show().
		SQL("APPLICATION PACKAGES").
		OptionalLike().
		OptionalStartsWith().
		OptionalLimit(),
)
