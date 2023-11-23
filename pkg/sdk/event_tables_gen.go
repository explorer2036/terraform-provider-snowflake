package sdk

import "context"

type EventTables interface {
	Create(ctx context.Context, request *CreateEventTableRequest) error
	Show(ctx context.Context, request *ShowEventTableRequest) ([]EventTable, error)
}

// CreateEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-event-table.
type CreateEventTableOptions struct {
	create                     bool                   `ddl:"static" sql:"CREATE"`
	OrReplace                  *bool                  `ddl:"keyword" sql:"OR REPLACE"`
	eventTable                 bool                   `ddl:"static" sql:"EVENT TABLE"`
	IfNotExists                *bool                  `ddl:"keyword" sql:"IF NOT EXISTS"`
	name                       SchemaObjectIdentifier `ddl:"identifier"`
	ClusterBy                  []string               `ddl:"parameter,parentheses,no_equals" sql:"CLUSTER BY"`
	DataRetentionTimeInDays    *int                   `ddl:"parameter" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *int                   `ddl:"parameter" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool                  `ddl:"parameter" sql:"CHANGE_TRACKING"`
	DefaultDdlCollation        *string                `ddl:"parameter,single_quotes" sql:"DEFAULT_DDL_COLLATION"`
	CopyGrants                 *bool                  `ddl:"keyword" sql:"COPY GRANTS"`
	Comment                    *string                `ddl:"parameter,single_quotes" sql:"COMMENT"`
	RowAccessPolicy            *RowAccessPolicy       `ddl:"keyword"`
	Tag                        []TagAssociation       `ddl:"keyword,parentheses" sql:"TAG"`
}

// ShowEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-event-tables.
type ShowEventTableOptions struct {
	show        bool    `ddl:"static" sql:"SHOW"`
	eventTables bool    `ddl:"static" sql:"EVENT TABLES"`
	Like        *Like   `ddl:"keyword" sql:"LIKE"`
	In          *In     `ddl:"keyword" sql:"IN"`
	StartsWith  *string `ddl:"parameter,single_quotes,no_equals" sql:"STARTS WITH"`
	Limit       *int    `ddl:"parameter" sql:"LIMIT"`
	From        *string `ddl:"parameter,single_quotes,no_equals" sql:"FROM"`
}

type eventTableRow struct {
	CreatedOn     string `db:"created_on"`
	Name          string `db:"name"`
	DatabaseName  string `db:"database_name"`
	SchemaName    string `db:"schema_name"`
	Owner         string `db:"owner"`
	Comment       string `db:"comment"`
	OwnerRoleType string `db:"owner_role_type"`
}

type EventTable struct {
	CreatedOn     string
	Name          string
	DatabaseName  string
	SchemaName    string
	Owner         string
	Comment       string
	OwnerRoleType string
}
