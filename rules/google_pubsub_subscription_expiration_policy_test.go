package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GooglePubsubSubscriptionExpirationPolicy(t *testing.T) {
	rule := NewGooglePubsubSubscriptionExpirationPolicyRule()

	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		// Missing expiration_policy block, should trigger an issue
		{
			Name: "no_expiration_policy",
			Content: `
resource "google_pubsub_subscription" "this" {
  name  = "this-subscription"
  topic = "this-topic"
}
`,
			Expected: helper.Issues{
				{
					Rule: rule,
					Message: "set expiration_policy explicitly or disable it with \"ttl = ''\", otherwise it defaults to 31 days",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},

		// expiration_policy block is present, no issue should be triggered
		{
			Name: "valid_expiration_policy_with_ttl",
			Content: `
resource "google_pubsub_subscription" "valid" {
  name  = "valid-subscription"
  topic = "valid-topic"

  expiration_policy {
    ttl = "3600s"
  }
}
`,
			Expected: helper.Issues{},
		},

		// explicitly disabling expiration_policy, no issue should be triggered
		{
			Name: "valid_expiration_policy_disabled",
			Content: `
resource "google_pubsub_subscription" "disabled" {
  name  = "disabled-subscription"
  topic = "disabled-topic"

  expiration_policy {
    ttl = ""
  }
}
`,
			Expected: helper.Issues{},
		},

		// multiple resources, 2nd one is missing expiration_policy, should trigger an issue
		{
			Name: "mixed_valid_invalid_resources",
			Content: `
resource "google_pubsub_subscription" "valid" {
  name  = "valid-subscription"
  topic = "valid-topic"

  expiration_policy {
    ttl = "3600s"
  }
}

resource "google_pubsub_subscription" "invalid" {
  name  = "invalid-subscription"
  topic = "invalid-topic"
}
`,
			Expected: helper.Issues{
				{
					Rule: rule,
					Message: "set expiration_policy explicitly or disable it with \"ttl = ''\", otherwise it defaults to 31 days",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 11, Column: 1},
						End:      hcl.Pos{Line: 11, Column: 48},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
