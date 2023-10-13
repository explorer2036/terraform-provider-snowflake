package sdk

var (
	_ optionsProvider[createEventTableOptions]   = (*CreateEventTableRequest)(nil)
	_ optionsProvider[alterEventTableOptions]    = (*AlterEventTableRequest)(nil)
	_ optionsProvider[describeEventTableOptions] = (*DescribeEventTableRequest)(nil)
	_ optionsProvider[showEventTableOptions]     = (*ShowEventTableRequest)(nil)
)

type CreateEventTableRequest struct {
	orReplace   bool
	ifNotExists bool

	name SchemaObjectIdentifier // required

	clusterBy                  []string
	dataRetentionTimeInDays    *int
	maxDataExtensionTimeInDays *int
	changeTracking             *bool
	defaultDDLCollation        *string
	copyGrants                 *bool
	comment                    *string
	rowAccessPolicy            *RowAccessPolicyRequest
	tag                        []*TagAssociationRequest
}

type AlterEventTableRequest struct {
	ifExists bool
	name     SchemaObjectIdentifier // required

	// One of
	clusteringAction         *ClusteringActionRequest
	searchOptimizationAction *SearchOptimizationActionRequest
	addRowAccessPolicy       *EventTableAddRowAccessPolicy
	dropRowAccessPolicy      *EventTableDropRowAccessPolicy
	dropAllRowAccessPolicies *bool
	set                      *EventTableSetRequest
	unset                    *EventTableUnsetRequest
	rename                   *RenameSchemaObjectIdentifier
}

type ClusteringActionRequest struct {
	clusterBy         *[]string
	suspendRecluster  *bool
	resumeRecluster   *bool
	dropClusteringKey *bool
}

func (s *ClusteringActionRequest) toOpts() *ClusteringAction {
	return &ClusteringAction{
		ClusterBy:         s.clusterBy,
		SuspendRecluster:  s.suspendRecluster,
		ResumeRecluster:   s.resumeRecluster,
		DropClusteringKey: s.dropClusteringKey,
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
	dataRetentionTimeInDays    *int
	maxDataExtensionTimeInDays *int
	changeTracking             *bool
	comment                    *string
	tag                        []*TagAssociationRequest
}

func (s *EventTableSetRequest) toOpts() *EventTableSet {
	opts := &EventTableSet{}
	if s.dataRetentionTimeInDays != nil {
		opts.DataRetentionTimeInDays = s.dataRetentionTimeInDays
	}
	if s.maxDataExtensionTimeInDays != nil {
		opts.MaxDataExtensionTimeInDays = s.maxDataExtensionTimeInDays
	}
	if s.changeTracking != nil {
		opts.ChangeTracking = s.changeTracking
	}
	if s.comment != nil {
		opts.Comment = s.comment
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
	TagNames                   *[]string
}

func (s *EventTableUnsetRequest) toOpts() *EventTableUnset {
	opts := &EventTableUnset{
		DataRetentionTimeInDays:    s.DataRetentionTimeInDays,
		MaxDataExtensionTimeInDays: s.MaxDataExtensionTimeInDays,
		ChangeTracking:             s.ChangeTracking,
		Comment:                    s.Comment,
		TagNames:                   s.TagNames,
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
