func (r applicationPackageRow) convert() *ApplicationPackage {
	e := &ApplicationPackage{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		IsDefault:     r.IsDefault == "Y",
		IsCurrent:     r.IsCurrent == "Y",
		Distribution:  r.Distribution,
		Owner:         r.Owner,
		Comment:       r.Comment,
		RetentionTime: r.RetentionTime,
		Options:       r.Options,
	}
	if r.DroppedOn.Valid {
		e.DroppedOn = r.DroppedOn.String
	}
	if r.ApplicationClass.Valid {
		e.ApplicationClass = r.ApplicationClass.String
	}
	return e
}

func (v *applicationPackages) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*ApplicationPackage, error) {
	request := NewShowApplicationPackageRequest().WithLike(&Like{String(id.Name())})
	applicationPackages, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(applicationPackages, func(r ApplicationPackage) bool { return r.Name == id.Name() })
}
