package testint

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/require"
)

func TestInt_ReplicationGroups(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	cleanupReplicationGroupHandle := func(id sdk.AccountObjectIdentifier) func() {
		return func() {
			err := client.ReplicationGroups.Drop(ctx, sdk.NewDropReplicationGroupRequest(id).WithIfExists(sdk.Bool(true)))
			require.NoError(t, err)
		}
	}

	t.Run("create replication group", func(t *testing.T) {
		db1, cleanupDb1 := createDatabase(t, client)
		t.Cleanup(cleanupDb1)
		db2, cleanupDb2 := createDatabase(t, client)
		t.Cleanup(cleanupDb2)
		share1, cleanupShare1 := createShare(t, client)
		t.Cleanup(cleanupShare1)
		share2, cleanupShare2 := createShare(t, client)
		t.Cleanup(cleanupShare2)

		id := sdk.NewAccountObjectIdentifier(random.AlphaN(5))
		ot := sdk.NewReplicationGroupObjectTypesRequest().
			WithDatabases(sdk.Bool(true)).
			WithShares(sdk.Bool(true))
		drs := []sdk.AccountObjectIdentifier{
			sdk.NewAccountObjectIdentifier(db1.Name),
			sdk.NewAccountObjectIdentifier(db2.Name),
		}
		srs := []sdk.AccountObjectIdentifier{
			sdk.NewAccountObjectIdentifier(share1.Name.Name()),
			sdk.NewAccountObjectIdentifier(share2.Name.Name()),
		}
		as := []sdk.AccountObjectIdentifier{
			sdk.NewAccountObjectIdentifier("SFDEVREL.CLOUD_ENGINEERING2"),
		}
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
	})
}
