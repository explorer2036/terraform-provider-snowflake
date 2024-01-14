package testint

import (
	"errors"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/require"
)

func TestInt_Sequences(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	databaseTest, schemaTest := testDb(t), testSchema(t)

	cleanupSequenceHandle := func(t *testing.T, id sdk.SchemaObjectIdentifier) func() {
		t.Helper()
		return func() {
			err := client.Sequences.Drop(ctx, sdk.NewDropSequenceRequest(id))
			if errors.Is(err, sdk.ErrObjectNotExistOrAuthorized) {
				return
			}
			require.NoError(t, err)
		}
	}

	t.Run("create sequence", func(t *testing.T) {
		name := random.StringN(4)
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		comment := random.StringN(4)
		request := sdk.NewCreateSequenceRequest(id).
			WithWith(sdk.Bool(true)).
			WithStart(sdk.Int(1)).
			WithIncrement(sdk.Int(1)).
			WithIfNotExists(sdk.Bool(true)).
			WithValuesBehavior(sdk.ValuesBehaviorPointer(sdk.ValuesBehaviorOrder)).
			WithComment(&comment)
		err := client.Sequences.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupSequenceHandle(t, id))
	})
}
