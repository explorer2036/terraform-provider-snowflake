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

func (v *eventTables) Show(ctx context.Context, request *ShowEventTableRequest) ([]EventTable, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[eventTableRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[eventTableRow, EventTable](dbRows)
	return resultList, nil
}

func (r *CreateEventTableRequest) toOpts() *CreateEventTableOptions {
	opts := &CreateEventTableOptions{
		OrReplace:                  r.OrReplace,
		IfNotExists:                r.IfNotExists,
		name:                       r.name,
		ClusterBy:                  r.ClusterBy,
		DataRetentionTimeInDays:    r.DataRetentionTimeInDays,
		MaxDataExtensionTimeInDays: r.MaxDataExtensionTimeInDays,
		ChangeTracking:             r.ChangeTracking,
		DefaultDdlCollation:        r.DefaultDdlCollation,
		CopyGrants:                 r.CopyGrants,
		Comment:                    r.Comment,
		RowAccessPolicy:            r.RowAccessPolicy,
		Tag:                        r.Tag,
	}
	return opts
}

func (r *ShowEventTableRequest) toOpts() *ShowEventTableOptions {
	opts := &ShowEventTableOptions{
		Like:       r.Like,
		In:         r.In,
		StartsWith: r.StartsWith,
		Limit:      r.Limit,
		From:       r.From,
	}
	return opts
}

func (r eventTableRow) convert() *EventTable {
	// TODO: Mapping
	return &EventTable{}
}
