// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewCreateEventTableRequest(
	name SchemaObjectIdentifier,
) *CreateEventTableRequest {
	s := CreateEventTableRequest{}
	s.name = name
	return &s
}

func (s *CreateEventTableRequest) WithOrReplace(OrReplace *bool) *CreateEventTableRequest {
	s.OrReplace = OrReplace
	return s
}

func (s *CreateEventTableRequest) WithIfNotExists(IfNotExists *bool) *CreateEventTableRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *CreateEventTableRequest) WithClusterBy(ClusterBy []string) *CreateEventTableRequest {
	s.ClusterBy = ClusterBy
	return s
}

func (s *CreateEventTableRequest) WithDataRetentionTimeInDays(DataRetentionTimeInDays *int) *CreateEventTableRequest {
	s.DataRetentionTimeInDays = DataRetentionTimeInDays
	return s
}

func (s *CreateEventTableRequest) WithMaxDataExtensionTimeInDays(MaxDataExtensionTimeInDays *int) *CreateEventTableRequest {
	s.MaxDataExtensionTimeInDays = MaxDataExtensionTimeInDays
	return s
}

func (s *CreateEventTableRequest) WithChangeTracking(ChangeTracking *bool) *CreateEventTableRequest {
	s.ChangeTracking = ChangeTracking
	return s
}

func (s *CreateEventTableRequest) WithDefaultDdlCollation(DefaultDdlCollation *string) *CreateEventTableRequest {
	s.DefaultDdlCollation = DefaultDdlCollation
	return s
}

func (s *CreateEventTableRequest) WithCopyGrants(CopyGrants *bool) *CreateEventTableRequest {
	s.CopyGrants = CopyGrants
	return s
}

func (s *CreateEventTableRequest) WithComment(Comment *string) *CreateEventTableRequest {
	s.Comment = Comment
	return s
}

func (s *CreateEventTableRequest) WithRowAccessPolicy(RowAccessPolicy *RowAccessPolicy) *CreateEventTableRequest {
	s.RowAccessPolicy = RowAccessPolicy
	return s
}

func (s *CreateEventTableRequest) WithTag(Tag []TagAssociation) *CreateEventTableRequest {
	s.Tag = Tag
	return s
}

func NewShowEventTableRequest() *ShowEventTableRequest {
	return &ShowEventTableRequest{}
}

func (s *ShowEventTableRequest) WithLike(Like *Like) *ShowEventTableRequest {
	s.Like = Like
	return s
}

func (s *ShowEventTableRequest) WithIn(In *In) *ShowEventTableRequest {
	s.In = In
	return s
}

func (s *ShowEventTableRequest) WithStartsWith(StartsWith *string) *ShowEventTableRequest {
	s.StartsWith = StartsWith
	return s
}

func (s *ShowEventTableRequest) WithLimit(Limit *LimitFrom) *ShowEventTableRequest {
	s.Limit = Limit
	return s
}

func NewDescribeEventTableRequest(
	name SchemaObjectIdentifier,
) *DescribeEventTableRequest {
	s := DescribeEventTableRequest{}
	s.name = name
	return &s
}

func NewAlterEventTableRequest(
	name SchemaObjectIdentifier,
) *AlterEventTableRequest {
	s := AlterEventTableRequest{}
	s.name = name
	return &s
}

func (s *AlterEventTableRequest) WithIfNotExists(IfNotExists *bool) *AlterEventTableRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *AlterEventTableRequest) WithSet(Set *EventTableSetRequest) *AlterEventTableRequest {
	s.Set = Set
	return s
}

func (s *AlterEventTableRequest) WithUnset(Unset *EventTableUnsetRequest) *AlterEventTableRequest {
	s.Unset = Unset
	return s
}

func (s *AlterEventTableRequest) WithAddRowAccessPolicy(AddRowAccessPolicy *EventTableAddRowAccessPolicyRequest) *AlterEventTableRequest {
	s.AddRowAccessPolicy = AddRowAccessPolicy
	return s
}

func (s *AlterEventTableRequest) WithDropRowAccessPolicy(DropRowAccessPolicy *EventTableDropRowAccessPolicyRequest) *AlterEventTableRequest {
	s.DropRowAccessPolicy = DropRowAccessPolicy
	return s
}

func (s *AlterEventTableRequest) WithDropAndAddRowAccessPolicy(DropAndAddRowAccessPolicy *EventTableDropAndAddRowAccessPolicyRequest) *AlterEventTableRequest {
	s.DropAndAddRowAccessPolicy = DropAndAddRowAccessPolicy
	return s
}

func (s *AlterEventTableRequest) WithDropAllRowAccessPolicies(DropAllRowAccessPolicies *bool) *AlterEventTableRequest {
	s.DropAllRowAccessPolicies = DropAllRowAccessPolicies
	return s
}

func (s *AlterEventTableRequest) WithClusteringAction(ClusteringAction *EventTableClusteringActionRequest) *AlterEventTableRequest {
	s.ClusteringAction = ClusteringAction
	return s
}

func (s *AlterEventTableRequest) WithSearchOptimizationAction(SearchOptimizationAction *EventTableSearchOptimizationActionRequest) *AlterEventTableRequest {
	s.SearchOptimizationAction = SearchOptimizationAction
	return s
}

func (s *AlterEventTableRequest) WithSetTags(SetTags []TagAssociation) *AlterEventTableRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterEventTableRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterEventTableRequest {
	s.UnsetTags = UnsetTags
	return s
}

func (s *AlterEventTableRequest) WithRenameTo(RenameTo *SchemaObjectIdentifier) *AlterEventTableRequest {
	s.RenameTo = RenameTo
	return s
}

func NewEventTableSetRequest() *EventTableSetRequest {
	return &EventTableSetRequest{}
}

func (s *EventTableSetRequest) WithDataRetentionTimeInDays(DataRetentionTimeInDays *int) *EventTableSetRequest {
	s.DataRetentionTimeInDays = DataRetentionTimeInDays
	return s
}

func (s *EventTableSetRequest) WithMaxDataExtensionTimeInDays(MaxDataExtensionTimeInDays *int) *EventTableSetRequest {
	s.MaxDataExtensionTimeInDays = MaxDataExtensionTimeInDays
	return s
}

func (s *EventTableSetRequest) WithChangeTracking(ChangeTracking *bool) *EventTableSetRequest {
	s.ChangeTracking = ChangeTracking
	return s
}

func (s *EventTableSetRequest) WithComment(Comment *string) *EventTableSetRequest {
	s.Comment = Comment
	return s
}

func NewEventTableUnsetRequest() *EventTableUnsetRequest {
	return &EventTableUnsetRequest{}
}

func (s *EventTableUnsetRequest) WithDataRetentionTimeInDays(DataRetentionTimeInDays *bool) *EventTableUnsetRequest {
	s.DataRetentionTimeInDays = DataRetentionTimeInDays
	return s
}

func (s *EventTableUnsetRequest) WithMaxDataExtensionTimeInDays(MaxDataExtensionTimeInDays *bool) *EventTableUnsetRequest {
	s.MaxDataExtensionTimeInDays = MaxDataExtensionTimeInDays
	return s
}

func (s *EventTableUnsetRequest) WithChangeTracking(ChangeTracking *bool) *EventTableUnsetRequest {
	s.ChangeTracking = ChangeTracking
	return s
}

func (s *EventTableUnsetRequest) WithComment(Comment *bool) *EventTableUnsetRequest {
	s.Comment = Comment
	return s
}

func NewEventTableAddRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
	On []string,
) *EventTableAddRowAccessPolicyRequest {
	s := EventTableAddRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	s.On = On
	return &s
}

func NewEventTableDropRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
) *EventTableDropRowAccessPolicyRequest {
	s := EventTableDropRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	return &s
}

func NewEventTableDropAndAddRowAccessPolicyRequest(
	Drop EventTableDropRowAccessPolicyRequest,
	Add EventTableAddRowAccessPolicyRequest,
) *EventTableDropAndAddRowAccessPolicyRequest {
	s := EventTableDropAndAddRowAccessPolicyRequest{}
	s.Drop = Drop
	s.Add = Add
	return &s
}

func NewEventTableClusteringActionRequest() *EventTableClusteringActionRequest {
	return &EventTableClusteringActionRequest{}
}

func (s *EventTableClusteringActionRequest) WithClusterBy(ClusterBy *[]string) *EventTableClusteringActionRequest {
	s.ClusterBy = ClusterBy
	return s
}

func (s *EventTableClusteringActionRequest) WithSuspendRecluster(SuspendRecluster *bool) *EventTableClusteringActionRequest {
	s.SuspendRecluster = SuspendRecluster
	return s
}

func (s *EventTableClusteringActionRequest) WithResumeRecluster(ResumeRecluster *bool) *EventTableClusteringActionRequest {
	s.ResumeRecluster = ResumeRecluster
	return s
}

func (s *EventTableClusteringActionRequest) WithDropClusteringKey(DropClusteringKey *bool) *EventTableClusteringActionRequest {
	s.DropClusteringKey = DropClusteringKey
	return s
}

func NewEventTableSearchOptimizationActionRequest() *EventTableSearchOptimizationActionRequest {
	return &EventTableSearchOptimizationActionRequest{}
}

func (s *EventTableSearchOptimizationActionRequest) WithAdd(Add *SearchOptimizationRequest) *EventTableSearchOptimizationActionRequest {
	s.Add = Add
	return s
}

func (s *EventTableSearchOptimizationActionRequest) WithDrop(Drop *SearchOptimizationRequest) *EventTableSearchOptimizationActionRequest {
	s.Drop = Drop
	return s
}

func NewSearchOptimizationRequest() *SearchOptimizationRequest {
	return &SearchOptimizationRequest{}
}

func (s *SearchOptimizationRequest) WithOn(On []string) *SearchOptimizationRequest {
	s.On = On
	return s
}
