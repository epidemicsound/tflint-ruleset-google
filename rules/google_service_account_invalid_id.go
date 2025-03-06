package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleServiceAccountInvalidIDRule checks whether google_service_account resource has an invalid id
type GoogleServiceAccountInvalidIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleServiceAccountInvalidIDRule returns new rule with default attributes
func NewGoogleServiceAccountInvalidIDRule() *GoogleServiceAccountInvalidIDRule {
	return &GoogleServiceAccountInvalidIDRule{
		resourceType:  "google_service_account",
		attributeName: "account_id",
	}
}

// Name returns the rule name
func (r *GoogleServiceAccountInvalidIDRule) Name() string {
	return "google_service_account_invalid_id"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleServiceAccountInvalidIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleServiceAccountInvalidIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleServiceAccountInvalidIDRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether google_service_account id matches [a-z]([-a-z0-9]*[a-z0-9]) and length is between 6 and 30
func (r *GoogleServiceAccountInvalidIDRule) Check(runner tflint.Runner) error {

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		// Check if name contains only lowercase letters, numbers, and hyphens
		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			// regex to check if name contains only lowercase letters, numbers, and hyphens
			validateFunc := validateRegexp(`^[a-z]([-a-z0-9]*[a-z0-9])$`)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}

		// Check if name length is between 6 and 30
		err = runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validateLength(6, 30)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
