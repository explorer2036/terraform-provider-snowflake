package testint

import (
	"context"
	"errors"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt_ApplicationPackages(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	cleanupApplicationPackageHandle := func(id sdk.AccountObjectIdentifier) func() {
		return func() {
			err := client.ApplicationPackages.Drop(ctx, sdk.NewDropApplicationPackageRequest(id))
			if errors.Is(err, sdk.ErrObjectNotExistOrAuthorized) {
				return
			}
			require.NoError(t, err)
		}
	}

	createTagHandle := func(t *testing.T, client *sdk.Client) *sdk.Tag {
		t.Helper()
		schema, schemaCleanup := createSchema(t, client, testDb(t))
		t.Cleanup(schemaCleanup)
		tag, tagCleanup := createTag(t, client, testDb(t), schema)
		t.Cleanup(tagCleanup)
		return tag
	}

	createApplicationPackageHandle := func(t *testing.T, client *sdk.Client) *sdk.ApplicationPackage {
		t.Helper()

		id := sdk.NewAccountObjectIdentifier(random.String())
		err := client.ApplicationPackages.Create(ctx, sdk.NewCreateApplicationPackageRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationPackageHandle(id))

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		return e
	}

	t.Run("create application package", func(t *testing.T) {
		tag := createTagHandle(t, client)

		name := random.String()
		id := sdk.NewAccountObjectIdentifier(name)
		comment := random.String()
		request := sdk.NewCreateApplicationPackageRequest(id).
			WithComment(&comment).
			WithTag([]sdk.TagAssociation{
				{
					Name:  tag.ID(),
					Value: "abc",
				},
			}).
			WithDistribution(sdk.String("INTERNAL"))
		err := client.ApplicationPackages.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationPackageHandle(id))

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, comment, e.Comment)
	})

	t.Run("create application package: no optionals", func(t *testing.T) {
		name := random.String()
		id := sdk.NewAccountObjectIdentifier(name)
		err := client.ApplicationPackages.Create(ctx, sdk.NewCreateApplicationPackageRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationPackageHandle(id))

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, "", e.Comment)
	})

	t.Run("drop application package: existing", func(t *testing.T) {
		e := createApplicationPackageHandle(t, client)

		id := sdk.NewAccountObjectIdentifier(e.Name)
		err := client.ApplicationPackages.Drop(ctx, sdk.NewDropApplicationPackageRequest(id))
		require.NoError(t, err)
	})

	t.Run("drop application package: no-existing", func(t *testing.T) {
		id := sdk.NewAccountObjectIdentifier(random.String())
		err := client.ApplicationPackages.Drop(ctx, sdk.NewDropApplicationPackageRequest(id))
		assert.ErrorIs(t, err, sdk.ErrObjectNotExistOrAuthorized)
	})

	t.Run("alter application package: set and unset comment", func(t *testing.T) {
		e := createApplicationPackageHandle(t, client)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		comment := random.Comment()
		set := sdk.NewApplicationPackageSetRequest().WithComment(&comment)
		err := client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithSet(set))
		require.NoError(t, err)

		res, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, comment, res.Comment)

		unset := sdk.NewApplicationPackageUnsetRequest().WithComment(sdk.Bool(true))
		err = client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithUnset(unset))
		require.NoError(t, err)

		res, err = client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, "", res.Comment)
	})

	t.Run("alter application package: set distribution", func(t *testing.T) {
		e := createApplicationPackageHandle(t, client)
		id := sdk.NewAccountObjectIdentifier(e.Name)

		distribution := "EXTERNAL"
		set := sdk.NewApplicationPackageSetRequest().WithDistribution(&distribution)
		err := client.ApplicationPackages.Alter(ctx, sdk.NewAlterApplicationPackageRequest(id).WithSet(set))
		require.NoError(t, err)

		res, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, distribution, res.Distribution)
	})

	t.Run("show application package: without like", func(t *testing.T) {
		e := createApplicationPackageHandle(t, client)

		res, err := client.ApplicationPackages.Show(ctx, sdk.NewShowApplicationPackageRequest())
		require.NoError(t, err)

		assert.Equal(t, 1, len(res))
		assert.Contains(t, res, *e)
	})

	t.Run("show application package: with like", func(t *testing.T) {
		e1 := createApplicationPackageHandle(t, client)
		e2 := createApplicationPackageHandle(t, client)

		res, err := client.ApplicationPackages.Show(ctx, sdk.NewShowApplicationPackageRequest().WithLike(e1.Name))
		require.NoError(t, err)

		assert.Equal(t, 1, len(res))
		assert.Contains(t, res, *e1)
		assert.NotContains(t, res, *e2)
	})

	t.Run("show application package: no matches", func(t *testing.T) {
		res, err := client.ApplicationPackages.Show(ctx, sdk.NewShowApplicationPackageRequest().WithLike("no match"))
		require.NoError(t, err)
		assert.Equal(t, 0, len(res))
	})
}
