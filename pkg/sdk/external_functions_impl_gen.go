package sdk

import (
	"context"
)

var _ ExternalFunctions = (*externalFunctions)(nil)

type externalFunctions struct {
	client *Client
}

func (v *externalFunctions) Create(ctx context.Context, request *CreateExternalFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateExternalFunctionRequest) toOpts() *CreateExternalFunctionOptions {
	opts := &CreateExternalFunctionOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		ResultDataType:        r.ResultDataType,
		ReturnNullValues:      r.ReturnNullValues,
		NullInputBehavior:     r.NullInputBehavior,
		ReturnResultsBehavior: r.ReturnResultsBehavior,
		Comment:               r.Comment,
		ApiIntegration:        r.ApiIntegration,

		MaxBatchRows:       r.MaxBatchRows,
		Compression:        r.Compression,
		RequestTranslator:  r.RequestTranslator,
		ResponseTranslator: r.ResponseTranslator,
		As:                 r.As,
	}
	if r.Arguments != nil {
		s := make([]ExternalFunctionArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = ExternalFunctionArgument{
				ArgName:     v.ArgName,
				ArgDataType: v.ArgDataType,
			}
		}
		opts.Arguments = s
	}
	if r.Headers != nil {
		s := make([]ExternalFunctionHeader, len(r.Headers))
		for i, v := range r.Headers {
			s[i] = ExternalFunctionHeader{
				Name:  v.Name,
				Value: v.Value,
			}
		}
		opts.Headers = s
	}
	if r.ContextHeaders != nil {
		s := make([]ExternalFunctionContextHeader, len(r.ContextHeaders))
		for i, v := range r.ContextHeaders {
			s[i] = ExternalFunctionContextHeader{
				ContextFunction: v.ContextFunction,
			}
		}
		opts.ContextHeaders = s
	}
	return opts
}
