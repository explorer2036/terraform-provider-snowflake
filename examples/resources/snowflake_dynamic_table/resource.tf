resource "snowflake_dynamic_table" "current" {
  warehouse  = "warehouse"
  query      = "select id from product"
  target_lag = "2 minutes"
  comment    = "comment"
}
