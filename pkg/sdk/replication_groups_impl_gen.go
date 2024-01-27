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

		AllowedDatabases: r.AllowedDatabases,
		AllowedShares:    r.AllowedShares,

		AllowedAccounts:    r.AllowedAccounts,
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
	if r.AllowedIntegrationTypes != nil {
		s := make([]ReplicationGroupIntegrationType, len(r.AllowedIntegrationTypes))
		for i, v := range r.AllowedIntegrationTypes {
			s[i] = ReplicationGroupIntegrationType(v)
		}
		opts.AllowedIntegrationTypes = s
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
			AllowedDatabases: r.Set.AllowedDatabases,
			AllowedShares:    r.Set.AllowedShares,

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
		opts.AddDatabases = &ReplicationGroupAddDatabases{
			Add: r.AddDatabases.Add,
		}
	}
	if r.RemoveDatabases != nil {
		opts.RemoveDatabases = &ReplicationGroupRemoveDatabases{
			Remove: r.RemoveDatabases.Remove,
		}
	}
	if r.MoveDatabases != nil {
		opts.MoveDatabases = &ReplicationGroupMoveDatabases{
			MoveDatabases: r.MoveDatabases.MoveDatabases,
			MoveTo:        r.MoveDatabases.MoveTo,
		}
	}
	if r.AddShares != nil {
		opts.AddShares = &ReplicationGroupAddShares{
			Add: r.AddShares.Add,
		}
	}
	if r.RemoveShares != nil {
		opts.RemoveShares = &ReplicationGroupRemoveShares{
			Remove: r.RemoveShares.Remove,
		}
	}
	if r.MoveShares != nil {
		opts.MoveShares = &ReplicationGroupMoveShares{
			MoveShares: r.MoveShares.MoveShares,
			MoveTo:     r.MoveShares.MoveTo,
		}
	}
	if r.AddAccounts != nil {
		opts.AddAccounts = &ReplicationGroupAddAccounts{
			Add:                r.AddAccounts.Add,
			IgnoreEditionCheck: r.AddAccounts.IgnoreEditionCheck,
		}
	}
	if r.RemoveAccounts != nil {
		opts.RemoveAccounts = &ReplicationGroupRemoveAccounts{
			Remove: r.RemoveAccounts.Remove,
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
