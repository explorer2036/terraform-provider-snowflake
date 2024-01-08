package sdk

import "context"

type Applications interface {
	Create(ctx context.Context, request *CreateApplicationRequest) error
	Drop(ctx context.Context, request *DropApplicationRequest) error
	Show(ctx context.Context, request *ShowApplicationRequest) ([]ApplicationPackage, error)
	ShowByID(ctx context.Context, id AccountObjectIdentifier) (*Application, error)
	Describe(ctx context.Context, id AccountObjectIdentifier) ([]ApplicationDetail, error)
}

// CreateApplicationOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-application.
type CreateApplicationOptions struct {
	create                 bool                    `ddl:"static" sql:"CREATE"`
	application            bool                    `ddl:"static" sql:"APPLICATION"`
	name                   AccountObjectIdentifier `ddl:"identifier"`
	fromApplicationPackage bool                    `ddl:"static" sql:"FROM APPLICATION PACKAGE"`
	PackageName            AccountObjectIdentifier `ddl:"identifier"`
	Version                *ApplicationVersion     `ddl:"keyword" sql:"USING"`
	DebugMode              *bool                   `ddl:"parameter" sql:"DEBUG_MODE"`
	Comment                *string                 `ddl:"parameter,single_quotes" sql:"COMMENT"`
	Tag                    []TagAssociation        `ddl:"keyword,parentheses" sql:"TAG"`
}

type ApplicationVersion struct {
	VersionDirectory *string          `ddl:"keyword,single_quotes"`
	VersionAndPatch  *VersionAndPatch `ddl:"keyword,no_quotes"`
}

type VersionAndPatch struct {
	Version string `ddl:"parameter,no_quotes,no_equals" sql:"VERSION"`
	Patch   *int   `ddl:"parameter,no_equals" sql:"PATCH"`
}

// DropApplicationOptions is based on https://docs.snowflake.com/en/sql-reference/sql/drop-application.
type DropApplicationOptions struct {
	drop        bool                    `ddl:"static" sql:"DROP"`
	application bool                    `ddl:"static" sql:"APPLICATION"`
	IfExists    *bool                   `ddl:"keyword" sql:"IF EXISTS"`
	name        AccountObjectIdentifier `ddl:"identifier"`
	Cascade     *bool                   `ddl:"keyword" sql:"CASCADE"`
}

// ShowApplicationOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-applications.
type ShowApplicationOptions struct {
	show         bool       `ddl:"static" sql:"SHOW"`
	applications bool       `ddl:"static" sql:"APPLICATIONS"`
	Like         *Like      `ddl:"keyword" sql:"LIKE"`
	StartsWith   *string    `ddl:"parameter,single_quotes,no_equals" sql:"STARTS WITH"`
	Limit        *LimitFrom `ddl:"keyword" sql:"LIMIT"`
}

type applicationPackageRow struct {
	CreatedOn     string `db:"created_on"`
	Name          string `db:"name"`
	IsDefault     string `db:"is_default"`
	IsCurrent     string `db:"is_current"`
	SourceType    string `db:"source_type"`
	Source        string `db:"source"`
	Owner         string `db:"owner"`
	Comment       string `db:"comment"`
	Version       string `db:"version"`
	Label         string `db:"label"`
	Patch         int    `db:"patch"`
	Options       string `db:"options"`
	RetentionTime int    `db:"retention_time"`
}

type ApplicationPackage struct {
	CreatedOn     string
	Name          string
	IsDefault     bool
	IsCurrent     bool
	SourceType    string
	Source        string
	Owner         string
	Comment       string
	Version       string
	Label         string
	Patch         int
	Options       string
	RetentionTime int
}

// DescribeApplicationOptions is based on https://docs.snowflake.com/en/sql-reference/sql/desc-application.
type DescribeApplicationOptions struct {
	describe    bool                    `ddl:"static" sql:"DESCRIBE"`
	application bool                    `ddl:"static" sql:"APPLICATION"`
	name        AccountObjectIdentifier `ddl:"identifier"`
}

type applicationDetailRow struct {
	Property string `db:"property"`
	Value    string `db:"value"`
}

type ApplicationDetail struct {
	Property string
	Value    string
}
