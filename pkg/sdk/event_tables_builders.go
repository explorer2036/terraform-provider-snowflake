package sdk

func NewCreateEventTableRequest(name SchemaObjectIdentifier) *CreateEventTableRequest {
	return &CreateEventTableRequest{
		name: name,
	}
}

func (s *CreateEventTableRequest) WithOrReplace(value bool) *CreateEventTableRequest {
	s.orReplace = value
	return s
}

func (s *CreateEventTableRequest) WithIfNotExists(value bool) *CreateEventTableRequest {
	s.ifNotExists = value
	return s
}

func (s *CreateEventTableRequest) WithClusterBy(value []string) *CreateEventTableRequest {
	s.clusterBy = value
	return s
}

func (s *CreateEventTableRequest) WithCopyGrants(value bool) *CreateEventTableRequest {
	s.copyGrants = &value
	return s
}

func (s *CreateEventTableRequest) WithDataRetentionTimeInDays(value uint) *CreateEventTableRequest {
	s.dataRetentionTimeInDays = &value
	return s
}

func (s *CreateEventTableRequest) WithMaxDataExtensionTimeInDays(value uint) *CreateEventTableRequest {
	s.maxDataExtensionTimeInDays = &value
	return s
}

func (s *CreateEventTableRequest) WithChangeTracking(value bool) *CreateEventTableRequest {
	s.changeTracking = &value
	return s
}

func (s *CreateEventTableRequest) WithDefaultDDLCollation(value string) *CreateEventTableRequest {
	s.defaultDDLCollation = &value
	return s
}

func (s *CreateEventTableRequest) WithComment(value string) *CreateEventTableRequest {
	s.comment = &value
	return s
}

func (s *CreateEventTableRequest) WithRowAccessPolicy(value *RowAccessPolicyRequest) *CreateEventTableRequest {
	s.rowAccessPolicy = value
	return s
}

func (s *CreateEventTableRequest) WithTag(value []TagAssociationRequest) *CreateEventTableRequest {
	s.tag = value
	return s
}

func NewAlterEventTableRequest(name SchemaObjectIdentifier) *AlterEventTableRequest {
	return &AlterEventTableRequest{
		name: name,
	}
}

func (s *AlterEventTableRequest) WithRename(name SchemaObjectIdentifier) *AlterEventTableRequest {
	s.rename = &EventTableRename{
		Name: name,
	}
	return s
}

func (s *AlterEventTableRequest) WithSet(value *EventTableSetRequest) *AlterEventTableRequest {
	s.set = value
	return s
}

func (s *AlterEventTableRequest) WithUnset(value *EventTableUnsetRequest) *AlterEventTableRequest {
	s.unset = value
	return s
}

func (s *AlterEventTableRequest) WithDropAllRowAccessPolicies(value bool) *AlterEventTableRequest {
	s.dropAllRowAccessPolicies = &value
	return s
}

func (s *AlterEventTableRequest) WithDropRowAccessPolicy(name SchemaObjectIdentifier) *AlterEventTableRequest {
	s.dropRowAccessPolicy = &EventTableDropRowAccessPolicy{
		Name: name,
	}
	return s
}

func (s *AlterEventTableRequest) WithAddRowAccessPolicy(policy *RowAccessPolicyRequest) *AlterEventTableRequest {
	s.addRowAccessPolicy = &EventTableAddRowAccessPolicy{
		RowAccessPolicy: policy.toOpts(),
	}
	return s
}

func NewClusteringActionRequest() *ClusteringActionRequest {
	return &ClusteringActionRequest{}
}

func (s *ClusteringActionRequest) WithClusterBy(value []string) *ClusteringActionRequest {
	if len(value) > 0 {
		s.clusterBy = &value
	}
	return s
}

func (s *ClusteringActionRequest) WithSuspend(value bool) *ClusteringActionRequest {
	s.suspend = &value
	return s
}

func (s *ClusteringActionRequest) WithResume(value bool) *ClusteringActionRequest {
	s.resume = &value
	return s
}

func (s *ClusteringActionRequest) WithDrop(value bool) *ClusteringActionRequest {
	s.drop = &value
	return s
}

func NewSearchOptimizationActionRequest() *SearchOptimizationActionRequest {
	return &SearchOptimizationActionRequest{}
}

func (s *SearchOptimizationActionRequest) WithAdd(on []string) *SearchOptimizationActionRequest {
	if len(on) > 0 {
		s.add = &AddSearchOptimization{
			On: on,
		}
	}
	return s
}

func (s *SearchOptimizationActionRequest) WithDrop(on []string) *SearchOptimizationActionRequest {
	if len(on) > 0 {
		s.drop = &DropSearchOptimization{
			On: on,
		}
	}
	return s
}

func NewEventTableSetRequest() *EventTableSetRequest {
	return &EventTableSetRequest{}
}

func (s *EventTableSetRequest) WithDataRetentionTimeInDays(value uint) *EventTableSetRequest {
	s.dataRetentionTimeInDays = &value
	return s
}

func (s *EventTableSetRequest) WithMaxDataExtensionTimeInDays(value uint) *EventTableSetRequest {
	s.maxDataExtensionTimeInDays = &value
	return s
}

func (s *EventTableSetRequest) WithChangeTracking(value bool) *EventTableSetRequest {
	s.changeTracking = &value
	return s
}

func (s *EventTableSetRequest) WithComment(value string) *EventTableSetRequest {
	s.comment = &value
	return s
}

func (s *EventTableSetRequest) WithTag(value []TagAssociationRequest) *EventTableSetRequest {
	s.tag = value
	return s
}

func NewEventTableUnsetRequest() *EventTableUnsetRequest {
	return &EventTableUnsetRequest{}
}

func (s *EventTableUnsetRequest) WithDataRetentionTimeInDays(value bool) *EventTableUnsetRequest {
	s.DataRetentionTimeInDays = &value
	return s
}

func (s *EventTableUnsetRequest) WithMaxDataExtensionTimeInDays(value bool) *EventTableUnsetRequest {
	s.MaxDataExtensionTimeInDays = &value
	return s
}

func (s *EventTableUnsetRequest) WithChangeTracking(value bool) *EventTableUnsetRequest {
	s.ChangeTracking = &value
	return s
}

func (s *EventTableUnsetRequest) WithComment(value bool) *EventTableUnsetRequest {
	s.Comment = &value
	return s
}

func (s *EventTableUnsetRequest) WithTag(value []TagAssociationRequest) *EventTableUnsetRequest {
	s.tag = value
	return s
}

func NewShowEventTableRequest() *ShowEventTableRequest {
	return &ShowEventTableRequest{}
}

func (s *ShowEventTableRequest) WithLike(value string) *ShowEventTableRequest {
	s.like = &Like{
		Pattern: String(value),
	}
	return s
}

func (s *ShowEventTableRequest) WithIn(in *In) *ShowEventTableRequest {
	s.in = in
	return s
}

func (s *ShowEventTableRequest) WithStartsWith(value string) *ShowEventTableRequest {
	s.startsWith = &value
	return s
}

func (s *ShowEventTableRequest) WithLimit(limit *LimitFrom) *ShowEventTableRequest {
	s.limit = limit
	return s
}
