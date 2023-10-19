package testint

import (
	"context"
	"errors"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/require"
)

func TestInt_ApplicationPackages(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	cleanupApplicationPacketHandle := func(id sdk.AccountObjectIdentifier) func() {
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

	t.Run("create application package", func(t *testing.T) {
		tag := createTagHandle(t, client)

		name := random.String()
		id := sdk.NewAccountObjectIdentifier(name)

		comment := random.String()
		tags := []sdk.TagAssociation{
			{
				Name:  tag.ID(),
				Value: "abc",
			},
		}
		request := sdk.NewCreateApplicationPackageRequest(id).
			WithComment(&comment).
			WithTag(tags).
			WithDistribution(sdk.String("INTERNAL"))

		err := client.ApplicationPackages.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationPacketHandle(id))

		e, err := client.ApplicationPackages.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, comment, e.Comment)
	})
}
