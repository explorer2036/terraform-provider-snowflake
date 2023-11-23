package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateEventTableOptions] = new(CreateEventTableRequest)
	_ optionsProvider[ShowEventTableOptions]   = new(ShowEventTableRequest)
)

type CreateEventTableRequest struct {
	OrReplace                  *bool
	IfNotExists                *bool
	name                       SchemaObjectIdentifier // required
	ClusterBy                  []string
	DataRetentionTimeInDays    *int
	MaxDataExtensionTimeInDays *int
	ChangeTracking             *bool
	DefaultDdlCollation        *string
	CopyGrants                 *bool
	Comment                    *string
	RowAccessPolicy            *RowAccessPolicy
	Tag                        []TagAssociation
}

type ShowEventTableRequest struct {
	Like       *Like
	In         *In
	StartsWith *string
	Limit      *int
	From       *string
}
