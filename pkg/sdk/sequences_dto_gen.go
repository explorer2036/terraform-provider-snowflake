package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateSequenceOptions]   = new(CreateSequenceRequest)
	_ optionsProvider[ShowSequenceOptions]     = new(ShowSequenceRequest)
	_ optionsProvider[DescribeSequenceOptions] = new(DescribeSequenceRequest)
	_ optionsProvider[DropSequenceOptions]     = new(DropSequenceRequest)
)

type CreateSequenceRequest struct {
	OrReplace      *bool
	IfNotExists    *bool
	name           SchemaObjectIdentifier // required
	With           *bool
	Start          *int
	Increment      *int
	ValuesBehavior *ValuesBehavior
	Comment        *string
}

type ShowSequenceRequest struct {
	Like *Like
	In   *In
}

type DescribeSequenceRequest struct {
	name SchemaObjectIdentifier // required
}

type DropSequenceRequest struct {
	IfExists *bool
	name     SchemaObjectIdentifier // required
}
