package main

import (
	"github.com/staranto/tflint-ruleset-elements-of-style/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "type-echo",
			Version: "0.2.1",
			Rules: []tflint.Rule{
				rules.NewLengthRule(),
				rules.NewShoutRule(),
				rules.NewTypeEchoRule(),
			},
		},
	})
}
