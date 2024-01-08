package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateApplicationOptions]   = new(CreateApplicationRequest)
	_ optionsProvider[DropApplicationOptions]     = new(DropApplicationRequest)
	_ optionsProvider[ShowApplicationOptions]     = new(ShowApplicationRequest)
	_ optionsProvider[DescribeApplicationOptions] = new(DescribeApplicationRequest)
)

type CreateApplicationRequest struct {
	name        AccountObjectIdentifier // required
	PackageName AccountObjectIdentifier // required
	Version     *ApplicationVersionRequest
	DebugMode   *bool
	Comment     *string
	Tag         []TagAssociation
}

type ApplicationVersionRequest struct {
	VersionDirectory *string // required
	VersionAndPatch  *VersionAndPatchRequest
}

type VersionAndPatchRequest struct {
	Version string // required
	Patch   *int   // required
}

type DropApplicationRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
	Cascade  *bool
}

type ShowApplicationRequest struct {
	Like       *Like
	StartsWith *string
	Limit      *LimitFrom
}

type DescribeApplicationRequest struct {
	name AccountObjectIdentifier // required
}
