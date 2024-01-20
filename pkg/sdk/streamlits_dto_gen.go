package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateStreamlitOptions] = new(CreateStreamlitRequest)
	_ optionsProvider[AlterStreamlitOptions]  = new(AlterStreamlitRequest)
)

type CreateStreamlitRequest struct {
	OrReplace    *bool
	IfNotExists  *bool
	name         SchemaObjectIdentifier // required
	RootLocation string                 // required
	MainFile     string                 // required
	Warehouse    *AccountObjectIdentifier
	Comment      *string
}

type AlterStreamlitRequest struct {
	IfExists *bool
	name     SchemaObjectIdentifier // required
	Set      *StreamlitsSetRequest
	RenameTo *SchemaObjectIdentifier
}

type StreamlitsSetRequest struct {
	RootLocation *string // required
	Warehouse    *AccountObjectIdentifier
	MainFile     *string // required
	Comment      *string
}
