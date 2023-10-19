package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var ApplicationPackagesDef = g.NewInterface(
	"ApplicationPackages",
	"ApplicationPackage",
	g.KindOfT[SchemaObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/create-application-package",
	g.QueryStruct("CreateApplicationPackage").
		Create().
		SQL("APPLICATION PACKAGE").
		IfNotExists().
		Name().
		OptionalNumberAssignment("DATA_RETENTION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
		OptionalNumberAssignment("MAX_DATA_EXTENSION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
		OptionalTextAssignment("DEFAULT_DDL_COLLATION", g.ParameterOptions().SingleQuotes()).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		WithTags().
		OptionalTextAssignment("DISTRIBUTION", g.ParameterOptions().NoQuotes()).
		WithValidation(g.ValidIdentifier, "name"),
).AlterOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/alter-application-package",
	g.QueryStruct("AlterApplicationPackage").
		Alter().
		SQL("APPLICATION PACKAGE").
		IfExists().
		Name().
		OptionalQueryStructField(
			"Set",
			g.QueryStruct("ApplicationPackageSet").
				OptionalNumberAssignment("DATA_RETENTION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
				OptionalNumberAssignment("MAX_DATA_EXTENSION_TIME_IN_DAYS", g.ParameterOptions().NoQuotes()).
				OptionalTextAssignment("DEFAULT_DDL_COLLATION", g.ParameterOptions().SingleQuotes()).
				OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
				OptionalTextAssignment("DISTRIBUTION", g.ParameterOptions().NoQuotes()),
			g.KeywordOptions().SQL("SET"),
		).
		OptionalQueryStructField(
			"Unset",
			g.QueryStruct("ApplicationPackageUnset").
				OptionalSQL("DATA_RETENTION_TIME_IN_DAYS").
				OptionalSQL("MAX_DATA_EXTENSION_TIME_IN_DAYS").
				OptionalSQL("DEFAULT_DDL_COLLATION").
				OptionalSQL("COMMENT").
				OptionalSQL("DISTRIBUTION"),
			g.KeywordOptions().SQL("UNSET"),
		),
).DropOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/drop-application-package",
	g.QueryStruct("DropApplicationPackage").
		Drop().
		SQL("APPLICATION PACKAGE").
		Name().
		WithValidation(g.ValidIdentifier, "name"),
).ShowOperation(
	"https://docs.snowflake.com/en/sql-reference/sql/show-application-packages",
	g.DbStruct("applicationPackageRow").
		Field("created_on", "string").
		Field("name", "string").
		Field("distribution", "string").
		Field("owner", "string").
		Field("comment", "string").
		Field("options", "string"),
	g.PlainStruct("ApplicationPackage").
		Field("CreatedOn", "string").
		Field("Name", "string").
		Field("Distribution", "string").
		Field("Owner", "string").
		Field("Comment", "string").
		Field("Options", "string"),
	g.QueryStruct("ShowApplicationPackages").
		Show().
		SQL("APPLICATION PACKAGES").
		OptionalLike().
		OptionalIn().
		OptionalStartsWith().
		OptionalLimit(),
)
