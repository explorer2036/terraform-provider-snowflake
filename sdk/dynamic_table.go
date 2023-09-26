package sdk

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	_ validatable = new(createDynamicTableOptions)
	_ validatable = new(alterDynamicTableOptions)
	_ validatable = new(dropDynamicTableOptions)
	_ validatable = new(ShowDynamicTableOptions)
	_ validatable = new(describeDynamicTableOptions)
	_ validatable = new(DynamicTableSet)
)

type DynamicTables interface {
	Create(ctx context.Context, request *CreateDynamicTableRequest) error
	Alter(ctx context.Context, request *AlterDynamicTableRequest) error
	Describe(ctx context.Context, request *DescribeDynamicTableRequest) (*DynamicTableDetails, error)
	Drop(ctx context.Context, request *DropDynamicTableRequest) error
	Show(ctx context.Context, opts *ShowDynamicTableOptions) ([]*DynamicTable, error)
}

// createDynamicTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/create-dynamic-table
type createDynamicTableOptions struct {
	create       bool                    `ddl:"static" sql:"CREATE"`
	OrReplace    *bool                   `ddl:"keyword" sql:"OR REPLACE"`
	dynamicTable bool                    `ddl:"static" sql:"DYNAMIC TABLE"`
	name         AccountObjectIdentifier `ddl:"identifier"`
	targetLag    string                  `ddl:"parameter,no_quotes" sql:"TARGET_LAG"`
	warehouse    AccountObjectIdentifier `ddl:"identifier,equals" sql:"WAREHOUSE"`
	Comment      *string                 `ddl:"parameter,single_quotes" sql:"COMMENT"`
	query        string                  `ddl:"parameter,no_equals,no_quotes" sql:"AS"`
}

// target log format: '<num> { seconds | minutes | hours | days }' | DOWNSTREAM
func validateAndSingleQuoteTargetLag(s *string) error {
	if *s == "DOWNSTREAM" {
		return nil
	}
	parts := strings.Split(strings.TrimSpace(*s), " ")
	if len(parts) != 2 {
		return errors.New("The string format is invalid")
	}
	if _, err := strconv.Atoi(parts[0]); err != nil {
		return errors.New("The number value is invalid")
	}
	switch parts[1] {
	case "second", "seconds":
	case "minute", "minutes":
	case "hour", "hours":
	case "day", "days":
	default:
		return errors.New("The unit is invalid")
	}
	*s = fmt.Sprintf("'%s'", *s)
	return nil
}

func (opts *createDynamicTableOptions) validate() error {
	if opts == nil {
		return errNilOptions
	}
	if !validObjectidentifier(opts.name) {
		return ErrInvalidObjectIdentifier
	}
	if !validObjectidentifier(opts.warehouse) {
		return ErrInvalidObjectIdentifier
	}
	if err := validateAndSingleQuoteTargetLag(&opts.targetLag); err != nil {
		return err
	}
	return nil
}

type DynamicTableSet struct {
	TargetLag *string                  `ddl:"parameter,no_quotes" sql:"TARGET_LAG"`
	Warehouse *AccountObjectIdentifier `ddl:"identifier,equals" sql:"WAREHOUSE"`
}

func (dts *DynamicTableSet) validate() error {
	if valueSet(dts.TargetLag) {
		if err := validateAndSingleQuoteTargetLag(dts.TargetLag); err != nil {
			return err
		}
	}
	if valueSet(dts.Warehouse) {
		if !validObjectidentifier(*dts.Warehouse) {
			return ErrInvalidObjectIdentifier
		}
	}
	return nil
}

// alterDynamicTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/alter-dynamic-table
type alterDynamicTableOptions struct {
	alter        bool                    `ddl:"static" sql:"ALTER"`
	dynamicTable bool                    `ddl:"static" sql:"DYNAMIC TABLE"`
	name         AccountObjectIdentifier `ddl:"identifier"`

	Suspend *bool            `ddl:"keyword" sql:"SUSPEND"`
	Resume  *bool            `ddl:"keyword" sql:"RESUME"`
	Refresh *bool            `ddl:"keyword" sql:"REFRESH"`
	Set     *DynamicTableSet `ddl:"keyword" sql:"SET"`
}

func (opts *alterDynamicTableOptions) validate() error {
	if opts == nil {
		return errNilOptions
	}
	if !validObjectidentifier(opts.name) {
		return ErrInvalidObjectIdentifier
	}
	if ok := exactlyOneValueSet(
		opts.Suspend,
		opts.Resume,
		opts.Refresh,
		opts.Set,
	); !ok {
		return fmt.Errorf("exactly one of Suspend, Resume, Refresh, Set must be set")
	}
	if everyValueSet(opts.Suspend, opts.Resume) && (*opts.Suspend && *opts.Resume) {
		return fmt.Errorf("Suspend and Resume cannot both be true")
	}
	if valueSet(opts.Set) && valueSet(opts.Set.TargetLag) {
		if err := validateAndSingleQuoteTargetLag(opts.Set.TargetLag); err != nil {
			return err
		}
	}
	return nil
}

// dropDynamicTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/drop-dynamic-table
type dropDynamicTableOptions struct {
	drop         bool                    `ddl:"static" sql:"DROP"`
	dynamicTable bool                    `ddl:"static" sql:"DYNAMIC TABLE"`
	name         AccountObjectIdentifier `ddl:"identifier"`
}

func (opts *dropDynamicTableOptions) validate() error {
	if opts == nil {
		return errNilOptions
	}
	if !validObjectidentifier(opts.name) {
		return ErrInvalidObjectIdentifier
	}
	return nil
}

// ShowDynamicTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/show-dynamic-tables
type ShowDynamicTableOptions struct {
	show         bool       `ddl:"static" sql:"SHOW"`
	dynamicTable bool       `ddl:"static" sql:"DYNAMIC TABLES"`
	Like         *Like      `ddl:"keyword" sql:"LIKE"`
	In           *In        `ddl:"keyword" sql:"IN"`
	StartsWith   *string    `ddl:"parameter,single_quotes,no_equals" sql:"STARTS WITH"`
	Limit        *LimitFrom `ddl:"keyword" sql:"LIMIT"`
}

func (opts *ShowDynamicTableOptions) validate() error {
	if opts == nil {
		return errNilOptions
	}
	if valueSet(opts.Like) && !valueSet(opts.Like.Pattern) {
		return errPatternRequiredForLikeKeyword
	}
	if valueSet(opts.In) && !exactlyOneValueSet(opts.In.Account, opts.In.Database, opts.In.Schema) {
		return errScopeRequiredForInKeyword
	}
	return nil
}

type DynamicTableRefreshMode string

const (
	DynamicTableRefreshModeIncremental DynamicTableRefreshMode = "INCREMENTAL"
	DynamicTableRefreshModeFull        DynamicTableRefreshMode = "FULL"
)

type DynamicTableSchedulingState string

const (
	DynamicTableSchedulingStateRunning   DynamicTableSchedulingState = "RUNNING"
	DynamicTableSchedulingStateSuspended DynamicTableSchedulingState = "SUSPENDED"
)

type DynamicTable struct {
	CreatedOn           time.Time
	Name                string
	Reserved            string
	DatabaseName        string
	SchemaName          string
	ClusterBy           string
	Rows                int
	Bytes               int
	Owner               string
	TargetLag           string
	RefreshMode         DynamicTableRefreshMode
	RefreshModeReason   string
	Warehouse           string
	Comment             string
	Text                string
	AutomaticClustering bool
	SchedulingState     DynamicTableSchedulingState
	LastSuspendedOn     time.Time
	IsClone             bool
	IsReplica           bool
	DataTimestamp       time.Time
}

func (dt *DynamicTable) ID() AccountObjectIdentifier {
	return NewAccountObjectIdentifier(dt.Name)
}

type dynamicTableRow struct {
	CreatedOn           time.Time      `db:"created_on"`
	Name                string         `db:"name"`
	Reserved            string         `db:"reserved"`
	DatabaseName        string         `db:"database_name"`
	SchemaName          string         `db:"schema_name"`
	ClusterBy           string         `db:"cluster_by"`
	Rows                int            `db:"rows"`
	Bytes               int            `db:"bytes"`
	Owner               string         `db:"owner"`
	TargetLag           string         `db:"target_lag"`
	RefreshMode         string         `db:"refresh_mode"`
	RefreshModeReason   sql.NullString `db:"refresh_mode_reason"`
	Warehouse           string         `db:"warehouse"`
	Comment             string         `db:"comment"`
	Text                string         `db:"text"`
	AutomaticClustering string         `db:"automatic_clustering"`
	SchedulingState     string         `db:"scheduling_state"`
	LastSuspendedOn     sql.NullTime   `db:"last_suspended_on"`
	IsClone             bool           `db:"is_clone"`
	IsReplica           bool           `db:"is_replica"`
	DataTimestamp       time.Time      `db:"data_timestamp"`
}

func (dtr *dynamicTableRow) toDynamicTable() *DynamicTable {
	dt := &DynamicTable{
		CreatedOn:           dtr.CreatedOn,
		Name:                dtr.Name,
		Reserved:            dtr.Reserved,
		DatabaseName:        dtr.DatabaseName,
		SchemaName:          dtr.SchemaName,
		ClusterBy:           dtr.ClusterBy,
		Rows:                dtr.Rows,
		Bytes:               dtr.Bytes,
		Owner:               dtr.Owner,
		TargetLag:           dtr.TargetLag,
		RefreshMode:         DynamicTableRefreshMode(dtr.RefreshMode),
		Warehouse:           dtr.Warehouse,
		Comment:             dtr.Comment,
		Text:                dtr.Text,
		AutomaticClustering: dtr.AutomaticClustering == "ON", // "ON" or "OFF
		SchedulingState:     DynamicTableSchedulingState(dtr.SchedulingState),
		IsClone:             dtr.IsClone,
		IsReplica:           dtr.IsReplica,
		DataTimestamp:       dtr.DataTimestamp,
	}
	if dtr.RefreshModeReason.Valid {
		dt.RefreshModeReason = dtr.RefreshModeReason.String
	}
	if dtr.LastSuspendedOn.Valid {
		dt.LastSuspendedOn = dtr.LastSuspendedOn.Time
	}
	return dt
}

func (dt *dynamicTables) Show(ctx context.Context, opts *ShowDynamicTableOptions) ([]*DynamicTable, error) {
	if opts == nil {
		opts = &ShowDynamicTableOptions{}
	}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	sql, err := structToSQL(opts)
	if err != nil {
		return nil, err
	}
	rows := []*dynamicTableRow{}
	if err := dt.client.query(ctx, &rows, sql); err != nil {
		return nil, err
	}
	entities := make([]*DynamicTable, len(rows))
	for i, row := range rows {
		entities[i] = row.toDynamicTable()
	}
	return entities, nil
}

// describeDynamicTableOptions is based on https://docs.snowflake.com/en/sql-reference/sql/desc-dynamic-table
type describeDynamicTableOptions struct {
	describe     bool                    `ddl:"static" sql:"DESCRIBE"`
	dynamicTable bool                    `ddl:"static" sql:"DYNAMIC TABLE"`
	name         AccountObjectIdentifier `ddl:"identifier"`
}

func (opts *describeDynamicTableOptions) validate() error {
	if !validObjectidentifier(opts.name) {
		return ErrInvalidObjectIdentifier
	}
	return nil
}

type DynamicTableDetails struct {
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

type dynamicTableDetailsRow struct {
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

func (row dynamicTableDetailsRow) convert() *DynamicTableDetails {
	typ, _ := ToDataType(row.Type)
	dtd := &DynamicTableDetails{
		Name:       row.Name,
		Type:       typ,
		Kind:       row.Kind,
		IsNull:     row.IsNull == "Y",
		PrimaryKey: row.PrimaryKey,
		UniqueKey:  row.UniqueKey,
	}
	if row.Default.Valid {
		dtd.Default = row.Default.String
	}
	if row.Check.Valid {
		dtd.Check = row.Check.String
	}
	if row.Expression.Valid {
		dtd.Expression = row.Expression.String
	}
	if row.Comment.Valid {
		dtd.Comment = row.Comment.String
	}
	if row.PolicyName.Valid {
		dtd.PolicyName = row.PolicyName.String
	}
	return dtd
}
