package sdk

var (
	_ validatable = new(CreateExternalFunctionOptions)
	_ validatable = new(AlterExternalFunctionOptions)
	_ validatable = new(ShowExternalFunctionOptions)
)

func (opts *CreateExternalFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *AlterExternalFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !exactlyOneValueSet(opts.Set, opts.Unset) {
		errs = append(errs, errExactlyOneOf("AlterExternalFunctionOptions", "Set", "Unset"))
	}
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if valueSet(opts.Set) {
		if !exactlyOneValueSet(opts.Set.ApiIntegration, opts.Set.Headers, opts.Set.ContextHeaders, opts.Set.MaxBatchRows, opts.Set.Compression, opts.Set.RequestTranslator, opts.Set.ResponseTranslator) {
			errs = append(errs, errExactlyOneOf("AlterExternalFunctionOptions.Set", "ApiIntegration", "Headers", "ContextHeaders", "MaxBatchRows", "Compression", "RequestTranslator", "ResponseTranslator"))
		}
	}
	return JoinErrors(errs...)
}

func (opts *ShowExternalFunctionOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}
