package sdk

import (
	"context"
)

var _ ApplicationPackages = (*applicationPackages)(nil)

type applicationPackages struct {
	client *Client
}

func (v *applicationPackages) Create(ctx context.Context, request *CreateApplicationPackageRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *applicationPackages) Alter(ctx context.Context, request *AlterApplicationPackageRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *applicationPackages) Drop(ctx context.Context, request *DropApplicationPackageRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *applicationPackages) Show(ctx context.Context, request *ShowApplicationPackageRequest) ([]ApplicationPackage, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[applicationPackageRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[applicationPackageRow, ApplicationPackage](dbRows)
	return resultList, nil
}

func (r *CreateApplicationPackageRequest) toOpts() *CreateApplicationPackageOptions {
	opts := &CreateApplicationPackageOptions{
		IfNotExists:                r.IfNotExists,
		name:                       r.name,
		DataRetentionTimeInDays:    r.DataRetentionTimeInDays,
		MaxDataExtensionTimeInDays: r.MaxDataExtensionTimeInDays,
		DefaultDdlCollation:        r.DefaultDdlCollation,
		Comment:                    r.Comment,
		Tag:                        r.Tag,
		Distribution:               r.Distribution,
	}
	return opts
}

func (r *AlterApplicationPackageRequest) toOpts() *AlterApplicationPackageOptions {
	opts := &AlterApplicationPackageOptions{
		IfExists: r.IfExists,
		name:     r.name,
	}
	if r.Set != nil {
		opts.Set = &ApplicationPackageSet{
			DataRetentionTimeInDays:    r.Set.DataRetentionTimeInDays,
			MaxDataExtensionTimeInDays: r.Set.MaxDataExtensionTimeInDays,
			DefaultDdlCollation:        r.Set.DefaultDdlCollation,
			Comment:                    r.Set.Comment,
			Distribution:               r.Set.Distribution,
		}
	}
	if r.Unset != nil {
		opts.Unset = &ApplicationPackageUnset{
			DataRetentionTimeInDays:    r.Unset.DataRetentionTimeInDays,
			MaxDataExtensionTimeInDays: r.Unset.MaxDataExtensionTimeInDays,
			DefaultDdlCollation:        r.Unset.DefaultDdlCollation,
			Comment:                    r.Unset.Comment,
			Distribution:               r.Unset.Distribution,
		}
	}
	return opts
}

func (r *DropApplicationPackageRequest) toOpts() *DropApplicationPackageOptions {
	opts := &DropApplicationPackageOptions{
		name: r.name,
	}
	return opts
}

func (r *ShowApplicationPackageRequest) toOpts() *ShowApplicationPackageOptions {
	opts := &ShowApplicationPackageOptions{
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
