package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ ApplicationRoles = (*applicationRoles)(nil)

type applicationRoles struct {
	client *Client
}

func (v *applicationRoles) Show(ctx context.Context, request *ShowApplicationRoleRequest) ([]ApplicationRole, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[applicationRoleDbRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[applicationRoleDbRow, ApplicationRole](dbRows)
	return resultList, nil
}

func (v *applicationRoles) ShowByID(ctx context.Context, applicationName AccountObjectIdentifier, id DatabaseObjectIdentifier) (*ApplicationRole, error) {
	request := NewShowApplicationRoleRequest().WithApplicationName(applicationName)
	applicationRoles, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(applicationRoles, func(r ApplicationRole) bool { return r.Name == id.Name() })
}

func (r *ShowApplicationRoleRequest) toOpts() *ShowApplicationRoleOptions {
	opts := &ShowApplicationRoleOptions{
		ApplicationName: r.ApplicationName,
		Limit:           r.Limit,
	}
	return opts
}

func (r applicationRoleDbRow) convert() *ApplicationRole {
	return &ApplicationRole{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		Owner:         r.Owner,
		Comment:       r.Comment,
		OwnerRoleType: r.OwnerRoleType,
	}
}
