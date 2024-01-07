package sdk

var (
	_ validatable = new(CreateApplicationOptions)
)

func (opts *CreateApplicationOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if valueSet(opts.ApplicationVersion) {
		if !exactlyOneValueSet(opts.ApplicationVersion.VersionDirectory, opts.ApplicationVersion.VersionAndPatch) {
			errs = append(errs, errExactlyOneOf("CreateApplicationOptions.ApplicationVersion", "VersionDirectory", "VersionAndPatch"))
		}
	}
	return JoinErrors(errs...)
}
