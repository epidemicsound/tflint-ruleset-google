# google_service_account_invalid_id

The rule checks that the id of the service account is 6-30 characters long, and matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])` to comply with RFC1035.

## Example

```hcl
resource "google_service_account" "foo" {
  account_id = "journey_4_gha"
}
```

```
$ tflint

Error: "account_id" ("journey_4_gha-gha") doesn't match regexp "^[a-z]([-a-z0-9]*[a-z0-9])$" (google_service_account_invalid_id)

  on journey_4.tf line 71:
  71:   app        = "journey_4_gha"
```

## Why

It's a requirement from GCP side (https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_service_account#account_id-1)

## How To Fix

Provide the correct `account_id`
