package sdk

//go:generate go run ./dto-builder-generator/main.go

var _ optionsProvider[createDynamicTableOptions] = new(CreateDynamicTableRequest)

type CreateDynamicTableRequest struct {
	orReplace bool

	name      AccountObjectIdentifier // required
	warehouse AccountObjectIdentifier // required
	targetLag string                  // required
	query     string                  // required

	comment *string
}
