package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateReplicationGroupOptions] = new(CreateReplicationGroupRequest)
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
