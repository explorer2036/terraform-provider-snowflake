package testint

import (
	"errors"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
todo: add tests for:
  - Modifies the properties of an existing application package: https://docs.snowflake.com/en/sql-reference/sql/alter-application-package-release-directive
  - Modifies the versioning of an existing application package: https://docs.snowflake.com/en/sql-reference/sql/alter-application-package-version
*/

func TestInt_ApplicationPackages(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	databaseTest, schemaTest := testDb(t), testSchema(t)
	tagTest, tagCleanup := createTag(t, client, databaseTest, schemaTest)
	t.Cleanup(tagCleanup)

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

	assertApplicationPackage := func(t *testing.T, id sdk.AccountObjectIdentifier) {
		t.Helper()

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)

		assert.NotEmpty(t, e.CreatedOn)
		assert.Equal(t, id.Name(), e.Name)
		assert.Equal(t, false, e.IsDefault)
		assert.Equal(t, true, e.IsCurrent)
		assert.Equal(t, sdk.DistributionInternal, sdk.Distribution(e.Distribution))
		assert.Equal(t, "ACCOUNTADMIN", e.Owner)
		assert.Empty(t, e.Comment)
		assert.Equal(t, 1, e.RetentionTime)
		assert.Empty(t, e.Options)
		assert.Empty(t, e.DroppedOn)
		assert.Empty(t, e.ApplicationClass)
	}

	t.Run("create application package", func(t *testing.T) {
		id := sdk.NewAccountObjectIdentifier(random.StringN(4))
		comment := random.StringN(4)
		request := sdk.NewCreateApplicationPackageRequest(id).
			WithComment(&comment).
			// todo: insufficient privileges for the following three fields
			// WithDataRetentionTimeInDays(sdk.Int(1)).
			// WithMaxDataExtensionTimeInDays(sdk.Int(1)).
			// WithDefaultDdlCollation(sdk.String("en_US")).
			WithTag([]sdk.TagAssociation{
				{
					Name:  tagTest.ID(),
					Value: "v1",
				},
			}).
			WithDistribution(sdk.DistributionPointer(sdk.DistributionExternal))
		err := client.ApplicationPackages.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationPackageHandle(id))

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, id.Name(), e.Name)
		require.Equal(t, sdk.DistributionExternal, sdk.Distribution(e.Distribution))
		require.Equal(t, "ACCOUNTADMIN", e.Owner)
		require.Equal(t, comment, e.Comment)
		require.Equal(t, 1, e.RetentionTime)
	})

	t.Run("alter application package: set", func(t *testing.T) {
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
		require.NoError(t, err)
		assert.Equal(t, *distribution, sdk.Distribution(o.Distribution))
		assert.Equal(t, 2, o.RetentionTime)
		assert.Equal(t, "test", o.Comment)
	})

	t.Run("alter application package: unset", func(t *testing.T) {
		e := createApplicationPackageHandle(t, client)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		// unset comment
		err := client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithUnsetComment(sdk.Bool(true)))
		require.NoError(t, err)
		o, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Empty(t, o.Comment)

		// unset distribution
		err = client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithUnsetDistribution(sdk.Bool(true)))
		require.NoError(t, err)
		o, err = client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, sdk.DistributionInternal, sdk.Distribution(o.Distribution))
	})

	t.Run("alter application package: set and unset tags", func(t *testing.T) {
		f := createApplicationPackageHandle(t, client)

		id := sdk.NewAccountObjectIdentifier(f.Name)
		setTags := []sdk.TagAssociation{
			{
				Name:  tagTest.ID(),
				Value: "v1",
			},
		}
		err := client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithSetTags(setTags))
		require.NoError(t, err)
		assertApplicationPackage(t, id)

		unsetTags := []sdk.ObjectIdentifier{
			tagTest.ID(),
		}
		err = client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithUnsetTags(unsetTags))
		require.NoError(t, err)
		assertApplicationPackage(t, id)
	})

	t.Run("show application package for SQL: with like", func(t *testing.T) {
		p := createApplicationPackageHandle(t, client)

		packages, err := client.ApplicationPackages.Show(ctx, sdk.NewShowApplicationPackageRequest().WithLike(&sdk.Like{Pattern: &p.Name}))
		require.NoError(t, err)

		require.Equal(t, 1, len(packages))
		require.Equal(t, *p, packages[0])
	})
}
