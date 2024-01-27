package testint

import (
	"strings"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/require"
)

/*
 * todo
 * 1. add tests for ALLOWED_INTEGRATION_TYPES, it Requires Business Critical Edition (or higher)
 * 2. CREATE REPLICATION GROUP [ IF NOT EXISTS ] <secondary_name> AS REPLICA OF <org_name>.<source_account_name>.<name>
 */

func TestInt_ReplicationGroups(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	defaultAllowedAccount := "SFDEVREL.CLOUD_ENGINEERING2"

	cleanupReplicationGroupHandle := func(id sdk.AccountObjectIdentifier) func() {
		return func() {
			err := client.ReplicationGroups.Drop(ctx, sdk.NewDropReplicationGroupRequest(id).WithIfExists(sdk.Bool(true)))
			require.NoError(t, err)
		}
	}

	createReplicationGroupHandle := func(t *testing.T, db *sdk.Database, share *sdk.Share, account string, minute int) *sdk.ReplicationGroup {
		t.Helper()

		if db == nil {
			d, cleanupDatabase := createDatabase(t, client)
			t.Cleanup(cleanupDatabase)
			db = d
		}
		if share == nil {
			s, cleanupShare := createShare(t, client)
			t.Cleanup(cleanupShare)
			share = s
		}

		id := sdk.NewAccountObjectIdentifier(random.AlphaN(4))
		ot := sdk.NewReplicationGroupObjectTypesRequest().WithDatabases(sdk.Bool(true)).WithShares(sdk.Bool(true))
		drs := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(db.Name)}
		srs := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(share.Name.Name())}
		as := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(account)}
		schedule := sdk.NewReplicationGroupScheduleRequest().WithIntervalMinutes(sdk.NewScheduleIntervalMinutesRequest(minute))
		request := sdk.NewCreateReplicationGroupRequest(id).
			WithObjectTypes(*ot).
			WithAllowedDatabases(drs).
			WithAllowedShares(srs).
			WithAllowedAccounts(as).
			WithReplicationSchedule(schedule)
		err := client.ReplicationGroups.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupReplicationGroupHandle(id))

		e, err := client.ReplicationGroups.ShowByID(ctx, id)
		require.NoError(t, err)
		return e
	}

	t.Run("create replication group", func(t *testing.T) {
		db, cleanupDb := createDatabase(t, client)
		t.Cleanup(cleanupDb)
		share, cleanupShare := createShare(t, client)
		t.Cleanup(cleanupShare)

		id := sdk.NewAccountObjectIdentifier(random.AlphaN(5))
		ot := sdk.NewReplicationGroupObjectTypesRequest().WithDatabases(sdk.Bool(true)).WithShares(sdk.Bool(true))
		drs := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(db.Name)}
		srs := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(share.Name.Name())}
		as := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(defaultAllowedAccount)}
		schedule := sdk.NewReplicationGroupScheduleRequest().WithIntervalMinutes(sdk.NewScheduleIntervalMinutesRequest(10))
		request := sdk.NewCreateReplicationGroupRequest(id).
			WithObjectTypes(*ot).
			WithAllowedDatabases(drs).
			WithAllowedShares(srs).
			WithAllowedAccounts(as).
			WithReplicationSchedule(schedule)
		err := client.ReplicationGroups.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupReplicationGroupHandle(id))

		e, err := client.ReplicationGroups.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, id.Name(), e.Name)
		require.Equal(t, "REPLICATION", e.Type)
		require.Equal(t, false, e.IsPrimary)
		require.Equal(t, "DATABASES, SHARES", e.ObjectTypes)
		require.Equal(t, "10 MINUTES", e.ReplicationSchedule)
		require.Equal(t, "ACCOUNTADMIN", e.Owner)
	})

	t.Run("alter replication group: rename", func(t *testing.T) {
		e := createReplicationGroupHandle(t, nil, nil, defaultAllowedAccount, 10)

		id := sdk.NewAccountObjectIdentifier(e.Name)
		nid := sdk.NewAccountObjectIdentifier(random.StringN(3))
		request := sdk.NewAlterReplicationGroupRequest(id).WithRenameTo(&nid)
		err := client.ReplicationGroups.Alter(ctx, request)
		if err != nil {
			t.Cleanup(cleanupReplicationGroupHandle(id))
		} else {
			t.Cleanup(cleanupReplicationGroupHandle(nid))
		}
		require.NoError(t, err)

		_, err = client.ReplicationGroups.ShowByID(ctx, id)
		require.ErrorIs(t, err, collections.ErrObjectNotFound)

		e, err = client.ReplicationGroups.ShowByID(ctx, nid)
		require.NoError(t, err)
		require.Equal(t, nid.Name(), e.Name)
	})

	t.Run("alter replication group: add databases and add shares", func(t *testing.T) {
		e := createReplicationGroupHandle(t, nil, nil, defaultAllowedAccount, 10)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		db, cleanupDb := createDatabase(t, client)
		t.Cleanup(cleanupDb)
		addDatabase := sdk.NewReplicationGroupAddDatabasesRequest().WithAdd([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(db.Name)})
		err := client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithAddDatabases(addDatabase))
		require.NoError(t, err)
		databases, err := client.ReplicationGroups.ShowDatabases(ctx, sdk.NewShowDatabasesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 2, len(databases))

		share, cleanupShare := createShare(t, client)
		t.Cleanup(cleanupShare)
		addShare := sdk.NewReplicationGroupAddSharesRequest().WithAdd([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(share.Name.Name())})
		err = client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithAddShares(addShare))
		require.NoError(t, err)
		shares, err := client.ReplicationGroups.ShowShares(ctx, sdk.NewShowSharesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 2, len(shares))
	})

	t.Run("alter replication group: remove databases and add shares", func(t *testing.T) {
		db, cleanupDatabase := createDatabase(t, client)
		t.Cleanup(cleanupDatabase)
		share, cleanupShare := createShare(t, client)
		t.Cleanup(cleanupShare)

		e := createReplicationGroupHandle(t, db, share, defaultAllowedAccount, 10)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		removeDatabase := sdk.NewReplicationGroupRemoveDatabasesRequest().WithRemove([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(db.Name)})
		err := client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithRemoveDatabases(removeDatabase))
		require.NoError(t, err)
		databases, err := client.ReplicationGroups.ShowDatabases(ctx, sdk.NewShowDatabasesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 0, len(databases))

		removeShare := sdk.NewReplicationGroupRemoveSharesRequest().WithRemove([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(share.Name.Name())})
		err = client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithRemoveShares(removeShare))
		require.NoError(t, err)
		shares, err := client.ReplicationGroups.ShowShares(ctx, sdk.NewShowSharesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 0, len(shares))
	})

	t.Run("alter replication group: add and remove account", func(t *testing.T) {
		e := createReplicationGroupHandle(t, nil, nil, defaultAllowedAccount, 10)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		added := "SFDEVREL.CLOUD_ENGINEERING3"
		ar := sdk.NewReplicationGroupAddAccountsRequest().WithAdd([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifierFromFullyQualifiedName(added)})
		err := client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithAddAccounts(ar))
		require.NoError(t, err)
		e, err = client.ReplicationGroups.ShowByID(ctx, id)
		require.NoError(t, err)

		pairs := make(map[string]bool)
		for _, item := range strings.Split(e.AllowedAccounts, ",") {
			pairs[strings.TrimSpace(item)] = true
		}
		require.True(t, pairs[added])

		removed := "SFDEVREL.CLOUD_ENGINEERING3"
		rr := sdk.NewReplicationGroupRemoveAccountsRequest().WithRemove([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifierFromFullyQualifiedName(removed)})
		err = client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithRemoveAccounts(rr))
		require.NoError(t, err)
		e, err = client.ReplicationGroups.ShowByID(ctx, id)
		require.NoError(t, err)

		pairs = make(map[string]bool)
		for _, item := range strings.Split(e.AllowedAccounts, ",") {
			pairs[strings.TrimSpace(item)] = true
		}
		require.False(t, pairs[removed])
	})

	t.Run("alter replication group: move database", func(t *testing.T) {
		db1, cleanupDatabase1 := createDatabase(t, client)
		t.Cleanup(cleanupDatabase1)
		share1, cleanupShare1 := createShare(t, client)
		t.Cleanup(cleanupShare1)
		e1 := createReplicationGroupHandle(t, db1, share1, defaultAllowedAccount, 10)

		db2, cleanupDatabase2 := createDatabase(t, client)
		t.Cleanup(cleanupDatabase2)
		share2, cleanupShare2 := createShare(t, client)
		t.Cleanup(cleanupShare2)
		e2 := createReplicationGroupHandle(t, db2, share2, defaultAllowedAccount, 10)

		to := sdk.NewAccountObjectIdentifier(e2.Name)
		mr := sdk.NewReplicationGroupMoveDatabasesRequest().WithMoveDatabases([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(db1.Name)}).WithMoveTo(&to)
		err := client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(sdk.NewAccountObjectIdentifier(e1.Name)).WithMoveDatabases(mr))
		require.NoError(t, err)

		databases, err := client.ReplicationGroups.ShowDatabases(ctx, sdk.NewShowDatabasesInReplicationGroupRequest(sdk.NewAccountObjectIdentifier(e1.Name)))
		require.NoError(t, err)
		require.Equal(t, 0, len(databases))

		databases, err = client.ReplicationGroups.ShowDatabases(ctx, sdk.NewShowDatabasesInReplicationGroupRequest(sdk.NewAccountObjectIdentifier(e2.Name)))
		require.NoError(t, err)
		require.Equal(t, 2, len(databases))
	})

	t.Run("alter replication group: move share", func(t *testing.T) {
		db1, cleanupDatabase1 := createDatabase(t, client)
		t.Cleanup(cleanupDatabase1)
		share1, cleanupShare1 := createShare(t, client)
		t.Cleanup(cleanupShare1)
		e1 := createReplicationGroupHandle(t, db1, share1, defaultAllowedAccount, 10)

		db2, cleanupDatabase2 := createDatabase(t, client)
		t.Cleanup(cleanupDatabase2)
		share2, cleanupShare2 := createShare(t, client)
		t.Cleanup(cleanupShare2)
		e2 := createReplicationGroupHandle(t, db2, share2, defaultAllowedAccount, 10)

		to := sdk.NewAccountObjectIdentifier(e2.Name)
		mr := sdk.NewReplicationGroupMoveSharesRequest().WithMoveShares([]sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(share1.Name.Name())}).WithMoveTo(&to)
		err := client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(sdk.NewAccountObjectIdentifier(e1.Name)).WithMoveShares(mr))
		require.NoError(t, err)

		shares, err := client.ReplicationGroups.ShowShares(ctx, sdk.NewShowSharesInReplicationGroupRequest(sdk.NewAccountObjectIdentifier(e1.Name)))
		require.NoError(t, err)
		require.Equal(t, 0, len(shares))

		shares, err = client.ReplicationGroups.ShowShares(ctx, sdk.NewShowSharesInReplicationGroupRequest(sdk.NewAccountObjectIdentifier(e2.Name)))
		require.NoError(t, err)
		require.Equal(t, 2, len(shares))
	})

	t.Run("alter replication group: set", func(t *testing.T) {
		db1, cleanupDatabase1 := createDatabase(t, client)
		t.Cleanup(cleanupDatabase1)
		share1, cleanupShare1 := createShare(t, client)
		t.Cleanup(cleanupShare1)
		e := createReplicationGroupHandle(t, db1, share1, defaultAllowedAccount, 10)

		id := sdk.NewAccountObjectIdentifier(e.Name)
		databases, err := client.ReplicationGroups.ShowDatabases(ctx, sdk.NewShowDatabasesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 1, len(databases))
		shares, err := client.ReplicationGroups.ShowShares(ctx, sdk.NewShowSharesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 1, len(shares))

		db2, cleanupDatabase2 := createDatabase(t, client)
		t.Cleanup(cleanupDatabase2)
		share2, cleanupShare2 := createShare(t, client)
		t.Cleanup(cleanupShare2)

		ds := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(db1.Name), sdk.NewAccountObjectIdentifier(db2.Name)}
		ss := []sdk.AccountObjectIdentifier{sdk.NewAccountObjectIdentifier(share1.Name.Name()), sdk.NewAccountObjectIdentifier(share2.Name.Name())}
		set := sdk.NewReplicationGroupSetRequest().WithAllowedDatabases(ds).WithAllowedShares(ss)
		err = client.ReplicationGroups.Alter(ctx, sdk.NewAlterReplicationGroupRequest(id).WithSet(set))
		require.NoError(t, err)

		databases, err = client.ReplicationGroups.ShowDatabases(ctx, sdk.NewShowDatabasesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 2, len(databases))
		shares, err = client.ReplicationGroups.ShowShares(ctx, sdk.NewShowSharesInReplicationGroupRequest(id))
		require.NoError(t, err)
		require.Equal(t, 2, len(shares))
	})
}
