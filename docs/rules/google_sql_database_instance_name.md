# google_sql_database_instance_name

Validate if `name` of `google_sql_database_instance` contains only lowercase letters, numbers and hyphens.

## Example of invalid resource

```hcl
resource "google_sql_database_instance" "instance" {
  name = "my_instance"
}
```

```
$ tflint

Error: "name" ("postgres_eu_we1-journey-4") doesn't match regexp "^[a-z0-9-]+$" (google_sql_database_instance_name)

  on journey_4.tf line 8:
   8:   name                = "postgres_eu_we1-journey-4"

Reference: https://github.com/terraform-linters/tflint-ruleset-google/blob/v0.31.0/docs/rules/google_sql_database_instance_name.md

```

## Why

That's the limitation from GCP API.

## How To Fix

Use the name that contains only lowercase letters, numbers and hyphens.
