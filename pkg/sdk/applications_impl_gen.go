package sdk

import (
	"context"
)

var _ Applications = (*applications)(nil)

type applications struct {
	client *Client
}

func (v *applications) Create(ctx context.Context, request *CreateApplicationRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateApplicationRequest) toOpts() *CreateApplicationOptions {
	opts := &CreateApplicationOptions{
		name:        r.name,
		PackageName: r.PackageName,

		DebugMode: r.DebugMode,
		Comment:   r.Comment,
		Tag:       r.Tag,
	}
	if r.ApplicationVersion != nil {
		opts.ApplicationVersion = &ApplicationVersion{
			VersionDirectory: r.ApplicationVersion.VersionDirectory,
		}
		if r.ApplicationVersion.VersionAndPatch != nil {
			opts.ApplicationVersion.VersionAndPatch = &VersionAndPatch{
				Version: r.ApplicationVersion.VersionAndPatch.Version,
				Patch:   r.ApplicationVersion.VersionAndPatch.Patch,
			}
		}
	}
	return opts
}
