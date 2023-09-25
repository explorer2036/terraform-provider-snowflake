package sdk

func NewCreateDynamicTableRequest(
	name AccountObjectIdentifier,
	warehouse AccountObjectIdentifier,
	targetLag string,
	query string,
) *CreateDynamicTableRequest {
	s := CreateDynamicTableRequest{}
	s.name = name
	s.warehouse = warehouse
	s.targetLag = targetLag
	s.query = query
	return &s
}

func (s *CreateDynamicTableRequest) WithOrReplace(orReplace bool) *CreateDynamicTableRequest {
	s.orReplace = orReplace
	return s
}

func (s *CreateDynamicTableRequest) WithComment(comment *string) *CreateDynamicTableRequest {
	s.comment = comment
	return s
}
