package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ Sequences = (*sequences)(nil)

type sequences struct {
	client *Client
}

func (v *sequences) Create(ctx context.Context, request *CreateSequenceRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *sequences) Show(ctx context.Context, request *ShowSequenceRequest) ([]Sequence, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[sequenceRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[sequenceRow, Sequence](dbRows)
	return resultList, nil
}

func (v *sequences) ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*Sequence, error) {
	request := NewShowSequenceRequest().WithIn(&In{Database: NewAccountObjectIdentifier(id.DatabaseName())}).WithLike(&Like{String(id.Name())})
	sequences, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(sequences, func(r Sequence) bool { return r.Name == id.Name() })
}

func (v *sequences) Describe(ctx context.Context, id SchemaObjectIdentifier) ([]SequenceDetail, error) {
	opts := &DescribeSequenceOptions{
		name: id,
	}
	rows, err := validateAndQuery[sequenceDetailRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertRows[sequenceDetailRow, SequenceDetail](rows), nil
}

func (v *sequences) Drop(ctx context.Context, request *DropSequenceRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateSequenceRequest) toOpts() *CreateSequenceOptions {
	opts := &CreateSequenceOptions{
		OrReplace:      r.OrReplace,
		IfNotExists:    r.IfNotExists,
		name:           r.name,
		With:           r.With,
		Start:          r.Start,
		Increment:      r.Increment,
		ValuesBehavior: r.ValuesBehavior,
		Comment:        r.Comment,
	}
	return opts
}

func (r *ShowSequenceRequest) toOpts() *ShowSequenceOptions {
	opts := &ShowSequenceOptions{
		Like: r.Like,
		In:   r.In,
	}
	return opts
}

func (r sequenceRow) convert() *Sequence {
	return &Sequence{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		SchemaName:    r.SchemaName,
		DatabaseName:  r.DatabaseName,
		NextValue:     r.NextValue,
		Interval:      r.Interval,
		Owner:         r.Owner,
		OwnerRoleType: r.OwnerRoleType,
		Comment:       r.Comment,
		Ordered:       r.Ordered == "Y",
	}
}

func (r *DescribeSequenceRequest) toOpts() *DescribeSequenceOptions {
	opts := &DescribeSequenceOptions{
		name: r.name,
	}
	return opts
}

func (r sequenceDetailRow) convert() *SequenceDetail {
	return &SequenceDetail{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		SchemaName:    r.SchemaName,
		DatabaseName:  r.DatabaseName,
		NextValue:     r.NextValue,
		Interval:      r.Interval,
		Owner:         r.Owner,
		OwnerRoleType: r.OwnerRoleType,
		Comment:       r.Comment,
		Ordered:       r.Ordered == "Y",
	}
}

func (r *DropSequenceRequest) toOpts() *DropSequenceOptions {
	opts := &DropSequenceOptions{
		IfExists: r.IfExists,
		name:     r.name,
	}
	return opts
}
