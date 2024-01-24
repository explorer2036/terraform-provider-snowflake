package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateExternalFunctionOptions] = new(CreateExternalFunctionRequest)
)

type CreateExternalFunctionRequest struct {
	OrReplace             *bool
	Secure                *bool
	name                  SchemaObjectIdentifier // required
	Arguments             []ExternalFunctionArgumentRequest
	ResultDataType        DataType // required
	ReturnNullValues      *ReturnNullValues
	NullInputBehavior     *NullInputBehavior
	ReturnResultsBehavior *ReturnResultsBehavior
	Comment               *string
	ApiIntegration        *AccountObjectIdentifier
	Headers               []ExternalFunctionHeaderRequest
	ContextHeaders        []ExternalFunctionContextHeaderRequest
	MaxBatchRows          *int
	Compression           *string
	RequestTranslator     *SchemaObjectIdentifier
	ResponseTranslator    *SchemaObjectIdentifier
	As                    string // required
}

type ExternalFunctionArgumentRequest struct {
	ArgName     string   // required
	ArgDataType DataType // required
}

type ExternalFunctionHeaderRequest struct {
	Name  string // required
	Value string // required
}

type ExternalFunctionContextHeaderRequest struct {
	ContextFunction string
}
