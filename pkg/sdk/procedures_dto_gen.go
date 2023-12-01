package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateProcedureForJavaProcedureOptions]       = new(CreateProcedureForJavaProcedureRequest)
	_ optionsProvider[CreateProcedureForJavaScriptProcedureOptions] = new(CreateProcedureForJavaScriptProcedureRequest)
	_ optionsProvider[CreateProcedureForPythonProcedureOptions]     = new(CreateProcedureForPythonProcedureRequest)
	_ optionsProvider[CreateProcedureForScalaProcedureOptions]      = new(CreateProcedureForScalaProcedureRequest)
	_ optionsProvider[CreateProcedureForSQLProcedureOptions]        = new(CreateProcedureForSQLProcedureRequest)
	_ optionsProvider[AlterProcedureOptions]                        = new(AlterProcedureRequest)
	_ optionsProvider[DropProcedureOptions]                         = new(DropProcedureRequest)
	_ optionsProvider[ShowProcedureOptions]                         = new(ShowProcedureRequest)
	_ optionsProvider[DescribeProcedureOptions]                     = new(DescribeProcedureRequest)
)

type CreateProcedureForJavaProcedureRequest struct {
	OrReplace                  *bool
	Secure                     *bool
	name                       SchemaObjectIdentifier // required
	Arguments                  []ProcedureArgumentRequest
	CopyGrants                 *bool
	Returns                    ProcedureReturnsRequest
	RuntimeVersion             string                    // required
	Packages                   []ProcedurePackageRequest // required
	Imports                    []ProcedureImportRequest
	Handler                    string // required
	ExternalAccessIntegrations []AccountObjectIdentifier
	Secrets                    []Secret
	TargetPath                 *string
	NullInputBehavior          *NullInputBehavior
	Comment                    *string
	ExecuteAs                  *ExecuteAs
	ProcedureDefinition        *string
}

type ProcedureArgumentRequest struct {
	ArgName      string   // required
	ArgDataType  DataType // required
	DefaultValue *string
}

type ProcedureReturnsRequest struct {
	ResultDataType *ProcedureReturnsResultDataTypeRequest
	Table          *ProcedureReturnsTableRequest
}

type ProcedureReturnsResultDataTypeRequest struct {
	ResultDataType DataType // required
	Null           *bool
	NotNull        *bool
}

type ProcedureReturnsTableRequest struct {
	Columns []ProcedureColumnRequest
}

type ProcedureColumnRequest struct {
	ColumnName     string
	ColumnDataType DataType // required
}

type ProcedurePackageRequest struct {
	Package string
}

type ProcedureImportRequest struct {
	Import string
}

type CreateProcedureForJavaScriptProcedureRequest struct {
	OrReplace           *bool
	Secure              *bool
	name                SchemaObjectIdentifier // required
	Arguments           []ProcedureArgumentRequest
	CopyGrants          *bool
	Returns             *ProcedureJavascriptReturnsRequest
	NullInputBehavior   *NullInputBehavior
	Comment             *string
	ExecuteAs           *ExecuteAs
	ProcedureDefinition string
}

type ProcedureJavascriptReturnsRequest struct {
	ResultDataType DataType // required
	NotNull        *bool
}

type CreateProcedureForPythonProcedureRequest struct {
	OrReplace                  *bool
	Secure                     *bool
	name                       SchemaObjectIdentifier // required
	Arguments                  []ProcedureArgumentRequest
	CopyGrants                 *bool
	Returns                    *ProcedureReturnsRequest
	RuntimeVersion             *string
	Packages                   []ProcedurePackageRequest
	Imports                    []ProcedureImportRequest
	Handler                    string // required
	ExternalAccessIntegrations []AccountObjectIdentifier
	Secrets                    []Secret
	NullInputBehavior          *NullInputBehavior
	Comment                    *string
	ExecuteAs                  *ExecuteAs
	ProcedureDefinition        *string
}

type CreateProcedureForScalaProcedureRequest struct {
	OrReplace           *bool
	Secure              *bool
	name                SchemaObjectIdentifier // required
	Arguments           []ProcedureArgumentRequest
	CopyGrants          *bool
	Returns             *ProcedureReturnsRequest
	RuntimeVersion      *string
	Packages            []ProcedurePackageRequest
	Imports             []ProcedureImportRequest
	Handler             string // required
	TargetPath          *string
	NullInputBehavior   *NullInputBehavior
	Comment             *string
	ExecuteAs           *ExecuteAs
	ProcedureDefinition *string
}

type CreateProcedureForSQLProcedureRequest struct {
	OrReplace           *bool
	Secure              *bool
	name                SchemaObjectIdentifier // required
	Arguments           []ProcedureArgumentRequest
	CopyGrants          *bool
	Returns             *ProcedureSQLReturnsRequest
	NullInputBehavior   *NullInputBehavior
	Comment             *string
	ExecuteAs           *ExecuteAs
	ProcedureDefinition string
}

type ProcedureSQLReturnsRequest struct {
	ResultDataType *ProcedureReturnsResultDataTypeRequest
	Table          *ProcedureReturnsTableRequest
	NotNull        *bool
}

type AlterProcedureRequest struct {
	IfExists          *bool
	name              SchemaObjectIdentifier // required
	ArgumentDataTypes []DataType             // required
	RenameTo          *SchemaObjectIdentifier
	SetComment        *string
	SetLogLevel       *string
	SetTraceLevel     *string
	UnsetComment      *bool
	SetTags           []TagAssociation
	UnsetTags         []ObjectIdentifier
	ExecuteAs         *ExecuteAs
}

type DropProcedureRequest struct {
	IfExists          *bool
	name              SchemaObjectIdentifier // required
	ArgumentDataTypes []DataType             // required
}

type ShowProcedureRequest struct {
	Like *Like
	In   *In
}

type DescribeProcedureRequest struct {
	name              SchemaObjectIdentifier // required
	ArgumentDataTypes []DataType             // required
}
