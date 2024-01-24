package sdk

import "testing"

func TestExternalFunctions_Create(t *testing.T) {
	id := RandomSchemaObjectIdentifier()

	defaultOpts := func() *CreateExternalFunctionOptions {
		return &CreateExternalFunctionOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateExternalFunctionOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.OrReplace = Bool(true)
		opts.Secure = Bool(true)
		opts.Arguments = []ExternalFunctionArgument{
			{
				ArgName:     "id",
				ArgDataType: DataTypeNumber,
			},
			{
				ArgName:     "name",
				ArgDataType: DataTypeVARCHAR,
			},
		}
		opts.ResultDataType = DataTypeVARCHAR
		opts.ReturnNullValues = ReturnNullValuesPointer(ReturnNullValuesNotNull)
		opts.NullInputBehavior = NullInputBehaviorPointer(NullInputBehaviorCalledOnNullInput)
		opts.ReturnResultsBehavior = ReturnResultsBehaviorPointer(ReturnResultsBehaviorImmutable)
		opts.Comment = String("comment")
		integration := NewAccountObjectIdentifier("api_integration")
		opts.ApiIntegration = &integration
		opts.Headers = []ExternalFunctionHeader{
			{
				Name:  "header1",
				Value: "value1",
			},
			{
				Name:  "header2",
				Value: "value2",
			},
		}
		opts.ContextHeaders = []ExternalFunctionContextHeader{
			{
				ContextFunction: "CURRENT_ACCOUNT()",
			},
			{
				ContextFunction: "CURRENT_USER()",
			},
		}
		opts.MaxBatchRows = Int(100)
		opts.Compression = String("GZIP")
		rt := NewSchemaObjectIdentifier("db", "schema", "request_translator")
		opts.RequestTranslator = &rt
		rs := NewSchemaObjectIdentifier("db", "schema", "response_translator")
		opts.ResponseTranslator = &rs
		opts.As = "https://xyz.execute-api.us-west-2.amazonaws.com/prod/remote_echo"
		assertOptsValidAndSQLEquals(t, opts, `CREATE OR REPLACE SECURE EXTERNAL FUNCTION %s (id NUMBER, name VARCHAR) RETURNS VARCHAR NOT NULL CALLED ON NULL INPUT IMMUTABLE COMMENT = 'comment' API_INTEGRATION = "api_integration" HEADERS = ('header1' = 'value1', 'header2' = 'value2') CONTEXT_HEADERS = (CURRENT_ACCOUNT(), CURRENT_USER()) MAX_BATCH_ROWS = 100 COMPRESSION = GZIP REQUEST_TRANSLATOR = %s RESPONSE_TRANSLATOR = %s AS = 'https://xyz.execute-api.us-west-2.amazonaws.com/prod/remote_echo'`, id.FullyQualifiedName(), rt.FullyQualifiedName(), rs.FullyQualifiedName())
	})
}
