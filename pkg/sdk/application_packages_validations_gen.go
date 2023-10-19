package sdk

import "errors"

var (
	_ validatable = new(CreateApplicationPackageOptions)
	_ validatable = new(AlterApplicationPackageOptions)
	_ validatable = new(DropApplicationPackageOptions)
	_ validatable = new(ShowApplicationPackageOptions)
	_ validatable = (*ApplicationPackageSet)(nil)
	_ validatable = (*ApplicationPackageUnset)(nil)
)

func (opts *CreateApplicationPackageOptions) validate() error {
	if opts == nil {
		return errors.Join(ErrNilOptions)
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if valueSet(opts.DataRetentionTimeInDays) {
		if !validateIntGreaterThanOrEqual(*opts.DataRetentionTimeInDays, 0) {
			errs = append(errs, errors.New("DataRetentionTimeInDays must be greater than or equal to 0"))
		}
	}
	if valueSet(opts.MaxDataExtensionTimeInDays) {
		if !validateIntInRange(*opts.MaxDataExtensionTimeInDays, 0, 90) {
			errs = append(errs, errors.New("MaxDataExtensionTimeInDays must be between 0 and 90"))
		}
	}
	return errors.Join(errs...)
}

func (v *ApplicationPackageSet) validate() error {
	if v == nil {
		return nil
	}
	var errs []error
	if ok := anyValueSet(
		v.DataRetentionTimeInDays,
		v.MaxDataExtensionTimeInDays,
		v.DefaultDdlCollation,
		v.Comment,
		v.Distribution,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	if valueSet(v.DataRetentionTimeInDays) {
		if !validateIntGreaterThanOrEqual(*v.DataRetentionTimeInDays, 0) {
			errs = append(errs, errors.New("DataRetentionTimeInDays must be greater than or equal to 0"))
		}
	}
	if valueSet(v.MaxDataExtensionTimeInDays) {
		if !validateIntInRange(*v.MaxDataExtensionTimeInDays, 0, 90) {
			errs = append(errs, errors.New("MaxDataExtensionTimeInDays must be between 0 and 90"))
		}
	}
	return errors.Join(errs...)
}

func (v *ApplicationPackageUnset) validate() error {
	if v == nil {
		return nil
	}
	var errs []error
	if ok := anyValueSet(
		v.DataRetentionTimeInDays,
		v.MaxDataExtensionTimeInDays,
		v.DefaultDdlCollation,
		v.Comment,
		v.Distribution,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	return errors.Join(errs...)
}

func (opts *AlterApplicationPackageOptions) validate() error {
	if opts == nil {
		return errors.Join(ErrNilOptions)
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if ok := anyValueSet(opts.Set, opts.Unset); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	if opts.Set != nil {
		if err := opts.Set.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if opts.Unset != nil {
		if err := opts.Unset.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (opts *DropApplicationPackageOptions) validate() error {
	if opts == nil {
		return errors.Join(ErrNilOptions)
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return errors.Join(errs...)
}

func (opts *ShowApplicationPackageOptions) validate() error {
	if opts == nil {
		return errors.Join(ErrNilOptions)
	}
	var errs []error
	return errors.Join(errs...)
}
