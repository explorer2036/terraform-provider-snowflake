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
	AllowedDatabases        []ReplicationGroupDatabaseRequest
	AllowedShares           []ReplicationGroupShareRequest
	AllowedIntegrationTypes []ReplicationGroupIntegrationTypeRequest
	AllowedAccounts         []ReplicationGroupAccountRequest
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
	AllowedDatabases     []ReplicationGroupDatabaseRequest
	AllowedShares        []ReplicationGroupShareRequest
	ReplicationSchedule  *ReplicationGroupScheduleRequest
	EnableEtlReplication *bool
}

type ReplicationGroupSetIntegrationRequest struct {
	ObjectTypes             *ReplicationGroupObjectTypesRequest
	AllowedIntegrationTypes []ReplicationGroupIntegrationTypeRequest
	ReplicationSchedule     *ReplicationGroupScheduleRequest
}

type ReplicationGroupAddDatabasesRequest struct {
	Databases []ReplicationGroupDatabaseRequest
}

type ReplicationGroupRemoveDatabasesRequest struct {
	Databases []ReplicationGroupDatabaseRequest
}

type ReplicationGroupMoveDatabasesRequest struct {
	Databases []ReplicationGroupDatabaseRequest
	MoveTo    *AccountObjectIdentifier
}

type ReplicationGroupAddSharesRequest struct {
	Shares []ReplicationGroupShareRequest
}

type ReplicationGroupRemoveSharesRequest struct {
	Shares []ReplicationGroupShareRequest
}

type ReplicationGroupMoveSharesRequest struct {
	Shares []ReplicationGroupShareRequest
	MoveTo *AccountObjectIdentifier
}

type ReplicationGroupAddAccountsRequest struct {
	Accounts           []ReplicationGroupAccountRequest
	IgnoreEditionCheck *bool
}

type ReplicationGroupRemoveAccountsRequest struct {
	Accounts []ReplicationGroupAccountRequest
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
