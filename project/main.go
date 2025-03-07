package project

import "fmt"

// Version is ruleset version
const Version string = "0.31.2"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/epidemicsound/tflint-ruleset-google/blob/v%s/docs/rules/%s.md", Version, name)
}
