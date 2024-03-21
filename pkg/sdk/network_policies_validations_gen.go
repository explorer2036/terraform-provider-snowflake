package sdk

var (
	_ validatable = new(CreateNetworkPolicyOptions)
	_ validatable = new(AlterNetworkPolicyOptions)
	_ validatable = new(DropNetworkPolicyOptions)
	_ validatable = new(ShowNetworkPolicyOptions)
	_ validatable = new(DescribeNetworkPolicyOptions)
)

func (opts *CreateNetworkPolicyOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *AlterNetworkPolicyOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !exactlyOneValueSet(opts.Set, opts.UnsetComment, opts.RenameTo, opts.Add, opts.Remove) {
		errs = append(errs, errExactlyOneOf("AlterNetworkPolicyOptions", "Set", "UnsetComment", "RenameTo", "Add", "Remove"))
	}
	if opts.RenameTo != nil && !ValidObjectIdentifier(opts.RenameTo) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if valueSet(opts.Set) {
		if !anyValueSet(opts.Set.AllowedIpList, opts.Set.BlockedIpList, opts.Set.Comment, opts.Set.AllowedNetworkRuleList, opts.Set.BlockedNetworkRuleList) {
			errs = append(errs, errAtLeastOneOf("AlterNetworkPolicyOptions.Set", "AllowedIpList", "BlockedIpList", "Comment", "AllowedNetworkRuleList", "BlockedNetworkRuleList"))
		}
	}
	if valueSet(opts.Add) {
		if !exactlyOneValueSet(opts.Add.AddAllowedNetworkRule, opts.Add.AddBlockedNetworkRule) {
			errs = append(errs, errExactlyOneOf("AlterNetworkPolicyOptions.Add", "AddAllowedNetworkRule", "AddBlockedNetworkRule"))
		}
		if opts.Add.AddAllowedNetworkRule != nil && !ValidObjectIdentifier(opts.Add.AddAllowedNetworkRule) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
		if opts.Add.AddBlockedNetworkRule != nil && !ValidObjectIdentifier(opts.Add.AddBlockedNetworkRule) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
	}
	if valueSet(opts.Remove) {
		if !exactlyOneValueSet(opts.Remove.RemoveAllowedNetworkRule, opts.Remove.RemoveBlockedNetworkRule) {
			errs = append(errs, errExactlyOneOf("AlterNetworkPolicyOptions.Remove", "RemoveAllowedNetworkRule", "RemoveBlockedNetworkRule"))
		}
		if opts.Remove.RemoveAllowedNetworkRule != nil && !ValidObjectIdentifier(opts.Remove.RemoveAllowedNetworkRule) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
		if opts.Remove.RemoveBlockedNetworkRule != nil && !ValidObjectIdentifier(opts.Remove.RemoveBlockedNetworkRule) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
	}
	return JoinErrors(errs...)
}

func (opts *DropNetworkPolicyOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *ShowNetworkPolicyOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *DescribeNetworkPolicyOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}
