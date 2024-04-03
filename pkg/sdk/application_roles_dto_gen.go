package sdk

//go:generate go run ./dto-builder-generator/main.go

var _ optionsProvider[ShowApplicationRoleOptions] = new(ShowApplicationRoleRequest)

type ShowApplicationRoleRequest struct {
	ApplicationName AccountObjectIdentifier
	Limit           *LimitFrom
}
