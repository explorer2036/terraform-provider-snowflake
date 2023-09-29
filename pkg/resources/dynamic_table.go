package resources

import (
	"context"
	"database/sql"
	"log"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dynamicTableShema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Specifies the identifier for the dynamic table.",
		ForceNew:    true,
	},
	"database": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The database in which to create the dynamic table.",
		ForceNew:    true,
	},
	"schema": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The schema in which to create the dynamic table.",
		ForceNew:    true,
	},
	"warehouse": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The warehouse in which to create the dynamic table.",
		ForceNew:    true,
	},
	"query": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Specifies the query to use to populate the dynamic table.",
	},
	"or_replace": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Specifies whether to replace the dynamic table if it already exists.",
		Default:     false,
	},
	"target_lag": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Specifies the target lag time for the dynamic table.",
	},
	"comment": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Specifies a comment for the dynamic table.",
	},
}

// DynamicTable returns a pointer to the resource representing a dynamic table.
func DynamicTable() *schema.Resource {
	return &schema.Resource{
		Create: CreateDynamicTable,
		Read:   ReadDynamicTable,
		Update: UpdateDynamicTable,
		Delete: DeleteDynamicTable,

		Schema: dynamicTableShema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// ReadDynamicTable implements schema.ReadFunc.
func ReadDynamicTable(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.AccountObjectIdentifier)
	dynamicTable, err := client.DynamicTables.ShowByID(context.Background(), id)
	if err != nil {
		log.Printf("[DEBUG] dynamic table (%s) not found", d.Id())
		d.SetId("")
		return nil
	}
	if err := d.Set("name", dynamicTable.Name); err != nil {
		return err
	}
	if err := d.Set("warehouse", dynamicTable.Warehouse); err != nil {
		return err
	}
	if err := d.Set("comment", dynamicTable.Comment); err != nil {
		return err
	}
	if err := d.Set("scheduling_state", dynamicTable.SchedulingState); err != nil {
		return err
	}
	return nil
}

// CreateDynamicTable implements schema.CreateFunc.
func CreateDynamicTable(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	databaseName := d.Get("database").(string)
	schemaName := d.Get("schema").(string)
	name := d.Get("name").(string)
	id := sdk.NewSchemaObjectIdentifier(databaseName, schemaName, name)

	warehouseName := d.Get("warehouse").(string)
	warehouse := sdk.NewAccountObjectIdentifier(warehouseName)
	query := d.Get("query").(string)

	var tl sdk.TargetLag
	targetLag := d.Get("target_lag").(string)
	if targetLag == "DOWNSTREAM" {
		tl.Downstream = sdk.Bool(true)
	} else {
		tl.Lagtime = sdk.String(targetLag)
	}
	request := sdk.NewCreateDynamicTableRequest(id, warehouse, tl, query)
	if v, ok := d.GetOk("comment"); ok {
		request.WithComment(sdk.String(v.(string)))
	}
	if v, ok := d.GetOk("or_replace"); ok && v.(bool) {
		request.WithOrReplace(true)
	}
	if err := client.DynamicTables.Create(context.Background(), request); err != nil {
		return err
	}
	d.SetId(helpers.EncodeSnowflakeID(name))

	return ReadDynamicTable(d, meta)
}

// UpdateDynamicTable implements schema.UpdateFunc.
func UpdateDynamicTable(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.SchemaObjectIdentifier)
	request := sdk.NewAlterDynamicTableRequest(id)
	switch {
	case d.HasChange("suspend"):
		_, newVal := d.GetChange("suspend")
		request.WithSuspend(sdk.Bool(newVal.(bool)))
	case d.HasChange("resume"):
		_, newVal := d.GetChange("resume")
		request.WithResume(sdk.Bool(newVal.(bool)))
	case d.HasChange("refresh"):
		_, newVal := d.GetChange("refresh")
		request.WithRefresh(sdk.Bool(newVal.(bool)))
	}
	if err := client.DynamicTables.Alter(context.Background(), request); err != nil {
		return err
	}
	return ReadDynamicTable(d, meta)
}

// DeleteDynamicTable implements schema.DeleteFunc.
func DeleteDynamicTable(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.SchemaObjectIdentifier)
	if err := client.DynamicTables.Drop(context.Background(), sdk.NewDropDynamicTableRequest(id)); err != nil {
		return err
	}
	d.SetId("")

	return nil
}
