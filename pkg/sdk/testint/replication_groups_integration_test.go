package testint

import (
	"fmt"
	"strings"
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

	assertReplicationGroup := func(t *testing.T, id sdk.AccountObjectIdentifier, objectTypes []string, org string, account string, schedule string) {
		t.Helper()

		e, err := client.ReplicationGroups.ShowByID(ctx, id)
		require.NoError(t, err)
		primary := fmt.Sprintf("%s.%s.%s", org, account, id.Name())
		require.NotEmpty(t, e.CreatedOn)
		require.Equal(t, "AWS_US_WEST_2", e.SnowflakeRegion)
		require.Equal(t, account, e.AccountName)
		require.Equal(t, id.Name(), e.Name)
		require.Equal(t, "REPLICATION", e.Type)
		require.Empty(t, e.Comment)
		require.Equal(t, false, e.IsPrimary)
		require.Equal(t, primary, e.Primary)
		require.Equal(t, strings.Join(objectTypes, ", "), e.ObjectTypes)
		require.Empty(t, e.AllowedIntegrationTypes)
		require.NotEmpty(t, e.AllowedAccounts)
		require.Equal(t, org, e.OrganizationName)
		require.NotEmpty(t, e.AccountLocator)
		require.Equal(t, schedule, e.ReplicationSchedule)
		require.Empty(t, e.SecondaryState)
		require.Empty(t, e.NextScheduledRefresh)
		require.Equal(t, "ACCOUNTADMIN", e.Owner)
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
		objectTypes := []string{"DATABASES", "SHARES"}
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
		org, account := "SFDEVREL", "CLOUD_ENGINEERING"
		as := []sdk.AccountObjectIdentifier{
			sdk.NewAccountObjectIdentifier(org + "." + account),
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

		assertReplicationGroup(t, id, objectTypes, org, account, "10 MINUTES")
	})
}
