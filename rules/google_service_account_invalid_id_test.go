package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleServiceAccountInvalidID(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "regex mismatch",
			Content: `
resource "google_service_account" "this" {
  account_id = "my_account"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleServiceAccountInvalidIDRule(),
					Message: `"account_id" ("my_account") doesn't match regexp "^[a-z]([-a-z0-9]*[a-z0-9])$"`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 16},
						End:      hcl.Pos{Line: 3, Column: 28},
					},
				},
			},
		},
		{
			Name: "incorrect length",
			Content: `
resource "google_service_account" "this" {
  account_id = "x123456789x123456789x123456789x"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleServiceAccountInvalidIDRule(),
					Message: `"account_id" ("x123456789x123456789x123456789x") must be 6-30 characters`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 16},
						End:      hcl.Pos{Line: 3, Column: 49},
					},
				},
			},
		},
	}

	rule := NewGoogleServiceAccountInvalidIDRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
