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

  validation {
    condition     = alltrue([for k, s in var.stream_configs : length(s.name) > 0])
    error_message = "Stream name must not be empty."
  }

  validation {
    condition     = alltrue([for k, s in var.stream_configs : length(s.database) > 0])
    error_message = "Database name must not be empty."
  }

  validation {
    condition     = alltrue([for k, s in var.stream_configs : length(s.schema) > 0])
    error_message = "Schema name must not be empty."
  }

  validation {
    condition     = alltrue([for k, s in var.stream_configs : length(s.table) > 0])
    error_message = "Table name must not be empty."
  }

  validation {
    condition     = alltrue([for k, s in var.stream_configs : s.append_only == null || contains(["true", "false"], s.append_only)])
    error_message = "append_only must be 'true', 'false', or null."
  }

  validation {
    condition     = alltrue([for k, s in var.stream_configs : s.show_initial_rows == null || contains(["true", "false"], s.show_initial_rows)])
    error_message = "show_initial_rows must be 'true', 'false', or null."
  }
}
