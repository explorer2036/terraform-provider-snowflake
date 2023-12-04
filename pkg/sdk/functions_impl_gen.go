package sdk

import (
	"context"
)

var _ Functions = (*functions)(nil)

type functions struct {
	client *Client
}

func (v *functions) CreateForJava(ctx context.Context, request *CreateForJavaFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) CreateForJavascript(ctx context.Context, request *CreateForJavascriptFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) CreateForPython(ctx context.Context, request *CreateForPythonFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) CreateForScala(ctx context.Context, request *CreateForScalaFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) CreateForSQL(ctx context.Context, request *CreateForSQLFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) Alter(ctx context.Context, request *AlterFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) Drop(ctx context.Context, request *DropFunctionRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *functions) Show(ctx context.Context, request *ShowFunctionRequest) ([]Function, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[functionRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[functionRow, Function](dbRows)
	return resultList, nil
}

func (v *functions) Describe(ctx context.Context, id SchemaObjectIdentifier) ([]FunctionDetail, error) {
	opts := &DescribeFunctionOptions{
		name: id,
	}
	rows, err := validateAndQuery[functionDetailRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertRows[functionDetailRow, FunctionDetail](rows), nil
}

func (r *CreateForJavaFunctionRequest) toOpts() *CreateForJavaFunctionOptions {
	opts := &CreateForJavaFunctionOptions{
		OrReplace:   r.OrReplace,
		Temporary:   r.Temporary,
		Secure:      r.Secure,
		IfNotExists: r.IfNotExists,
		name:        r.name,

		CopyGrants: r.CopyGrants,

		ReturnNullValues:      r.ReturnNullValues,
		NullInputBehavior:     r.NullInputBehavior,
		ReturnResultsBehavior: r.ReturnResultsBehavior,
		RuntimeVersion:        r.RuntimeVersion,
		Comment:               r.Comment,

		Handler:                    r.Handler,
		ExternalAccessIntegrations: r.ExternalAccessIntegrations,
		Secrets:                    r.Secrets,
		TargetPath:                 r.TargetPath,
		FunctionDefinition:         r.FunctionDefinition,
	}
	if r.Arguments != nil {
		s := make([]FunctionArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = FunctionArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &FunctionReturns{
			ResultDataType: r.Returns.ResultDataType,
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &FunctionReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]FunctionColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = FunctionColumn{
						ColumnName:     v.ColumnName,
						ColumnDataType: v.ColumnDataType,
					}
				}
				opts.Returns.Table.Columns = s
			}
		}
	}
	if r.Imports != nil {
		s := make([]FunctionImports, len(r.Imports))
		for i, v := range r.Imports {
			s[i] = FunctionImports{
				Import: v.Import,
			}
		}
		opts.Imports = s
	}
	if r.Packages != nil {
		s := make([]FunctionPackages, len(r.Packages))
		for i, v := range r.Packages {
			s[i] = FunctionPackages{
				Package: v.Package,
			}
		}
		opts.Packages = s
	}
	return opts
}

func (r *CreateForJavascriptFunctionRequest) toOpts() *CreateForJavascriptFunctionOptions {
	opts := &CreateForJavascriptFunctionOptions{
		OrReplace: r.OrReplace,
		Temporary: r.Temporary,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		ReturnNullValues:      r.ReturnNullValues,
		NullInputBehavior:     r.NullInputBehavior,
		ReturnResultsBehavior: r.ReturnResultsBehavior,
		Comment:               r.Comment,
		FunctionDefinition:    r.FunctionDefinition,
	}
	if r.Arguments != nil {
		s := make([]FunctionArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = FunctionArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &FunctionReturns{
			ResultDataType: r.Returns.ResultDataType,
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &FunctionReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]FunctionColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = FunctionColumn{
						ColumnName:     v.ColumnName,
						ColumnDataType: v.ColumnDataType,
					}
				}
				opts.Returns.Table.Columns = s
			}
		}
	}
	return opts
}

func (r *CreateForPythonFunctionRequest) toOpts() *CreateForPythonFunctionOptions {
	opts := &CreateForPythonFunctionOptions{
		OrReplace:   r.OrReplace,
		Temporary:   r.Temporary,
		Secure:      r.Secure,
		IfNotExists: r.IfNotExists,
		name:        r.name,

		CopyGrants: r.CopyGrants,

		ReturnNullValues:      r.ReturnNullValues,
		NullInputBehavior:     r.NullInputBehavior,
		ReturnResultsBehavior: r.ReturnResultsBehavior,
		RuntimeVersion:        r.RuntimeVersion,
		Comment:               r.Comment,

		Handler:                    r.Handler,
		ExternalAccessIntegrations: r.ExternalAccessIntegrations,
		Secrets:                    r.Secrets,
		FunctionDefinition:         r.FunctionDefinition,
	}
	if r.Arguments != nil {
		s := make([]FunctionArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = FunctionArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &FunctionReturns{
			ResultDataType: r.Returns.ResultDataType,
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &FunctionReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]FunctionColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = FunctionColumn{
						ColumnName:     v.ColumnName,
						ColumnDataType: v.ColumnDataType,
					}
				}
				opts.Returns.Table.Columns = s
			}
		}
	}
	if r.Imports != nil {
		s := make([]FunctionImports, len(r.Imports))
		for i, v := range r.Imports {
			s[i] = FunctionImports{
				Import: v.Import,
			}
		}
		opts.Imports = s
	}
	if r.Packages != nil {
		s := make([]FunctionPackages, len(r.Packages))
		for i, v := range r.Packages {
			s[i] = FunctionPackages{
				Package: v.Package,
			}
		}
		opts.Packages = s
	}
	return opts
}

func (r *CreateForScalaFunctionRequest) toOpts() *CreateForScalaFunctionOptions {
	opts := &CreateForScalaFunctionOptions{
		OrReplace:   r.OrReplace,
		Temporary:   r.Temporary,
		Secure:      r.Secure,
		IfNotExists: r.IfNotExists,
		name:        r.name,

		CopyGrants:            r.CopyGrants,
		ResultDataType:        r.ResultDataType,
		ReturnNullValues:      r.ReturnNullValues,
		NullInputBehavior:     r.NullInputBehavior,
		ReturnResultsBehavior: r.ReturnResultsBehavior,
		RuntimeVersion:        r.RuntimeVersion,
		Comment:               r.Comment,

		Handler:            r.Handler,
		TargetPath:         r.TargetPath,
		FunctionDefinition: r.FunctionDefinition,
	}
	if r.Arguments != nil {
		s := make([]FunctionArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = FunctionArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Imports != nil {
		s := make([]FunctionImports, len(r.Imports))
		for i, v := range r.Imports {
			s[i] = FunctionImports{
				Import: v.Import,
			}
		}
		opts.Imports = s
	}
	if r.Packages != nil {
		s := make([]FunctionPackages, len(r.Packages))
		for i, v := range r.Packages {
			s[i] = FunctionPackages{
				Package: v.Package,
			}
		}
		opts.Packages = s
	}
	return opts
}

func (r *CreateForSQLFunctionRequest) toOpts() *CreateForSQLFunctionOptions {
	opts := &CreateForSQLFunctionOptions{
		OrReplace: r.OrReplace,
		Temporary: r.Temporary,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		ReturnNullValues:      r.ReturnNullValues,
		ReturnResultsBehavior: r.ReturnResultsBehavior,
		Memoizable:            r.Memoizable,
		Comment:               r.Comment,
		FunctionDefinition:    r.FunctionDefinition,
	}
	if r.Arguments != nil {
		s := make([]FunctionArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = FunctionArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &FunctionReturns{
			ResultDataType: r.Returns.ResultDataType,
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &FunctionReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]FunctionColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = FunctionColumn{
						ColumnName:     v.ColumnName,
						ColumnDataType: v.ColumnDataType,
					}
				}
				opts.Returns.Table.Columns = s
			}
		}
	}
	return opts
}

func (r *AlterFunctionRequest) toOpts() *AlterFunctionOptions {
	opts := &AlterFunctionOptions{
		IfExists:          r.IfExists,
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
		RenameTo:          r.RenameTo,
		SetComment:        r.SetComment,
		SetLogLevel:       r.SetLogLevel,
		SetTraceLevel:     r.SetTraceLevel,
		SetSecure:         r.SetSecure,
		UnsetSecure:       r.UnsetSecure,
		UnsetLogLevel:     r.UnsetLogLevel,
		UnsetTraceLevel:   r.UnsetTraceLevel,
		UnsetComment:      r.UnsetComment,
		SetTags:           r.SetTags,
		UnsetTags:         r.UnsetTags,
	}
	return opts
}

func (r *DropFunctionRequest) toOpts() *DropFunctionOptions {
	opts := &DropFunctionOptions{
		IfExists:          r.IfExists,
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
	}
	return opts
}

func (r *ShowFunctionRequest) toOpts() *ShowFunctionOptions {
	opts := &ShowFunctionOptions{
		Like: r.Like,
		In:   r.In,
	}
	return opts
}

func (r functionRow) convert() *Function {
	// TODO: Mapping
	return &Function{}
}

func (r *DescribeFunctionRequest) toOpts() *DescribeFunctionOptions {
	opts := &DescribeFunctionOptions{
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
	}
	return opts
}

func (r functionDetailRow) convert() *FunctionDetail {
	// TODO: Mapping
	return &FunctionDetail{}
}
