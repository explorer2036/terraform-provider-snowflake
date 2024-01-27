package sdk

var (
	_ validatable = new(CreateReplicationGroupOptions)
	_ validatable = new(CreateSecondaryReplicationGroupOptions)
	_ validatable = new(AlterReplicationGroupOptions)
	_ validatable = new(ShowReplicationGroupOptions)
	_ validatable = new(DropReplicationGroupOptions)
)

func (opts *CreateReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !valueSet(opts.ObjectTypes) {
		errs = append(errs, errNotSet("CreateReplicationGroupOptions", "ObjectTypes"))
	}
	if !valueSet(opts.AllowedAccounts) {
		errs = append(errs, errNotSet("CreateReplicationGroupOptions", "AllowedAccounts"))
	}
	if valueSet(opts.ObjectTypes) {
		if !anyValueSet(opts.ObjectTypes.AccountParameters, opts.ObjectTypes.Databases, opts.ObjectTypes.Integrations, opts.ObjectTypes.NetworkPolicies, opts.ObjectTypes.ResourceMonitors, opts.ObjectTypes.Roles, opts.ObjectTypes.Shares, opts.ObjectTypes.Users, opts.ObjectTypes.Warehouses) {
			errs = append(errs, errAtLeastOneOf("CreateReplicationGroupOptions.ObjectTypes", "AccountParameters", "Databases", "Integrations", "NetworkPolicies", "ResourceMonitors", "Roles", "Shares", "Users", "Warehouses"))
		}
	}
	if valueSet(opts.ReplicationSchedule) {
		if !exactlyOneValueSet(opts.ReplicationSchedule.IntervalMinutes, opts.ReplicationSchedule.CronExpression) {
			errs = append(errs, errExactlyOneOf("CreateReplicationGroupOptions.ReplicationSchedule", "IntervalMinutes", "CronExpression"))
		}
	}
	return JoinErrors(errs...)
}

func (opts *CreateSecondaryReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !ValidObjectIdentifier(opts.Primary) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *AlterReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if opts.RenameTo != nil && !ValidObjectIdentifier(opts.RenameTo) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !exactlyOneValueSet(opts.RenameTo, opts.Set, opts.SetIntegration, opts.AddDatabases, opts.RemoveDatabases, opts.MoveDatabases, opts.AddShares, opts.RemoveShares, opts.MoveShares, opts.AddAccounts, opts.RemoveAccounts, opts.Refresh, opts.Suspend, opts.Resume) {
		errs = append(errs, errExactlyOneOf("AlterReplicationGroupOptions", "RenameTo", "Set", "SetIntegration", "AddDatabases", "RemoveDatabases", "MoveDatabases", "AddShares", "RemoveShares", "MoveShares", "AddAccounts", "RemoveAccounts", "Refresh", "Suspend", "Resume"))
	}
	if valueSet(opts.Set) {
		if valueSet(opts.Set.ObjectTypes) {
			if !anyValueSet(opts.Set.ObjectTypes.AccountParameters, opts.Set.ObjectTypes.Databases, opts.Set.ObjectTypes.Integrations, opts.Set.ObjectTypes.NetworkPolicies, opts.Set.ObjectTypes.ResourceMonitors, opts.Set.ObjectTypes.Roles, opts.Set.ObjectTypes.Shares, opts.Set.ObjectTypes.Users, opts.Set.ObjectTypes.Warehouses) {
				errs = append(errs, errAtLeastOneOf("AlterReplicationGroupOptions.Set.ObjectTypes", "AccountParameters", "Databases", "Integrations", "NetworkPolicies", "ResourceMonitors", "Roles", "Shares", "Users", "Warehouses"))
			}
		}
		if valueSet(opts.Set.ReplicationSchedule) {
			if !exactlyOneValueSet(opts.Set.ReplicationSchedule.IntervalMinutes, opts.Set.ReplicationSchedule.CronExpression) {
				errs = append(errs, errExactlyOneOf("AlterReplicationGroupOptions.Set.ReplicationSchedule", "IntervalMinutes", "CronExpression"))
			}
		}
	}
	if valueSet(opts.SetIntegration) {
		if valueSet(opts.SetIntegration.ObjectTypes) {
			if !anyValueSet(opts.SetIntegration.ObjectTypes.AccountParameters, opts.SetIntegration.ObjectTypes.Databases, opts.SetIntegration.ObjectTypes.Integrations, opts.SetIntegration.ObjectTypes.NetworkPolicies, opts.SetIntegration.ObjectTypes.ResourceMonitors, opts.SetIntegration.ObjectTypes.Roles, opts.SetIntegration.ObjectTypes.Shares, opts.SetIntegration.ObjectTypes.Users, opts.SetIntegration.ObjectTypes.Warehouses) {
				errs = append(errs, errAtLeastOneOf("AlterReplicationGroupOptions.SetIntegration.ObjectTypes", "AccountParameters", "Databases", "Integrations", "NetworkPolicies", "ResourceMonitors", "Roles", "Shares", "Users", "Warehouses"))
			}
		}
		if valueSet(opts.SetIntegration.ReplicationSchedule) {
			if !exactlyOneValueSet(opts.SetIntegration.ReplicationSchedule.IntervalMinutes, opts.SetIntegration.ReplicationSchedule.CronExpression) {
				errs = append(errs, errExactlyOneOf("AlterReplicationGroupOptions.SetIntegration.ReplicationSchedule", "IntervalMinutes", "CronExpression"))
			}
		}
	}
	return JoinErrors(errs...)
}

func (opts *ShowReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *ShowDatabasesInReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *ShowSharesInReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *DropReplicationGroupOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}
