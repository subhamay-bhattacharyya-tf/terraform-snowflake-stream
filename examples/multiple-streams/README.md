# Multiple Streams Example

This example demonstrates how to create multiple Snowflake streams on tables using the module with a map of configurations.

## Usage

```hcl
module "streams" {
  source = "../../modules/snowflake-stream"

  stream_configs = {
    "orders_stream" = {
      name              = "SN_ORDERS_STREAM"
      database          = "SALES_DB"
      schema            = "PUBLIC"
      table             = "ORDERS"
      append_only       = "false"
      show_initial_rows = "false"
      copy_grants       = false
      comment           = "Stream for tracking changes on orders table"
    }
    "customers_stream" = {
      name              = "SN_CUSTOMERS_STREAM"
      database          = "SALES_DB"
      schema            = "PUBLIC"
      table             = "CUSTOMERS"
      append_only       = "true"
      show_initial_rows = "false"
      copy_grants       = false
      comment           = "Append-only stream for customers table"
    }
    "products_stream" = {
      name              = "SN_PRODUCTS_STREAM"
      database          = "SALES_DB"
      schema            = "PUBLIC"
      table             = "PRODUCTS"
      append_only       = "false"
      show_initial_rows = "true"
      copy_grants       = false
      comment           = "Stream for products table with initial rows"
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
| streams | All stream resources |
