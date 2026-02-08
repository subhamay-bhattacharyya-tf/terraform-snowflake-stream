output "stream_names" {
  description = "The names of the created streams."
  value       = { for k, v in snowflake_stream_on_table.this : k => v.name }
}

output "stream_fully_qualified_names" {
  description = "The fully qualified names of the streams."
  value       = { for k, v in snowflake_stream_on_table.this : k => v.fully_qualified_name }
}

output "stream_databases" {
  description = "The databases of the streams."
  value       = { for k, v in snowflake_stream_on_table.this : k => v.database }
}

output "stream_schemas" {
  description = "The schemas of the streams."
  value       = { for k, v in snowflake_stream_on_table.this : k => v.schema }
}

output "stream_tables" {
  description = "The source tables of the streams."
  value       = { for k, v in snowflake_stream_on_table.this : k => v.table }
}

output "streams" {
  description = "All stream resources."
  value       = snowflake_stream_on_table.this
}
