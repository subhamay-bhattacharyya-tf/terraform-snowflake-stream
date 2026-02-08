# Changelog

All notable changes to this project will be documented in this file.

## 1.0.0 (2026-02-08)

### ⚠ BREAKING CHANGES

* This module now manages Snowflake streams instead of warehouses.
All warehouse-related resources, variables, and outputs have been replaced with
stream equival

- Replace snowflake-warehouse module with new snowflake-stream module
- Update CI/CD pipeline to test stream examples instead of warehouse examples
- Migrate examples from basic/multiple-warehouses to simple-stream/multiple-streams
- Remove warehouse-specific design and requirements documentation
- Update test suite with TestSingleStream and TestMultipleStreams
- Consolidate release workflow into main CI pipeline with semantic-release plugins
- Remove legacy utility scripts and warehouse module files
- Update dependencies in package.json and Go modules
- Refactor module to focus on Snowflake stream management instead of warehouse provisioning

### Features

* replace snowflake-warehouse module with snowflake-stream module ([572fdf9](https://github.com/subhamay-bhattacharyya-tf/terraform-snowflake-stream/commit/572fdf97c65c172a03f2520a06c38f7fe82c4fc8))
* **snowflake-stream:** use fully qualified table name in stream resource ([d521253](https://github.com/subhamay-bhattacharyya-tf/terraform-snowflake-stream/commit/d5212530b9a84bfcf9cb3e6401038852c3058db7))

## [unreleased]

### 🚀 Features

- [**breaking**] Replace snowflake-warehouse module with snowflake-stream module
- *(snowflake-stream)* Use fully qualified table name in stream resource

### 📚 Documentation

- *(readme)* Add custom endpoint badge and improve formatting
- Update CHANGELOG.md [skip ci]
