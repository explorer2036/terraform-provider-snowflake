package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ Procedures = (*procedures)(nil)

type procedures struct {
	client *Client
}

func (v *procedures) CreateProcedureForJava(ctx context.Context, request *CreateProcedureForJavaProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateProcedureForJavaScript(ctx context.Context, request *CreateProcedureForJavaScriptProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateProcedureForPython(ctx context.Context, request *CreateProcedureForPythonProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateProcedureForScala(ctx context.Context, request *CreateProcedureForScalaProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateProcedureForSQL(ctx context.Context, request *CreateProcedureForSQLProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) Alter(ctx context.Context, request *AlterProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) Drop(ctx context.Context, request *DropProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) Show(ctx context.Context, request *ShowProcedureRequest) ([]Procedure, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[procedureRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[procedureRow, Procedure](dbRows)
	return resultList, nil
}

func (v *procedures) ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*Procedure, error) {
	// TODO: adjust request if e.g. LIKE is supported for the resource
	procedures, err := v.Show(ctx, NewShowProcedureRequest())
	if err != nil {
		return nil, err
	}
	return collections.FindOne(procedures, func(r Procedure) bool { return r.Name == id.Name() })
}

func (v *procedures) Describe(ctx context.Context, id SchemaObjectIdentifier) ([]ProcedureDetail, error) {
	opts := &DescribeProcedureOptions{
		name: id,
	}
	rows, err := validateAndQuery[procedureDetailRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertRows[procedureDetailRow, ProcedureDetail](rows), nil
}

func (r *CreateProcedureForJavaProcedureRequest) toOpts() *CreateProcedureForJavaProcedureOptions {
	opts := &CreateProcedureForJavaProcedureOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		RuntimeVersion: r.RuntimeVersion,

		Handler:                    r.Handler,
		ExternalAccessIntegrations: r.ExternalAccessIntegrations,
		Secrets:                    r.Secrets,
		TargetPath:                 r.TargetPath,
		NullInputBehavior:          r.NullInputBehavior,
		Comment:                    r.Comment,
		ExecuteAs:                  r.ExecuteAs,
		ProcedureDefinition:        r.ProcedureDefinition,
	}
	if r.Arguments != nil {
		s := make([]ProcedureArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = ProcedureArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	opts.Returns = ProcedureReturns{}
	if r.Returns.ResultDataType != nil {
		opts.Returns.ResultDataType = &ProcedureReturnsResultDataType{
			ResultDataType: r.Returns.ResultDataType.ResultDataType,
			Null:           r.Returns.ResultDataType.Null,
			NotNull:        r.Returns.ResultDataType.NotNull,
		}
	}
	if r.Returns.Table != nil {
		opts.Returns.Table = &ProcedureReturnsTable{}
		if r.Returns.Table.Columns != nil {
			s := make([]ProcedureColumn, len(r.Returns.Table.Columns))
			for i, v := range r.Returns.Table.Columns {
				s[i] = ProcedureColumn{
					ColumnName:     v.ColumnName,
					ColumnDataType: v.ColumnDataType,
				}
			}
			opts.Returns.Table.Columns = s
		}
	}
	if r.Packages != nil {
		s := make([]ProcedurePackage, len(r.Packages))
		for i, v := range r.Packages {
			s[i] = ProcedurePackage{
				Package: v.Package,
			}
		}
		opts.Packages = s
	}
	if r.Imports != nil {
		s := make([]ProcedureImport, len(r.Imports))
		for i, v := range r.Imports {
			s[i] = ProcedureImport{
				Import: v.Import,
			}
		}
		opts.Imports = s
	}
	return opts
}

func (r *CreateProcedureForJavaScriptProcedureRequest) toOpts() *CreateProcedureForJavaScriptProcedureOptions {
	opts := &CreateProcedureForJavaScriptProcedureOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		NullInputBehavior:   r.NullInputBehavior,
		Comment:             r.Comment,
		ExecuteAs:           r.ExecuteAs,
		ProcedureDefinition: r.ProcedureDefinition,
	}
	if r.Arguments != nil {
		s := make([]ProcedureArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = ProcedureArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &ProcedureJavascriptReturns{
			ResultDataType: r.Returns.ResultDataType,
			NotNull:        r.Returns.NotNull,
		}
	}
	return opts
}

func (r *CreateProcedureForPythonProcedureRequest) toOpts() *CreateProcedureForPythonProcedureOptions {
	opts := &CreateProcedureForPythonProcedureOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		RuntimeVersion: r.RuntimeVersion,

		Handler:                    r.Handler,
		ExternalAccessIntegrations: r.ExternalAccessIntegrations,
		Secrets:                    r.Secrets,
		NullInputBehavior:          r.NullInputBehavior,
		Comment:                    r.Comment,
		ExecuteAs:                  r.ExecuteAs,
		ProcedureDefinition:        r.ProcedureDefinition,
	}
	if r.Arguments != nil {
		s := make([]ProcedureArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = ProcedureArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &ProcedureReturns{}
		if r.Returns.ResultDataType != nil {
			opts.Returns.ResultDataType = &ProcedureReturnsResultDataType{
				ResultDataType: r.Returns.ResultDataType.ResultDataType,
				Null:           r.Returns.ResultDataType.Null,
				NotNull:        r.Returns.ResultDataType.NotNull,
			}
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &ProcedureReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]ProcedureColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = ProcedureColumn{
						ColumnName:     v.ColumnName,
						ColumnDataType: v.ColumnDataType,
					}
				}
				opts.Returns.Table.Columns = s
			}
		}
	}
	if r.Packages != nil {
		s := make([]ProcedurePackage, len(r.Packages))
		for i, v := range r.Packages {
			s[i] = ProcedurePackage{
				Package: v.Package,
			}
		}
		opts.Packages = s
	}
	if r.Imports != nil {
		s := make([]ProcedureImport, len(r.Imports))
		for i, v := range r.Imports {
			s[i] = ProcedureImport{
				Import: v.Import,
			}
		}
		opts.Imports = s
	}
	return opts
}

func (r *CreateProcedureForScalaProcedureRequest) toOpts() *CreateProcedureForScalaProcedureOptions {
	opts := &CreateProcedureForScalaProcedureOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		RuntimeVersion: r.RuntimeVersion,

		Handler:             r.Handler,
		TargetPath:          r.TargetPath,
		NullInputBehavior:   r.NullInputBehavior,
		Comment:             r.Comment,
		ExecuteAs:           r.ExecuteAs,
		ProcedureDefinition: r.ProcedureDefinition,
	}
	if r.Arguments != nil {
		s := make([]ProcedureArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = ProcedureArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &ProcedureReturns{}
		if r.Returns.ResultDataType != nil {
			opts.Returns.ResultDataType = &ProcedureReturnsResultDataType{
				ResultDataType: r.Returns.ResultDataType.ResultDataType,
				Null:           r.Returns.ResultDataType.Null,
				NotNull:        r.Returns.ResultDataType.NotNull,
			}
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &ProcedureReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]ProcedureColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = ProcedureColumn{
						ColumnName:     v.ColumnName,
						ColumnDataType: v.ColumnDataType,
					}
				}
				opts.Returns.Table.Columns = s
			}
		}
	}
	if r.Packages != nil {
		s := make([]ProcedurePackage, len(r.Packages))
		for i, v := range r.Packages {
			s[i] = ProcedurePackage{
				Package: v.Package,
			}
		}
		opts.Packages = s
	}
	if r.Imports != nil {
		s := make([]ProcedureImport, len(r.Imports))
		for i, v := range r.Imports {
			s[i] = ProcedureImport{
				Import: v.Import,
			}
		}
		opts.Imports = s
	}
	return opts
}

func (r *CreateProcedureForSQLProcedureRequest) toOpts() *CreateProcedureForSQLProcedureOptions {
	opts := &CreateProcedureForSQLProcedureOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants: r.CopyGrants,

		NullInputBehavior:   r.NullInputBehavior,
		Comment:             r.Comment,
		ExecuteAs:           r.ExecuteAs,
		ProcedureDefinition: r.ProcedureDefinition,
	}
	if r.Arguments != nil {
		s := make([]ProcedureArgument, len(r.Arguments))
		for i, v := range r.Arguments {
			s[i] = ProcedureArgument{
				ArgName:      v.ArgName,
				ArgDataType:  v.ArgDataType,
				DefaultValue: v.DefaultValue,
			}
		}
		opts.Arguments = s
	}
	if r.Returns != nil {
		opts.Returns = &ProcedureSQLReturns{
			NotNull: r.Returns.NotNull,
		}
		if r.Returns.ResultDataType != nil {
			opts.Returns.ResultDataType = &ProcedureReturnsResultDataType{
				ResultDataType: r.Returns.ResultDataType.ResultDataType,
			}
		}
		if r.Returns.Table != nil {
			opts.Returns.Table = &ProcedureReturnsTable{}
			if r.Returns.Table.Columns != nil {
				s := make([]ProcedureColumn, len(r.Returns.Table.Columns))
				for i, v := range r.Returns.Table.Columns {
					s[i] = ProcedureColumn{
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

func (r *AlterProcedureRequest) toOpts() *AlterProcedureOptions {
	opts := &AlterProcedureOptions{
		IfExists:          r.IfExists,
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
		RenameTo:          r.RenameTo,
		SetComment:        r.SetComment,
		SetLogLevel:       r.SetLogLevel,
		SetTraceLevel:     r.SetTraceLevel,
		UnsetComment:      r.UnsetComment,
		SetTags:           r.SetTags,
		UnsetTags:         r.UnsetTags,
		ExecuteAs:         r.ExecuteAs,
	}
	return opts
}

func (r *DropProcedureRequest) toOpts() *DropProcedureOptions {
	opts := &DropProcedureOptions{
		IfExists:          r.IfExists,
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
	}
	return opts
}

func (r *ShowProcedureRequest) toOpts() *ShowProcedureOptions {
	opts := &ShowProcedureOptions{
		Like: r.Like,
		In:   r.In,
	}
	return opts
}

func (r procedureRow) convert() *Procedure {
	// TODO: Mapping
	return &Procedure{}
}

func (r *DescribeProcedureRequest) toOpts() *DescribeProcedureOptions {
	opts := &DescribeProcedureOptions{
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
	}
	return opts
}

func (r procedureDetailRow) convert() *ProcedureDetail {
	// TODO: Mapping
	return &ProcedureDetail{}
}
