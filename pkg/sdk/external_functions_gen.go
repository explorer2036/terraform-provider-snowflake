package sdk

import "context"

type ExternalFunctions interface {
	Create(ctx context.Context, request *CreateExternalFunctionRequest) error
}

// CreateExternalFunctionOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-external-function.
type CreateExternalFunctionOptions struct {
	create                bool                            `ddl:"static" sql:"CREATE"`
	OrReplace             *bool                           `ddl:"keyword" sql:"OR REPLACE"`
	Secure                *bool                           `ddl:"keyword" sql:"SECURE"`
	externalFunction      bool                            `ddl:"static" sql:"EXTERNAL FUNCTION"`
	name                  SchemaObjectIdentifier          `ddl:"identifier"`
	Arguments             []ExternalFunctionArgument      `ddl:"list,must_parentheses"`
	ResultDataType        DataType                        `ddl:"parameter,no_equals" sql:"RETURNS"`
	ReturnNullValues      *ReturnNullValues               `ddl:"keyword"`
	NullInputBehavior     *NullInputBehavior              `ddl:"keyword"`
	ReturnResultsBehavior *ReturnResultsBehavior          `ddl:"keyword"`
	Comment               *string                         `ddl:"parameter,single_quotes" sql:"COMMENT"`
	ApiIntegration        *AccountObjectIdentifier        `ddl:"identifier" sql:"API_INTEGRATION ="`
	Headers               []ExternalFunctionHeader        `ddl:"parameter,parentheses" sql:"HEADERS"`
	ContextHeaders        []ExternalFunctionContextHeader `ddl:"parameter,parentheses" sql:"CONTEXT_HEADERS"`
	MaxBatchRows          *int                            `ddl:"parameter" sql:"MAX_BATCH_ROWS"`
	Compression           *string                         `ddl:"parameter" sql:"COMPRESSION"`
	RequestTranslator     *SchemaObjectIdentifier         `ddl:"identifier" sql:"REQUEST_TRANSLATOR ="`
	ResponseTranslator    *SchemaObjectIdentifier         `ddl:"identifier" sql:"RESPONSE_TRANSLATOR ="`
	As                    string                          `ddl:"parameter,single_quotes" sql:"AS"`
}

type ExternalFunctionArgument struct {
	ArgName     string   `ddl:"keyword,no_quotes"`
	ArgDataType DataType `ddl:"keyword,no_quotes"`
}

type ExternalFunctionHeader struct {
	Name  string `ddl:"keyword,single_quotes"`
	Value string `ddl:"parameter,single_quotes"`
}

type ExternalFunctionContextHeader struct {
	ContextFunction string `ddl:"keyword,no_quotes"`
}
