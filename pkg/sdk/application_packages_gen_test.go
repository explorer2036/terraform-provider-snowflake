package sdk

import "testing"

func TestApplicationPackages_Create(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid CreateApplicationPackageOptions
	defaultOpts := func() *CreateApplicationPackageOptions {
		return &CreateApplicationPackageOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateApplicationPackageOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `CREATE APPLICATION PACKAGE %s`, id.FullyQualifiedName())
	})

	t.Run("distribution", func(t *testing.T) {
		opts := defaultOpts()
		opts.Distribution = String("INTERNAL")
		assertOptsValidAndSQLEquals(t, opts, `CREATE APPLICATION PACKAGE %s DISTRIBUTION = INTERNAL`, id.FullyQualifiedName())
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.DataRetentionTimeInDays = Int(1)
		opts.MaxDataExtensionTimeInDays = Int(1)
		opts.DefaultDdlCollation = String("default_ddl_collation")
		opts.Comment = String("comment")
		tagName := RandomSchemaObjectIdentifier()
		opts.Tag = []TagAssociation{
			{
				Name:  tagName,
				Value: "tag_value",
			},
		}
		opts.Distribution = String("INTERNAL")
		assertOptsValidAndSQLEquals(t, opts, `CREATE APPLICATION PACKAGE IF NOT EXISTS %s DATA_RETENTION_TIME_IN_DAYS = 1 MAX_DATA_EXTENSION_TIME_IN_DAYS = 1 DEFAULT_DDL_COLLATION = 'default_ddl_collation' COMMENT = 'comment' TAG (%s = 'tag_value') DISTRIBUTION = INTERNAL`, id.FullyQualifiedName(), tagName.FullyQualifiedName())
	})
}

func TestApplicationPackages_Alter(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid AlterApplicationPackageOptions
	defaultOpts := func() *AlterApplicationPackageOptions {
		return &AlterApplicationPackageOptions{
			name:     id,
			IfExists: Bool(true),
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *AlterApplicationPackageOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("alter: set options", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &ApplicationPackageSet{
			DataRetentionTimeInDays:    Int(1),
			MaxDataExtensionTimeInDays: Int(1),
			DefaultDdlCollation:        String("default_ddl_collation"),
			Comment:                    String("comment"),
			Distribution:               String("INTERNAL"),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER APPLICATION PACKAGE IF EXISTS %s SET DATA_RETENTION_TIME_IN_DAYS = 1 MAX_DATA_EXTENSION_TIME_IN_DAYS = 1 DEFAULT_DDL_COLLATION = 'default_ddl_collation' COMMENT = 'comment' DISTRIBUTION = INTERNAL`, id.FullyQualifiedName())
	})

	t.Run("alter: unset options", func(t *testing.T) {
		opts := defaultOpts()
		opts.Unset = &ApplicationPackageUnset{
			DataRetentionTimeInDays:    Bool(true),
			MaxDataExtensionTimeInDays: Bool(true),
			DefaultDdlCollation:        Bool(true),
			Comment:                    Bool(true),
			Distribution:               Bool(true),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER APPLICATION PACKAGE IF EXISTS %s UNSET DATA_RETENTION_TIME_IN_DAYS MAX_DATA_EXTENSION_TIME_IN_DAYS DEFAULT_DDL_COLLATION COMMENT DISTRIBUTION`, id.FullyQualifiedName())
	})
}

func TestApplicationPackages_Drop(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid DropApplicationPackageOptions
	defaultOpts := func() *DropApplicationPackageOptions {
		return &DropApplicationPackageOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *DropApplicationPackageOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `DROP APPLICATION PACKAGE %s`, id.FullyQualifiedName())
	})
}

func TestApplicationPackages_Show(t *testing.T) {
	// Minimal valid ShowApplicationPackageOptions
	defaultOpts := func() *ShowApplicationPackageOptions {
		return &ShowApplicationPackageOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *ShowApplicationPackageOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, `SHOW APPLICATION PACKAGES`)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{
			Pattern: String("pattern"),
		}
		opts.StartsWith = String("A")
		opts.Limit = &LimitFrom{
			Rows: Int(1),
			From: String("B"),
		}
		assertOptsValidAndSQLEquals(t, opts, `SHOW APPLICATION PACKAGES LIKE 'pattern' STARTS WITH 'A' LIMIT 1 FROM 'B'`)
	})
}
