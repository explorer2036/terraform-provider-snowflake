package sdk

import "context"

type Sequences interface {
	Create(ctx context.Context, request *CreateSequenceRequest) error
	Show(ctx context.Context, request *ShowSequenceRequest) ([]Sequence, error)
	ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*Sequence, error)
	Describe(ctx context.Context, id SchemaObjectIdentifier) ([]SequenceDetail, error)
	Drop(ctx context.Context, request *DropSequenceRequest) error
}

// CreateSequenceOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-sequence.
type CreateSequenceOptions struct {
	create         bool                   `ddl:"static" sql:"CREATE"`
	OrReplace      *bool                  `ddl:"keyword" sql:"OR REPLACE"`
	sequence       bool                   `ddl:"static" sql:"SEQUENCE"`
	IfNotExists    *bool                  `ddl:"keyword" sql:"IF NOT EXISTS"`
	name           SchemaObjectIdentifier `ddl:"identifier"`
	With           *bool                  `ddl:"keyword" sql:"WITH"`
	Start          *int                   `ddl:"parameter,no_quotes" sql:"START"`
	Increment      *int                   `ddl:"parameter,no_quotes" sql:"INCREMENT"`
	ValuesBehavior *ValuesBehavior        `ddl:"keyword"`
	Comment        *string                `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

// ShowSequenceOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-sequences.
type ShowSequenceOptions struct {
	show      bool  `ddl:"static" sql:"SHOW"`
	sequences bool  `ddl:"static" sql:"SEQUENCES"`
	Like      *Like `ddl:"keyword" sql:"LIKE"`
	In        *In   `ddl:"keyword" sql:"IN"`
}

type sequenceRow struct {
	CreatedOn     string `db:"created_on"`
	Name          string `db:"name"`
	SchemaName    string `db:"schema_name"`
	DatabaseName  string `db:"database_name"`
	NextValue     int    `db:"next_value"`
	Interval      int    `db:"interval"`
	Owner         string `db:"owner"`
	OwnerRoleType string `db:"owner_role_type"`
	Comment       string `db:"comment"`
	Ordered       string `db:"ordered"`
}

type Sequence struct {
	CreatedOn     string
	Name          string
	SchemaName    string
	DatabaseName  string
	NextValue     int
	Interval      int
	Owner         string
	OwnerRoleType string
	Comment       string
	Ordered       bool
}

// DescribeSequenceOptions is based on https://docs.snowflake.com/en/sql-reference/sql/desc-sequence.
type DescribeSequenceOptions struct {
	describe bool                   `ddl:"static" sql:"DESCRIBE"`
	sequence bool                   `ddl:"static" sql:"SEQUENCE"`
	name     SchemaObjectIdentifier `ddl:"identifier"`
}

type sequenceDetailRow struct {
	CreatedOn     string `db:"created_on"`
	Name          string `db:"name"`
	SchemaName    string `db:"schema_name"`
	DatabaseName  string `db:"database_name"`
	NextValue     int    `db:"next_value"`
	Interval      int    `db:"interval"`
	Owner         string `db:"owner"`
	OwnerRoleType string `db:"owner_role_type"`
	Comment       string `db:"comment"`
	Ordered       string `db:"ordered"`
}

type SequenceDetail struct {
	CreatedOn     string
	Name          string
	SchemaName    string
	DatabaseName  string
	NextValue     int
	Interval      int
	Owner         string
	OwnerRoleType string
	Comment       string
	Ordered       bool
}

// DropSequenceOptions is based on https://docs.snowflake.com/en/sql-reference/sql/drop-sequence.
type DropSequenceOptions struct {
	drop     bool                   `ddl:"static" sql:"DROP"`
	sequence bool                   `ddl:"static" sql:"SEQUENCE"`
	IfExists *bool                  `ddl:"keyword" sql:"IF EXISTS"`
	name     SchemaObjectIdentifier `ddl:"identifier"`
}
