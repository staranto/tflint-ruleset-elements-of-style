// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"unicode"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// shoutRuleConfig represents the configuration for the ShoutRule.
type shoutRuleConfig struct {
	// Ignore provider prefix
}

// ShoutRule checks whether a block's type is shouted in its name.
type ShoutRule struct {
	tflint.DefaultRule
	Config shoutRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *ShoutRule) Check(runner tflint.Runner) error {
	return CheckBlocksAndLocals(runner, allLintableBlocks, r, checkForShout)
}

// checkForShout checks if the type is shouted in the name.
func checkForShout(runner tflint.Runner, r *ShoutRule, block *hclext.Block, _ string, name string, _ string) {
	hasAlpha := false
	allUpper := true

	for _, ch := range name {
		if unicode.IsLetter(ch) {
			hasAlpha = true
			if !unicode.IsUpper(ch) {
				allUpper = false
			}
		}
	}

	if hasAlpha && allUpper {
		message := fmt.Sprintf("'%s' should not be all uppercase", name)
		runner.EmitIssue(r, message, block.DefRange)
		logger.Debug(message)
	}
}

// Enabled returns whether the rule is enabled by default
func (r *ShoutRule) Enabled() bool {
	return true
}

// Link returns the rule reference link
func (r *ShoutRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_shout.md"
}

// Name returns the rule name.
func (r *ShoutRule) Name() string {
	return "eos_shout"
}

// Severity returns the rule severity
func (r *ShoutRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// NewShoutRule returns a new rule.
func NewShoutRule() *ShoutRule {
	return &ShoutRule{}
}
