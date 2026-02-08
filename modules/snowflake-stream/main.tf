# Snowflake Stream Resource
# Creates and manages one or more Snowflake streams on tables based on the stream_configs map

resource "snowflake_stream_on_table" "this" {
  for_each = var.stream_configs

  name              = each.value.name
  database          = each.value.database
  schema            = each.value.schema
  table             = "${each.value.database}.${each.value.schema}.${each.value.table}"
  append_only       = each.value.append_only
  show_initial_rows = each.value.show_initial_rows
  copy_grants       = each.value.copy_grants
  comment           = each.value.comment
}
