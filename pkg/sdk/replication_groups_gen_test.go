package sdk

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
)

func TestReplicationGroups_Create(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	defaultOpts := func() *CreateReplicationGroupOptions {
		return &CreateReplicationGroupOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateReplicationGroupOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: at least one field required", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsInvalidJoinedErrors(t, opts, errAtLeastOneOf("CreateReplicationGroupOptions.ObjectTypes", "AccountParameters", "Databases", "Integrations", "NetworkPolicies", "ResourceMonitors", "Roles", "Shares", "Users", "Warehouses"))
	})

	t.Run("validation: exactly one field should be present", func(t *testing.T) {
		opts := defaultOpts()
		opts.ReplicationSchedule = &ReplicationGroupSchedule{
			IntervalMinutes: &ScheduleIntervalMinutes{
				Minutes: 10,
			},
			CronExpression: &ScheduleCronExpression{
				Expression: "10 * * * *",
				TimeZone:   "America/New_York",
			},
		}
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("CreateReplicationGroupOptions.ReplicationSchedule", "IntervalMinutes", "CronExpression"))
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.ObjectTypes = ReplicationGroupObjectTypes{
			Databases: Bool(true),
			Shares:    Bool(true),
		}
		opts.AllowedDatabases = []ReplicationGroupDatabase{
			{
				Database: "db1",
			},
			{
				Database: "db2",
			},
		}
		opts.AllowedShares = []ReplicationGroupShare{
			{
				Share: "share1",
			},
			{
				Share: "share2",
			},
		}
		opts.AllowedIntegrationTypes = []ReplicationGroupIntegrationType{
			{
				IntegrationType: "SECURITY INTEGRATIONS",
			},
			{
				IntegrationType: "API INTEGRATIONS",
			},
		}
		opts.AllowedAccounts = []ReplicationGroupAccount{
			{
				Account: "org.acct1",
			},
			{
				Account: "org.acct2",
			},
		}
		opts.IgnoreEditionCheck = Bool(true)
		opts.ReplicationSchedule = &ReplicationGroupSchedule{
			CronExpression: &ScheduleCronExpression{
				Expression: "0 0 10-20 * TUE,THU",
				TimeZone:   "UTC",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE REPLICATION GROUP IF NOT EXISTS %s OBJECT_TYPES = DATABASES, SHARES ALLOWED_DATABASES = db1, db2 ALLOWED_SHARES = share1, share2 ALLOWED_INTEGRATION_TYPES = SECURITY INTEGRATIONS, API INTEGRATIONS ALLOWED_ACCOUNTS = org.acct1, org.acct2 IGNORE EDITION CHECK REPLICATION_SCHEDULE = 'USING CRON 0 0 10-20 * TUE,THU UTC'`, id.FullyQualifiedName())
	})

	t.Run("create secondary replication group", func(t *testing.T) {
		id := RandomAccountObjectIdentifier()
		primary := NewExternalObjectIdentifierFromFullyQualifiedName("myorg.myaccount.fg1")
		opts := &CreateSecondaryReplicationGroupOptions{
			name:    id,
			Primary: &primary,
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE REPLICATION GROUP %s AS REPLICA OF %s`, id.FullyQualifiedName(), primary.FullyQualifiedName())
	})
}

func TestReplicationGroups_Alter(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	defaultOpts := func() *AlterReplicationGroupOptions {
		return &AlterReplicationGroupOptions{
			IfExists: Bool(true),
			name:     id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *AlterReplicationGroupOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: exactly one field should be present", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("AlterReplicationGroupOptions", "RenameTo", "Set", "SetIntegration", "AddDatabases", "RemoveDatabases", "MoveDatabases", "AddShares", "RemoveShares", "MoveShares", "AddAccounts", "RemoveAccounts", "Refresh", "Suspend", "Resume"))
	})

	t.Run("validation: exactly one field should be present", func(t *testing.T) {
		opts := defaultOpts()
		opts.Refresh = Bool(true)
		opts.Suspend = Bool(true)
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("AlterReplicationGroupOptions", "RenameTo", "Set", "SetIntegration", "AddDatabases", "RemoveDatabases", "MoveDatabases", "AddShares", "RemoveShares", "MoveShares", "AddAccounts", "RemoveAccounts", "Refresh", "Suspend", "Resume"))
	})

	t.Run("alter: rename to", func(t *testing.T) {
		opts := defaultOpts()
		target := NewAccountObjectIdentifier(random.StringN(4))
		opts.RenameTo = &target
		assertOptsValidAndSQLEquals(t, opts, `ALTER REPLICATION GROUP IF EXISTS %s RENAME TO %s`, id.FullyQualifiedName(), target.FullyQualifiedName())
	})

	t.Run("alter: set options", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &ReplicationGroupSet{
			ObjectTypes: &ReplicationGroupObjectTypes{
				Databases: Bool(true),
				Shares:    Bool(true),
			},
			AllowedDatabases: []ReplicationGroupDatabase{
				{
					Database: "db1",
				},
				{
					Database: "db2",
				},
			},
			AllowedShares: []ReplicationGroupShare{
				{
					Share: "share1",
				},
				{
					Share: "share2",
				},
			},
			ReplicationSchedule: &ReplicationGroupSchedule{
				CronExpression: &ScheduleCronExpression{
					Expression: "0 0 10-20 * TUE,THU",
					TimeZone:   "UTC",
				},
			},
			EnableEtlReplication: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER REPLICATION GROUP IF EXISTS %s SET OBJECT_TYPES = DATABASES, SHARES ALLOWED_DATABASES = db1, db2 ALLOWED_SHARES = share1, share2 REPLICATION_SCHEDULE = 'USING CRON 0 0 10-20 * TUE,THU UTC' ENABLE_ETL_REPLICATION = true`, id.FullyQualifiedName())
	})

	t.Run("alter: set integration options", func(t *testing.T) {
		opts := defaultOpts()
		opts.SetIntegration = &ReplicationGroupSetIntegration{
			ObjectTypes: &ReplicationGroupObjectTypes{
				Databases: Bool(true),
				Shares:    Bool(true),
			},
			AllowedIntegrationTypes: []ReplicationGroupIntegrationType{
				{
					IntegrationType: "SECURITY INTEGRATIONS",
				},
				{
					IntegrationType: "API INTEGRATIONS",
				},
			},
			ReplicationSchedule: &ReplicationGroupSchedule{
				CronExpression: &ScheduleCronExpression{
					Expression: "0 0 10-20 * TUE,THU",
					TimeZone:   "UTC",
				},
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER REPLICATION GROUP IF EXISTS %s SET OBJECT_TYPES = DATABASES, SHARES ALLOWED_INTEGRATION_TYPES = SECURITY INTEGRATIONS, API INTEGRATIONS REPLICATION_SCHEDULE = 'USING CRON 0 0 10-20 * TUE,THU UTC'`, id.FullyQualifiedName())
	})
}

func TestReplicationGroups_Drop(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	defaultOpts := func() *DropReplicationGroupOptions {
		return &DropReplicationGroupOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *DropReplicationGroupOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfExists = Bool(true)
		assertOptsValidAndSQLEquals(t, opts, `DROP REPLICATION GROUP IF EXISTS %s`, id.FullyQualifiedName())
	})
}

func TestReplicationGroups_Show(t *testing.T) {
	defaultOpts := func() *ShowReplicationGroupOptions {
		return &ShowReplicationGroupOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *ShowReplicationGroupOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `SHOW REPLICATION GROUPS`)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		account := RandomAccountObjectIdentifier()
		opts.InAccount = &account
		assertOptsValidAndSQLEquals(t, opts, `SHOW REPLICATION GROUPS IN ACCOUNT %s`, account.FullyQualifiedName())
	})

	t.Run("show databases in replication group", func(t *testing.T) {
		id := RandomAccountObjectIdentifier()
		opts := &ShowDatabasesInReplicationGroupOptions{
			name: id,
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW DATABASES IN REPLICATION GROUP %s`, id.FullyQualifiedName())
	})

	t.Run("show shares in replication group", func(t *testing.T) {
		id := RandomAccountObjectIdentifier()
		opts := &ShowSharesInReplicationGroupOptions{
			name: id,
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW SHARES IN REPLICATION GROUP %s`, id.FullyQualifiedName())
	})
}
