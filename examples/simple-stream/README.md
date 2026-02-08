# Basic Example - Single Stream

This example demonstrates how to create a single Snowflake stream on a table using the module.

## Usage

```hcl
module "stream" {
  source = "../../modules/snowflake-stream"

  stream_configs = {
    "my_stream" = {
      name              = "MY_STREAM"
      database          = "MY_DATABASE"
      schema            = "PUBLIC"
      table             = "MY_TABLE"
      append_only       = "false"
      show_initial_rows = "false"
      copy_grants       = false
      comment           = "Stream for tracking changes on my table"
    }
  }
}
```

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|----------|
| stream_configs | Map of stream configuration objects | map(object) | yes |

## Outputs

| Name | Description |
|------|-------------|
| stream_names | The names of the created streams |
| stream_fully_qualified_names | The fully qualified names of the streams |
| stream_databases | The databases of the streams |
| stream_schemas | The schemas of the streams |
| stream_tables | The source tables of the streams |
