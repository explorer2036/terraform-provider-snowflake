package sdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt_EventTables(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	databaseTest, databaseCleanup := createDatabase(t, client)
	t.Cleanup(databaseCleanup)
	schemaTest, schemaCleanup := createSchema(t, client, databaseTest)
	t.Cleanup(schemaCleanup)

	assertEventTableHandle := func(t *testing.T, et *EventTable, expectedName string, expectedComment string, expectedAllowedValues []string) {
		t.Helper()
		assert.NotEmpty(t, et.CreatedOn)
		assert.Equal(t, expectedName, et.Name)
		assert.Equal(t, "ACCOUNTADMIN", et.Owner)
		assert.Equal(t, expectedComment, et.Comment)
	}

	cleanupTableHandle := func(t *testing.T, id SchemaObjectIdentifier) func() {
		return func() {
			_, err := client.exec(ctx, fmt.Sprintf("DROP TABLE \"%s\".\"%s\".\"%s\"", id.DatabaseName(), id.SchemaName(), id.Name()))
			require.NoError(t, err)
		}
	}

	createEventTableHandle := func(t *testing.T) *EventTable {
		t.Helper()

		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, randomString(t))
		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		return et
	}

	t.Run("create event table: comment", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)
		comment := randomComment(t)

		request := NewCreateEventTableRequest(id).WithComment(comment)
		err := client.EventTables.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, et, name, comment, nil)
	})

	t.Run("create event table: properties", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		request := NewCreateEventTableRequest(id).
			WithChangeTracking(true).
			WithDefaultDDLCollation("en_US").
			WithDataRetentionTimeInDays(1).
			WithMaxDataExtensionTimeInDays(2)
		err := client.EventTables.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, true, et.ChangeTracking)
	})

	t.Run("create event table: copy grants", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		request := NewCreateEventTableRequest(id).
			WithOrReplace(true).
			WithCopyGrants(true)
		err := client.EventTables.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		_, err = client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
	})

	t.Run("create event table: tag", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		tag, tagCleanup := createTag(t, client, databaseTest, schemaTest)
		t.Cleanup(tagCleanup)

		request := NewCreateEventTableRequest(id).
			WithTag([]TagAssociationRequest{
				{
					name:  tag.ID(),
					value: "value1",
				},
			})
		err := client.EventTables.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))
	})

	t.Run("create event table: no optionals", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		tag, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, tag, name, "", nil)
	})

	t.Run("alter event table: set and unset comment", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		comment := randomComment(t)
		set := NewEventTableSetRequest().WithComment(comment)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithSet(set))
		require.NoError(t, err)

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, et, name, comment, nil)

		unset := NewEventTableUnsetRequest().WithComment(true)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithUnset(unset))
		require.NoError(t, err)

		et, err = client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, et, name, "", nil)
	})

	t.Run("alter event table: set and unset change tacking", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		set := NewEventTableSetRequest().WithChangeTracking(true)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithSet(set))
		require.NoError(t, err)

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, true, et.ChangeTracking)

		unset := NewEventTableUnsetRequest().WithChangeTracking(true)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithUnset(unset))
		require.NoError(t, err)

		et, err = client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, false, et.ChangeTracking)
	})

	t.Run("alter event table: set and unset tag", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		tag, tagCleanup := createTag(t, client, databaseTest, schemaTest)
		t.Cleanup(tagCleanup)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		tr := []TagAssociationRequest{
			{
				name:  tag.ID(),
				value: "value1",
			},
		}
		set := NewEventTableSetRequest().WithTag(tr)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithSet(set))
		require.NoError(t, err)

		unset := NewEventTableUnsetRequest().WithTag([]string{tag.ID().FullyQualifiedName()})
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithUnset(unset))
		require.NoError(t, err)
	})

	t.Run("alter event table: rename", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)

		nid := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, randomString(t))
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithRename(nid))
		if err != nil {
			t.Cleanup(cleanupTableHandle(t, id))
		} else {
			t.Cleanup(cleanupTableHandle(t, nid))
		}
		require.NoError(t, err)

		_, err = client.EventTables.ShowByID(ctx, id)
		assert.ErrorIs(t, err, errObjectNotExistOrAuthorized)

		_, err = client.EventTables.ShowByID(ctx, nid)
		require.NoError(t, err)
	})

	t.Run("alter event table: clustering action with drop", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		action := NewClusteringActionRequest().WithDrop(true)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithClusteringAction(action))
		require.NoError(t, err)
	})

	t.Run("show event table: without like", func(t *testing.T) {
		et1 := createEventTableHandle(t)
		et2 := createEventTableHandle(t)

		tables, err := client.EventTables.Show(ctx, NewShowEventTableRequest())
		require.NoError(t, err)

		assert.Equal(t, 2, len(tables))
		assert.Contains(t, tables, *et1)
		assert.Contains(t, tables, *et2)
	})

	t.Run("show event table: with like", func(t *testing.T) {
		et1 := createEventTableHandle(t)
		et2 := createEventTableHandle(t)

		tables, err := client.EventTables.Show(ctx, NewShowEventTableRequest().WithLike(et1.Name))
		require.NoError(t, err)
		assert.Equal(t, 1, len(tables))
		assert.Contains(t, tables, *et1)
		assert.NotContains(t, tables, *et2)
	})

	t.Run("show event table: no matches", func(t *testing.T) {
		tables, err := client.EventTables.Show(ctx, NewShowEventTableRequest().WithLike("non-existent"))
		require.NoError(t, err)
		assert.Equal(t, 0, len(tables))
	})

	t.Run("describe event table", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)
		t.Cleanup(cleanupTableHandle(t, id))

		details, err := client.EventTables.Describe(ctx, NewDescribeEventTableRequest(id))
		require.NoError(t, err)
		assert.Equal(t, "TIMESTAMP", details.Name)
	})
}
