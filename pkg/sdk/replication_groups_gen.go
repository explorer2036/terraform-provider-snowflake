package sdk

import (
	"context"
	"database/sql"
)

type ReplicationGroups interface {
	Create(ctx context.Context, request *CreateReplicationGroupRequest) error
	CreateSecondary(ctx context.Context, request *CreateSecondaryReplicationGroupRequest) error
	Alter(ctx context.Context, request *AlterReplicationGroupRequest) error
	Show(ctx context.Context, request *ShowReplicationGroupRequest) ([]ReplicationGroup, error)
	ShowByID(ctx context.Context, id AccountObjectIdentifier) (*ReplicationGroup, error)
	ShowDatabases(ctx context.Context, request *ShowDatabasesInReplicationGroupRequest) ([]DatabaseInReplicationGroup, error)
	ShowShares(ctx context.Context, request *ShowSharesInReplicationGroupRequest) ([]ShareInReplicationGroup, error)
	Drop(ctx context.Context, request *DropReplicationGroupRequest) error
}

// CreateReplicationGroupOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-replication-group.
type CreateReplicationGroupOptions struct {
	create                  bool                              `ddl:"static" sql:"CREATE"`
	replicationGroup        bool                              `ddl:"static" sql:"REPLICATION GROUP"`
	IfNotExists             *bool                             `ddl:"keyword" sql:"IF NOT EXISTS"`
	name                    AccountObjectIdentifier           `ddl:"identifier"`
	ObjectTypes             ReplicationGroupObjectTypes       `ddl:"list,no_parentheses" sql:"OBJECT_TYPES ="`
	AllowedDatabases        []AccountObjectIdentifier         `ddl:"parameter,no_parentheses" sql:"ALLOWED_DATABASES"`
	AllowedShares           []AccountObjectIdentifier         `ddl:"parameter,no_parentheses" sql:"ALLOWED_SHARES"`
	AllowedIntegrationTypes []ReplicationGroupIntegrationType `ddl:"parameter,no_parentheses" sql:"ALLOWED_INTEGRATION_TYPES"`
	AllowedAccounts         []AccountObjectIdentifier         `ddl:"parameter,no_parentheses" sql:"ALLOWED_ACCOUNTS"`
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

type ReplicationGroupIntegrationType struct {
	IntegrationType string `ddl:"keyword"`
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

// CreateSecondaryReplicationGroupOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-replication-group.
type CreateSecondaryReplicationGroupOptions struct {
	create           bool                      `ddl:"static" sql:"CREATE"`
	replicationGroup bool                      `ddl:"static" sql:"REPLICATION GROUP"`
	IfNotExists      *bool                     `ddl:"keyword" sql:"IF NOT EXISTS"`
	name             AccountObjectIdentifier   `ddl:"identifier"`
	asReplicaOf      bool                      `ddl:"static" sql:"AS REPLICA OF"`
	Primary          *ExternalObjectIdentifier `ddl:"identifier"`
}

// AlterReplicationGroupOptions is based on https://docs.snowflake.com/en/sql-reference/sql/alter-replication-group.
type AlterReplicationGroupOptions struct {
	alter            bool                             `ddl:"static" sql:"ALTER"`
	replicationGroup bool                             `ddl:"static" sql:"REPLICATION GROUP"`
	IfExists         *bool                            `ddl:"keyword" sql:"IF EXISTS"`
	name             AccountObjectIdentifier          `ddl:"identifier"`
	RenameTo         *AccountObjectIdentifier         `ddl:"identifier" sql:"RENAME TO"`
	Set              *ReplicationGroupSet             `ddl:"keyword" sql:"SET"`
	SetIntegration   *ReplicationGroupSetIntegration  `ddl:"keyword" sql:"SET"`
	AddDatabases     *ReplicationGroupAddDatabases    `ddl:"keyword"`
	RemoveDatabases  *ReplicationGroupRemoveDatabases `ddl:"keyword"`
	MoveDatabases    *ReplicationGroupMoveDatabases   `ddl:"keyword"`
	AddShares        *ReplicationGroupAddShares       `ddl:"keyword"`
	RemoveShares     *ReplicationGroupRemoveShares    `ddl:"keyword"`
	MoveShares       *ReplicationGroupMoveShares      `ddl:"keyword"`
	AddAccounts      *ReplicationGroupAddAccounts     `ddl:"keyword"`
	RemoveAccounts   *ReplicationGroupRemoveAccounts  `ddl:"keyword"`
	Refresh          *bool                            `ddl:"keyword" sql:"REFRESH"`
	Suspend          *bool                            `ddl:"keyword" sql:"SUSPEND"`
	Resume           *bool                            `ddl:"keyword" sql:"RESUME"`
}

type ReplicationGroupSet struct {
	ObjectTypes          *ReplicationGroupObjectTypes `ddl:"list,no_parentheses" sql:"OBJECT_TYPES ="`
	AllowedDatabases     []AccountObjectIdentifier    `ddl:"parameter,no_parentheses" sql:"ALLOWED_DATABASES"`
	AllowedShares        []AccountObjectIdentifier    `ddl:"parameter,no_parentheses" sql:"ALLOWED_SHARES"`
	ReplicationSchedule  *ReplicationGroupSchedule    `ddl:"parameter,single_quotes" sql:"REPLICATION_SCHEDULE"`
	EnableEtlReplication *bool                        `ddl:"parameter" sql:"ENABLE_ETL_REPLICATION"`
}

type ReplicationGroupSetIntegration struct {
	ObjectTypes             *ReplicationGroupObjectTypes      `ddl:"list,no_parentheses" sql:"OBJECT_TYPES ="`
	AllowedIntegrationTypes []ReplicationGroupIntegrationType `ddl:"parameter,no_parentheses" sql:"ALLOWED_INTEGRATION_TYPES"`
	ReplicationSchedule     *ReplicationGroupSchedule         `ddl:"parameter,single_quotes" sql:"REPLICATION_SCHEDULE"`
}

type ReplicationGroupAddDatabases struct {
	Add                []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"ADD"`
	toAllowedDatabases bool                      `ddl:"static" sql:"TO ALLOWED_DATABASES"`
}

type ReplicationGroupRemoveDatabases struct {
	Remove               []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"REMOVE"`
	fromAllowedDatabases bool                      `ddl:"static" sql:"FROM ALLOWED_DATABASES"`
}

type ReplicationGroupMoveDatabases struct {
	MoveDatabases []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"MOVE DATABASES"`
	MoveTo        *AccountObjectIdentifier  `ddl:"identifier" sql:"TO REPLICATION GROUP"`
}

type ReplicationGroupAddShares struct {
	Add             []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"ADD"`
	toAllowedShares bool                      `ddl:"static" sql:"TO ALLOWED_SHARES"`
}

type ReplicationGroupRemoveShares struct {
	Remove            []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"REMOVE"`
	fromAllowedShares bool                      `ddl:"static" sql:"FROM ALLOWED_SHARES"`
}

type ReplicationGroupMoveShares struct {
	MoveShares []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"MOVE SHARES"`
	MoveTo     *AccountObjectIdentifier  `ddl:"identifier" sql:"TO REPLICATION GROUP"`
}

type ReplicationGroupAddAccounts struct {
	Add                []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"ADD"`
	toAllowedAccounts  bool                      `ddl:"static" sql:"TO ALLOWED_ACCOUNTS"`
	IgnoreEditionCheck *bool                     `ddl:"keyword" sql:"IGNORE EDITION CHECK"`
}

type ReplicationGroupRemoveAccounts struct {
	Remove              []AccountObjectIdentifier `ddl:"parameter,no_parentheses,no_equals" sql:"REMOVE"`
	fromAllowedAccounts bool                      `ddl:"static" sql:"FROM ALLOWED_ACCOUNTS"`
}

// ShowReplicationGroupOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-replication-groups.
type ShowReplicationGroupOptions struct {
	show              bool                     `ddl:"static" sql:"SHOW"`
	replicationGroups bool                     `ddl:"static" sql:"REPLICATION GROUPS"`
	InAccount         *AccountObjectIdentifier `ddl:"identifier" sql:"IN ACCOUNT"`
}

type ShowDatabasesInReplicationGroupOptions struct {
	show              bool                    `ddl:"static" sql:"SHOW"`
	databases         bool                    `ddl:"static" sql:"DATABASES"`
	in                bool                    `ddl:"static" sql:"IN"`
	replicationGroups bool                    `ddl:"static" sql:"REPLICATION GROUP"`
	name              AccountObjectIdentifier `ddl:"identifier"`
}

type ShowSharesInReplicationGroupOptions struct {
	show              bool                    `ddl:"static" sql:"SHOW"`
	shares            bool                    `ddl:"static" sql:"SHARES"`
	in                bool                    `ddl:"static" sql:"IN"`
	replicationGroups bool                    `ddl:"static" sql:"REPLICATION GROUP"`
	name              AccountObjectIdentifier `ddl:"identifier"`
}

type replicationGroupRow struct {
	SnowflakeRegion         string         `db:"snowflake_region"`
	CreatedOn               string         `db:"created_on"`
	AccountName             string         `db:"account_name"`
	Name                    string         `db:"name"`
	Type                    string         `db:"type"`
	Comment                 sql.NullString `db:"comment"`
	IsPrimary               string         `db:"is_primary"`
	Primary                 string         `db:"primary"`
	ObjectTypes             string         `db:"object_types"`
	AllowedIntegrationTypes string         `db:"allowed_integration_types"`
	AllowedAccounts         string         `db:"allowed_accounts"`
	OrganizationName        string         `db:"organization_name"`
	AccountLocator          string         `db:"account_locator"`
	ReplicationSchedule     sql.NullString `db:"replication_schedule"`
	SecondaryState          sql.NullString `db:"secondary_state"`
	NextScheduledRefresh    sql.NullString `db:"next_scheduled_refresh"`
	Owner                   string         `db:"owner"`
}

type ReplicationGroup struct {
	SnowflakeRegion         string
	CreatedOn               string
	AccountName             string
	Name                    string
	Type                    string
	Comment                 string
	IsPrimary               bool
	Primary                 string
	ObjectTypes             string
	AllowedIntegrationTypes string
	AllowedAccounts         string
	OrganizationName        string
	AccountLocator          string
	ReplicationSchedule     string
	SecondaryState          string
	NextScheduledRefresh    string
	Owner                   string
}

type databaseInReplicationGroupRow struct {
	CreatedOn     string         `db:"created_on"`
	Name          string         `db:"name"`
	IsDefault     string         `db:"is_default"`
	IsCurrent     string         `db:"is_current"`
	Origin        string         `db:"origin"`
	Owner         string         `db:"owner"`
	Comment       string         `db:"comment"`
	Options       string         `db:"options"`
	RetentionTime int            `db:"retention_time"`
	Kind          string         `db:"kind"`
	Budget        sql.NullString `db:"budget"`
	OwnerRoleType string         `db:"owner_role_type"`
}

type DatabaseInReplicationGroup struct {
	CreatedOn     string
	Name          string
	IsDefault     bool
	IsCurrent     bool
	Origin        string
	Owner         string
	Comment       string
	Options       string
	RetentionTime int
	Kind          string
	Budget        string
	OwnerRoleType string
}

type shareInReplicationGroupRow struct {
	CreatedOn         string `db:"created_on"`
	Kind              string `db:"kind"`
	OwnerAccount      string `db:"owner_account"`
	Name              string `db:"name"`
	DatabaseName      string `db:"database_name"`
	To                string `db:"to"`
	Owner             string `db:"owner"`
	Comment           string `db:"comment"`
	ListingGlobalName string `db:"listing_global_name"`
}

type ShareInReplicationGroup struct {
	CreatedOn         string
	Kind              string
	OwnerAccount      string
	Name              string
	DatabaseName      string
	To                string
	Owner             string
	Comment           string
	ListingGlobalName string
}

// DropReplicationGroupOptions is based on https://docs.snowflake.com/en/sql-reference/sql/drop-replication-group.
type DropReplicationGroupOptions struct {
	drop             bool                    `ddl:"static" sql:"DROP"`
	replicationGroup bool                    `ddl:"static" sql:"REPLICATION GROUP"`
	IfExists         *bool                   `ddl:"keyword" sql:"IF EXISTS"`
	name             AccountObjectIdentifier `ddl:"identifier"`
}
