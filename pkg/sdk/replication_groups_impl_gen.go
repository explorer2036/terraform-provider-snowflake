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
	if r.Databases != nil {
		s := make([]ReplicationGroupDatabase, len(r.Databases))
		for i, v := range r.Databases {
			s[i] = ReplicationGroupDatabase{
				Database: v.Database,
			}
		}
		opts.Databases = s
	}
	if r.Shares != nil {
		s := make([]ReplicationGroupShare, len(r.Shares))
		for i, v := range r.Shares {
			s[i] = ReplicationGroupShare{
				Share: v.Share,
			}
		}
		opts.Shares = s
	}
	if r.IntegrationTypes != nil {
		s := make([]ReplicationGroupIntegrationType, len(r.IntegrationTypes))
		for i, v := range r.IntegrationTypes {
			s[i] = ReplicationGroupIntegrationType{
				IntegrationType: v.IntegrationType,
			}
		}
		opts.IntegrationTypes = s
	}
	if r.Accounts != nil {
		s := make([]ReplicationGroupAccount, len(r.Accounts))
		for i, v := range r.Accounts {
			s[i] = ReplicationGroupAccount{
				Account: v.Account,
			}
		}
		opts.Accounts = s
	}
	return opts
}
