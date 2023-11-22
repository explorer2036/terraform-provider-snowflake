package sdk

import (
	"context"
)

var _ EventTables = (*eventTables)(nil)

type eventTables struct {
	client *Client
}

func (v *eventTables) Create(ctx context.Context, request *CreateEventTableRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateEventTableRequest) toOpts() *CreateEventTableOptions {
	opts := &CreateEventTableOptions{
		OrReplace:   r.OrReplace,
		IfNotExists: r.IfNotExists,
		name:        r.name,
		ClusterBy:   r.ClusterBy,
	}
	return opts
}
