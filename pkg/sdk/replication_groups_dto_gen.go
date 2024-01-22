package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateReplicationGroupOptions] = new(CreateReplicationGroupRequest)
)

type CreateReplicationGroupRequest struct {
	IfNotExists        *bool
	name               AccountObjectIdentifier // required
	ObjectTypes        ReplicationGroupObjectTypesRequest
	Databases          []ReplicationGroupDatabaseRequest
	Shares             []ReplicationGroupShareRequest
	IntegrationTypes   []ReplicationGroupIntegrationTypeRequest
	Accounts           []ReplicationGroupAccountRequest
	IgnoreEditionCheck *bool
}

type ReplicationGroupObjectTypesRequest struct {
	AccountParameters *bool
	Databases         *bool
	Integrations      *bool
	NetworkPolicies   *bool
	ResourceMonitors  *bool
	Roles             *bool
	Shares            *bool
	Users             *bool
	Warehouses        *bool
}

type ReplicationGroupDatabaseRequest struct {
	Database string
}

type ReplicationGroupShareRequest struct {
	Share string
}

type ReplicationGroupIntegrationTypeRequest struct {
	IntegrationType string
}

type ReplicationGroupAccountRequest struct {
	Account string
}
