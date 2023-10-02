package sdk

import "context"

type EventTables interface {
	Create(ctx context.Context, request *CreateEventTableRequest) error
}

// createEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-event-table
type createEventTableOptions struct {
	create      bool   `ddl:"static" sql:"CREATE"`
	OrReplace   *bool  `ddl:"keyword" sql:"OR REPLACE"`
	eventTable  string `ddl:"static" sql:"EVENT TABLE"`
	IfNotExists *bool  `ddl:"keyword" sql:"IF NOT EXISTS"`

	name                       SchemaObjectIdentifier `ddl:"identifier"`
	ClusterBy                  []string               `ddl:"keyword,parentheses" sql:"CLUSTER BY"`
	DataRetentionTimeInDays    *uint                  `ddl:"parameter" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *uint                  `ddl:"parameter" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool                  `ddl:"parameter" sql:"CHANGE_TRACKING"`
	DefaultDDLCollation        *string                `ddl:"parameter,single_quotes" sql:"DEFAULT_DDL_COLLATION"`
	CopyGrants                 *bool                  `ddl:"keyword" sql:"COPY_GRANTS"`
	Comment                    *string                `ddl:"parameter,single_quotes" sql:"COMMENT"`
	RowAccessPolicy            *RowAccessPolicy       `ddl:"keyword"`
	Tag                        []TagAssociation       `ddl:"keyword,parentheses" sql:"TAG"`
}
