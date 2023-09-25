package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[createDynamicTableOptions] = new(CreateDynamicTableRequest)
	_ optionsProvider[alterDynamicTableOptions]  = new(AlterDynamicTableRequest)
)

type CreateDynamicTableRequest struct {
	orReplace bool

	name      AccountObjectIdentifier // required
	warehouse AccountObjectIdentifier // required
	targetLag string                  // required
	query     string                  // required

	comment *string
}

type AlterDynamicTableRequest struct {
	name AccountObjectIdentifier // required

	// One of
	suspend *bool
	resume  *bool
	refresh *bool
	set     *DynamicTableSetRequest
}

type DynamicTableSetRequest struct {
	targetLag  *string
	warehourse *AccountObjectIdentifier
}
