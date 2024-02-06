resource "snowflake_procedure" "p" {
  database            = var.database
  schema              = var.schema
  name                = var.name
  language            = "SQL"
  return_type         = "VARCHAR"
  execute_as          = "CALLER"
  comment             = var.comment
  statement           = <<EOT
    BEGIN
			RETURN message;
		END;
  EOT
}
