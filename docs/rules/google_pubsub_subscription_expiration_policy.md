# google_pubsub_subscription_expiration_policy

By default if no expiration policy is set to pubsub subscriptions, they expire after 31 days with no activity.
We recommend setting this value explicitly to avoid confusion in the future.

## Example

```hcl
resource "google_pubsub_subscription" "this" {
  name  = "this-subscription"
  topic = "this-topic"
}
```

```
$ tflint

Error: set expiration_policy explicitly or disable it with "ttl = ''", otherwise it defaults to 31 days (google_pubsub_subscription_expiration_policy)

  on this.tf line 21:
 291: resource "google_pubsub_subscription" "this" {

Reference: https://github.com/epidemicsound/tflint-ruleset-google/blob/v0.31.2/docs/rules/google_pubsub_subscription_expiration_policy.md

```

## Why

It's not clear that the default `ttl` is 31 days if the value is unset. We enforce the setting to be set explicitly to avoid confusion in the future.

## How To Fix

Set the `expiration_policy`.
