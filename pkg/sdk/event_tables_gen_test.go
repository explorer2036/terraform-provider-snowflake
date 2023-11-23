package sdk

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
)

func TestEventTables_Create(t *testing.T) {
	id := RandomSchemaObjectIdentifier()

	defaultOpts := func() *CreateEventTableOptions {
		return &CreateEventTableOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.OrReplace = Bool(true)
		opts.IfNotExists = Bool(true)
		opts.ClusterBy = []string{"a", "b"}
		opts.DataRetentionTimeInDays = Int(1)
		opts.MaxDataExtensionTimeInDays = Int(2)
		opts.ChangeTracking = Bool(true)
		opts.DefaultDdlCollation = String("en_US")
		opts.CopyGrants = Bool(true)
		opts.Comment = String("test")
		pn := NewSchemaObjectIdentifier(random.StringN(4), random.StringN(4), random.StringN(4))
		opts.RowAccessPolicy = &RowAccessPolicy{
			Name: pn,
			On:   []string{"c1", "c2"},
		}
		tn := NewSchemaObjectIdentifier(random.StringN(4), random.StringN(4), random.StringN(4))
		opts.Tag = []TagAssociation{
			{
				Name:  tn,
				Value: "v1",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE OR REPLACE EVENT TABLE IF NOT EXISTS %s CLUSTER BY (a, b) DATA_RETENTION_TIME_IN_DAYS = 1 MAX_DATA_EXTENSION_TIME_IN_DAYS = 2 CHANGE_TRACKING = true DEFAULT_DDL_COLLATION = 'en_US' COPY GRANTS COMMENT = 'test' ROW ACCESS POLICY %s ON (c1, c2) TAG (%s = 'v1')`, id.FullyQualifiedName(), pn.FullyQualifiedName(), tn.FullyQualifiedName())
	})
}

func TestEventTables_Show(t *testing.T) {
	id := RandomSchemaObjectIdentifier()
	defaultOpts := func() *ShowEventTableOptions {
		return &ShowEventTableOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *ShowEventTableOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
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
