func (v *eventTables) ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*EventTable, error) {
	request := NewShowEventTableRequest().WithIn(&In{Database: NewAccountObjectIdentifier(id.DatabaseName())}).WithLike(&Like{String(id.Name())})
	eventTables, err := v.Show(ctx, request)
	if err != nil {
		return nil, err
	}
	return collections.FindOne(eventTables, func(r EventTable) bool { return r.Name == id.Name() })
}

func (r eventTableRow) convert() *EventTable {
	t := &EventTable{
		CreatedOn:    r.CreatedOn,
		Name:         r.Name,
		DatabaseName: r.DatabaseName,
		SchemaName:   r.SchemaName,
	}
	if r.Owner.Valid {
		t.Owner = r.Owner.String
	}
	if r.Comment.Valid {
		t.Comment = r.Comment.String
	}
	if r.OwnerRoleType.Valid {
		t.OwnerRoleType = r.OwnerRoleType.String
	}
	return t
}

func (r eventTableDetailsRow) convert() *EventTableDetails {
	return &EventTableDetails{
		Name:    r.Name,
		Kind:    r.Kind,
		Comment: r.Comment,
	}
}


if r.AddRowAccessPolicy != nil {
	opts.AddRowAccessPolicy = &EventTableAddRowAccessPolicy{
		RowAccessPolicy: r.AddRowAccessPolicy.RowAccessPolicy,
		On:              r.AddRowAccessPolicy.On,
	}
}
if r.DropRowAccessPolicy != nil {
	opts.DropRowAccessPolicy = &EventTableDropRowAccessPolicy{
		RowAccessPolicy: r.DropRowAccessPolicy.RowAccessPolicy,
	}
}
if r.DropAndAddRowAccessPolicy != nil {
	opts.DropAndAddRowAccessPolicy = &EventTableDropAndAddRowAccessPolicy{}
	opts.DropAndAddRowAccessPolicy.Drop = EventTableDropRowAccessPolicy{
		RowAccessPolicy: r.DropAndAddRowAccessPolicy.Drop.RowAccessPolicy,
	}
	opts.DropAndAddRowAccessPolicy.Add = EventTableAddRowAccessPolicy{
		RowAccessPolicy: r.DropAndAddRowAccessPolicy.Add.RowAccessPolicy,
		On:              r.DropAndAddRowAccessPolicy.Add.On,
	}
}


