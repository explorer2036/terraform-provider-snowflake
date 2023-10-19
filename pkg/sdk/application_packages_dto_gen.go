package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateApplicationPackageOptions] = new(CreateApplicationPackageRequest)
	_ optionsProvider[AlterApplicationPackageOptions]  = new(AlterApplicationPackageRequest)
	_ optionsProvider[DropApplicationPackageOptions]   = new(DropApplicationPackageRequest)
	_ optionsProvider[ShowApplicationPackageOptions]   = new(ShowApplicationPackageRequest)
)

type CreateApplicationPackageRequest struct {
	IfNotExists                *bool
	name                       AccountObjectIdentifier // required
	DataRetentionTimeInDays    *int
	MaxDataExtensionTimeInDays *int
	DefaultDdlCollation        *string
	Comment                    *string
	Tag                        []TagAssociation
	Distribution               *string
}

type AlterApplicationPackageRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
	Set      *ApplicationPackageSetRequest
	Unset    *ApplicationPackageUnsetRequest
}

type ApplicationPackageSetRequest struct {
	DataRetentionTimeInDays    *int
	MaxDataExtensionTimeInDays *int
	DefaultDdlCollation        *string
	Comment                    *string
	Distribution               *string
}

type ApplicationPackageUnsetRequest struct {
	DataRetentionTimeInDays    *bool
	MaxDataExtensionTimeInDays *bool
	DefaultDdlCollation        *bool
	Comment                    *bool
	Distribution               *bool
}

type DropApplicationPackageRequest struct {
	name AccountObjectIdentifier // required
}

type ShowApplicationPackageRequest struct {
	Like       *Like
	In         *In
	StartsWith *string
	Limit      *LimitFrom
}
