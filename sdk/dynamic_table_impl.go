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
