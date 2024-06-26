package sdk

import (
	"fmt"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/helpers/random"
)

func TestTagCreate(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *createTagOptions {
		return &createTagOptions{
			name: id,
		}
	}

	t.Run("create with all optional", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.OrReplace = Bool(false)
		opts.Comment = String("comment")
		opts.AllowedValues = &AllowedValues{
			Values: []AllowedValue{
				{
					Value: "value1",
				},
				{
					Value: "value2",
				},
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE TAG IF NOT EXISTS %s ALLOWED_VALUES 'value1', 'value2' COMMENT = 'comment'`, id.FullyQualifiedName())
	})

	t.Run("validation: nil options", func(t *testing.T) {
		opts := (*createTagOptions)(nil)
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: allowed values count", func(t *testing.T) {
		opts := defaultOpts()
		opts.AllowedValues = &AllowedValues{
			Values: []AllowedValue{},
		}
		assertOptsInvalidJoinedErrors(t, opts, errIntBetween("AllowedValues", "Values", 1, 300))
	})

	t.Run("validation: both ifNotExists and orReplace present", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.OrReplace = Bool(true)
		assertOptsInvalidJoinedErrors(t, opts, errOneOf("createTagOptions", "OrReplace", "IfNotExists"))
	})

	t.Run("validation: multiple errors", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		opts.IfNotExists = Bool(true)
		opts.OrReplace = Bool(true)
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier, errOneOf("createTagOptions", "OrReplace", "IfNotExists"))
	})
}

func TestTagDrop(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *dropTagOptions {
		return &dropTagOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *dropTagOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("drop with name", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `DROP TAG %s`, id.FullyQualifiedName())
	})

	t.Run("drop with if exists", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfExists = Bool(true)
		assertOptsValidAndSQLEquals(t, opts, `DROP TAG IF EXISTS %s`, id.FullyQualifiedName())
	})
}

func TestTagUndrop(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *undropTagOptions {
		return &undropTagOptions{
			name: id,
		}
	}
	t.Run("validation: nil options", func(t *testing.T) {
		var opts *dropTagOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("undrop with name", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `UNDROP TAG %s`, id.FullyQualifiedName())
	})
}

func TestTagShow(t *testing.T) {
	defaultOpts := func() *showTagOptions {
		return &showTagOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *showTagOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: empty like", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{}
		assertOptsInvalidJoinedErrors(t, opts, ErrPatternRequiredForLikeKeyword)
	})

	t.Run("validation: empty in", func(t *testing.T) {
		opts := defaultOpts()
		opts.In = &In{}
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("showTagOptions.In", "Account", "Database", "Schema"))
	})

	t.Run("show with empty options", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `SHOW TAGS`)
	})

	t.Run("show with like", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{Pattern: String("test")}
		assertOptsValidAndSQLEquals(t, opts, `SHOW TAGS LIKE 'test'`)
	})

	t.Run("show with in", func(t *testing.T) {
		opts := defaultOpts()
		opts.In = &In{
			Account: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW TAGS IN ACCOUNT`)
	})
}

func TestTagAlter(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *alterTagOptions {
		return &alterTagOptions{
			name: id,
		}
	}
	defaultAllowedValues := func() *AllowedValues {
		return &AllowedValues{
			Values: []AllowedValue{
				{
					Value: "value1",
				},
				{
					Value: "value2",
				},
			},
		}
	}
	mp1ID := NewSchemaObjectIdentifier(id.DatabaseName(), id.SchemaName(), "policy1")
	mp2ID := NewSchemaObjectIdentifier(id.DatabaseName(), id.SchemaName(), "policy2")
	defaultMaskingPolicies := func() []TagMaskingPolicy {
		return []TagMaskingPolicy{
			{
				Name: mp1ID,
			},
			{
				Name: mp2ID,
			},
		}
	}

	t.Run("alter with rename to", func(t *testing.T) {
		opts := defaultOpts()
		opts.Rename = &TagRename{Name: NewSchemaObjectIdentifier(id.DatabaseName(), id.SchemaName(), random.StringN(12))}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s RENAME TO %s`, id.FullyQualifiedName(), opts.Rename.Name.FullyQualifiedName())
	})

	t.Run("alter with add", func(t *testing.T) {
		opts := defaultOpts()
		opts.Add = &TagAdd{AllowedValues: defaultAllowedValues()}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s ADD ALLOWED_VALUES 'value1', 'value2'`, id.FullyQualifiedName())
	})

	t.Run("alter with drop", func(t *testing.T) {
		opts := defaultOpts()
		opts.Drop = &TagDrop{AllowedValues: defaultAllowedValues()}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s DROP ALLOWED_VALUES 'value1', 'value2'`, id.FullyQualifiedName())
	})

	t.Run("alter with unset allowed values", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &TagUnset{AllowedValues: Bool(true)}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s UNSET ALLOWED_VALUES`, id.FullyQualifiedName())
	})

	t.Run("alter with set masking policies", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &TagSet{
			MaskingPolicies: &TagSetMaskingPolicies{
				MaskingPolicies: defaultMaskingPolicies(),
				Force:           Bool(true),
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s SET MASKING POLICY %s, MASKING POLICY %s FORCE`, id.FullyQualifiedName(), mp1ID.FullyQualifiedName(), mp2ID.FullyQualifiedName())
	})

	t.Run("alter with unset masking policies", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &TagUnset{
			MaskingPolicies: &TagUnsetMaskingPolicies{
				MaskingPolicies: defaultMaskingPolicies(),
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s UNSET MASKING POLICY %s, MASKING POLICY %s`, id.FullyQualifiedName(), mp1ID.FullyQualifiedName(), mp2ID.FullyQualifiedName())
	})

	t.Run("alter with set comment", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &TagSet{Comment: String("comment")}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s SET COMMENT = 'comment'`, id.FullyQualifiedName())
	})

	t.Run("alter with unset comment", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &TagUnset{Comment: Bool(true)}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TAG %s UNSET COMMENT`, id.FullyQualifiedName())
	})

	t.Run("validation: nil options", func(t *testing.T) {
		opts := (*createTagOptions)(nil)
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: no alter action", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("alterTagOptions", "Add", "Drop", "Set", "Unset", "Rename"))
	})

	t.Run("validation: multiple alter actions", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &TagSet{
			Comment: String("comment"),
		}
		opts.Unset = &TagUnset{
			AllowedValues: Bool(true),
		}
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("alterTagOptions", "Add", "Drop", "Set", "Unset", "Rename"))
	})

	t.Run("validation: invalid new name", func(t *testing.T) {
		opts := defaultOpts()
		opts.Rename = &TagRename{
			Name: NewSchemaObjectIdentifier("", "", ""),
		}
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: new name from different db", func(t *testing.T) {
		newId := NewSchemaObjectIdentifier(id.DatabaseName()+random.StringN(1), random.StringN(12), random.StringN(12))

		opts := defaultOpts()
		opts.Rename = &TagRename{
			Name: newId,
		}
		assertOptsValid(t, opts)
	})

	t.Run("validation: no property to unset", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &TagUnset{}
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("TagUnset", "MaskingPolicies", "AllowedValues", "Comment"))
	})

	t.Run("validation: allowed values count", func(t *testing.T) {
		opts := defaultOpts()
		opts.Add = &TagAdd{
			AllowedValues: &AllowedValues{
				Values: []AllowedValue{},
			},
		}
		assertOptsInvalidJoinedErrors(t, opts, errIntBetween("AllowedValues", "Values", 1, 300))
	})
}

func TestTagSet(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *setTagOptions {
		return &setTagOptions{
			objectType: ObjectTypeStage,
			objectName: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		opts := (*setTagOptions)(nil)
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.objectName = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("set with all optional", func(t *testing.T) {
		opts := defaultOpts()
		opts.SetTags = []TagAssociation{
			{
				Name:  NewAccountObjectIdentifier("tag1"),
				Value: "value1",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER %s %s SET TAG "tag1" = 'value1'`, opts.objectType, id.FullyQualifiedName())
	})

	t.Run("set with column", func(t *testing.T) {
		objectName := NewTableColumnIdentifier("db1", "schema1", "table1", "column1")
		tableName := NewSchemaObjectIdentifier("db1", "schema1", "table1")
		request := NewSetTagRequest(ObjectTypeColumn, objectName).WithSetTags([]TagAssociation{
			{
				Name:  NewAccountObjectIdentifier("tag1"),
				Value: "value1",
			},
		})
		opts := request.toOpts()
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s MODIFY COLUMN "%s" SET TAG "tag1" = 'value1'`, tableName.FullyQualifiedName(), objectName.columnName)
	})
}

func TestTagUnset(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *unsetTagOptions {
		return &unsetTagOptions{
			objectType: ObjectTypeStage,
			objectName: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		opts := (*unsetTagOptions)(nil)
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.objectName = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("unset with all optional", func(t *testing.T) {
		opts := defaultOpts()
		opts.UnsetTags = []ObjectIdentifier{
			NewAccountObjectIdentifier("tag1"),
			NewAccountObjectIdentifier("tag2"),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER %s %s UNSET TAG "tag1", "tag2"`, opts.objectType, id.FullyQualifiedName())
	})

	t.Run("unset with column", func(t *testing.T) {
		table, column := NewSchemaObjectIdentifier("db1", "schema1", "table1"), "column1"
		objectName := NewObjectIdentifierFromFullyQualifiedName(fmt.Sprintf("%s.%s.%s.%s", table.DatabaseName(), table.SchemaName(), table.Name(), column))
		request := UnsetTagRequest{
			objectType: ObjectTypeColumn,
			objectName: objectName,
			UnsetTags: []ObjectIdentifier{
				NewAccountObjectIdentifier("tag1"),
				NewAccountObjectIdentifier("tag2"),
			},
		}
		opts := request.toOpts()
		assertOptsValidAndSQLEquals(t, opts, `ALTER %s %s MODIFY COLUMN "%s" UNSET TAG "tag1", "tag2"`, opts.objectType, table.FullyQualifiedName(), column)
	})
}
