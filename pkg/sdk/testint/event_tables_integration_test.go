package testint

import (
	"context"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt_EventTables(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	databaseTest, schemaTest := testDb(t), testSchema(t)
	tagTest, tagCleaup := createTag(t, client, databaseTest, schemaTest)
	t.Cleanup(tagCleaup)

	assertEventTableHandle := func(t *testing.T, et *sdk.EventTable, expectedName string, expectedComment string, expectedAllowedValues []string) {
		t.Helper()
		assert.NotEmpty(t, et.CreatedOn)
		assert.Equal(t, expectedName, et.Name)
		assert.Equal(t, "ACCOUNTADMIN", et.Owner)
		assert.Equal(t, expectedComment, et.Comment)
	}

	cleanupTableHandle := func(t *testing.T, id sdk.SchemaObjectIdentifier) func() {
		t.Helper()
		return func() {
			err := client.EventTables.Drop(ctx, sdk.NewDropEventTableRequest(id).WithIfExists(sdk.Bool(true)))
			require.NoError(t, err)
		}
	}

	createEventTableHandle := func(t *testing.T) *sdk.EventTable {
		t.Helper()

		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, random.StringN(4))
		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		return et
	}

	t.Run("create event tables: all options", func(t *testing.T) {
		name := random.StringN(4)
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		request := sdk.NewCreateEventTableRequest(id).
			WithChangeTracking(sdk.Bool(true)).
			WithDefaultDdlCollation(sdk.String("en_US")).
			WithDataRetentionTimeInDays(sdk.Int(1)).
			WithMaxDataExtensionTimeInDays(sdk.Int(2)).
			WithComment(sdk.String("test")).
			WithIfNotExists(sdk.Bool(true)).
			WithTag([]sdk.TagAssociation{
				{
					Name:  tagTest.ID(),
					Value: "v1",
				},
			})
		err := client.EventTables.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))
	})

	t.Run("show event table: without like", func(t *testing.T) {
		et1 := createEventTableHandle(t)
		et2 := createEventTableHandle(t)

		tables, err := client.EventTables.Show(ctx, sdk.NewShowEventTableRequest())
		require.NoError(t, err)

		assert.Equal(t, 2, len(tables))
		assert.Contains(t, tables, *et1)
		assert.Contains(t, tables, *et2)
	})

	t.Run("show event table: with like", func(t *testing.T) {
		et1 := createEventTableHandle(t)
		et2 := createEventTableHandle(t)

		tables, err := client.EventTables.Show(ctx, sdk.NewShowEventTableRequest().WithLike(&sdk.Like{Pattern: &et1.Name}))
		require.NoError(t, err)
		assert.Equal(t, 1, len(tables))
		assert.Contains(t, tables, *et1)
		assert.NotContains(t, tables, *et2)
	})

	t.Run("show event table: no matches", func(t *testing.T) {
		tables, err := client.EventTables.Show(ctx, sdk.NewShowEventTableRequest().WithLike(&sdk.Like{Pattern: sdk.String("non-existent")}))
		require.NoError(t, err)
		assert.Equal(t, 0, len(tables))
	})

	t.Run("describe event table", func(t *testing.T) {
		name := random.String()
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		details, err := client.EventTables.Describe(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, "TIMESTAMP", details.Name)
	})

	t.Run("alter event table: set and unset comment", func(t *testing.T) {
		name := random.String()
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		comment := random.Comment()
		set := sdk.NewEventTableSetRequest().WithComment(&comment)
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithSet(set))
		require.NoError(t, err)

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, et, name, comment, nil)

		unset := sdk.NewEventTableUnsetRequest().WithComment(sdk.Bool(true))
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithUnset(unset))
		require.NoError(t, err)

		et, err = client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, et, name, "", nil)
	})

	t.Run("alter event table: set and unset change tacking", func(t *testing.T) {
		name := random.String()
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		set := sdk.NewEventTableSetRequest().WithChangeTracking(sdk.Bool(true))
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithSet(set))
		require.NoError(t, err)

		unset := sdk.NewEventTableUnsetRequest().WithChangeTracking(sdk.Bool(true))
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithUnset(unset))
		require.NoError(t, err)
	})

	t.Run("alter event table: set and unset tag", func(t *testing.T) {
		name := random.String()
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		set := []sdk.TagAssociation{
			{
				Name:  tagTest.ID(),
				Value: "v1",
			},
		}
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithSetTags(set))
		require.NoError(t, err)

		unset := []sdk.ObjectIdentifier{tagTest.ID()}
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithUnsetTags(unset))
		require.NoError(t, err)
	})

	t.Run("alter event table: rename", func(t *testing.T) {
		name := random.String()
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)

		nid := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, random.String())
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithRenameTo(&nid))
		if err != nil {
			t.Cleanup(cleanupTableHandle(t, id))
		} else {
			t.Cleanup(cleanupTableHandle(t, nid))
		}
		require.NoError(t, err)

		_, err = client.EventTables.ShowByID(ctx, id)
		assert.ErrorIs(t, err, collections.ErrObjectNotFound)

		_, err = client.EventTables.ShowByID(ctx, nid)
		require.NoError(t, err)
	})

	t.Run("alter event table: clustering action with drop", func(t *testing.T) {
		name := random.String()
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, sdk.NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		action := sdk.NewEventTableClusteringActionRequest().WithDropClusteringKey(sdk.Bool(true))
		err = client.EventTables.Alter(ctx, sdk.NewAlterEventTableRequest(id).WithClusteringAction(action))
		require.NoError(t, err)
	})

	// alter view: add and drop row access policies
	t.Run("alter event table: add and drop row access policies", func(t *testing.T) {
		rowAccessPolicyId, rowAccessPolicyCleanup := createRowAccessPolicy(t, client, testSchema(t))
		t.Cleanup(rowAccessPolicyCleanup)
		rowAccessPolicy2Id, rowAccessPolicy2Cleanup := createRowAccessPolicy(t, client, testSchema(t))
		t.Cleanup(rowAccessPolicy2Cleanup)

		table, tableCleanup := createTable(t, client, databaseTest, schemaTest)
		t.Cleanup(tableCleanup)
		id := sdk.NewSchemaObjectIdentifier(table.DatabaseName, table.SchemaName, table.Name)

		// add policy
		alterRequest := sdk.NewAlterEventTableRequest(id).WithAddRowAccessPolicy(sdk.NewEventTableAddRowAccessPolicyRequest(rowAccessPolicyId, []string{"id"}))
		err := client.EventTables.Alter(ctx, alterRequest)
		require.NoError(t, err)

		e, err := getRowAccessPolicyFor(t, client, table.ID(), sdk.ObjectTypeTable)
		require.NoError(t, err)
		assert.Equal(t, rowAccessPolicyId.Name(), e.PolicyName)
		assert.Equal(t, "ROW_ACCESS_POLICY", e.PolicyKind)
		assert.Equal(t, table.ID().Name(), e.RefEntityName)
		assert.Equal(t, "TABLE", e.RefEntityDomain)
		assert.Equal(t, "ACTIVE", e.PolicyStatus)

		// remove policy
		alterRequest = sdk.NewAlterEventTableRequest(id).WithDropRowAccessPolicy(sdk.NewEventTableDropRowAccessPolicyRequest(rowAccessPolicyId))
		err = client.EventTables.Alter(ctx, alterRequest)
		require.NoError(t, err)

		_, err = getRowAccessPolicyFor(t, client, table.ID(), sdk.ObjectTypeTable)
		require.Error(t, err, "no rows in result set")

		// add policy again
		alterRequest = sdk.NewAlterEventTableRequest(id).WithAddRowAccessPolicy(sdk.NewEventTableAddRowAccessPolicyRequest(rowAccessPolicyId, []string{"id"}))
		err = client.EventTables.Alter(ctx, alterRequest)
		require.NoError(t, err)

		e, err = getRowAccessPolicyFor(t, client, table.ID(), sdk.ObjectTypeTable)
		require.NoError(t, err)
		assert.Equal(t, rowAccessPolicyId.Name(), e.PolicyName)

		// drop and add other policy simultaneously
		alterRequest = sdk.NewAlterEventTableRequest(id).WithDropAndAddRowAccessPolicy(sdk.NewEventTableDropAndAddRowAccessPolicyRequest(
			*sdk.NewEventTableDropRowAccessPolicyRequest(rowAccessPolicyId),
			*sdk.NewEventTableAddRowAccessPolicyRequest(rowAccessPolicy2Id, []string{"id"}),
		))
		err = client.EventTables.Alter(ctx, alterRequest)
		require.NoError(t, err)

		e, err = getRowAccessPolicyFor(t, client, table.ID(), sdk.ObjectTypeTable)
		require.NoError(t, err)
		assert.Equal(t, rowAccessPolicy2Id.Name(), e.PolicyName)

		// drop all policies
		alterRequest = sdk.NewAlterEventTableRequest(id).WithDropAllRowAccessPolicies(sdk.Bool(true))
		err = client.EventTables.Alter(ctx, alterRequest)
		require.NoError(t, err)

		_, err = getRowAccessPolicyFor(t, client, table.ID(), sdk.ObjectTypeView)
		require.Error(t, err, "no rows in result set")
	})
}
