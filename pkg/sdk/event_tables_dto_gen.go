package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateEventTableOptions] = new(CreateEventTableRequest)
)

type CreateEventTableRequest struct {
	OrReplace   *bool
	IfNotExists *bool
	name        SchemaObjectIdentifier // required
	ClusterBy   []string
}
