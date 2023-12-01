## procedures_impl_gen.go

func (v *procedures) ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*Procedure, error) {
	request := NewShowProcedureRequest().WithIn(&In{Database: NewAccountObjectIdentifier(id.DatabaseName())}).WithLike(&Like{String(id.Name())})
	procedures, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(procedures, func(r Procedure) bool { return r.Name == id.Name() })
}

func (r procedureRow) convert() *Procedure {
	return &Procedure{
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
		IsSecure:           r.IsSecure == "Y",
	}
}

func (r procedureDetailRow) convert() *ProcedureDetail {
	return &ProcedureDetail{
		Property: r.Property,
		Value:    r.Value,
	}
}

## procedures_validations_gen.go

<!-- CreateProcedureForJavaProcedureOptions and CreateProcedureForScalaProcedureOptions-->
if opts.ProcedureDefinition == nil && opts.TargetPath != nil {
	errs = append(errs, errors.New("TARGET_PATH must be nil when AS is nil"))
}
