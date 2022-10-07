package resources

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/snowflake"
	snowflakeValidation "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/validation"
)

const (
	tagAttachmentPolicyIDDelimiter = '|'
)

var mpAttachmentPolicySchema = map[string]*schema.Schema{
	"tag_id": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  "Specifies the identifier for the tag. Note: format must follow: \"databaseName\".\"schemaName\".\"tagName\" or \"databaseName.schemaName.tagName\" or \"databaseName|schemaName.tagName\" (snowflake_tag.tag.id)",
		ValidateFunc: snowflakeValidation.ValidateFullyQualifiedObjectID,
		ForceNew:     true,
	},
	"masking_policy_id": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "The the resource id of the masking policy",
	},
}

type attachmentID struct {
	TagDatabaseName           string
	TagSchemaName             string
	TagName                   string
	MaskingPolicyDatabaseName string
	MaskingPolicySchemaName   string
	MaskingPolicyName         string
}

// String() takes in a schemaID object and returns a pipe-delimited string:
// TagDatabaseName | TagSchemaName | TagName | MaskingPolicyDatabaseName | MaskingPolicySchemaName | MaskingPolicyName
func (ti *attachmentID) String() (string, error) {
	var buf bytes.Buffer
	csvWriter := csv.NewWriter(&buf)
	csvWriter.Comma = tagAttachmentPolicyIDDelimiter
	dataIdentifiers := [][]string{{ti.TagDatabaseName, ti.TagSchemaName, ti.TagName, ti.MaskingPolicyDatabaseName, ti.MaskingPolicySchemaName, ti.MaskingPolicyName}}
	err := csvWriter.WriteAll(dataIdentifiers)
	if err != nil {
		return "", err
	}
	strTagID := strings.TrimSpace(buf.String())
	return strTagID, nil
}

// attachedPolicyIDFromString() takes in a pipe-delimited string: TagDatabaseName | TagSchemaName | TagName | MaskingPolicyDatabaseName | MaskingPolicySchemaName | MaskingPolicyName
// and returns a attachmentID object
func attachedPolicyIDFromString(stringID string) (*attachmentID, error) {
	reader := csv.NewReader(strings.NewReader(stringID))
	reader.Comma = tagAttachmentPolicyIDDelimiter
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("not CSV compatible")
	}

	if len(lines) != 1 {
		return nil, fmt.Errorf("1 line per schema")
	}
	if len(lines[0]) != 6 {
		return nil, fmt.Errorf("6 fields allowed")
	}

	attachmentResult := &attachmentID{
		TagDatabaseName:           lines[0][0],
		TagSchemaName:             lines[0][1],
		TagName:                   lines[0][2],
		MaskingPolicyDatabaseName: lines[0][3],
		MaskingPolicySchemaName:   lines[0][4],
		MaskingPolicyName:         lines[0][5],
	}
	return attachmentResult, nil
}

// Schema returns a pointer to the resource representing a schema
func TagMaskingPolicyAssociation() *schema.Resource {
	return &schema.Resource{
		Create: CreateTagMaskingPolicyAssociation,
		Read:   ReadTagMaskingPolicyAssociation,
		Delete: DeleteTagMaskingPolicyAssociation,

		Schema: mpAttachmentPolicySchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CreateTagMaskingPolicyAssociation implements schema.CreateFunc
func CreateTagMaskingPolicyAssociation(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	tagId := d.Get("tag_id").(string)
	tagIdStruct, idErr := tagIDFromString(tagId)
	if idErr != nil {
		return idErr
	}
	tagDb := tagIdStruct.DatabaseName
	tagSchema := tagIdStruct.SchemaName
	tagName := tagIdStruct.TagName

	mpId := d.Get("masking_policy_id").(string)
	mpIdStruct, mpIdErr := maskingPolicyIDFromString(mpId)
	if mpIdErr != nil {
		return mpIdErr
	}
	mpDb := mpIdStruct.DatabaseName
	mpSchema := mpIdStruct.SchemaName
	mpName := mpIdStruct.MaskingPolicyName

	mP := snowflake.MaskingPolicy(mpName, mpDb, mpSchema)
	builder := snowflake.Tag(tagName).WithDB(tagDb).WithSchema(tagSchema).WithMaskingPolicy(mP)

	q := builder.AddMaskingPolicy()

	err := snowflake.Exec(db, q)

	if err != nil {
		return errors.Wrapf(err, "error attaching masking policy %v to tag %v", mpName, tagName)
	}

	mpID := &attachmentID{
		TagDatabaseName:           tagDb,
		TagSchemaName:             tagSchema,
		TagName:                   tagName,
		MaskingPolicyDatabaseName: mpDb,
		MaskingPolicySchemaName:   mpSchema,
		MaskingPolicyName:         mpName,
	}
	dataIDInput, err := mpID.String()
	if err != nil {
		return err
	}
	d.SetId(dataIDInput)

	return ReadTagMaskingPolicyAssociation(d, meta)
}

// ReadTagTagMaskingPolicyAssociation implements schema.ReadFunc
func ReadTagMaskingPolicyAssociation(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	attachementID, err := attachedPolicyIDFromString(d.Id())
	if err != nil {
		return err
	}

	tagDbName := attachementID.TagDatabaseName
	tagSchemaName := attachementID.TagSchemaName
	tagName := attachementID.TagName
	mpDbName := attachementID.MaskingPolicyDatabaseName
	mpSchameName := attachementID.MaskingPolicySchemaName
	mpName := attachementID.MaskingPolicyName

	mP := snowflake.MaskingPolicy(mpName, mpDbName, mpSchameName)
	builder := snowflake.Tag(tagName).WithDB(tagDbName).WithSchema(tagSchemaName).WithMaskingPolicy(mP)

	row := snowflake.QueryRow(db, builder.ShowAttachedPolicy())
	t, err := snowflake.ScanTagPolicy(row)
	if err == sql.ErrNoRows {
		// If not found, mark resource to be removed from statefile during apply or refresh
		log.Printf("[DEBUG] attached policy (%s) not found", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return err
	}

	tagId := TagID{
		DatabaseName: t.RefDb.String,
		SchemaName:   t.RefSchema.String,
		TagName:      t.RefEntity.String,
	}

	tagIdString, err := tagId.String()
	if err != nil {
		return err
	}

	mpId := maskingPolicyID{
		DatabaseName:      t.PolicyDb.String,
		SchemaName:        t.PolicySchema.String,
		MaskingPolicyName: t.PolicyName.String,
	}

	mpIdString, err := mpId.String()
	if err != nil {
		return err
	}
	err = d.Set("tag_id", tagIdString)

	if err != nil {
		return err
	}

	err = d.Set("masking_policy_id", mpIdString)

	if err != nil {
		return err
	}

	return nil
}

// DeleteTagMaskingPolicyAssociation implements schema.DeleteFunc
func DeleteTagMaskingPolicyAssociation(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	attachementID, err := attachedPolicyIDFromString(d.Id())
	if err != nil {
		return err
	}

	tagDbName := attachementID.TagDatabaseName
	tagSchemaName := attachementID.TagSchemaName
	tagName := attachementID.TagName
	mpDbName := attachementID.MaskingPolicyDatabaseName
	mpSchameName := attachementID.MaskingPolicySchemaName
	mpName := attachementID.MaskingPolicyName

	mP := snowflake.MaskingPolicy(mpName, mpDbName, mpSchameName)

	builder := snowflake.Tag(tagName).WithDB(tagDbName).WithSchema(tagSchemaName).WithMaskingPolicy(mP)

	err = snowflake.Exec(db, builder.RemoveMaskingPolicy())
	if err != nil {
		return errors.Wrapf(err, "error unattaching masking policy for %v", d.Id())
	}

	d.SetId("")

	return nil
}
