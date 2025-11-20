// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"

	"github.com/staranto/tflint-ruleset-elements-of-style/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	log.SetFlags(0)
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "elements-of-style",
			Version: "0.2.1",
			Rules: []tflint.Rule{
				rules.NewHungarianRule(),
				rules.NewLengthRule(),
				rules.NewReminderRule(),
				rules.NewShoutRule(),
				rules.NewTypeEchoRule(),
			},
		},
	})
}
