variable "database" {
  type = string
}

variable "schema" {
  type = string
}

variable "name" {
  type = string
}

variable "api_allowed_prefixes" {
  type = list(string)
}

variable "url_of_proxy_and_resource" {
  type = string
}

variable "comment" {
  type = string
}

resource "snowflake_api_integration" "test_api_int" {
  name                 = var.name
  api_provider         = "aws_api_gateway"
  api_aws_role_arn     = "arn:aws:iam::000000000001:/role/test"
  api_allowed_prefixes = var.api_allowed_prefixes
  enabled              = true
}

resource "snowflake_external_function" "external_function" {
  name                      = var.name
  database                  = var.database
  schema                    = var.schema
  comment                   = var.comment
  return_type               = "VARIANT"
  return_behavior           = "IMMUTABLE"
  api_integration           = snowflake_api_integration.test_api_int.name
  url_of_proxy_and_resource = var.url_of_proxy_and_resource
}
