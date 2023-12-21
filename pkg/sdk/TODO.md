## procedures_impl_gen.go

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

func (r procedureDetailRow) convert() *ProcedureDetail {
	return &ProcedureDetail{
		Property: r.Property,
		Value:    r.Value,
	}
}

## procedures_validations_gen.go

<!-- CreateForJavaProcedureOptions and CreateForScalaProcedureOptions-->
if opts.ProcedureDefinition == nil && opts.TargetPath != nil {
	errs = append(errs, NewError("TARGET_PATH must be nil when AS is nil"))
}

<!-- Call and CreateAndCall functions -->
if valueSet(opts.Positions) && valueSet(opts.Names) {
	errs = append(errs, errOneOf("CallProcedureOptions", "Positions", "Names"))
}

## procedures_gen.go

Describe(ctx context.Context, request *DescribeProcedureRequest) ([]ProcedureDetail, error)
