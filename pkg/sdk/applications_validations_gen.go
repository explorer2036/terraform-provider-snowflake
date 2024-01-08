package sdk

var (
	_ validatable = new(CreateApplicationOptions)
	_ validatable = new(DropApplicationOptions)
	_ validatable = new(ShowApplicationOptions)
	_ validatable = new(DescribeApplicationOptions)
)

func (opts *CreateApplicationOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !ValidObjectIdentifier(opts.PackageName) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if valueSet(opts.Version) {
		if !exactlyOneValueSet(opts.Version.VersionDirectory, opts.Version.VersionAndPatch) {
			errs = append(errs, errExactlyOneOf("CreateApplicationOptions.Version", "VersionDirectory", "VersionAndPatch"))
		}
	}
	return JoinErrors(errs...)
}

func (opts *DropApplicationOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *ShowApplicationOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *DescribeApplicationOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}
