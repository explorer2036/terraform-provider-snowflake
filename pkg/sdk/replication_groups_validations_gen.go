package sdk

var (
	_ validatable = new(CreateReplicationGroupOptions)
)

func (opts *CreateReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if valueSet(opts.ObjectTypes) {
		if !anyValueSet(opts.ObjectTypes.AccountParameters, opts.ObjectTypes.Databases, opts.ObjectTypes.Integrations, opts.ObjectTypes.NetworkPolicies, opts.ObjectTypes.ResourceMonitors, opts.ObjectTypes.Roles, opts.ObjectTypes.Shares, opts.ObjectTypes.Users, opts.ObjectTypes.Warehouses) {
			errs = append(errs, errAtLeastOneOf("CreateReplicationGroupOptions.ObjectTypes", "AccountParameters", "Databases", "Integrations", "NetworkPolicies", "ResourceMonitors", "Roles", "Shares", "Users", "Warehouses"))
		}
	}
	return JoinErrors(errs...)
}
