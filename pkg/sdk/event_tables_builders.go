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
