package sdk

import "context"

type Applications interface {
	Create(ctx context.Context, request *CreateApplicationRequest) error
}

// CreateApplicationOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-application.
type CreateApplicationOptions struct {
	create                 bool                    `ddl:"static" sql:"CREATE"`
	application            bool                    `ddl:"static" sql:"APPLICATION"`
	name                   AccountObjectIdentifier `ddl:"identifier"`
	fromApplicationPackage bool                    `ddl:"static" sql:"FROM APPLICATION PACKAGE"`
	PackageName            AccountObjectIdentifier `ddl:"identifier"`
	ApplicationVersion     *ApplicationVersion     `ddl:"keyword" sql:"USING"`
	DebugMode              *bool                   `ddl:"parameter" sql:"DEBUG_MODE"`
	Comment                *string                 `ddl:"parameter,single_quotes" sql:"COMMENT"`
	Tag                    []TagAssociation        `ddl:"keyword,parentheses" sql:"TAG"`
}

type ApplicationVersion struct {
	VersionDirectory *string          `ddl:"keyword,no_quotes"`
	VersionAndPatch  *VersionAndPatch `ddl:"keyword,no_quotes"`
}

type VersionAndPatch struct {
	Version string `ddl:"parameter,no_quotes,no_equals" sql:"VERSION"`
	Patch   int    `ddl:"parameter,no_equals" sql:"PATCH"`
}
