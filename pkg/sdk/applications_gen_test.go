package sdk

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
)

func TestApplications_Create(t *testing.T) {
	id := RandomAccountObjectIdentifier()
	pid := RandomAccountObjectIdentifier()

	defaultOpts := func() *CreateApplicationOptions {
		return &CreateApplicationOptions{
			name:        id,
			PackageName: pid,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateApplicationOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: exactly one field should be present", func(t *testing.T) {
		opts := defaultOpts()
		opts.Version = &ApplicationVersion{
			VersionAndPatch: &VersionAndPatch{
				Version: "1.0",
				Patch:   Int(1),
			},
			VersionDirectory: String("@test"),
		}
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("CreateApplicationOptions.Version", "VersionDirectory", "VersionAndPatch"))
	})

	t.Run("all options", func(t *testing.T) {
		tid := NewSchemaObjectIdentifier(random.StringN(4), random.StringN(4), random.StringN(4))

		opts := defaultOpts()
		opts.Comment = String("test")
		opts.Tag = []TagAssociation{
			{
				Name:  tid,
				Value: "v1",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE APPLICATION %s FROM APPLICATION PACKAGE %s COMMENT = 'test' TAG (%s = 'v1')`, id.FullyQualifiedName(), pid.FullyQualifiedName(), tid.FullyQualifiedName())

		opts = defaultOpts()
		opts.Comment = String("test")
		opts.Version = &ApplicationVersion{
			VersionDirectory: String("@test"),
		}
		opts.DebugMode = Bool(true)
		opts.Tag = []TagAssociation{
			{
				Name:  tid,
				Value: "v1",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE APPLICATION %s FROM APPLICATION PACKAGE %s USING '@test' DEBUG_MODE = true COMMENT = 'test' TAG (%s = 'v1')`, id.FullyQualifiedName(), pid.FullyQualifiedName(), tid.FullyQualifiedName())

		opts = defaultOpts()
		opts.Comment = String("test")
		opts.Version = &ApplicationVersion{
			VersionAndPatch: &VersionAndPatch{
				Version: "V001",
				Patch:   Int(1),
			},
		}
		opts.DebugMode = Bool(true)
		opts.Tag = []TagAssociation{
			{
				Name:  tid,
				Value: "v1",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `CREATE APPLICATION %s FROM APPLICATION PACKAGE %s USING VERSION V001 PATCH 1 DEBUG_MODE = true COMMENT = 'test' TAG (%s = 'v1')`, id.FullyQualifiedName(), pid.FullyQualifiedName(), tid.FullyQualifiedName())
	})
}
