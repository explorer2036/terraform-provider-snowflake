package sdk

import (
	"context"
)

var _ ReplicationGroups = (*replicationGroups)(nil)

type replicationGroups struct {
	client *Client
}

func (v *replicationGroups) Create(ctx context.Context, request *CreateReplicationGroupRequest) error {
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
			s[i] = ReplicationGroupDatabase{
				Database: v.Database,
			}
		}
		opts.AllowedDatabases = s
	}
	if r.AllowedShares != nil {
		s := make([]ReplicationGroupShare, len(r.AllowedShares))
		for i, v := range r.AllowedShares {
			s[i] = ReplicationGroupShare{
				Share: v.Share,
			}
		}
		opts.AllowedShares = s
	}
	if r.AllowedIntegrationTypes != nil {
		s := make([]ReplicationGroupIntegrationType, len(r.AllowedIntegrationTypes))
		for i, v := range r.AllowedIntegrationTypes {
			s[i] = ReplicationGroupIntegrationType{
				IntegrationType: v.IntegrationType,
			}
		}
		opts.AllowedIntegrationTypes = s
	}
	if r.AllowedAccounts != nil {
		s := make([]ReplicationGroupAccount, len(r.AllowedAccounts))
		for i, v := range r.AllowedAccounts {
			s[i] = ReplicationGroupAccount{
				Account: v.Account,
			}
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
