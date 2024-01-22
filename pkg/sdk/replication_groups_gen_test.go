package sdk

import "testing"

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

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.ObjectTypes = ReplicationGroupObjectTypes{
			Databases: Bool(true),
			Shares:    Bool(true),
		}
		opts.Databases = []ReplicationGroupDatabase{
			{
				Database: "db1",
			},
			{
				Database: "db2",
			},
		}
		opts.Shares = []ReplicationGroupShare{
			{
				Share: "share1",
			},
			{
				Share: "share2",
			},
		}
		opts.Accounts = []ReplicationGroupAccount{
			{
				Account: "org.acct1",
			},
			{
				Account: "org.acct2",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE REPLICATION GROUP IF NOT EXISTS %s OBJECT_TYPES = DATABASES, SHARES ALLOWED_DATABASES = db1, db2 ALLOWED_SHARES = share1, share2 ALLOWED_ACCOUNTS = org.acct1, org.acct2`, id.FullyQualifiedName())
	})
}
