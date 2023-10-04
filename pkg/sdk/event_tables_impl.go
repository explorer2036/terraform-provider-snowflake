package sdk

import "context"

var _ EventTables = (*eventTables)(nil)

type eventTables struct {
	client *Client
}

func (v *eventTables) Create(ctx context.Context, request *CreateEventTableRequest) error {
	return validateAndExec(v.client, ctx, request.toOpts())
}

func (v *eventTables) Alter(ctx context.Context, request *AlterEventTableRequest) error {
	return validateAndExec(v.client, ctx, request.toOpts())
}

func (v *eventTables) Describe(ctx context.Context, request *DescribeEventTableRequest) (*EventTableDetails, error) {
	row, err := validateAndQueryOne[eventTableDetailsRow](v.client, ctx, request.toOpts())
	if err != nil {
		return nil, err
	}
	return row.convert(), nil
}

func (v *eventTables) Show(ctx context.Context, request *ShowEventTableRequest) ([]EventTable, error) {
	rows, err := validateAndQuery[eventTableRow](v.client, ctx, request.toOpts())
	if err != nil {
		return nil, err
	}
	result := convertRows[eventTableRow, EventTable](rows)
	return result, nil
}

func (v *eventTables) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*EventTable, error) {
	request := NewShowEventTableRequest().WithLike(id.Name())
	eventTables, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return findOne(eventTables, func(r EventTable) bool { return r.Name == id.Name() })
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

func (s *AlterEventTableRequest) toOpts() *alterEventTableOptions {
	opts := &alterEventTableOptions{
		name: s.name,
	}
	if s.clusteringAction != nil {
		opts.ClusteringAction = s.clusteringAction.toOpts()
	}
	if s.searchOptimizationAction != nil {
		opts.SearchOptimizationAction = s.searchOptimizationAction.toOpts()
	}
	if s.addRowAccessPolicy != nil {
		opts.AddRowAccessPolicy = s.addRowAccessPolicy
	}
	if s.dropRowAccessPolicy != nil {
		opts.DropRowAccessPolicy = s.dropRowAccessPolicy
	}
	if s.dropAllRowAccessPolicies != nil {
		opts.DropAllRowAccessPolicies = s.dropAllRowAccessPolicies
	}
	if s.set != nil {
		opts.Set = s.set.toOpts()
	}
	if s.unset != nil {
		opts.Unset = s.unset.toOpts()
	}
	if s.rename != nil {
		opts.Rename = s.rename
	}
	return opts
}

func (s *DescribeEventTableRequest) toOpts() *describeEventTableOptions {
	return &describeEventTableOptions{
		name: s.name,
	}
}

func (s *ShowEventTableRequest) toOpts() *showEventTableOptions {
	return &showEventTableOptions{
		Like:       s.like,
		In:         s.in,
		StartsWith: s.startsWith,
		Limit:      s.limit,
	}
}
