package sdk

import (
	"context"
	"database/sql"
	"time"
)

type EventTables interface {
	Create(ctx context.Context, request *CreateEventTableRequest) error
	Alter(ctx context.Context, request *AlterEventTableRequest) error
	Describe(ctx context.Context, request *DescribeEventTableRequest) (*EventTableDetails, error)
	Show(ctx context.Context, opts *ShowEventTableRequest) ([]EventTable, error)
	ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*EventTable, error)
}

// createEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-event-table
type createEventTableOptions struct {
	create      bool   `ddl:"static" sql:"CREATE"`
	OrReplace   *bool  `ddl:"keyword" sql:"OR REPLACE"`
	eventTable  string `ddl:"static" sql:"EVENT TABLE"`
	IfNotExists *bool  `ddl:"keyword" sql:"IF NOT EXISTS"`

	name                       SchemaObjectIdentifier `ddl:"identifier"`
	ClusterBy                  []string               `ddl:"keyword,parentheses" sql:"CLUSTER BY"`
	DataRetentionTimeInDays    *uint                  `ddl:"parameter" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *uint                  `ddl:"parameter" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool                  `ddl:"parameter" sql:"CHANGE_TRACKING"`
	DefaultDDLCollation        *string                `ddl:"parameter,single_quotes" sql:"DEFAULT_DDL_COLLATION"`
	CopyGrants                 *bool                  `ddl:"keyword" sql:"COPY GRANTS"`
	Comment                    *string                `ddl:"parameter,single_quotes" sql:"COMMENT"`
	RowAccessPolicy            *RowAccessPolicy       `ddl:"keyword"`
	Tag                        []TagAssociation       `ddl:"keyword,parentheses" sql:"TAG"`
}

type EventTableAddRowAccessPolicy struct {
	RowAccessPolicy *RowAccessPolicy `ddl:"keyword"`
}

type EventTableDropRowAccessPolicy struct {
	rowAccessPolicy bool                   `ddl:"static" sql:"ROW ACCESS POLICY"`
	Name            SchemaObjectIdentifier `ddl:"identifier"`
}

type EventTableSetProperties struct {
	DataRetentionTimeInDays    *uint   `ddl:"parameter" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *uint   `ddl:"parameter" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool   `ddl:"parameter" sql:"CHANGE_TRACKING"`
	Comment                    *string `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

type EventTableSet struct {
	Properties *EventTableSetProperties `ddl:"keyword"`
	Tag        *[]TagAssociation        `ddl:"keyword" sql:"TAG"`
}

type EventTableUnset struct {
	DataRetentionTimeInDays    *bool     `ddl:"keyword" sql:"DATA_RETENTION_TIME_IN_DAYS"`
	MaxDataExtensionTimeInDays *bool     `ddl:"keyword" sql:"MAX_DATA_EXTENSION_TIME_IN_DAYS"`
	ChangeTracking             *bool     `ddl:"keyword" sql:"CHANGE_TRACKING"`
	Comment                    *bool     `ddl:"keyword" sql:"COMMENT"`
	TagNames                   *[]string `ddl:"keyword" sql:"TAG"`
}

type EventTableRename struct {
	Name SchemaObjectIdentifier `ddl:"identifier"`
}

type ClusteringAction struct {
	ClusterBy *[]string `ddl:"keyword,parentheses" sql:"CLUSTER BY"`
	Suspend   *bool     `ddl:"keyword" sql:"SUSPEND RECLUSTER"`
	Resume    *bool     `ddl:"keyword" sql:"RESUME RECLUSTER"`
	Drop      *bool     `ddl:"keyword" sql:"DROP CLUSTERING KEY"`
}

type AddSearchOptimization struct {
	add bool     `ddl:"static" sql:"ADD SEARCH OPTIMIZATION"`
	On  []string `ddl:"keyword,parentheses" sql:"ON"`
}

type DropSearchOptimization struct {
	drop bool     `ddl:"static" sql:"DROP SEARCH OPTIMIZATION"`
	On   []string `ddl:"keyword,parentheses" sql:"ON"`
}

type SearchOptimizationAction struct {
	Add  *AddSearchOptimization  `ddl:"keyword"`
	Drop *DropSearchOptimization `ddl:"keyword"`
}

// alterEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/alter-table-event-table
type alterEventTableOptions struct {
	alter      bool                   `ddl:"static" sql:"ALTER"`
	eventTable string                 `ddl:"static" sql:"TABLE"`
	name       SchemaObjectIdentifier `ddl:"identifier"`

	// One of
	ClusteringAction         *ClusteringAction              `ddl:"keyword"`
	SearchOptimizationAction *SearchOptimizationAction      `ddl:"keyword"`
	AddRowAccessPolicy       *EventTableAddRowAccessPolicy  `ddl:"keyword" sql:"ADD"`
	DropRowAccessPolicy      *EventTableDropRowAccessPolicy `ddl:"keyword" sql:"DROP"`
	DropAllRowAccessPolicies *bool                          `ddl:"keyword" sql:"DROP ALL ROW ACCESS POLICIES"`
	Set                      *EventTableSet                 `ddl:"keyword" sql:"SET"`
	Unset                    *EventTableUnset               `ddl:"keyword" sql:"UNSET"`
	Rename                   *EventTableRename              `ddl:"keyword" sql:"RENAME TO"`
}

type EventTable struct {
	CreatedOn      time.Time
	Name           string
	DatabaseName   string
	SchemaName     string
	ClusterBy      string
	Owner          string
	Comment        string
	ChangeTracking bool
}

type eventTableRow struct {
	CreatedOn      time.Time `db:"created_on"`
	Name           string    `db:"name"`
	DatabaseName   string    `db:"database_name"`
	SchemaName     string    `db:"schema_name"`
	ClusterBy      string    `db:"cluster_by"`
	Owner          string    `db:"owner"`
	Comment        string    `db:"comment"`
	ChangeTracking string    `db:"change_tracking"`
}

func (r eventTableRow) convert() *EventTable {
	return &EventTable{
		CreatedOn:      r.CreatedOn,
		Name:           r.Name,
		DatabaseName:   r.DatabaseName,
		SchemaName:     r.SchemaName,
		ClusterBy:      r.ClusterBy,
		Owner:          r.Owner,
		Comment:        r.Comment,
		ChangeTracking: r.ChangeTracking == "ON",
	}
}

// showEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-event-tables
type showEventTableOptions struct {
	show       bool       `ddl:"static" sql:"SHOW"`
	eventTable bool       `ddl:"static" sql:"EVENT TABLES"`
	Like       *Like      `ddl:"keyword" sql:"LIKE"`
	In         *In        `ddl:"keyword" sql:"IN"`
	StartsWith *string    `ddl:"parameter,single_quotes,no_equals" sql:"STARTS WITH"`
	Limit      *LimitFrom `ddl:"keyword" sql:"LIMIT"`
}

type EventTableDetails struct {
	Name       string
	Type       DataType
	Kind       string
	IsNull     bool
	Default    string
	PrimaryKey string
	UniqueKey  string
	Check      string
	Expression string
	Comment    string
	PolicyName string
}

type eventTableDetailsRow struct {
	Name       string         `db:"name"`
	Type       string         `db:"type"`
	Kind       string         `db:"kind"`
	IsNull     string         `db:"null?"`
	Default    sql.NullString `db:"default"`
	PrimaryKey string         `db:"primary key"`
	UniqueKey  string         `db:"unique key"`
	Check      sql.NullString `db:"check"`
	Expression sql.NullString `db:"expression"`
	Comment    sql.NullString `db:"comment"`
	PolicyName sql.NullString `db:"policy name"`
}

func (r eventTableDetailsRow) convert() *EventTableDetails {
	typ, _ := ToDataType(r.Type)
	d := &EventTableDetails{
		Name:       r.Name,
		Type:       typ,
		Kind:       r.Kind,
		IsNull:     r.IsNull == "Y",
		PrimaryKey: r.PrimaryKey,
		UniqueKey:  r.UniqueKey,
	}
	if r.Default.Valid {
		d.Default = r.Default.String
	}
	if r.Check.Valid {
		d.Check = r.Check.String
	}
	if r.Expression.Valid {
		d.Expression = r.Expression.String
	}
	if r.Comment.Valid {
		d.Comment = r.Comment.String
	}
	if r.PolicyName.Valid {
		d.PolicyName = r.PolicyName.String
	}
	return d
}

// describeEventTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/desc-event-table
type describeEventTableOptions struct {
	describe   bool                   `ddl:"static" sql:"DESCRIBE"`
	eventTable bool                   `ddl:"static" sql:"EVENT TABLE"`
	name       SchemaObjectIdentifier `ddl:"identifier"`
}
