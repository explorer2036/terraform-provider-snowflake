package sdk

import "context"

var _ DynamicTables = (*dynamicTables)(nil)

type dynamicTables struct {
	client *Client
}

func (v *dynamicTables) Create(ctx context.Context, request *CreateDynamicTableRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *dynamicTables) Alter(ctx context.Context, request *AlterDynamicTableRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (s *CreateDynamicTableRequest) toOpts() *createDynamicTableOptions {
	return &createDynamicTableOptions{
		OrReplace: Bool(s.orReplace),
		name:      s.name,
		warehouse: s.warehouse,
		targetLag: s.targetLag,
		query:     s.query,
		Comment:   s.comment,
	}
}

func (s *AlterDynamicTableRequest) toOpts() *alterDynamicTableOptions {
	opts := alterDynamicTableOptions{
		name: s.name,
	}
	if s.suspend != nil {
		opts.Suspend = s.suspend
	}
	if s.resume != nil {
		opts.Resume = s.resume
	}
	if s.refresh != nil {
		opts.Refresh = s.refresh
	}
	if s.set != nil {
		opts.Set = &DynamicTableSet{s.set.targetLag, s.set.warehourse}
	}
	return &opts
}
