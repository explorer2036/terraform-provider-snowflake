## functions_impl_gen.go

func (v *functions) ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*Function, error) {
	request := NewShowFunctionRequest().WithIn(&In{Database: NewAccountObjectIdentifier(id.DatabaseName())}).WithLike(&Like{String(id.Name())})
	functions, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(functions, func(r Function) bool { return r.Name == id.Name() })
}

func (r functionRow) convert() *Function {
	e := &Function{
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
		IsExternalFunction: r.IsExternalFunction == "Y",
		Language:           r.Language,
	}
	if r.IsSecure.Valid {
		e.IsSecure = r.IsSecure.String == "Y"
	}
	if r.IsMemoizable.Valid {
		e.IsMemoizable = r.IsMemoizable.String == "Y"
	}
	return e
}

func (r functionDetailRow) convert() *FunctionDetail {
	return &FunctionDetail{
		Property: r.Property,
		Value:    r.Value,
	}
}

func (v *functions) Describe(ctx context.Context, request *DescribeFunctionRequest) ([]FunctionDetail, error) {
	opts := request.toOpts()
	rows, err := validateAndQuery[functionDetailRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertRows[functionDetailRow, FunctionDetail](rows), nil
}

## functions_gen.go

Describe(ctx context.Context, request *DescribeFunctionRequest) ([]FunctionDetail, error)

type DropFunctionOptions struct {
	drop              bool                   `ddl:"static" sql:"DROP"`
	function          bool                   `ddl:"static" sql:"FUNCTION"`
	IfExists          *bool                  `ddl:"keyword" sql:"IF EXISTS"`
	name              SchemaObjectIdentifier `ddl:"identifier"`
	ArgumentDataTypes []DataType             `ddl:"keyword,must_parentheses"`
}

## functions_validations_gen.go

<!-- CreateForJavaFunctionOptions -->
if opts.FunctionDefinition == nil {
	if opts.TargetPath != nil {
		errs = append(errs, NewError("TARGET_PATH must be nil when AS is nil"))
	}
	if len(opts.Packages) > 0 {
		errs = append(errs, NewError("PACKAGES must be empty when AS is nil"))
	}
	if len(opts.Imports) == 0 {
		errs = append(errs, NewError("IMPORTS must not be empty when AS is nil"))
	}
}

<!-- CreateForScalaFunctionOptions -->
if opts.FunctionDefinition == nil {
	if opts.TargetPath != nil {
		errs = append(errs, NewError("TARGET_PATH must be nil when AS is nil"))
	}
	if len(opts.Packages) > 0 {
		errs = append(errs, NewError("PACKAGES must be empty when AS is nil"))
	}
	if len(opts.Imports) == 0 {
		errs = append(errs, NewError("IMPORTS must not be empty when AS is nil"))
	}
}

<!-- CreateForPythonFunctionOptions -->
if opts.FunctionDefinition == nil {
	if len(opts.Imports) == 0 {
		errs = append(errs, NewError("IMPORTS must not be empty when AS is nil"))
	}
}
