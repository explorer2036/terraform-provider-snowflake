package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ Procedures = (*procedures)(nil)

type procedures struct {
	client *Client
}

func (v *procedures) CreateForJava(ctx context.Context, request *CreateForJavaProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateForJavaScript(ctx context.Context, request *CreateForJavaScriptProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateForPython(ctx context.Context, request *CreateForPythonProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateForScala(ctx context.Context, request *CreateForScalaProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateForSQL(ctx context.Context, request *CreateForSQLProcedureRequest) error {
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
	request := NewShowProcedureRequest().WithIn(&In{Database: NewAccountObjectIdentifier(id.DatabaseName())}).WithLike(&Like{String(id.Name())})
	procedures, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(procedures, func(r Procedure) bool { return r.Name == id.Name() })
}

func (v *procedures) Describe(ctx context.Context, request *DescribeProcedureRequest) ([]ProcedureDetail, error) {
	opts := request.toOpts()
	rows, err := validateAndQuery[procedureDetailRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertRows[procedureDetailRow, ProcedureDetail](rows), nil
}

func (v *procedures) Call(ctx context.Context, request *CallProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateAndCallForJava(ctx context.Context, request *CreateAndCallForJavaProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateAndCallForScala(ctx context.Context, request *CreateAndCallForScalaProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateAndCallForJavaScript(ctx context.Context, request *CreateAndCallForJavaScriptProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateAndCallForPython(ctx context.Context, request *CreateAndCallForPythonProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *procedures) CreateAndCallForSQL(ctx context.Context, request *CreateAndCallForSQLProcedureRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateForJavaProcedureRequest) toOpts() *CreateForJavaProcedureOptions {
	opts := &CreateForJavaProcedureOptions{
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

func (r *CreateForJavaScriptProcedureRequest) toOpts() *CreateForJavaScriptProcedureOptions {
	opts := &CreateForJavaScriptProcedureOptions{
		OrReplace: r.OrReplace,
		Secure:    r.Secure,
		name:      r.name,

		CopyGrants:          r.CopyGrants,
		ResultDataType:      r.ResultDataType,
		NotNull:             r.NotNull,
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
	return opts
}

func (r *CreateForPythonProcedureRequest) toOpts() *CreateForPythonProcedureOptions {
	opts := &CreateForPythonProcedureOptions{
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

func (r *CreateForScalaProcedureRequest) toOpts() *CreateForScalaProcedureOptions {
	opts := &CreateForScalaProcedureOptions{
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

func (r *CreateForSQLProcedureRequest) toOpts() *CreateForSQLProcedureOptions {
	opts := &CreateForSQLProcedureOptions{
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
	opts.Returns = ProcedureSQLReturns{
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
	e := &Procedure{
		CreatedOn:          r.CreatedOn,
		Name:               r.Name,
		SchemaName:         r.SchemaName,
		IsBuiltin:          r.IsBuiltin == "Y",
		IsAggregate:        r.IsAggregate == "Y",
		IsAnsi:             r.IsAnsi == "Y",
		MinNumArguments:    r.MinNumArguments,
		MaxNumArguments:    r.MaxNumArguments,
		Arguments:          r.Arguments,
		Description:        r.Description,
		CatalogName:        r.CatalogName,
		IsTableFunction:    r.IsTableFunction == "Y",
		ValidForClustering: r.ValidForClustering == "Y",
	}
	if r.IsSecure.Valid {
		e.IsSecure = r.IsSecure.String == "Y"
	}
	return e
}

func (r *DescribeProcedureRequest) toOpts() *DescribeProcedureOptions {
	opts := &DescribeProcedureOptions{
		name:              r.name,
		ArgumentDataTypes: r.ArgumentDataTypes,
	}
	return opts
}

func (r procedureDetailRow) convert() *ProcedureDetail {
	return &ProcedureDetail{
		Property: r.Property,
		Value:    r.Value,
	}
}

func (r *CallProcedureRequest) toOpts() *CallProcedureOptions {
	opts := &CallProcedureOptions{
		name: r.name,

		ScriptingVariable: r.ScriptingVariable,
	}
	if r.Positions != nil {
		s := make([]ProcedureCallArgumentPosition, len(r.Positions))
		for i, v := range r.Positions {
			s[i] = ProcedureCallArgumentPosition{
				Position: v.Position,
			}
		}
		opts.Positions = s
	}
	if r.Names != nil {
		s := make([]ProcedureCallArgumentName, len(r.Names))
		for i, v := range r.Names {
			s[i] = ProcedureCallArgumentName{
				Name:     v.Name,
				Position: v.Position,
			}
		}
		opts.Names = s
	}
	return opts
}

func (r *CreateAndCallForJavaProcedureRequest) toOpts() *CreateAndCallForJavaProcedureOptions {
	opts := &CreateAndCallForJavaProcedureOptions{
		Name: r.Name,

		RuntimeVersion: r.RuntimeVersion,

		Handler:             r.Handler,
		NullInputBehavior:   r.NullInputBehavior,
		ProcedureDefinition: r.ProcedureDefinition,

		ProcedureName: r.ProcedureName,

		ScriptingVariable: r.ScriptingVariable,
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
	if r.WithClauses != nil {
		s := make([]ProcedureWithClause, len(r.WithClauses))
		for i, v := range r.WithClauses {
			s[i] = ProcedureWithClause{
				CteName:    v.CteName,
				CteColumns: v.CteColumns,
				Statement:  v.Statement,
			}
		}
		opts.WithClauses = s
	}
	if r.Positions != nil {
		s := make([]ProcedureCallArgumentPosition, len(r.Positions))
		for i, v := range r.Positions {
			s[i] = ProcedureCallArgumentPosition{
				Position: v.Position,
			}
		}
		opts.Positions = s
	}
	if r.Names != nil {
		s := make([]ProcedureCallArgumentName, len(r.Names))
		for i, v := range r.Names {
			s[i] = ProcedureCallArgumentName{
				Name:     v.Name,
				Position: v.Position,
			}
		}
		opts.Names = s
	}
	return opts
}

func (r *CreateAndCallForScalaProcedureRequest) toOpts() *CreateAndCallForScalaProcedureOptions {
	opts := &CreateAndCallForScalaProcedureOptions{
		Name: r.Name,

		RuntimeVersion: r.RuntimeVersion,

		Handler:             r.Handler,
		NullInputBehavior:   r.NullInputBehavior,
		ProcedureDefinition: r.ProcedureDefinition,

		ProcedureName: r.ProcedureName,

		ScriptingVariable: r.ScriptingVariable,
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
	if r.WithClauses != nil {
		s := make([]ProcedureWithClause, len(r.WithClauses))
		for i, v := range r.WithClauses {
			s[i] = ProcedureWithClause{
				CteName:    v.CteName,
				CteColumns: v.CteColumns,
				Statement:  v.Statement,
			}
		}
		opts.WithClauses = s
	}
	if r.Positions != nil {
		s := make([]ProcedureCallArgumentPosition, len(r.Positions))
		for i, v := range r.Positions {
			s[i] = ProcedureCallArgumentPosition{
				Position: v.Position,
			}
		}
		opts.Positions = s
	}
	if r.Names != nil {
		s := make([]ProcedureCallArgumentName, len(r.Names))
		for i, v := range r.Names {
			s[i] = ProcedureCallArgumentName{
				Name:     v.Name,
				Position: v.Position,
			}
		}
		opts.Names = s
	}
	return opts
}

func (r *CreateAndCallForJavaScriptProcedureRequest) toOpts() *CreateAndCallForJavaScriptProcedureOptions {
	opts := &CreateAndCallForJavaScriptProcedureOptions{
		Name: r.Name,

		ResultDataType:      r.ResultDataType,
		NotNull:             r.NotNull,
		NullInputBehavior:   r.NullInputBehavior,
		ProcedureDefinition: r.ProcedureDefinition,

		ProcedureName: r.ProcedureName,

		ScriptingVariable: r.ScriptingVariable,
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
	if r.WithClauses != nil {
		s := make([]ProcedureWithClause, len(r.WithClauses))
		for i, v := range r.WithClauses {
			s[i] = ProcedureWithClause{
				CteName:    v.CteName,
				CteColumns: v.CteColumns,
				Statement:  v.Statement,
			}
		}
		opts.WithClauses = s
	}
	if r.Positions != nil {
		s := make([]ProcedureCallArgumentPosition, len(r.Positions))
		for i, v := range r.Positions {
			s[i] = ProcedureCallArgumentPosition{
				Position: v.Position,
			}
		}
		opts.Positions = s
	}
	if r.Names != nil {
		s := make([]ProcedureCallArgumentName, len(r.Names))
		for i, v := range r.Names {
			s[i] = ProcedureCallArgumentName{
				Name:     v.Name,
				Position: v.Position,
			}
		}
		opts.Names = s
	}
	return opts
}

func (r *CreateAndCallForPythonProcedureRequest) toOpts() *CreateAndCallForPythonProcedureOptions {
	opts := &CreateAndCallForPythonProcedureOptions{
		Name: r.Name,

		RuntimeVersion: r.RuntimeVersion,

		Handler:             r.Handler,
		NullInputBehavior:   r.NullInputBehavior,
		ProcedureDefinition: r.ProcedureDefinition,

		ProcedureName: r.ProcedureName,

		ScriptingVariable: r.ScriptingVariable,
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
	if r.WithClauses != nil {
		s := make([]ProcedureWithClause, len(r.WithClauses))
		for i, v := range r.WithClauses {
			s[i] = ProcedureWithClause{
				CteName:    v.CteName,
				CteColumns: v.CteColumns,
				Statement:  v.Statement,
			}
		}
		opts.WithClauses = s
	}
	if r.Positions != nil {
		s := make([]ProcedureCallArgumentPosition, len(r.Positions))
		for i, v := range r.Positions {
			s[i] = ProcedureCallArgumentPosition{
				Position: v.Position,
			}
		}
		opts.Positions = s
	}
	if r.Names != nil {
		s := make([]ProcedureCallArgumentName, len(r.Names))
		for i, v := range r.Names {
			s[i] = ProcedureCallArgumentName{
				Name:     v.Name,
				Position: v.Position,
			}
		}
		opts.Names = s
	}
	return opts
}

func (r *CreateAndCallForSQLProcedureRequest) toOpts() *CreateAndCallForSQLProcedureOptions {
	opts := &CreateAndCallForSQLProcedureOptions{
		Name: r.Name,

		NullInputBehavior:   r.NullInputBehavior,
		ProcedureDefinition: r.ProcedureDefinition,

		ProcedureName: r.ProcedureName,

		ScriptingVariable: r.ScriptingVariable,
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
	if r.WithClauses != nil {
		s := make([]ProcedureWithClause, len(r.WithClauses))
		for i, v := range r.WithClauses {
			s[i] = ProcedureWithClause{
				CteName:    v.CteName,
				CteColumns: v.CteColumns,
				Statement:  v.Statement,
			}
		}
		opts.WithClauses = s
	}
	if r.Positions != nil {
		s := make([]ProcedureCallArgumentPosition, len(r.Positions))
		for i, v := range r.Positions {
			s[i] = ProcedureCallArgumentPosition{
				Position: v.Position,
			}
		}
		opts.Positions = s
	}
	if r.Names != nil {
		s := make([]ProcedureCallArgumentName, len(r.Names))
		for i, v := range r.Names {
			s[i] = ProcedureCallArgumentName{
				Name:     v.Name,
				Position: v.Position,
			}
		}
		opts.Names = s
	}
	return opts
}
