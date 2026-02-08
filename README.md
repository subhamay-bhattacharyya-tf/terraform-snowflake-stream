# Terraform Snowflake Module - Stream

![Release](https://github.com/subhamay-bhattacharyya-tf/terraform-snowflake-stream/actions/workflows/ci.yaml/badge.svg)&nbsp;![Snowflake](https://img.shields.io/badge/Snowflake-29B5E8?logo=snowflake&logoColor=white)&nbsp;![Commit Activity](https://img.shields.io/github/commit-activity/t/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![Last Commit](https://img.shields.io/github/last-commit/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![Release Date](https://img.shields.io/github/release-date/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![Repo Size](https://img.shields.io/github/repo-size/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![File Count](https://img.shields.io/github/directory-file-count/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![Issues](https://img.shields.io/github/issues/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![Top Language](https://img.shields.io/github/languages/top/subhamay-bhattacharyya-tf/terraform-snowflake-stream)&nbsp;![Custom Endpoint](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/bsubhamay/1874e744480a2900d630c662ba60cf2e/raw/terraform-snowflake-stream.json?)


A Terraform module for creating and managing Snowflake streams on tables using a map of configuration objects. Supports creating single or multiple streams with a single module call.

## Features

- Map-based configuration for creating single or multiple streams
- Built-in input validation with descriptive error messages
- Sensible defaults for optional properties
- Outputs keyed by stream identifier for easy reference
- Support for append-only and standard stream modes

## Usage

### Single Stream

```hcl
module "stream" {
  source = "path/to/modules/snowflake-stream"

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

### Multiple Streams

```hcl
locals {
  streams = {
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

module "streams" {
  source = "path/to/modules/snowflake-stream"

  stream_configs = local.streams
}
```

## Examples

- [Simple Stream](examples/simple-stream) - Create a single stream
- [Multiple Streams](examples/multiple-streams) - Create multiple streams

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.3.0 |
| snowflake | >= 0.87.0 |

## Providers

| Name | Version |
|------|---------|
| snowflake | >= 0.87.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| stream_configs | Map of configuration objects for Snowflake streams | `map(object)` | `{}` | no |

### stream_configs Object Properties

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| name | string | - | Stream identifier (required) |
| database | string | - | Database where the stream will be created (required) |
| schema | string | - | Schema where the stream will be created (required) |
| table | string | - | Source table to monitor (required) |
| append_only | string | null | Whether this is an append-only stream ("true" or "false") |
| show_initial_rows | string | null | Whether to include initial rows ("true" or "false") |
| copy_grants | bool | false | Whether to copy grants from the source |
| comment | string | null | Description of the stream |

## Outputs

| Name | Description |
|------|-------------|
| stream_names | Map of stream names keyed by identifier |
| stream_fully_qualified_names | Map of fully qualified stream names |
| stream_databases | Map of stream databases |
| stream_schemas | Map of stream schemas |
| stream_tables | Map of source tables |
| streams | All stream resources |

## Validation

The module validates inputs and provides descriptive error messages for:

- Empty stream name
- Empty database name
- Empty schema name
- Empty table name
- Invalid append_only value (must be "true", "false", or null)
- Invalid show_initial_rows value (must be "true", "false", or null)

## Testing

The module includes Terratest-based integration tests:

```bash
cd test
go mod tidy
go test -v -timeout 30m
```

Required environment variables for testing:
- `SNOWFLAKE_ORGANIZATION_NAME` - Snowflake organization name
- `SNOWFLAKE_ACCOUNT_NAME` - Snowflake account name
- `SNOWFLAKE_USER` - Snowflake username
- `SNOWFLAKE_ROLE` - Snowflake role (e.g., "SYSADMIN")
- `SNOWFLAKE_PRIVATE_KEY` - Snowflake private key for key-pair authentication

## CI/CD Configuration

The CI workflow runs on:
- Push to `main`, `feature/**`, and `bug/**` branches (when `modules/**` changes)
- Pull requests to `main` (when `modules/**` changes)
- Manual workflow dispatch

The workflow includes:
- Terraform validation and format checking
- Examples validation
- Terratest integration tests (output displayed in GitHub Step Summary)
- Changelog generation (non-main branches)
- Semantic release (main branch only)

The CI workflow uses the following GitHub organization variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `TERRAFORM_VERSION` | Terraform version for CI jobs | `1.3.0` |
| `GO_VERSION` | Go version for Terratest | `1.21` |
| `SNOWFLAKE_ORGANIZATION_NAME` | Snowflake organization name | - |
| `SNOWFLAKE_ACCOUNT_NAME` | Snowflake account name | - |
| `SNOWFLAKE_USER` | Snowflake username | - |
| `SNOWFLAKE_ROLE` | Snowflake role (e.g., SYSADMIN) | - |

The following GitHub secrets are required for Terratest integration tests:

| Secret | Description | Required |
|--------|-------------|----------|
| `SNOWFLAKE_PRIVATE_KEY` | Snowflake private key for key-pair authentication | Yes |

## License

MIT License - See [LICENSE](LICENSE) for details.
