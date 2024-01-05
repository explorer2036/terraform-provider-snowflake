package testint

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt_ApplicationPackages(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	// databaseTest, schemaTest := testDb(t), testSchema(t)
	// tagTest, tagCleanup := createTag(t, client, databaseTest, schemaTest)
	// t.Cleanup(tagCleanup)

	cleanupApplicationPackageHandle := func(id sdk.AccountObjectIdentifier) func() {
		return func() {
			err := client.ApplicationPackages.Drop(ctx, sdk.NewDropApplicationPackageRequest(id))
			if errors.Is(err, sdk.ErrObjectNotExistOrAuthorized) {
				return
			}
			require.NoError(t, err)
		}
	}

	createApplicationPackageHandle := func(t *testing.T, client *sdk.Client) *sdk.ApplicationPackage {
		t.Helper()

		id := sdk.NewAccountObjectIdentifier(random.StringN(4))
		request := sdk.NewCreateApplicationPackageRequest(id).WithDistribution(sdk.DistributionPointer(sdk.DistributionInternal))
		err := client.ApplicationPackages.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationPackageHandle(id))

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		return e
	}

	// assertApplicationPackage := func(t *testing.T, id sdk.AccountObjectIdentifier) {
	// 	t.Helper()

	// 	e, err := client.ApplicationPackages.ShowByID(ctx, id)
	// 	require.NoError(t, err)

	// 	assert.NotEmpty(t, e.CreatedOn)
	// 	assert.Equal(t, id.Name(), e.Name)
	// 	assert.Equal(t, false, e.IsDefault)
	// 	assert.Equal(t, true, e.IsCurrent)
	// 	assert.Equal(t, sdk.DistributionInternal, sdk.Distribution(e.Distribution))
	// 	assert.Equal(t, "ACCOUNTADMIN", e.Owner)
	// 	assert.Empty(t, e.Comment)
	// 	assert.Equal(t, 1, e.RetentionTime)
	// 	assert.Empty(t, e.Options)
	// 	assert.Empty(t, e.DroppedOn)
	// 	assert.Empty(t, e.ApplicationClass)
	// }

	t.Run("alter application package: set options", func(t *testing.T) {
		e := createApplicationPackageHandle(t, client)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		distribution := sdk.DistributionPointer(sdk.DistributionExternal)
		set := sdk.NewApplicationPackageSetRequest().
			WithDistribution(distribution).
			WithComment(sdk.String("test")).
			WithDataRetentionTimeInDays(sdk.Int(2)).
			WithMaxDataExtensionTimeInDays(sdk.Int(2)).
			WithDefaultDdlCollation(sdk.String("utf8mb4_0900_ai_ci"))
		err := client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithSet(set))
		require.NoError(t, err)

		o, err := client.ApplicationPackages.ShowByID(ctx, id)
		b, _ := json.Marshal(o)
		t.Log(string(b))
		require.NoError(t, err)
		assert.Equal(t, *distribution, sdk.Distribution(o.Distribution))
		assert.Equal(t, 2, o.RetentionTime)
		assert.Equal(t, "test", o.Comment)
	})
}
