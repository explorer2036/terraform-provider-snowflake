package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateApplicationOptions] = new(CreateApplicationRequest)
)

type CreateApplicationRequest struct {
	name               AccountObjectIdentifier // required
	PackageName        AccountObjectIdentifier // required
	ApplicationVersion *ApplicationVersionRequest
	DebugMode          *bool
	Comment            *string
	Tag                []TagAssociation
}

type ApplicationVersionRequest struct {
	VersionDirectory *string
	VersionAndPatch  *VersionAndPatchRequest
}

type VersionAndPatchRequest struct {
	Version string // required
	Patch   int    // required
}
