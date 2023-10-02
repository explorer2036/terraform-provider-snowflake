package sdk

var _ optionsProvider[createEventTableOptions] = (*CreateEventTableRequest)(nil)

type CreateEventTableRequest struct {
	orReplace   bool
	ifNotExists bool

	name SchemaObjectIdentifier // required

	clusterBy                  []string
	dataRetentionTimeInDays    *uint
	maxDataExtensionTimeInDays *uint
	changeTracking             *bool
	defaultDDLCollation        *string
	copyGrants                 *bool
	comment                    *string
	rowAccessPolicy            *RowAccessPolicyRequest
	tag                        []TagAssociationRequest
}
