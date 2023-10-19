package sdk

import "context"

type ApplicationPackages interface {
	Create(ctx context.Context, request *CreateApplicationPackageRequest) error
	Alter(ctx context.Context, request *AlterApplicationPackageRequest) error
	Drop(ctx context.Context, request *DropApplicationPackageRequest) error
	Show(ctx context.Context, request *ShowApplicationPackageRequest) ([]ApplicationPackage, error)
	ShowByID(ctx context.Context, id AccountObjectIdentifier) (*ApplicationPackage, error)
}

// CreateApplicationPackageOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-application-package.
type CreateApplicationPackageOptions struct {
	create                     bool                    `ddl:"static" sql:"CREATE"`
	applicationPackage         bool                    `ddl:"static" sql:"APPLICATION PACKAGE"`
	IfNotExists                *bool                   `ddl:"keyword" sql:"IF NOT EXISTS"`
	name                       AccountObjectIdentifier `ddl:"identifier"`
	DataRetentionTimeInDays    *int                    `ddl:"parameter,no_quotes" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *int                    `ddl:"parameter,no_quotes" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	DefaultDdlCollation        *string                 `ddl:"parameter,single_quotes" sql:"DEFAULT_DDL_COLLATION"`
	Comment                    *string                 `ddl:"parameter,single_quotes" sql:"COMMENT"`
	Distribution               *string                 `ddl:"parameter,no_quotes" sql:"DISTRIBUTION"`
	Tag                        []TagAssociation        `ddl:"keyword,parentheses" sql:"TAG"`
}

// AlterApplicationPackageOptions is based on https://docs.snowflake.com/en/sql-reference/sql/alter-application-package.
type AlterApplicationPackageOptions struct {
	alter              bool                     `ddl:"static" sql:"ALTER"`
	applicationPackage bool                     `ddl:"static" sql:"APPLICATION PACKAGE"`
	IfExists           *bool                    `ddl:"keyword" sql:"IF EXISTS"`
	name               AccountObjectIdentifier  `ddl:"identifier"`
	Set                *ApplicationPackageSet   `ddl:"keyword" sql:"SET"`
	Unset              *ApplicationPackageUnset `ddl:"keyword" sql:"UNSET"`
}

type ApplicationPackageSet struct {
	DataRetentionTimeInDays    *int    `ddl:"parameter,no_quotes" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *int    `ddl:"parameter,no_quotes" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	DefaultDdlCollation        *string `ddl:"parameter,single_quotes" sql:"DEFAULT_DDL_COLLATION"`
	Comment                    *string `ddl:"parameter,single_quotes" sql:"COMMENT"`
	Distribution               *string `ddl:"parameter,no_quotes" sql:"DISTRIBUTION"`
}

type ApplicationPackageUnset struct {
	DataRetentionTimeInDays    *bool `ddl:"keyword" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *bool `ddl:"keyword" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	DefaultDdlCollation        *bool `ddl:"keyword" sql:"DEFAULT_DDL_COLLATION"`
	Comment                    *bool `ddl:"keyword" sql:"COMMENT"`
	Distribution               *bool `ddl:"keyword" sql:"DISTRIBUTION"`
}

// DropApplicationPackageOptions is based on https://docs.snowflake.com/en/sql-reference/sql/drop-application-package.
type DropApplicationPackageOptions struct {
	drop               bool                    `ddl:"static" sql:"DROP"`
	applicationPackage bool                    `ddl:"static" sql:"APPLICATION PACKAGE"`
	name               AccountObjectIdentifier `ddl:"identifier"`
}

// ShowApplicationPackageOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-application-packages.
type ShowApplicationPackageOptions struct {
	show                bool       `ddl:"static" sql:"SHOW"`
	applicationPackages bool       `ddl:"static" sql:"APPLICATION PACKAGES"`
	Like                *Like      `ddl:"keyword" sql:"LIKE"`
	StartsWith          *string    `ddl:"parameter,no_equals,single_quotes" sql:"STARTS WITH"`
	Limit               *LimitFrom `ddl:"keyword" sql:"LIMIT"`
}

type applicationPackageRow struct {
	CreatedOn    string `db:"created_on"`
	Name         string `db:"name"`
	Distribution string `db:"distribution"`
	Owner        string `db:"owner"`
	Comment      string `db:"comment"`
	Options      string `db:"options"`
}

type ApplicationPackage struct {
	CreatedOn    string
	Name         string
	Distribution string
	Owner        string
	Comment      string
	Options      string
}
