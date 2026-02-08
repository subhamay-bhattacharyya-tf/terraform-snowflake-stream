# Example: Multiple Snowflake Streams
#
# This example demonstrates how to use the snowflake-stream module
# to create multiple Snowflake streams using a map of configurations.

module "streams" {
  source = "../../modules/snowflake-stream"

  stream_configs = var.stream_configs
}
