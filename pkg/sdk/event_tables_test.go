package sdk

import (
	"errors"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
)

func TestEventTablesCreate(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *createEventTableOptions {
		return &createEventTableOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *createEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
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
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier, errOneOf("OrReplace", "IfNotExists"))
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
		opts.DataRetentionTimeInDays = Int(1)
		opts.MaxDataExtensionTimeInDays = Int(2)
		opts.ChangeTracking = Bool(true)
		opts.DefaultDDLCollation = String("default_ddl_collation")
		opts.CopyGrants = Bool(true)
		opts.Comment = String("comment")
		policyName := NewSchemaObjectIdentifier(random.StringN(8), random.StringN(8), random.StringN(8))
		opts.RowAccessPolicy = &RowAccessPolicy{
			Name: policyName,
			On:   []string{"column1", "column2"},
		}
		tagName := NewSchemaObjectIdentifier(random.StringN(8), random.StringN(8), random.StringN(8))
		opts.Tag = []TagAssociation{
			{
				Name:  tagName,
				Value: "tag_value",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE OR REPLACE EVENT TABLE %s CLUSTER BY (a, b) DATA_RETENTION_TIME_IN_DAYS = 1 MAX_DATA_EXTENSION_TIME_IN_DAYS = 2 CHANGE_TRACKING = true DEFAULT_DDL_COLLATION = 'default_ddl_collation' COPY GRANTS COMMENT = 'comment' ROW ACCESS POLICY %s ON (column1, column2) TAG (%s = 'tag_value')`, id.FullyQualifiedName(), policyName.FullyQualifiedName(), tagName.FullyQualifiedName())
	})
}

func TestEventTablesShow(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *showEventTableOptions {
		return &showEventTableOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *showEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: empty like", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{}
		assertOptsInvalidJoinedErrors(t, opts, ErrPatternRequiredForLikeKeyword)
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

func TestEventTablesDescribe(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *describeEventTableOptions {
		return &describeEventTableOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *describeEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("describe", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `DESCRIBE EVENT TABLE %s`, id.FullyQualifiedName())
	})
}

func TestEventTablesAlter(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *alterEventTableOptions {
		return &alterEventTableOptions{
			name: id,
		}
	}
	defaultTag := func(name SchemaObjectIdentifier) *[]TagAssociation {
		tag := []TagAssociation{
			{
				Name:  name,
				Value: "tag_value",
			},
		}
		return &tag
	}

	t.Run("rename to", func(t *testing.T) {
		opts := defaultOpts()
		opts.Rename = &RenameSchemaObjectIdentifier{Name: NewSchemaObjectIdentifier(id.DatabaseName(), id.SchemaName(), random.StringN(12))}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s RENAME TO %s`, id.FullyQualifiedName(), opts.Rename.Name.FullyQualifiedName())
	})

	t.Run("add row access policy", func(t *testing.T) {
		opts := defaultOpts()
		opts.AddRowAccessPolicy = &EventTableAddRowAccessPolicy{
			RowAccessPolicy: &RowAccessPolicy{
				Name: NewSchemaObjectIdentifier(random.StringN(8), random.StringN(8), random.StringN(8)),
				On:   []string{"column1", "column2"},
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s ADD ROW ACCESS POLICY %s ON (column1, column2)`,
			id.FullyQualifiedName(),
			opts.AddRowAccessPolicy.RowAccessPolicy.Name.FullyQualifiedName(),
		)
	})

	t.Run("drop row access policy", func(t *testing.T) {
		opts := defaultOpts()
		opts.DropRowAccessPolicy = &EventTableDropRowAccessPolicy{
			Name: NewSchemaObjectIdentifier(random.StringN(8), random.StringN(8), random.StringN(8)),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s DROP ROW ACCESS POLICY %s`,
			id.FullyQualifiedName(),
			opts.DropRowAccessPolicy.Name.FullyQualifiedName(),
		)
	})

	t.Run("drop all row access policies", func(t *testing.T) {
		opts := defaultOpts()
		opts.DropAllRowAccessPolicies = Bool(true)
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s DROP ALL ROW ACCESS POLICIES`, id.FullyQualifiedName())
	})

	t.Run("clustering action with cluster by", func(t *testing.T) {
		opts := defaultOpts()
		clusterBy := []string{"column1", "column2"}
		opts.ClusteringAction = &ClusteringAction{
			ClusterBy: &clusterBy,
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s CLUSTER BY (column1, column2)`, id.FullyQualifiedName())
	})

	t.Run("clustering action with suspend", func(t *testing.T) {
		opts := defaultOpts()
		opts.ClusteringAction = &ClusteringAction{
			SuspendRecluster: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s SUSPEND RECLUSTER`, id.FullyQualifiedName())
	})

	t.Run("clustering action with resume", func(t *testing.T) {
		opts := defaultOpts()
		opts.ClusteringAction = &ClusteringAction{
			ResumeRecluster: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s RESUME RECLUSTER`, id.FullyQualifiedName())
	})

	t.Run("clustering action with drop", func(t *testing.T) {
		opts := defaultOpts()
		opts.ClusteringAction = &ClusteringAction{
			DropClusteringKey: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s DROP CLUSTERING KEY`, id.FullyQualifiedName())
	})

	t.Run("search optimization action with add", func(t *testing.T) {
		opts := defaultOpts()
		opts.SearchOptimizationAction = &SearchOptimizationAction{
			Add: &AddSearchOptimization{
				On: []string{"column1", "column2"},
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s ADD SEARCH OPTIMIZATION ON (column1, column2)`, id.FullyQualifiedName())
	})

	t.Run("search optimization action with drop", func(t *testing.T) {
		opts := defaultOpts()
		opts.SearchOptimizationAction = &SearchOptimizationAction{
			Drop: &DropSearchOptimization{
				On: []string{"column1", "column2"},
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s DROP SEARCH OPTIMIZATION ON (column1, column2)`, id.FullyQualifiedName())
	})

	t.Run("set properties", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &EventTableSet{
			DataRetentionTimeInDays:    Int(1),
			MaxDataExtensionTimeInDays: Int(2),
			ChangeTracking:             Bool(true),
			Comment:                    String("comment"),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s SET DATA_RETENTION_TIME_IN_DAYS = 1 MAX_DATA_EXTENSION_TIME_IN_DAYS = 2 CHANGE_TRACKING = true COMMENT = 'comment'`, id.FullyQualifiedName())
	})

	t.Run("set tag", func(t *testing.T) {
		opts := defaultOpts()
		name := NewSchemaObjectIdentifier(random.String(), random.String(), random.String())
		opts.Set = &EventTableSet{
			Tag: defaultTag(name),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s SET TAG %s = 'tag_value'`, id.FullyQualifiedName(), name.FullyQualifiedName())
	})

	t.Run("unset data retention time in days", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &EventTableUnset{
			DataRetentionTimeInDays: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s UNSET DATA_RETENTION_TIME_IN_DAYS`, id.FullyQualifiedName())
	})

	t.Run("unset max data extension time in days", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &EventTableUnset{
			MaxDataExtensionTimeInDays: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s UNSET MAX_DATA_EXTENSION_TIME_IN_DAYS`, id.FullyQualifiedName())
	})

	t.Run("unset change tracking", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &EventTableUnset{
			ChangeTracking: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s UNSET CHANGE_TRACKING`, id.FullyQualifiedName())
	})

	t.Run("unset comment", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &EventTableUnset{
			Comment: Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s UNSET COMMENT`, id.FullyQualifiedName())
	})

	t.Run("unset tag", func(t *testing.T) {
		opts := defaultOpts()
		name := NewSchemaObjectIdentifier(random.String(), random.String(), random.String())
		tagNames := []string{name.FullyQualifiedName()}
		opts.Unset = &EventTableUnset{
			TagNames: &tagNames,
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER TABLE %s UNSET TAG %s`, id.FullyQualifiedName(), name.FullyQualifiedName())
	})

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *alterEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: no alter action", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsInvalidJoinedErrors(t, opts, errors.New("exactly one action of alterEventTableOptions must be set"))
	})

	t.Run("validation: multiple alter actions", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &EventTableSet{
			Tag: defaultTag(NewSchemaObjectIdentifier(random.String(), random.String(), random.String())),
		}
		name := NewSchemaObjectIdentifier(random.String(), random.String(), random.String())
		tagNames := []string{name.FullyQualifiedName()}
		opts.Unset = &EventTableUnset{
			TagNames: &tagNames,
		}
		assertOptsInvalidJoinedErrors(t, opts, errors.New("exactly one action of alterEventTableOptions must be set"))
	})

	t.Run("validation: invalid new name", func(t *testing.T) {
		opts := defaultOpts()
		opts.Rename = &RenameSchemaObjectIdentifier{
			Name: NewSchemaObjectIdentifier("", "", ""),
		}
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: no property to unset", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &EventTableUnset{}
		assertOptsInvalidJoinedErrors(t, opts, errAlterNeedsAtLeastOneProperty)
	})

	t.Run("validation: no property to set", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &EventTableSet{}
		assertOptsInvalidJoinedErrors(t, opts, errAlterNeedsAtLeastOneProperty)
	})

	t.Run("validation: invalid add row access policy name", func(t *testing.T) {
		opts := defaultOpts()
		opts.AddRowAccessPolicy = &EventTableAddRowAccessPolicy{
			RowAccessPolicy: &RowAccessPolicy{
				Name: NewSchemaObjectIdentifier("", "", ""),
			},
		}
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: invalid drop row access policy name", func(t *testing.T) {
		opts := defaultOpts()
		opts.DropRowAccessPolicy = &EventTableDropRowAccessPolicy{
			Name: NewSchemaObjectIdentifier("", "", ""),
		}
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: search optimization action with both add and drop", func(t *testing.T) {
		opts := defaultOpts()
		opts.SearchOptimizationAction = &SearchOptimizationAction{
			Add: &AddSearchOptimization{
				On: []string{"column1", "column2"},
			},
			Drop: &DropSearchOptimization{
				On: []string{"column1", "column2"},
			},
		}
		assertOptsInvalidJoinedErrors(t, opts, errors.New("exactly one action of SearchOptimizationAction must be set"))
	})

	t.Run("validation: search optimization action without property", func(t *testing.T) {
		opts := defaultOpts()
		opts.SearchOptimizationAction = &SearchOptimizationAction{}
		assertOptsInvalidJoinedErrors(t, opts, errors.New("exactly one action of SearchOptimizationAction must be set"))
	})

	t.Run("validation: clustering action with both resume and suspend", func(t *testing.T) {
		opts := defaultOpts()
		opts.ClusteringAction = &ClusteringAction{
			ResumeRecluster:  Bool(true),
			SuspendRecluster: Bool(true),
		}
		assertOptsInvalidJoinedErrors(t, opts, errors.New("exactly one action of ClusteringAction must be set"))
	})

	t.Run("validation: clustering action without property", func(t *testing.T) {
		opts := defaultOpts()
		opts.ClusteringAction = &ClusteringAction{}
		assertOptsInvalidJoinedErrors(t, opts, errors.New("exactly one action of ClusteringAction must be set"))
	})
}
