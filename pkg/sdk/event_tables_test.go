package sdk

import (
	"testing"
)

func TestEventTablesCreate(t *testing.T) {
	id := randomSchemaObjectIdentifier(t)
	defaultOpts := func() *createEventTableOptions {
		return &createEventTableOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *createEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, errNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, errInvalidObjectIdentifier)
	})

	t.Run("validation: both ifNotExists and orReplace present", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.OrReplace = Bool(true)
		assertOptsInvalidJoinedErrors(t, opts, errOneOf("OrReplace", "IfNotExists"))
	})

	t.Run("validation: multiple errors", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		opts.IfNotExists = Bool(true)
		opts.OrReplace = Bool(true)
		assertOptsInvalidJoinedErrors(t, opts, errInvalidObjectIdentifier, errOneOf("OrReplace", "IfNotExists"))
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `CREATE EVENT TABLE %s`, id.FullyQualifiedName())
	})

	t.Run("empty slice", func(t *testing.T) {
		opts := defaultOpts()
		opts.ClusterBy = []string{}
		opts.Tag = []TagAssociation{}
		assertOptsValidAndSQLEquals(t, opts, `CREATE EVENT TABLE %s`, id.FullyQualifiedName())
	})

	t.Run("all optional", func(t *testing.T) {
		opts := defaultOpts()
		opts.OrReplace = Bool(true)
		opts.ClusterBy = []string{"a", "b"}
		opts.DataRetentionTimeInDays = Uint(1)
		opts.MaxDataExtensionTimeInDays = Uint(2)
		opts.ChangeTracking = Bool(true)
		opts.DefaultDDLCollation = String("default_ddl_collation")
		opts.CopyGrants = Bool(true)
		opts.Comment = String("comment")
		opts.RowAccessPolicy = &RowAccessPolicy{
			Name: NewSchemaObjectIdentifier("access_policy_database", "access_policy_schema", "access_policy_name"),
			On:   []string{"column1", "column2"},
		}
		opts.Tag = []TagAssociation{
			{
				Name:  NewSchemaObjectIdentifier("tag_database", "tag_schema", "tag_name"),
				Value: "tag_value",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE OR REPLACE EVENT TABLE %s CLUSTER BY (a, b) DATA_RETENTION_TIME_IN_DAYS = 1 MAX_DATA_EXTENSION_TIME_IN_DAYS = 2 CHANGE_TRACKING = true DEFAULT_DDL_COLLATION = 'default_ddl_collation' COPY_GRANTS COMMENT = 'comment' ROW ACCESS POLICY "access_policy_database"."access_policy_schema"."access_policy_name" ON (column1, column2) TAG ("tag_database"."tag_schema"."tag_name" = 'tag_value')`, id.FullyQualifiedName())
	})
}

func TestEventTableShow(t *testing.T) {
	id := randomSchemaObjectIdentifier(t)
	defaultOpts := func() *showEventTableOptions {
		return &showEventTableOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *showEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, errNilOptions)
	})

	t.Run("validation: empty like", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{}
		assertOptsInvalidJoinedErrors(t, opts, errPatternRequiredForLikeKeyword)
	})

	t.Run("show with in", func(t *testing.T) {
		opts := defaultOpts()
		opts.In = &In{
			Database: NewAccountObjectIdentifier("database"),
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW EVENT TABLES IN DATABASE "database"`)
	})

	t.Run("show with like", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{
			Pattern: String(id.Name()),
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW EVENT TABLES LIKE '%s'`, id.Name())
	})

	t.Run("show with like and in", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{
			Pattern: String(id.Name()),
		}
		opts.In = &In{
			Database: NewAccountObjectIdentifier("database"),
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW EVENT TABLES LIKE '%s' IN DATABASE "database"`, id.Name())
	})
}

func TestEventTableDescribe(t *testing.T) {
	id := randomSchemaObjectIdentifier(t)
	defaultOpts := func() *describeEventTableOptions {
		return &describeEventTableOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *describeEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, errNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, errInvalidObjectIdentifier)
	})

	t.Run("describe", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `DESCRIBE EVENT TABLE %s`, id.FullyQualifiedName())
	})
}
