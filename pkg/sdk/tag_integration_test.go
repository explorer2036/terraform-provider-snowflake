package sdk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt_TagCreate(t *testing.T) {
	client := testClient(t)

	databaseTest, databaseCleanup := createDatabase(t, client)
	t.Cleanup(databaseCleanup)
	_, schemaCleanup := createSchema(t, client, databaseTest)
	t.Cleanup(schemaCleanup)

	ctx := context.Background()
	t.Run("create with comment", func(t *testing.T) {
		name := randomAccountObjectIdentifier(t)
		comment := randomComment(t)
		err := client.Tags.Create(ctx, NewCreateTagRequest(name).WithOrReplace(true).WithComment(&comment))
		require.NoError(t, err)
		t.Cleanup(func() {
			err = client.Tags.Drop(ctx, NewDropTagRequest(name))
			require.NoError(t, err)
		})
		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(name.Name()))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))

		entity := entities[0]
		require.Equal(t, name.Name(), entity.Name)
		require.Equal(t, comment, entity.Comment)
	})

	t.Run("create with one allowed value", func(t *testing.T) {
		name := randomAccountObjectIdentifier(t)
		values := []string{"value1"}
		err := client.Tags.Create(ctx, NewCreateTagRequest(name).WithOrReplace(true).WithAllowedValues(values))
		require.NoError(t, err)
		t.Cleanup(func() {
			err = client.Tags.Drop(ctx, NewDropTagRequest(name))
			require.NoError(t, err)
		})
		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(name.Name()))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))

		entity := entities[0]
		require.Equal(t, name.Name(), entity.Name)
		require.Equal(t, values, entity.AllowedValues)
	})

	t.Run("create with two allowed values", func(t *testing.T) {
		name := randomAccountObjectIdentifier(t)
		values := []string{"value1", "value2"}
		err := client.Tags.Create(ctx, NewCreateTagRequest(name).WithOrReplace(true).WithAllowedValues(values))
		require.NoError(t, err)
		t.Cleanup(func() {
			err = client.Tags.Drop(ctx, NewDropTagRequest(name))
			require.NoError(t, err)
		})
		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(name.Name()))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))

		entity := entities[0]
		require.Equal(t, name.Name(), entity.Name)
		require.Equal(t, values, entity.AllowedValues)
	})

	t.Run("create with comment and allowed values", func(t *testing.T) {
		name := randomAccountObjectIdentifier(t)
		comment := randomComment(t)
		values := []string{"value1"}
		err := client.Tags.Create(ctx, NewCreateTagRequest(name).WithOrReplace(true).WithComment(&comment).WithAllowedValues(values))
		expected := "fields [Comment AllowedValues] are incompatible and cannot be set at once"
		require.Equal(t, expected, err.Error())
	})
}

func TestInt_TagAlter(t *testing.T) {
	client := testClient(t)

	ctx := context.Background()
	t.Run("alter with set and unset comment", func(t *testing.T) {
		databaseTest, databaseCleanup := createDatabase(t, client)
		t.Cleanup(databaseCleanup)
		schemaTest, schemaCleanup := createSchema(t, client, databaseTest)
		t.Cleanup(schemaCleanup)
		tag, tagCleanup := createTag(t, client, databaseTest, schemaTest)
		t.Cleanup(tagCleanup)

		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, "", entities[0].Comment)

		comment := randomComment(t)
		set := NewTagSetRequest().WithComment(comment)
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithSet(set))
		require.NoError(t, err)

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, comment, entities[0].Comment)

		unset := NewTagUnsetRequest().WithComment(true)
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithUnset(unset))
		require.NoError(t, err)

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, "", entities[0].Comment)
	})

	t.Run("alter with set and unset masking policies", func(t *testing.T) {
		databaseTest, databaseCleanup := createDatabase(t, client)
		t.Cleanup(databaseCleanup)
		schemaTest, schemaCleanup := createSchema(t, client, databaseTest)
		t.Cleanup(schemaCleanup)
		policyTest, policyCleanup := createMaskingPolicy(t, client, databaseTest, schemaTest)
		t.Cleanup(policyCleanup)
		tag, tagCleanup := createTag(t, client, databaseTest, schemaTest)
		t.Cleanup(tagCleanup)

		policies := []string{policyTest.Name}
		set := NewTagSetRequest().WithMaskingPolicies(policies)
		err := client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithSet(set))
		require.NoError(t, err)

		unset := NewTagUnsetRequest().WithMaskingPolicies(policies)
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithUnset(unset))
		require.NoError(t, err)
	})

	t.Run("alter with add and drop allowed values", func(t *testing.T) {
		databaseTest, databaseCleanup := createDatabase(t, client)
		t.Cleanup(databaseCleanup)
		schemaTest, schemaCleanup := createSchema(t, client, databaseTest)
		t.Cleanup(schemaCleanup)
		tag, tagCleanup := createTag(t, client, databaseTest, schemaTest)
		t.Cleanup(tagCleanup)

		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, 0, len(entities[0].AllowedValues))

		values := []string{"value1"}
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithAdd(values))
		require.NoError(t, err)

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, values, entities[0].AllowedValues)

		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithDrop(values))
		require.NoError(t, err)

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, 0, len(entities[0].AllowedValues))
	})

	t.Run("alter with rename", func(t *testing.T) {
		databaseTest, databaseCleanup := createDatabase(t, client)
		t.Cleanup(databaseCleanup)
		schemaTest, schemaCleanup := createSchema(t, client, databaseTest)
		t.Cleanup(schemaCleanup)
		tag, _ := createTag(t, client, databaseTest, schemaTest)

		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, "", entities[0].Comment)

		id := randomAccountObjectIdentifier(t)
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithRename(id))
		require.NoError(t, err)
		t.Cleanup(func() {
			err = client.Tags.Drop(ctx, NewDropTagRequest(id))
			require.NoError(t, err)
		})

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(id.Name()))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
	})

	t.Run("alter with unset allowed values", func(t *testing.T) {
		databaseTest, databaseCleanup := createDatabase(t, client)
		t.Cleanup(databaseCleanup)
		schemaTest, schemaCleanup := createSchema(t, client, databaseTest)
		t.Cleanup(schemaCleanup)
		tag, tagCleanup := createTag(t, client, databaseTest, schemaTest)
		t.Cleanup(tagCleanup)

		entities, err := client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, 0, len(entities[0].AllowedValues))

		values := []string{"value1", "value2"}
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithAdd(values))
		require.NoError(t, err)

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, values, entities[0].AllowedValues)

		unset := NewTagUnsetRequest().WithAllowedValues(true)
		err = client.Tags.Alter(ctx, NewAlterTagRequest(NewAccountObjectIdentifier(tag.Name)).WithUnset(unset))
		require.NoError(t, err)

		entities, err = client.Tags.Show(ctx, NewShowTagRequest().WithLike(tag.Name))
		require.NoError(t, err)
		require.Equal(t, 1, len(entities))
		require.Equal(t, 0, len(entities[0].AllowedValues))
	})
}
