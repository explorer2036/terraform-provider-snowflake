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
		if !exactlyOneValueSet(opts.Add.AddToAllowedNetworkRuleList, opts.Add.AddToBlockedNetworkRuleList) {
			errs = append(errs, errExactlyOneOf("AlterNetworkPolicyOptions.Add", "AddToAllowedNetworkRuleList", "AddToBlockedNetworkRuleList"))
		}
		if opts.Add.AddToAllowedNetworkRuleList != nil && !ValidObjectIdentifier(opts.Add.AddToAllowedNetworkRuleList) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
		if opts.Add.AddToBlockedNetworkRuleList != nil && !ValidObjectIdentifier(opts.Add.AddToBlockedNetworkRuleList) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
	}
	if valueSet(opts.Remove) {
		if !exactlyOneValueSet(opts.Remove.RemoveFromAllowedNetworkRuleList, opts.Remove.RemoveFromBlockedNetworkRuleList) {
			errs = append(errs, errExactlyOneOf("AlterNetworkPolicyOptions.Remove", "RemoveFromAllowedNetworkRuleList", "RemoveFromBlockedNetworkRuleList"))
		}
		if opts.Remove.RemoveFromAllowedNetworkRuleList != nil && !ValidObjectIdentifier(opts.Remove.RemoveFromAllowedNetworkRuleList) {
			errs = append(errs, ErrInvalidObjectIdentifier)
		}
		if opts.Remove.RemoveFromBlockedNetworkRuleList != nil && !ValidObjectIdentifier(opts.Remove.RemoveFromBlockedNetworkRuleList) {
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
