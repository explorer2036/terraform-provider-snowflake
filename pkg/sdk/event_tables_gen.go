package sdk

import "context"

type EventTables interface {
	Create(ctx context.Context, request *CreateEventTableRequest) error
	Show(ctx context.Context, request *ShowEventTableRequest) ([]EventTable, error)
	Describe(ctx context.Context, id SchemaObjectIdentifier) (*EventTableDetails, error)
	Alter(ctx context.Context, request *AlterEventTableRequest) error
}

// CreateEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-event-table.
type CreateEventTableOptions struct {
	create                     bool                   `ddl:"static" sql:"CREATE"`
	OrReplace                  *bool                  `ddl:"keyword" sql:"OR REPLACE"`
	eventTable                 bool                   `ddl:"static" sql:"EVENT TABLE"`
	IfNotExists                *bool                  `ddl:"keyword" sql:"IF NOT EXISTS"`
	name                       SchemaObjectIdentifier `ddl:"identifier"`
	ClusterBy                  []string               `ddl:"keyword,parentheses" sql:"CLUSTER BY"`
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

// DescribeEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/describe-event-table.
type DescribeEventTableOptions struct {
	describe   bool                   `ddl:"static" sql:"DESCRIBE"`
	eventTable bool                   `ddl:"static" sql:"EVENT TABLE"`
	name       SchemaObjectIdentifier `ddl:"identifier"`
}

type eventTableDetailsRow struct {
	Name    string `db:"name"`
	Kind    string `db:"kind"`
	Comment string `db:"comment"`
}

type EventTableDetails struct {
	Name    string
	Kind    string
	Comment string
}

// AlterEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/alter-event-table.
type AlterEventTableOptions struct {
	alter                    bool                                `ddl:"static" sql:"ALTER"`
	eventTable               bool                                `ddl:"static" sql:"EVENT TABLE"`
	IfNotExists              *bool                               `ddl:"keyword" sql:"IF NOT EXISTS"`
	name                     SchemaObjectIdentifier              `ddl:"identifier"`
	Set                      *EventTableSet                      `ddl:"keyword" sql:"SET"`
	Unset                    *EventTableUnset                    `ddl:"keyword" sql:"UNSET"`
	AddRowAccessPolicy       *RowAccessPolicy                    `ddl:"keyword" sql:"ADD"`
	DropRowAccessPolicy      *EventTableDropRowAccessPolicy      `ddl:"keyword" sql:"DROP"`
	DropAllRowAccessPolicies *bool                               `ddl:"keyword" sql:"DROP ALL ROW ACCESS POLICIES"`
	ClusteringAction         *EventTableClusteringAction         `ddl:"keyword"`
	SearchOptimizationAction *EventTableSearchOptimizationAction `ddl:"keyword"`
	RenameTo                 *SchemaObjectIdentifier             `ddl:"identifier" sql:"RENAME TO"`
	SetTags                  []TagAssociation                    `ddl:"keyword"`
	UnsetTags                []ObjectIdentifier                  `ddl:"keyword"`
}

type EventTableSet struct {
	DataRetentionTimeInDays    *int    `ddl:"parameter" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *int    `ddl:"parameter" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool   `ddl:"parameter" sql:"CHANGE_TRACKING"`
	Comment                    *string `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

type EventTableUnset struct {
	DataRetentionTimeInDays    *bool `ddl:"keyword" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *bool `ddl:"keyword" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool `ddl:"keyword" sql:"CHANGE_TRACKING"`
	Comment                    *bool `ddl:"keyword" sql:"COMMENT"`
}

type EventTableDropRowAccessPolicy struct {
	rowAccessPolicy bool                   `ddl:"static" sql:"ROW ACCESS POLICY"`
	Name            SchemaObjectIdentifier `ddl:"identifier"`
}

type EventTableClusteringAction struct {
	ClusterBy         *[]string `ddl:"keyword,parentheses" sql:"CLUSTER BY"`
	SuspendRecluster  *bool     `ddl:"keyword" sql:"SUSPEND RECLUSTER"`
	ResumeRecluster   *bool     `ddl:"keyword" sql:"RESUME RECLUSTER"`
	DropClusteringKey *bool     `ddl:"keyword" sql:"DROP CLUSTERING KEY"`
}

type EventTableSearchOptimizationAction struct {
	Add  *SearchOptimization `ddl:"keyword" sql:"ADD"`
	Drop *SearchOptimization `ddl:"keyword" sql:"DROP"`
}

type SearchOptimization struct {
	searchOptimization bool     `ddl:"static" sql:"SEARCH OPTIMIZATION"`
	On                 []string `ddl:"keyword" sql:"ON"`
}
