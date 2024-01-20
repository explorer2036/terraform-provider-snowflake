package sdk

import (
	"context"
)

var _ Streamlits = (*streamlits)(nil)

type streamlits struct {
	client *Client
}

func (v *streamlits) Create(ctx context.Context, request *CreateStreamlitRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *streamlits) Alter(ctx context.Context, request *AlterStreamlitRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateStreamlitRequest) toOpts() *CreateStreamlitOptions {
	opts := &CreateStreamlitOptions{
		OrReplace:    r.OrReplace,
		IfNotExists:  r.IfNotExists,
		name:         r.name,
		RootLocation: r.RootLocation,
		MainFile:     r.MainFile,
		Warehouse:    r.Warehouse,
		Comment:      r.Comment,
	}
	return opts
}

func (r *AlterStreamlitRequest) toOpts() *AlterStreamlitOptions {
	opts := &AlterStreamlitOptions{
		IfExists: r.IfExists,
		name:     r.name,

		RenameTo: r.RenameTo,
	}
	if r.Set != nil {
		opts.Set = &StreamlitsSet{
			RootLocation: r.Set.RootLocation,
			Warehouse:    r.Set.Warehouse,
			MainFile:     r.Set.MainFile,
			Comment:      r.Set.Comment,
		}
	}
	return opts
}
