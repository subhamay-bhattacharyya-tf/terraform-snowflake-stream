output "stream_names" {
  description = "The names of the created streams"
  value       = module.stream.stream_names
}

output "stream_fully_qualified_names" {
  description = "The fully qualified names of the streams"
  value       = module.stream.stream_fully_qualified_names
}

output "stream_databases" {
  description = "The databases of the streams"
  value       = module.stream.stream_databases
}

output "stream_schemas" {
  description = "The schemas of the streams"
  value       = module.stream.stream_schemas
}

output "stream_tables" {
  description = "The source tables of the streams"
  value       = module.stream.stream_tables
}
