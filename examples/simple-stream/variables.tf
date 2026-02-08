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
  default = {}
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
