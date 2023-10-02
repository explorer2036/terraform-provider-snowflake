package sdk

var (
	_ optionsProvider[createEventTableOptions]   = (*CreateEventTableRequest)(nil)
	_ optionsProvider[describeEventTableOptions] = (*DescribeEventTableRequest)(nil)
	_ optionsProvider[showEventTableOptions]     = (*ShowEventTableRequest)(nil)
)

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

type ShowEventTableRequest struct {
	like       *Like
	in         *In
	startsWith *string
	limit      *LimitFrom
}

type DescribeEventTableRequest struct {
	name SchemaObjectIdentifier // required
}
