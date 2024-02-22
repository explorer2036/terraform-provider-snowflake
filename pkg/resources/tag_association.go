package resources

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	snowflakeValidation "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/validation"
)

var tagAssociationSchema = map[string]*schema.Schema{
	"object_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Specifies the object identifier for the tag association.",
		Deprecated:  "Use `object_identifier` instead",
		ForceNew:    true,
	},
	"object_identifier": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		Description: "Specifies the object identifier for the tag association.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					ForceNew:    true,
					Description: "Name of the object to associate the tag with.",
				},
				"database": {
					Type:        schema.TypeString,
					Optional:    true,
					ForceNew:    true,
					Description: "Name of the database that the object was created in.",
				},
				"schema": {
					Type:        schema.TypeString,
					Optional:    true,
					ForceNew:    true,
					Description: "Name of the schema that the object was created in.",
				},
			},
		},
	},
	"object_type": {
		Type:     schema.TypeString,
		Required: true,
		Description: "Specifies the type of object to add a tag to. ex: 'ACCOUNT', 'COLUMN', 'DATABASE', etc. " +
			"For more information: https://docs.snowflake.com/en/user-guide/object-tagging.html#supported-objects",
		ValidateFunc: validation.StringInSlice([]string{
			"ACCOUNT", "COLUMN", "DATABASE", "INTEGRATION", "PIPE", "ROLE", "SCHEMA", "STREAM", "SHARE", "STAGE",
			"TABLE", "TASK", "USER", "VIEW", "WAREHOUSE",
		}, true),
		ForceNew: true,
	},
	"tag_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Specifies the identifier for the tag. Note: format must follow: \"databaseName\".\"schemaName\".\"tagName\" or \"databaseName.schemaName.tagName\" or \"databaseName|schemaName.tagName\" (snowflake_tag.tag.id)",
		ForceNew:    true,
	},
	"tag_value": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Specifies the value of the tag, (e.g. 'finance' or 'engineering')",
		ForceNew:    true,
	},
	"skip_validation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, skips validation of the tag association.",
		Default:     true,
	},
}

// TagAssociation returns a pointer to the resource representing a schema.
func TagAssociation() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateContextTagAssociation,
		ReadContext:   ReadContextTagAssociation,
		UpdateContext: UpdateContextTagAssociation,
		DeleteContext: DeleteContextTagAssociation,

		Schema: tagAssociationSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(70 * time.Minute),
		},
	}
}

func tagIdentifierAndObjectIdentifier(d *schema.ResourceData) (sdk.SchemaObjectIdentifier, sdk.ObjectIdentifier, sdk.ObjectType) {
	tag := d.Get("tag_id").(string)
	objectIdentifier := d.Get("object_identifier")
	objectType := sdk.ObjectType(d.Get("object_type").(string))

	tagDatabase, tagSchema, tagName := snowflakeValidation.ParseFullyQualifiedObjectID(tag)
	tid := sdk.NewSchemaObjectIdentifier(tagDatabase, tagSchema, tagName)

	objectDatabase, objectSchema, objectName := expandObjectIdentifier(objectIdentifier)
	fullyQualifierObjectIdentifier := snowflakeValidation.FormatFullyQualifiedObjectID(objectDatabase, objectSchema, objectName)
	oid := sdk.NewObjectIdentifierFromFullyQualifiedName(fullyQualifierObjectIdentifier)

	return tid, oid, objectType
}

func CreateContextTagAssociation(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	tagValue := d.Get("tag_value").(string)

	tid, oid, ot := tagIdentifierAndObjectIdentifier(d)
	request := sdk.NewSetTagRequest(ot, oid).WithSetTags([]sdk.TagAssociation{
		{
			Name:  tid,
			Value: tagValue,
		},
	})
	if err := client.Tags.Set(ctx, request); err != nil {
		return diag.FromErr(err)
	}
	skipValidate := d.Get("skip_validation").(bool)
	if !skipValidate {
		log.Println("[DEBUG] validating tag creation")
		if err := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *retry.RetryError {
			tags, err := client.SystemFunctions.GetTags(ctx, tid, oid, ot)
			if err != nil {
				return retry.NonRetryableError(fmt.Errorf("error getting tags: %w", err))
			}
			// if length of response is zero, tag association was not found. retry for up to 70 minutes
			if len(tags) == 0 {
				return retry.RetryableError(fmt.Errorf("expected tag association to be created but not yet created"))
			}
			return nil
		}); err != nil {
			return diag.FromErr(fmt.Errorf("error validating tag creation: %w", err))
		}
	}
	d.SetId(helpers.EncodeSnowflakeID(tid.DatabaseName(), tid.SchemaName(), tid.Name()))
	return ReadContextTagAssociation(ctx, d, meta)
}

func ReadContextTagAssociation(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	tid, oid, ot := tagIdentifierAndObjectIdentifier(d)
	tagValue, err := client.SystemFunctions.GetTag(ctx, tid, oid, ot)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tag_value", tagValue); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func UpdateContextTagAssociation(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	tid, oid, ot := tagIdentifierAndObjectIdentifier(d)
	if d.HasChange("skip_validation") {
		o, n := d.GetChange("skip_validation")
		log.Printf("[DEBUG] skip_validation changed from %v to %v", o, n)
	}
	if d.HasChange("tag_value") {
		tagValue, ok := d.GetOk("tag_value")
		if ok {
			request := sdk.NewSetTagRequest(ot, oid).WithSetTags([]sdk.TagAssociation{
				{
					Name:  tid,
					Value: tagValue.(string),
				},
			})
			if err := client.Tags.Set(ctx, request); err != nil {
				return diag.FromErr(err)
			}
		} else {
			request := sdk.NewUnsetTagRequest(ot, oid).WithUnsetTags([]sdk.ObjectIdentifier{tid})
			if err := client.Tags.Unset(ctx, request); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return ReadContextTagAssociation(ctx, d, meta)
}

func DeleteContextTagAssociation(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)

	tid, oid, ot := tagIdentifierAndObjectIdentifier(d)
	request := sdk.NewUnsetTagRequest(ot, oid).WithUnsetTags([]sdk.ObjectIdentifier{tid})
	if err := client.Tags.Unset(ctx, request); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
