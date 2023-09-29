package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagCreate(t *testing.T) {
	t.Run("create with allowed values", func(t *testing.T) {
		opts := &createTagOptions{
			OrReplace: Bool(true),
			name: SchemaObjectIdentifier{
				databaseName: "db",
				schemaName:   "schema",
				name:         "tag",
			},
			AllowedValues: &AllowedValues{
				Values: []AllowedValue{
					{
						Value: "value1",
					},
					{
						Value: "value2",
					},
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `CREATE OR REPLACE TAG "db"."schema"."tag" ALLOWED_VALUES 'value1', 'value2'`
		assert.Equal(t, expected, actual)
	})

	t.Run("create with comment", func(t *testing.T) {
		opts := &createTagOptions{
			OrReplace: Bool(true),
			name: SchemaObjectIdentifier{
				databaseName: "db",
				schemaName:   "schema",
				name:         "tag",
			},
			Comment: String("comment"),
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `CREATE OR REPLACE TAG "db"."schema"."tag" COMMENT = 'comment'`
		assert.Equal(t, expected, actual)
	})

	t.Run("create with not exists", func(t *testing.T) {
		opts := &createTagOptions{
			IfNotExists: Bool(true),
			name: SchemaObjectIdentifier{
				databaseName: "db",
				schemaName:   "schema",
				name:         "tag",
			},
			Comment: String("comment"),
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `CREATE TAG IF NOT EXISTS "db"."schema"."tag" COMMENT = 'comment'`
		assert.Equal(t, expected, actual)
	})
}

func TestTagDrop(t *testing.T) {
	t.Run("drop with name", func(t *testing.T) {
		opts := &dropTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `DROP TAG "db"."schema"."test"`
		assert.Equal(t, expected, actual)
	})
}

func TestTagUndrop(t *testing.T) {
	t.Run("undrop with name", func(t *testing.T) {
		opts := &undropTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `UNDROP TAG "db"."schema"."test"`
		assert.Equal(t, expected, actual)
	})
}

func TestTagShow(t *testing.T) {
	t.Run("show with empty options", func(t *testing.T) {
		opts := &showTagOptions{}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `SHOW TAGS`
		assert.Equal(t, expected, actual)
	})

	t.Run("show with like", func(t *testing.T) {
		opts := &showTagOptions{
			Like: &Like{
				Pattern: String("test"),
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `SHOW TAGS LIKE 'test'`
		assert.Equal(t, expected, actual)
	})
}

func TestTagAlter(t *testing.T) {
	t.Run("alter with rename to", func(t *testing.T) {
		opts := &alterTagOptions{
			name:   NewSchemaObjectIdentifier("db", "schema", "test"),
			Rename: &TagRename{Name: NewSchemaObjectIdentifier("db", "schema", "test2")},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" RENAME TO "db"."schema"."test2"`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with add", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Add: &TagAdd{
				AllowedValues: &AllowedValues{
					Values: []AllowedValue{
						{
							Value: "value1",
						},
						{
							Value: "value2",
						},
					},
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" ADD ALLOWED_VALUES 'value1', 'value2'`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with drop", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Drop: &TagDrop{
				AllowedValues: &AllowedValues{
					Values: []AllowedValue{
						{
							Value: "value1",
						},
						{
							Value: "value2",
						},
					},
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" DROP ALLOWED_VALUES 'value1', 'value2'`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with unset allowed values", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Unset: &TagUnset{
				AllowedValues: Bool(true),
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" UNSET ALLOWED_VALUES`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with set masking policy", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Set: &TagSet{
				MaskingPolicies: &TagSetMaskingPolicies{
					MaskingPolicies: []TagMaskingPolicy{
						{
							Name: "policy1",
						},
					},
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" SET MASKING POLICY "policy1"`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with set masking policies", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Set: &TagSet{
				MaskingPolicies: &TagSetMaskingPolicies{
					MaskingPolicies: []TagMaskingPolicy{
						{
							Name: "policy1",
						},
						{
							Name: "policy2",
						},
					},
					Force: Bool(true),
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" SET MASKING POLICY "policy1", MASKING POLICY "policy2" FORCE`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with unset masking policy", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Unset: &TagUnset{
				MaskingPolicies: &TagUnsetMaskingPolicies{
					MaskingPolicies: []TagMaskingPolicy{
						{
							Name: "policy1",
						},
					},
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" UNSET MASKING POLICY "policy1"`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with unset masking policies", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Unset: &TagUnset{
				MaskingPolicies: &TagUnsetMaskingPolicies{
					MaskingPolicies: []TagMaskingPolicy{
						{
							Name: "policy1",
						},
						{
							Name: "policy2",
						},
					},
				},
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" UNSET MASKING POLICY "policy1", MASKING POLICY "policy2"`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with set comment", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Set: &TagSet{
				Comment: String("comment"),
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" SET COMMENT = 'comment'`
		assert.Equal(t, expected, actual)
	})

	t.Run("alter with unset comment", func(t *testing.T) {
		opts := &alterTagOptions{
			name: NewSchemaObjectIdentifier("db", "schema", "test"),
			Unset: &TagUnset{
				Comment: Bool(true),
			},
		}
		actual, err := structToSQL(opts)
		require.NoError(t, err)
		expected := `ALTER TAG "db"."schema"."test" UNSET COMMENT`
		assert.Equal(t, expected, actual)
	})
}
