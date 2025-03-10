package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GooglePubsubSubscriptionExpirationPolicyRule checks if expiration_policy is set in google_pubsub_subscription
type GooglePubsubSubscriptionExpirationPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGooglePubsubSubscriptionExpirationPolicyRule returns new rule with default attributes
func NewGooglePubsubSubscriptionExpirationPolicyRule() *GooglePubsubSubscriptionExpirationPolicyRule {
	return &GooglePubsubSubscriptionExpirationPolicyRule{
		resourceType:  "google_pubsub_subscription",
		attributeName: "expiration_policy",
	}
}

// Name returns the rule name
func (r *GooglePubsubSubscriptionExpirationPolicyRule) Name() string {
	return "google_pubsub_subscription_expiration_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *GooglePubsubSubscriptionExpirationPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GooglePubsubSubscriptionExpirationPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GooglePubsubSubscriptionExpirationPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Throw error if expiration_policy is missing in google_pubsub_subscription
func (r *GooglePubsubSubscriptionExpirationPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		expirationPolicyExists := false

		// Iterate over blocks to check if "expiration_policy" exists
		for _, block := range resource.Body.Blocks {
			if block.Type == r.attributeName {
				expirationPolicyExists = true
				break
			}
		}

		if !expirationPolicyExists {
			if err := runner.EmitIssue(
				r,
				"set expiration_policy explicitly or disable it with \"ttl = ''\", otherwise it defaults to 31 days",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
