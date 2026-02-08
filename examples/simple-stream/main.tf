# Example: Single Snowflake Stream
#
# This example demonstrates how to use the snowflake-stream module
# to create a single Snowflake stream on a table.

module "stream" {
  source = "../../modules/snowflake-stream"

  stream_configs = var.stream_configs
}
