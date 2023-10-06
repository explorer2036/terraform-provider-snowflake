package sdk

import (
	"context"
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

	t.Run("create event table: comment", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)
		comment := randomComment(t)

		request := NewCreateEventTableRequest(id).WithComment(comment)
		err := client.EventTables.Create(ctx, request)
		require.NoError(t, err)

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
	})

	t.Run("create event table: no optionals", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)

		tag, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, tag, name, "", nil)
	})

	t.Run("alter event table: set and unset comment", func(t *testing.T) {
		name := randomString(t)
		id := NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)

		err := client.EventTables.Create(ctx, NewCreateEventTableRequest(id))
		require.NoError(t, err)

		comment := randomComment(t)
		set := NewEventTableSetRequest().WithComment(comment)
		err = client.EventTables.Alter(ctx, NewAlterEventTableRequest(id).WithSet(set))
		require.NoError(t, err)

		et, err := client.EventTables.ShowByID(ctx, id)
		require.NoError(t, err)
		assertEventTableHandle(t, et, name, comment, nil)
	})
}
