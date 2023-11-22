package sdk

import "context"

type EventTables interface {
	Create(ctx context.Context, request *CreateEventTableRequest) error
}

// CreateEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-event-table.
type CreateEventTableOptions struct {
	create      bool                   `ddl:"static" sql:"CREATE"`
	OrReplace   *bool                  `ddl:"keyword" sql:"OR REPLACE"`
	eventTable  bool                   `ddl:"static" sql:"EVENT TABLE"`
	IfNotExists *bool                  `ddl:"keyword" sql:"IF NOT EXISTS"`
	name        SchemaObjectIdentifier `ddl:"identifier"`
	ClusterBy   []string               `ddl:"parameter,parentheses,no_equals" sql:"CLUSTER BY"`
}
