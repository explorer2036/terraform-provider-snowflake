package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateReplicationGroupOptions]          = new(CreateReplicationGroupRequest)
	_ optionsProvider[CreateSecondaryReplicationGroupOptions] = new(CreateSecondaryReplicationGroupRequest)
	_ optionsProvider[AlterReplicationGroupOptions]           = new(AlterReplicationGroupRequest)
	_ optionsProvider[ShowReplicationGroupOptions]            = new(ShowReplicationGroupRequest)
	_ optionsProvider[DropReplicationGroupOptions]            = new(DropReplicationGroupRequest)
)

type CreateReplicationGroupRequest struct {
	IfNotExists             *bool
	name                    AccountObjectIdentifier // required
	ObjectTypes             ReplicationGroupObjectTypesRequest
	AllowedDatabases        []AccountObjectIdentifier
	AllowedShares           []AccountObjectIdentifier
	AllowedIntegrationTypes []ReplicationGroupIntegrationTypeRequest
	AllowedAccounts         []AccountObjectIdentifier
	IgnoreEditionCheck      *bool
	ReplicationSchedule     *ReplicationGroupScheduleRequest
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

type ReplicationGroupIntegrationTypeRequest struct {
	IntegrationType string
}

type ReplicationGroupScheduleRequest struct {
	IntervalMinutes *ScheduleIntervalMinutesRequest
	CronExpression  *ScheduleCronExpressionRequest
}

type ScheduleIntervalMinutesRequest struct {
	Minutes int // required
}

type ScheduleCronExpressionRequest struct {
	Expression string // required
	TimeZone   string // required
}

type CreateSecondaryReplicationGroupRequest struct {
	IfNotExists *bool
	name        AccountObjectIdentifier   // required
	Primary     *ExternalObjectIdentifier // required
}

type AlterReplicationGroupRequest struct {
	IfExists        *bool
	name            AccountObjectIdentifier // required
	RenameTo        *AccountObjectIdentifier
	Set             *ReplicationGroupSetRequest
	SetIntegration  *ReplicationGroupSetIntegrationRequest
	AddDatabases    *ReplicationGroupAddDatabasesRequest
	RemoveDatabases *ReplicationGroupRemoveDatabasesRequest
	MoveDatabases   *ReplicationGroupMoveDatabasesRequest
	AddShares       *ReplicationGroupAddSharesRequest
	RemoveShares    *ReplicationGroupRemoveSharesRequest
	MoveShares      *ReplicationGroupMoveSharesRequest
	AddAccounts     *ReplicationGroupAddAccountsRequest
	RemoveAccounts  *ReplicationGroupRemoveAccountsRequest
	Refresh         *bool
	Suspend         *bool
	Resume          *bool
}

type ReplicationGroupSetRequest struct {
	ObjectTypes          *ReplicationGroupObjectTypesRequest
	AllowedDatabases     []AccountObjectIdentifier
	AllowedShares        []AccountObjectIdentifier
	ReplicationSchedule  *ReplicationGroupScheduleRequest
	EnableEtlReplication *bool
}

type ReplicationGroupSetIntegrationRequest struct {
	ObjectTypes             *ReplicationGroupObjectTypesRequest
	AllowedIntegrationTypes []ReplicationGroupIntegrationTypeRequest
	ReplicationSchedule     *ReplicationGroupScheduleRequest
}

type ReplicationGroupAddDatabasesRequest struct {
	Add []AccountObjectIdentifier
}

type ReplicationGroupRemoveDatabasesRequest struct {
	Remove []AccountObjectIdentifier
}

type ReplicationGroupMoveDatabasesRequest struct {
	MoveDatabases []AccountObjectIdentifier
	MoveTo        *AccountObjectIdentifier
}

type ReplicationGroupAddSharesRequest struct {
	Add []AccountObjectIdentifier
}

type ReplicationGroupRemoveSharesRequest struct {
	Remove []AccountObjectIdentifier
}

type ReplicationGroupMoveSharesRequest struct {
	MoveShares []AccountObjectIdentifier
	MoveTo     *AccountObjectIdentifier
}

type ReplicationGroupAddAccountsRequest struct {
	Add                []AccountObjectIdentifier
	IgnoreEditionCheck *bool
}

type ReplicationGroupRemoveAccountsRequest struct {
	Remove []AccountObjectIdentifier
}

type ShowReplicationGroupRequest struct {
	InAccount *AccountObjectIdentifier
}

type ShowDatabasesInReplicationGroupRequest struct {
	name AccountObjectIdentifier // required
}

type ShowSharesInReplicationGroupRequest struct {
	name AccountObjectIdentifier // required
}

type DropReplicationGroupRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
}
