package sdk

var (
	_ validatable = new(CreateForJavaFunctionOptions)
	_ validatable = new(CreateForJavascriptFunctionOptions)
	_ validatable = new(CreateForPythonFunctionOptions)
	_ validatable = new(CreateForScalaFunctionOptions)
	_ validatable = new(CreateForSQLFunctionOptions)
	_ validatable = new(AlterFunctionOptions)
	_ validatable = new(DropFunctionOptions)
	_ validatable = new(ShowFunctionOptions)
	_ validatable = new(DescribeFunctionOptions)
)

func (opts *CreateForJavaFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if everyValueSet(opts.OrReplace, opts.IfNotExists) {
		errs = append(errs, errOneOf("CreateForJavaFunctionOptions", "OrReplace", "IfNotExists"))
	}
	return JoinErrors(errs...)
}

func (opts *CreateForJavascriptFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *CreateForPythonFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if everyValueSet(opts.OrReplace, opts.IfNotExists) {
		errs = append(errs, errOneOf("CreateForPythonFunctionOptions", "OrReplace", "IfNotExists"))
	}
	return JoinErrors(errs...)
}

func (opts *CreateForScalaFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if everyValueSet(opts.OrReplace, opts.IfNotExists) {
		errs = append(errs, errOneOf("CreateForScalaFunctionOptions", "OrReplace", "IfNotExists"))
	}
	return JoinErrors(errs...)
}

func (opts *CreateForSQLFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *AlterFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !exactlyOneValueSet(opts.RenameTo, opts.SetComment, opts.SetLogLevel, opts.SetTraceLevel, opts.SetSecure, opts.UnsetLogLevel, opts.UnsetTraceLevel, opts.UnsetSecure, opts.UnsetComment, opts.SetTags, opts.UnsetTags) {
		errs = append(errs, errExactlyOneOf("AlterFunctionOptions", "RenameTo", "SetComment", "SetLogLevel", "SetTraceLevel", "SetSecure", "UnsetLogLevel", "UnsetTraceLevel", "UnsetSecure", "UnsetComment", "SetTags", "UnsetTags"))
	}
	return JoinErrors(errs...)
}

func (opts *DropFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *ShowFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *DescribeFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}
