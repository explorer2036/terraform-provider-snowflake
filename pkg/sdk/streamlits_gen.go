package sdk

import "context"

type Streamlits interface {
	Create(ctx context.Context, request *CreateStreamlitRequest) error
	Alter(ctx context.Context, request *AlterStreamlitRequest) error
}

// CreateStreamlitOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-streamlit.
type CreateStreamlitOptions struct {
	OrReplace    *bool                    `ddl:"keyword" sql:"OR REPLACE"`
	streamlit    bool                     `ddl:"static" sql:"STREAMLIT"`
	IfNotExists  *bool                    `ddl:"keyword" sql:"IF NOT EXISTS"`
	name         SchemaObjectIdentifier   `ddl:"identifier"`
	RootLocation string                   `ddl:"parameter,single_quotes" sql:"ROOT_LOCATION"`
	MainFile     string                   `ddl:"parameter,single_quotes" sql:"MAIN_FILE"`
	Warehouse    *AccountObjectIdentifier `ddl:"identifier,equals" sql:"QUERY_WAREHOUSE"`
	Comment      *string                  `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

// AlterStreamlitOptions is based on https://docs.snowflake.com/en/sql-reference/sql/alter-streamlit.
type AlterStreamlitOptions struct {
	alter     bool                    `ddl:"static" sql:"ALTER"`
	streamlit bool                    `ddl:"static" sql:"STREAMLIT"`
	IfExists  *bool                   `ddl:"keyword" sql:"IF EXISTS"`
	name      SchemaObjectIdentifier  `ddl:"identifier"`
	Set       *StreamlitsSet          `ddl:"keyword" sql:"SET"`
	RenameTo  *SchemaObjectIdentifier `ddl:"identifier" sql:"RENAME TO"`
}

type StreamlitsSet struct {
	RootLocation *string                  `ddl:"parameter,single_quotes" sql:"ROOT_LOCATION"`
	Warehouse    *AccountObjectIdentifier `ddl:"identifier,equals" sql:"QUERY_WAREHOUSE"`
	MainFile     *string                  `ddl:"parameter,single_quotes" sql:"MAIN_FILE"`
	Comment      *string                  `ddl:"parameter,single_quotes" sql:"COMMENT"`
}
