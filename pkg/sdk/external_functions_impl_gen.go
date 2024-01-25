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

func (v *externalFunctions) Alter(ctx context.Context, request *AlterExternalFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *externalFunctions) Show(ctx context.Context, request *ShowExternalFunctionRequest) ([]ExternalFunction, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[externalFunctionRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[externalFunctionRow, ExternalFunction](dbRows)
	return resultList, nil
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

func (r *AlterExternalFunctionRequest) toOpts() *AlterExternalFunctionOptions {
	opts := &AlterExternalFunctionOptions{
		IfExists:          r.IfExists,
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
	}
	if r.Set != nil {
		opts.Set = &ExternalFunctionSet{
			ApiIntegration: r.Set.ApiIntegration,

			MaxBatchRows:       r.Set.MaxBatchRows,
			Compression:        r.Set.Compression,
			RequestTranslator:  r.Set.RequestTranslator,
			ResponseTranslator: r.Set.ResponseTranslator,
		}
		if r.Set.Headers != nil {
			s := make([]ExternalFunctionHeader, len(r.Set.Headers))
			for i, v := range r.Set.Headers {
				s[i] = ExternalFunctionHeader{
					Name:  v.Name,
					Value: v.Value,
				}
			}
			opts.Set.Headers = s
		}
		if r.Set.ContextHeaders != nil {
			s := make([]ExternalFunctionContextHeader, len(r.Set.ContextHeaders))
			for i, v := range r.Set.ContextHeaders {
				s[i] = ExternalFunctionContextHeader{
					ContextFunction: v.ContextFunction,
				}
			}
			opts.Set.ContextHeaders = s
		}
	}
	if r.Unset != nil {
		opts.Unset = &ExternalFunctionUnset{
			Comment:            r.Unset.Comment,
			Headers:            r.Unset.Headers,
			ContextHeaders:     r.Unset.ContextHeaders,
			MaxBatchRows:       r.Unset.MaxBatchRows,
			Compression:        r.Unset.Compression,
			Secure:             r.Unset.Secure,
			RequestTranslator:  r.Unset.RequestTranslator,
			ResponseTranslator: r.Unset.ResponseTranslator,
		}
	}
	return opts
}

func (r *ShowExternalFunctionRequest) toOpts() *ShowExternalFunctionOptions {
	opts := &ShowExternalFunctionOptions{
		Like: r.Like,
	}
	return opts
}

func (r externalFunctionRow) convert() *ExternalFunction {
	// TODO: Mapping
	return &ExternalFunction{}
}
