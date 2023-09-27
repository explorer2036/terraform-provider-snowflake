package datasources

import (
	"context"
	"database/sql"
	"log"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dynamicTablesSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the dynamic table.",
	},
	"dynamic_tables": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "The list of dynamic tables.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The name of the dynamic table.",
				},
				"warehouse": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The warehouse of the dynamic table.",
				},
				"target_lag": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The target lag time of the dynamic table.",
				},
				"scheduling_state": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The scheduling state of the dynamic table.",
				},
				"comment": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The comment of the dynamic table.",
				},
			},
		},
	},
}

// DynamicTables Snowflake Dynamic Tables resource.
func DynamicTables() *schema.Resource {
	return &schema.Resource{
		Read:   ReadDynamicTables,
		Schema: dynamicTablesSchema,
	}
}

// ReadDynamicTables Reads the dynamic tables metadata information.
func ReadDynamicTables(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	d.SetId("dynamic_tables_read")

	name := d.Get("name").(string)
	request := sdk.NewShowDynamicTableRequest().WithLike(name)
	extractedDynamicTables, err := client.DynamicTables.Show(context.Background(), request)
	if err != nil {
		log.Printf("[DEBUG] unable to show dynamic tables with (%s)", name)
		d.SetId("")
		return err
	}
	dynamicTables := make([]map[string]any, 0, len(extractedDynamicTables))
	for _, dynamicTable := range extractedDynamicTables {
		item := map[string]any{}
		item["name"] = dynamicTable.Name
		item["comment"] = dynamicTable.Comment
		item["warehouse"] = dynamicTable.Warehouse
		item["target_lag"] = dynamicTable.TargetLag
		item["scheduling_state"] = dynamicTable.SchedulingState
		dynamicTables = append(dynamicTables, item)
	}
	return d.Set("dynamic_tables", dynamicTables)
}
