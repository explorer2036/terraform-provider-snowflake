resource "snowflake_procedure" "p" {
    database            = var.database
    schema              = var.schema
    name                = var.name
    language            = "SQL"
    return_type         = "VARCHAR"
    execute_as          = "CALLER"
    null_input_behavior = "RETURNS NULL ON NULL INPUT"
    comment             = var.comment
    statement           = <<EOT
        BEGIN
			RETURN message;
		END;
    EOT
}
