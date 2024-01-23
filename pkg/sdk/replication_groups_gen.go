package sdk

import "context"

type ReplicationGroups interface {
	Create(ctx context.Context, request *CreateReplicationGroupRequest) error
}

// CreateReplicationGroupOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-replication-group.
type CreateReplicationGroupOptions struct {
	create                  bool                              `ddl:"static" sql:"CREATE"`
	replicationGroup        bool                              `ddl:"static" sql:"REPLICATION GROUP"`
	IfNotExists             *bool                             `ddl:"keyword" sql:"IF NOT EXISTS"`
	name                    AccountObjectIdentifier           `ddl:"identifier"`
	ObjectTypes             ReplicationGroupObjectTypes       `ddl:"list,no_parentheses" sql:"OBJECT_TYPES ="`
	AllowedDatabases        []ReplicationGroupDatabase        `ddl:"parameter,no_parentheses" sql:"ALLOWED_DATABASES"`
	AllowedShares           []ReplicationGroupShare           `ddl:"parameter,no_parentheses" sql:"ALLOWED_SHARES"`
	AllowedIntegrationTypes []ReplicationGroupIntegrationType `ddl:"parameter,no_parentheses" sql:"ALLOWED_INTEGRATION_TYPES"`
	AllowedAccounts         []ReplicationGroupAccount         `ddl:"parameter,no_parentheses" sql:"ALLOWED_ACCOUNTS"`
	IgnoreEditionCheck      *bool                             `ddl:"keyword" sql:"IGNORE EDITION CHECK"`
	ReplicationSchedule     *ReplicationGroupSchedule         `ddl:"parameter,single_quotes" sql:"REPLICATION_SCHEDULE"`
}

type ReplicationGroupObjectTypes struct {
	AccountParameters *bool `ddl:"keyword" sql:"ACCOUNT PARAMETERS"`
	Databases         *bool `ddl:"keyword" sql:"DATABASES"`
	Integrations      *bool `ddl:"keyword" sql:"INTEGRATIONS"`
	NetworkPolicies   *bool `ddl:"keyword" sql:"NETWORK POLICIES"`
	ResourceMonitors  *bool `ddl:"keyword" sql:"RESOURCE MONITORS"`
	Roles             *bool `ddl:"keyword" sql:"ROLES"`
	Shares            *bool `ddl:"keyword" sql:"SHARES"`
	Users             *bool `ddl:"keyword" sql:"USERS"`
	Warehouses        *bool `ddl:"keyword" sql:"WAREHOUSES"`
}

type ReplicationGroupDatabase struct {
	Database string `ddl:"keyword"`
}

type ReplicationGroupShare struct {
	Share string `ddl:"keyword"`
}

type ReplicationGroupIntegrationType struct {
	IntegrationType string `ddl:"keyword"`
}

type ReplicationGroupAccount struct {
	Account string `ddl:"keyword"`
}

type ReplicationGroupSchedule struct {
	IntervalMinutes *ScheduleIntervalMinutes `ddl:"keyword"`
	CronExpression  *ScheduleCronExpression  `ddl:"keyword"`
}

type ScheduleIntervalMinutes struct {
	Minutes int  `ddl:"keyword"`
	minutes bool `ddl:"static" sql:"MINUTES"`
}

type ScheduleCronExpression struct {
	usingCron  bool   `ddl:"static" sql:"USING CRON"`
	Expression string `ddl:"keyword"`
	TimeZone   string `ddl:"keyword"`
}
