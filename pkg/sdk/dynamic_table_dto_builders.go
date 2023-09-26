package sdk

func NewCreateDynamicTableRequest(
	name AccountObjectIdentifier,
	warehouse AccountObjectIdentifier,
	targetLag TargetLag,
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

func NewAlterDynamicTableRequest(
	name AccountObjectIdentifier,
) *AlterDynamicTableRequest {
	s := AlterDynamicTableRequest{}
	s.name = name
	return &s
}

func (s *AlterDynamicTableRequest) WithSuspend(suspend *bool) *AlterDynamicTableRequest {
	s.suspend = suspend
	return s
}

func (s *AlterDynamicTableRequest) WithResume(resume *bool) *AlterDynamicTableRequest {
	s.resume = resume
	return s
}

func (s *AlterDynamicTableRequest) WithRefresh(refresh *bool) *AlterDynamicTableRequest {
	s.refresh = refresh
	return s
}

func (s *AlterDynamicTableRequest) WithSet(set *DynamicTableSetRequest) *AlterDynamicTableRequest {
	s.set = set
	return s
}

func NewDynamicTableSetRequest() *DynamicTableSetRequest {
	return &DynamicTableSetRequest{}
}

func (s *DynamicTableSetRequest) WithTargetLag(targetLag *TargetLag) *DynamicTableSetRequest {
	s.targetLag = targetLag
	return s
}

func (s *DynamicTableSetRequest) WithWarehourse(warehourse *AccountObjectIdentifier) *DynamicTableSetRequest {
	s.warehourse = warehourse
	return s
}

func NewDropDynamicTableRequest(
	name AccountObjectIdentifier,
) *DropDynamicTableRequest {
	s := DropDynamicTableRequest{}
	s.name = name
	return &s
}

func NewDescribeDynamicTableRequest(
	name AccountObjectIdentifier,
) *DescribeDynamicTableRequest {
	s := DescribeDynamicTableRequest{}
	s.name = name
	return &s
}

func NewShowDynamicTableRequest() *ShowDynamicTableRequest {
	return &ShowDynamicTableRequest{}
}

func (s *ShowDynamicTableRequest) WithLike(like *Like) *ShowDynamicTableRequest {
	s.like = like
	return s
}

func (s *ShowDynamicTableRequest) WithIn(in *In) *ShowDynamicTableRequest {
	s.in = in
	return s
}

func (s *ShowDynamicTableRequest) WithStartsWith(startsWith *string) *ShowDynamicTableRequest {
	s.startsWith = startsWith
	return s
}

func (s *ShowDynamicTableRequest) WithLimitFrom(limitFrom *LimitFrom) *ShowDynamicTableRequest {
	s.limitFrom = limitFrom
	return s
}
