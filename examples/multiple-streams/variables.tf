variable "stream_configs" {
  description = "Map of configuration objects for Snowflake streams on tables"
  type = map(object({
    name              = string
    database          = string
    schema            = string
    table             = string
    append_only       = optional(string, null)
    show_initial_rows = optional(string, null)
    copy_grants       = optional(bool, false)
    comment           = optional(string, null)
  }))
  default = {
    "orders_stream" = {
      name              = "SN_TEST_ORDERS_STREAM"
      database          = "TEST_DB"
      schema            = "PUBLIC"
      table             = "ORDERS"
      append_only       = "false"
      show_initial_rows = "false"
      copy_grants       = false
      comment           = "Stream for tracking changes on orders table"
    }
    "customers_stream" = {
      name              = "SN_TEST_CUSTOMERS_STREAM"
      database          = "TEST_DB"
      schema            = "PUBLIC"
      table             = "CUSTOMERS"
      append_only       = "true"
      show_initial_rows = "false"
      copy_grants       = false
      comment           = "Append-only stream for customers table"
    }
    "products_stream" = {
      name              = "SN_TEST_PRODUCTS_STREAM"
      database          = "TEST_DB"
      schema            = "PUBLIC"
      table             = "PRODUCTS"
      append_only       = "false"
      show_initial_rows = "true"
      copy_grants       = false
      comment           = "Stream for products table with initial rows"
    }
  }
}

# Snowflake authentication variables
variable "snowflake_organization_name" {
  description = "Snowflake organization name"
  type        = string
  default     = null
}

variable "snowflake_account_name" {
  description = "Snowflake account name"
  type        = string
  default     = null
}

variable "snowflake_user" {
  description = "Snowflake username"
  type        = string
  default     = null
}

variable "snowflake_role" {
  description = "Snowflake role"
  type        = string
  default     = null
}

variable "snowflake_private_key" {
  description = "Snowflake private key for key-pair authentication"
  type        = string
  sensitive   = true
  default     = null
}
