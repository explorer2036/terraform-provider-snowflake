package sdk

import "context"

var _ EventTables = (*eventTables)(nil)

type eventTables struct {
	client *Client
}

func (v *eventTables) Create(ctx context.Context, request *CreateEventTableRequest) error {
	return validateAndExec(v.client, ctx, request.toOpts())
}

func (v *CreateEventTableRequest) toOpts() *createEventTableOptions {
	opts := &createEventTableOptions{
		OrReplace:                  Bool(v.orReplace),
		IfNotExists:                Bool(v.ifNotExists),
		name:                       v.name,
		ClusterBy:                  v.clusterBy,
		DataRetentionTimeInDays:    v.dataRetentionTimeInDays,
		MaxDataExtensionTimeInDays: v.maxDataExtensionTimeInDays,
		ChangeTracking:             v.changeTracking,
		DefaultDDLCollation:        v.defaultDDLCollation,
		CopyGrants:                 v.copyGrants,
		Comment:                    v.comment,
	}
	if v.rowAccessPolicy != nil {
		opts.RowAccessPolicy = v.rowAccessPolicy.toOpts()
	}
	if len(v.tag) > 0 {
		tag := make([]TagAssociation, len(v.tag))
		for i, item := range v.tag {
			tag[i] = item.toOpts()
		}
		opts.Tag = tag
	}
	return opts
}
