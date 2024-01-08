func (v *applications) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*Application, error) {
	request := NewShowApplicationRequest().WithLike(&Like{String(id.Name())})
	applications, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(applications, func(r Application) bool { return r.Name == id.Name() })
}

func (r applicationDetailRow) convert() *ApplicationDetail {
	return &ApplicationDetail{
		Property: r.Property,
		Value:    r.Value,
	}
}

func (r applicationRow) convert() *Application {
	return &Application{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		IsDefault:     r.IsDefault == "Y",
		IsCurrent:     r.IsCurrent == "Y",
		SourceType:    r.SourceType,
		Source:        r.Source,
		Owner:         r.Owner,
		Comment:       r.Comment,
		Version:       r.Version,
		Label:         r.Label,
		Patch:         r.Patch,
		Options:       r.Options,
		RetentionTime: r.RetentionTime,
	}
}
