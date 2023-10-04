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

type AlterEventTableRequest struct {
	name SchemaObjectIdentifier // required

	// One of
	clusteringAction         *ClusteringActionRequest
	searchOptimizationAction *SearchOptimizationActionRequest
	addRowAccessPolicy       *EventTableAddRowAccessPolicy
	dropRowAccessPolicy      *EventTableDropRowAccessPolicy
	dropAllRowAccessPolicies *bool
	set                      *EventTableSetRequest
	unset                    *EventTableUnsetRequest
	rename                   *EventTableRename
}

type ClusteringActionRequest struct {
	clusterBy *[]string
	suspend   *bool
	resume    *bool
	drop      *bool
}

func (s *ClusteringActionRequest) toOpts() *ClusteringAction {
	return &ClusteringAction{
		ClusterBy: s.clusterBy,
		Suspend:   s.suspend,
		Resume:    s.resume,
		Drop:      s.drop,
	}
}

type SearchOptimizationActionRequest struct {
	add  *AddSearchOptimization
	drop *DropSearchOptimization
}

func (s *SearchOptimizationActionRequest) toOpts() *SearchOptimizationAction {
	action := &SearchOptimizationAction{}
	if s.add != nil {
		action.Add = &AddSearchOptimization{
			On: s.add.On,
		}
	}
	if s.drop != nil {
		action.Drop = &DropSearchOptimization{
			On: s.drop.On,
		}
	}
	return action
}

type EventTableSetRequest struct {
	dataRetentionTimeInDays    *uint
	maxDataExtensionTimeInDays *uint
	changeTracking             *bool
	comment                    *string
	tag                        []TagAssociationRequest
}

func (s *EventTableSetRequest) toOpts() *EventTableSet {
	opts := &EventTableSet{}
	if s.dataRetentionTimeInDays != nil || s.maxDataExtensionTimeInDays != nil || s.changeTracking != nil || s.comment != nil {
		opts.Properties = &EventTableSetProperties{}
		if s.dataRetentionTimeInDays != nil {
			opts.Properties.DataRetentionTimeInDays = s.dataRetentionTimeInDays
		}
		if s.maxDataExtensionTimeInDays != nil {
			opts.Properties.MaxDataExtensionTimeInDays = s.maxDataExtensionTimeInDays
		}
		if s.changeTracking != nil {
			opts.Properties.ChangeTracking = s.changeTracking
		}
		if s.comment != nil {
			opts.Properties.Comment = s.comment
		}
	}
	if len(s.tag) > 0 {
		tag := make([]TagAssociation, len(s.tag))
		for i, item := range s.tag {
			tag[i] = item.toOpts()
		}
		opts.Tag = &tag
	}
	return opts
}

type EventTableUnsetRequest struct {
	DataRetentionTimeInDays    *bool
	MaxDataExtensionTimeInDays *bool
	ChangeTracking             *bool
	Comment                    *bool
	tag                        []TagAssociationRequest
}

func (s *EventTableUnsetRequest) toOpts() *EventTableUnset {
	opts := &EventTableUnset{
		DataRetentionTimeInDays:    s.DataRetentionTimeInDays,
		MaxDataExtensionTimeInDays: s.MaxDataExtensionTimeInDays,
		ChangeTracking:             s.ChangeTracking,
		Comment:                    s.Comment,
	}
	if len(s.tag) > 0 {
		tag := make([]TagAssociation, len(s.tag))
		for i, item := range s.tag {
			tag[i] = item.toOpts()
		}
		opts.Tag = &tag
	}
	return opts
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
