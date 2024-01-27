package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ ReplicationGroups = (*replicationGroups)(nil)

type replicationGroups struct {
	client *Client
}

func (v *replicationGroups) Create(ctx context.Context, request *CreateReplicationGroupRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *replicationGroups) CreateSecondary(ctx context.Context, request *CreateSecondaryReplicationGroupRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *replicationGroups) Alter(ctx context.Context, request *AlterReplicationGroupRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *replicationGroups) Show(ctx context.Context, request *ShowReplicationGroupRequest) ([]ReplicationGroup, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[replicationGroupRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[replicationGroupRow, ReplicationGroup](dbRows)
	return resultList, nil
}

func (v *replicationGroups) ShowDatabases(ctx context.Context, request *ShowDatabasesInReplicationGroupRequest) ([]DatabaseInReplicationGroup, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[databaseInReplicationGroupRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[databaseInReplicationGroupRow, DatabaseInReplicationGroup](dbRows)
	return resultList, nil
}

func (v *replicationGroups) ShowShares(ctx context.Context, request *ShowSharesInReplicationGroupRequest) ([]ShareInReplicationGroup, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[shareInReplicationGroupRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[shareInReplicationGroupRow, ShareInReplicationGroup](dbRows)
	return resultList, nil
}

func (v *replicationGroups) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*ReplicationGroup, error) {
	replicationGroups, err := v.Show(ctx, NewShowReplicationGroupRequest())
	if err != nil {
		return nil, err
	}
	return collections.FindOne(replicationGroups, func(r ReplicationGroup) bool { return r.Name == id.Name() })
}

func (v *replicationGroups) Drop(ctx context.Context, request *DropReplicationGroupRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateReplicationGroupRequest) toOpts() *CreateReplicationGroupOptions {
	opts := &CreateReplicationGroupOptions{
		IfNotExists: r.IfNotExists,
		name:        r.name,

		IgnoreEditionCheck: r.IgnoreEditionCheck,
	}
	opts.ObjectTypes = ReplicationGroupObjectTypes{
		AccountParameters: r.ObjectTypes.AccountParameters,
		Databases:         r.ObjectTypes.Databases,
		Integrations:      r.ObjectTypes.Integrations,
		NetworkPolicies:   r.ObjectTypes.NetworkPolicies,
		ResourceMonitors:  r.ObjectTypes.ResourceMonitors,
		Roles:             r.ObjectTypes.Roles,
		Shares:            r.ObjectTypes.Shares,
		Users:             r.ObjectTypes.Users,
		Warehouses:        r.ObjectTypes.Warehouses,
	}
	if r.AllowedDatabases != nil {
		s := make([]ReplicationGroupDatabase, len(r.AllowedDatabases))
		for i, v := range r.AllowedDatabases {
			s[i] = ReplicationGroupDatabase(v)
		}
		opts.AllowedDatabases = s
	}
	if r.AllowedShares != nil {
		s := make([]ReplicationGroupShare, len(r.AllowedShares))
		for i, v := range r.AllowedShares {
			s[i] = ReplicationGroupShare(v)
		}
		opts.AllowedShares = s
	}
	if r.AllowedIntegrationTypes != nil {
		s := make([]ReplicationGroupIntegrationType, len(r.AllowedIntegrationTypes))
		for i, v := range r.AllowedIntegrationTypes {
			s[i] = ReplicationGroupIntegrationType(v)
		}
		opts.AllowedIntegrationTypes = s
	}
	if r.AllowedAccounts != nil {
		s := make([]ReplicationGroupAccount, len(r.AllowedAccounts))
		for i, v := range r.AllowedAccounts {
			s[i] = ReplicationGroupAccount(v)
		}
		opts.AllowedAccounts = s
	}
	if r.ReplicationSchedule != nil {
		opts.ReplicationSchedule = &ReplicationGroupSchedule{}
		if r.ReplicationSchedule.IntervalMinutes != nil {
			opts.ReplicationSchedule.IntervalMinutes = &ScheduleIntervalMinutes{
				Minutes: r.ReplicationSchedule.IntervalMinutes.Minutes,
			}
		}
		if r.ReplicationSchedule.CronExpression != nil {
			opts.ReplicationSchedule.CronExpression = &ScheduleCronExpression{
				Expression: r.ReplicationSchedule.CronExpression.Expression,
				TimeZone:   r.ReplicationSchedule.CronExpression.TimeZone,
			}
		}
	}
	return opts
}

func (r *CreateSecondaryReplicationGroupRequest) toOpts() *CreateSecondaryReplicationGroupOptions {
	opts := &CreateSecondaryReplicationGroupOptions{
		IfNotExists: r.IfNotExists,
		name:        r.name,
		Primary:     r.Primary,
	}
	return opts
}

func (r *AlterReplicationGroupRequest) toOpts() *AlterReplicationGroupOptions {
	opts := &AlterReplicationGroupOptions{
		IfExists: r.IfExists,
		name:     r.name,
		RenameTo: r.RenameTo,

		Refresh: r.Refresh,
		Suspend: r.Suspend,
		Resume:  r.Resume,
	}
	if r.Set != nil {
		opts.Set = &ReplicationGroupSet{
			EnableEtlReplication: r.Set.EnableEtlReplication,
		}
		if r.Set.ObjectTypes != nil {
			opts.Set.ObjectTypes = &ReplicationGroupObjectTypes{
				AccountParameters: r.Set.ObjectTypes.AccountParameters,
				Databases:         r.Set.ObjectTypes.Databases,
				Integrations:      r.Set.ObjectTypes.Integrations,
				NetworkPolicies:   r.Set.ObjectTypes.NetworkPolicies,
				ResourceMonitors:  r.Set.ObjectTypes.ResourceMonitors,
				Roles:             r.Set.ObjectTypes.Roles,
				Shares:            r.Set.ObjectTypes.Shares,
				Users:             r.Set.ObjectTypes.Users,
				Warehouses:        r.Set.ObjectTypes.Warehouses,
			}
		}
		if r.Set.AllowedDatabases != nil {
			s := make([]ReplicationGroupDatabase, len(r.Set.AllowedDatabases))
			for i, v := range r.Set.AllowedDatabases {
				s[i] = ReplicationGroupDatabase(v)
			}
			opts.Set.AllowedDatabases = s
		}
		if r.Set.AllowedShares != nil {
			s := make([]ReplicationGroupShare, len(r.Set.AllowedShares))
			for i, v := range r.Set.AllowedShares {
				s[i] = ReplicationGroupShare(v)
			}
			opts.Set.AllowedShares = s
		}
		if r.Set.ReplicationSchedule != nil {
			opts.Set.ReplicationSchedule = &ReplicationGroupSchedule{}
			if r.SetIntegration.ReplicationSchedule.IntervalMinutes != nil {
				opts.SetIntegration.ReplicationSchedule.IntervalMinutes = &ScheduleIntervalMinutes{
					Minutes: r.SetIntegration.ReplicationSchedule.IntervalMinutes.Minutes,
				}
			}
			if r.SetIntegration.ReplicationSchedule.CronExpression != nil {
				opts.SetIntegration.ReplicationSchedule.CronExpression = &ScheduleCronExpression{
					Expression: r.SetIntegration.ReplicationSchedule.CronExpression.Expression,
					TimeZone:   r.SetIntegration.ReplicationSchedule.CronExpression.TimeZone,
				}
			}
		}
	}
	if r.SetIntegration != nil {
		opts.SetIntegration = &ReplicationGroupSetIntegration{}
		if r.SetIntegration.ObjectTypes != nil {
			opts.SetIntegration.ObjectTypes = &ReplicationGroupObjectTypes{
				AccountParameters: r.SetIntegration.ObjectTypes.AccountParameters,
				Databases:         r.SetIntegration.ObjectTypes.Databases,
				Integrations:      r.SetIntegration.ObjectTypes.Integrations,
				NetworkPolicies:   r.SetIntegration.ObjectTypes.NetworkPolicies,
				ResourceMonitors:  r.SetIntegration.ObjectTypes.ResourceMonitors,
				Roles:             r.SetIntegration.ObjectTypes.Roles,
				Shares:            r.SetIntegration.ObjectTypes.Shares,
				Users:             r.SetIntegration.ObjectTypes.Users,
				Warehouses:        r.SetIntegration.ObjectTypes.Warehouses,
			}
		}
		if r.SetIntegration.AllowedIntegrationTypes != nil {
			s := make([]ReplicationGroupIntegrationType, len(r.SetIntegration.AllowedIntegrationTypes))
			for i, v := range r.SetIntegration.AllowedIntegrationTypes {
				s[i] = ReplicationGroupIntegrationType(v)
			}
			opts.SetIntegration.AllowedIntegrationTypes = s
		}
		if r.SetIntegration.ReplicationSchedule != nil {
			opts.SetIntegration.ReplicationSchedule = &ReplicationGroupSchedule{}
			if r.SetIntegration.ReplicationSchedule.IntervalMinutes != nil {
				opts.SetIntegration.ReplicationSchedule.IntervalMinutes = &ScheduleIntervalMinutes{
					Minutes: r.SetIntegration.ReplicationSchedule.IntervalMinutes.Minutes,
				}
			}
			if r.SetIntegration.ReplicationSchedule.CronExpression != nil {
				opts.SetIntegration.ReplicationSchedule.CronExpression = &ScheduleCronExpression{
					Expression: r.SetIntegration.ReplicationSchedule.CronExpression.Expression,
					TimeZone:   r.SetIntegration.ReplicationSchedule.CronExpression.TimeZone,
				}
			}
		}
	}
	if r.AddDatabases != nil {
		opts.AddDatabases = &ReplicationGroupAddDatabases{}
		if r.AddDatabases.Databases != nil {
			s := make([]ReplicationGroupDatabase, len(r.AddDatabases.Databases))
			for i, v := range r.AddDatabases.Databases {
				s[i] = ReplicationGroupDatabase(v)
			}
			opts.AddDatabases.Databases = s
		}
	}
	if r.RemoveDatabases != nil {
		opts.RemoveDatabases = &ReplicationGroupRemoveDatabases{}
		if r.RemoveDatabases.Databases != nil {
			s := make([]ReplicationGroupDatabase, len(r.RemoveDatabases.Databases))
			for i, v := range r.RemoveDatabases.Databases {
				s[i] = ReplicationGroupDatabase(v)
			}
			opts.RemoveDatabases.Databases = s
		}
	}
	if r.MoveDatabases != nil {
		opts.MoveDatabases = &ReplicationGroupMoveDatabases{
			MoveTo: r.MoveDatabases.MoveTo,
		}
		if r.MoveDatabases.Databases != nil {
			s := make([]ReplicationGroupDatabase, len(r.MoveDatabases.Databases))
			for i, v := range r.MoveDatabases.Databases {
				s[i] = ReplicationGroupDatabase(v)
			}
			opts.MoveDatabases.Databases = s
		}
	}
	if r.AddShares != nil {
		opts.AddShares = &ReplicationGroupAddShares{}
		if r.AddShares.Shares != nil {
			s := make([]ReplicationGroupShare, len(r.AddShares.Shares))
			for i, v := range r.AddShares.Shares {
				s[i] = ReplicationGroupShare(v)
			}
			opts.AddShares.Shares = s
		}
	}
	if r.RemoveShares != nil {
		opts.RemoveShares = &ReplicationGroupRemoveShares{}
		if r.RemoveShares.Shares != nil {
			s := make([]ReplicationGroupShare, len(r.RemoveShares.Shares))
			for i, v := range r.RemoveShares.Shares {
				s[i] = ReplicationGroupShare(v)
			}
			opts.RemoveShares.Shares = s
		}
	}
	if r.MoveShares != nil {
		opts.MoveShares = &ReplicationGroupMoveShares{
			MoveTo: r.MoveShares.MoveTo,
		}
		if r.MoveShares.Shares != nil {
			s := make([]ReplicationGroupShare, len(r.MoveShares.Shares))
			for i, v := range r.MoveShares.Shares {
				s[i] = ReplicationGroupShare(v)
			}
			opts.MoveShares.Shares = s
		}
	}
	if r.AddAccounts != nil {
		opts.AddAccounts = &ReplicationGroupAddAccounts{
			IgnoreEditionCheck: r.AddAccounts.IgnoreEditionCheck,
		}
		if r.AddAccounts.Accounts != nil {
			s := make([]ReplicationGroupAccount, len(r.AddAccounts.Accounts))
			for i, v := range r.AddAccounts.Accounts {
				s[i] = ReplicationGroupAccount(v)
			}
			opts.AddAccounts.Accounts = s
		}
	}
	if r.RemoveAccounts != nil {
		opts.RemoveAccounts = &ReplicationGroupRemoveAccounts{}
		if r.RemoveAccounts.Accounts != nil {
			s := make([]ReplicationGroupAccount, len(r.RemoveAccounts.Accounts))
			for i, v := range r.RemoveAccounts.Accounts {
				s[i] = ReplicationGroupAccount(v)
			}
			opts.RemoveAccounts.Accounts = s
		}
	}
	return opts
}

func (r *ShowReplicationGroupRequest) toOpts() *ShowReplicationGroupOptions {
	opts := &ShowReplicationGroupOptions{
		InAccount: r.InAccount,
	}
	return opts
}

func (r replicationGroupRow) convert() *ReplicationGroup {
	e := &ReplicationGroup{
		SnowflakeRegion:         r.SnowflakeRegion,
		CreatedOn:               r.CreatedOn,
		AccountName:             r.AccountName,
		Name:                    r.Name,
		Type:                    r.Type,
		IsPrimary:               r.IsPrimary == "Y",
		Primary:                 r.Primary,
		ObjectTypes:             r.ObjectTypes,
		AllowedIntegrationTypes: r.AllowedIntegrationTypes,
		AllowedAccounts:         r.AllowedAccounts,
		OrganizationName:        r.OrganizationName,
		AccountLocator:          r.AccountLocator,
		ReplicationSchedule:     r.ReplicationSchedule,
		Owner:                   r.Owner,
	}
	if r.Comment.Valid {
		e.Comment = r.Comment.String
	}
	if r.SecondaryState.Valid {
		e.SecondaryState = r.SecondaryState.String
	}
	if r.NextScheduledRefresh.Valid {
		e.NextScheduledRefresh = r.NextScheduledRefresh.String
	}
	return e
}

func (r *ShowDatabasesInReplicationGroupRequest) toOpts() *ShowDatabasesInReplicationGroupOptions {
	opts := &ShowDatabasesInReplicationGroupOptions{
		name: r.name,
	}
	return opts
}

func (r databaseInReplicationGroupRow) convert() *DatabaseInReplicationGroup {
	e := &DatabaseInReplicationGroup{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		IsDefault:     r.IsDefault == "Y",
		IsCurrent:     r.IsCurrent == "Y",
		Origin:        r.Origin,
		Owner:         r.Owner,
		Comment:       r.Comment,
		Options:       r.Options,
		RetentionTime: r.RetentionTime,
		Kind:          r.Kind,
		OwnerRoleType: r.OwnerRoleType,
	}
	if r.Budget.Valid {
		e.Budget = r.Budget.String
	}
	return e
}

func (r *ShowSharesInReplicationGroupRequest) toOpts() *ShowSharesInReplicationGroupOptions {
	opts := &ShowSharesInReplicationGroupOptions{
		name: r.name,
	}
	return opts
}

func (r shareInReplicationGroupRow) convert() *ShareInReplicationGroup {
	e := &ShareInReplicationGroup{
		CreatedOn:         r.CreatedOn,
		Kind:              r.Kind,
		OwnerAccount:      r.OwnerAccount,
		Name:              r.Name,
		DatabaseName:      r.DatabaseName,
		To:                r.To,
		Owner:             r.Owner,
		Comment:           r.Comment,
		ListingGlobalName: r.ListingGlobalName,
	}
	return e
}

func (r *DropReplicationGroupRequest) toOpts() *DropReplicationGroupOptions {
	opts := &DropReplicationGroupOptions{
		IfExists: r.IfExists,
		name:     r.name,
	}
	return opts
}
