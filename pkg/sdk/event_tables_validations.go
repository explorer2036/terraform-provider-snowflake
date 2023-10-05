package sdk

import (
	"errors"
)

var (
	_ validatable = (*createEventTableOptions)(nil)
	_ validatable = (*alterEventTableOptions)(nil)
	_ validatable = (*showEventTableOptions)(nil)
	_ validatable = (*describeEventTableOptions)(nil)
	_ validatable = (*ClusteringAction)(nil)
	_ validatable = (*SearchOptimizationAction)(nil)
	_ validatable = (*EventTableSet)(nil)
	_ validatable = (*EventTableUnset)(nil)
	_ validatable = (*EventTableAddRowAccessPolicy)(nil)
	_ validatable = (*EventTableDropRowAccessPolicy)(nil)
)

func (opts *createEventTableOptions) validate() error {
	if opts == nil {
		return errors.Join(errNilOptions)
	}
	var errs []error
	if !validObjectidentifier(opts.name) {
		errs = append(errs, errInvalidObjectIdentifier)
	}
	if everyValueSet(opts.OrReplace, opts.IfNotExists) && *opts.OrReplace && *opts.IfNotExists {
		errs = append(errs, errOneOf("OrReplace", "IfNotExists"))
	}
	if valueSet(opts.CopyGrants) && !valueSet(opts.OrReplace) {
		errs = append(errs, errors.New("CopyGrants requires OrReplace"))
	}
	return errors.Join(errs...)
}

func (v *ClusteringAction) validate() error {
	var errs []error
	if ok := exactlyOneValueSet(
		v.ClusterBy,
		v.Suspend,
		v.Resume,
		v.Drop,
	); !ok {
		errs = append(errs, errAlterNeedsExactlyOneAction)
	}
	if ok := anyValueSet(
		v.ClusterBy,
		v.Suspend,
		v.Resume,
		v.Drop,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	return errors.Join(errs...)
}

func (v *SearchOptimizationAction) validate() error {
	var errs []error
	if ok := exactlyOneValueSet(
		v.Add,
		v.Drop,
	); !ok {
		errs = append(errs, errAlterNeedsExactlyOneAction)
	}
	if ok := anyValueSet(
		v.Add,
		v.Drop,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	return errors.Join(errs...)
}

func (v *EventTableAddRowAccessPolicy) validate() error {
	var errs []error
	if valueSet(v.RowAccessPolicy) {
		if !validObjectidentifier(v.RowAccessPolicy.Name) {
			errs = append(errs, errInvalidObjectIdentifier)
		}
	}
	return errors.Join(errs...)
}

func (v *EventTableDropRowAccessPolicy) validate() error {
	var errs []error
	if !validObjectidentifier(v.Name) {
		errs = append(errs, errInvalidObjectIdentifier)
	}
	return errors.Join(errs...)
}

func (v *EventTableSet) validate() error {
	var errs []error
	if ok := exactlyOneValueSet(
		v.Properties,
		v.Tag,
	); !ok {
		errs = append(errs, errAlterNeedsExactlyOneAction)
	}
	if ok := anyValueSet(
		v.Properties,
		v.Tag,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	return errors.Join(errs...)
}

func (v *EventTableUnset) validate() error {
	var errs []error
	if ok := exactlyOneValueSet(
		v.DataRetentionTimeInDays,
		v.MaxDataExtensionTimeInDays,
		v.ChangeTracking,
		v.Comment,
		v.Tag,
	); !ok {
		errs = append(errs, errAlterNeedsExactlyOneAction)
	}
	if ok := anyValueSet(
		v.DataRetentionTimeInDays,
		v.MaxDataExtensionTimeInDays,
		v.ChangeTracking,
		v.Comment,
		v.Tag,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	return errors.Join(errs...)
}

func (opts *alterEventTableOptions) validate() error {
	if opts == nil {
		return errors.Join(errNilOptions)
	}
	var errs []error
	if !validObjectidentifier(opts.name) {
		errs = append(errs, errInvalidObjectIdentifier)
	}
	if ok := exactlyOneValueSet(
		opts.ClusteringAction,
		opts.SearchOptimizationAction,
		opts.AddRowAccessPolicy,
		opts.DropAllRowAccessPolicies,
		opts.DropRowAccessPolicy,
		opts.Set,
		opts.Unset,
		opts.Rename,
	); !ok {
		errs = append(errs, errAlterNeedsExactlyOneAction)
	}
	if ok := anyValueSet(
		opts.ClusteringAction,
		opts.SearchOptimizationAction,
		opts.AddRowAccessPolicy,
		opts.DropAllRowAccessPolicies,
		opts.DropRowAccessPolicy,
		opts.Set,
		opts.Unset,
		opts.Rename,
	); !ok {
		errs = append(errs, errAlterNeedsAtLeastOneProperty)
	}
	if valueSet(opts.ClusteringAction) {
		if err := opts.ClusteringAction.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if valueSet(opts.SearchOptimizationAction) {
		if err := opts.SearchOptimizationAction.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if valueSet(opts.AddRowAccessPolicy) {
		if err := opts.AddRowAccessPolicy.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if valueSet(opts.DropRowAccessPolicy) {
		if err := opts.DropRowAccessPolicy.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if valueSet(opts.Set) {
		if err := opts.Set.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if valueSet(opts.Unset) {
		if err := opts.Unset.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if valueSet(opts.Rename) {
		if !validObjectidentifier(opts.Rename.Name) {
			errs = append(errs, errInvalidObjectIdentifier)
		}
		if opts.name.DatabaseName() != opts.Rename.Name.DatabaseName() {
			errs = append(errs, errDifferentDatabase)
		}
	}
	return errors.Join(errs...)
}

func (opts *describeEventTableOptions) validate() error {
	if opts == nil {
		return errors.Join(errNilOptions)
	}
	var errs []error
	if !validObjectidentifier(opts.name) {
		errs = append(errs, errInvalidObjectIdentifier)
	}
	return errors.Join(errs...)
}

func (opts *showEventTableOptions) validate() error {
	if opts == nil {
		return errors.Join(errNilOptions)
	}
	var errs []error
	if valueSet(opts.Like) && !valueSet(opts.Like.Pattern) {
		errs = append(errs, errPatternRequiredForLikeKeyword)
	}
	if valueSet(opts.In) && !exactlyOneValueSet(opts.In.Account, opts.In.Database, opts.In.Schema) {
		errs = append(errs, errScopeRequiredForInKeyword)
	}
	return errors.Join(errs...)
}
