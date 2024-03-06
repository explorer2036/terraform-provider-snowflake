resource "snowflake_tag" "test" {
  name     = var.tag_name
  database = var.database
  schema   = var.schema
}

resource "snowflake_table" "test" {
  name     = var.table_name
  database = var.database
  schema   = var.schema

  column {
    name = "column_name"
    type = "VARIANT"
  }
}

resource "snowflake_tag_association" "test" {
  object_identifier {
    database = var.database
    schema   = var.schema
    name     = "${snowflake_table.test.name}.${snowflake_table.test.column[0].name}"
  }

  object_type = "COLUMN"
  tag_id      = snowflake_tag.test.id
  tag_value   = "TAG_VALUE"
}
