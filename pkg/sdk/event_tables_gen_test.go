package sdk

import "testing"

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
		assertOptsValidAndSQLEquals(t, opts, `CREATE OR REPLACE EVENT TABLE IF NOT EXISTS %s CLUSTER BY (a, b)`, id.FullyQualifiedName())
	})
}
