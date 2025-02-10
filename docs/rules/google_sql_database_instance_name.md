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

// TODO: Write the output when inspects the above code

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
