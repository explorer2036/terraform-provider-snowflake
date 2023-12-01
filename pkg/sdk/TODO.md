TODO

1. we could validate that TARGET_PATH  can be set only if AS is also set (two variants difference)
2. update function ShowByID in procedures_impl_gen.go
3. add validation for RenameTo

4. procedures_impl_gen.go

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

5. procedures_gen.go

SetTags           []TagAssociation        `ddl:"keyword" sql:"SET TAG"`
UnsetTags         []ObjectIdentifier      `ddl:"keyword" sql:"UNSET TAG"`


6. procedures_validations_gen.go

<!-- CreateProcedureForJavaProcedureOptions -->
if opts.ProcedureDefinition == nil && opts.TargetPath != nil {
	return fmt.Errorf("TARGET_PATH must be nil when AS is nil")
}
