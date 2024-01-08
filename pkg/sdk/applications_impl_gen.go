package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ Applications = (*applications)(nil)

type applications struct {
	client *Client
}

func (v *applications) Create(ctx context.Context, request *CreateApplicationRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *applications) Drop(ctx context.Context, request *DropApplicationRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *applications) Show(ctx context.Context, request *ShowApplicationRequest) ([]ApplicationPackage, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[applicationPackageRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[applicationPackageRow, ApplicationPackage](dbRows)
	return resultList, nil
}

func (v *applications) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*Application, error) {
	// TODO: adjust request if e.g. LIKE is supported for the resource
	applications, err := v.Show(ctx, NewShowApplicationRequest())
	if err != nil {
		return nil, err
	}
	return collections.FindOne(applications, func(r Application) bool { return r.Name == id.Name() })
}

func (v *applications) Describe(ctx context.Context, id AccountObjectIdentifier) ([]ApplicationDetail, error) {
	opts := &DescribeApplicationOptions{
		name: id,
	}
	rows, err := validateAndQuery[applicationDetailRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertRows[applicationDetailRow, ApplicationDetail](rows), nil
}

func (r *CreateApplicationRequest) toOpts() *CreateApplicationOptions {
	opts := &CreateApplicationOptions{
		name:        r.name,
		PackageName: r.PackageName,

		DebugMode: r.DebugMode,
		Comment:   r.Comment,
		Tag:       r.Tag,
	}
	if r.Version != nil {
		opts.Version = &ApplicationVersion{
			VersionDirectory: r.Version.VersionDirectory,
		}
		if r.Version.VersionAndPatch != nil {
			opts.Version.VersionAndPatch = &VersionAndPatch{
				Version: r.Version.VersionAndPatch.Version,
				Patch:   r.Version.VersionAndPatch.Patch,
			}
		}
	}
	return opts
}

func (r *DropApplicationRequest) toOpts() *DropApplicationOptions {
	opts := &DropApplicationOptions{
		IfExists: r.IfExists,
		name:     r.name,
		Cascade:  r.Cascade,
	}
	return opts
}

func (r *ShowApplicationRequest) toOpts() *ShowApplicationOptions {
	opts := &ShowApplicationOptions{
		Like:       r.Like,
		StartsWith: r.StartsWith,
		Limit:      r.Limit,
	}
	return opts
}

func (r applicationPackageRow) convert() *ApplicationPackage {
	// TODO: Mapping
	return &ApplicationPackage{}
}

func (r *DescribeApplicationRequest) toOpts() *DescribeApplicationOptions {
	opts := &DescribeApplicationOptions{
		name: r.name,
	}
	return opts
}

func (r applicationDetailRow) convert() *ApplicationDetail {
	// TODO: Mapping
	return &ApplicationDetail{}
}
