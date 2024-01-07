package sdk

import "testing"

func TestApplications_Create(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	defaultOpts := func() *CreateApplicationOptions {
		return &CreateApplicationOptions{
			name: id,
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
		opts.ApplicationVersion = &ApplicationVersion{
			VersionAndPatch: &VersionAndPatch{
				Version: "1.0",
				Patch:   1,
			},
			VersionDirectory: String("@test"),
		}
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("CreateApplicationOptions.ApplicationVersion", "VersionDirectory", "VersionAndPatch"))
	})

	// t.Run("all options", func(t *testing.T) {
	// 	opts := defaultOpts()
	// 	assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	// })
}
