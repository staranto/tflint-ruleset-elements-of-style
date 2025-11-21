// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// defaultLength is the default maximum length for names.
const defaultLength = 16

// lengthRuleConfig represents the configuration for the LengthRule.
type lengthRuleConfig struct {
	Length int `hclext:"length,optional"`
}

// LengthRule checks whether a block's name is excessively long.
type LengthRule struct {
	tflint.DefaultRule
	Config lengthRuleConfig
}

// Check checks whether the rule conditions are met.
func (rule *LengthRule) Check(runner tflint.Runner) error {
	if err := runner.DecodeRuleConfig(rule.Name(), &rule.Config); err != nil {
		return err
	}
	logger.Debug(fmt.Sprintf("rule.Config=%v", rule.Config))

	return CheckBlocksAndLocals(runner, allLintableBlocks, rule, checkForLength)
}

// checkForLength checks if the type is shouted in the name.
func checkForLength(runner tflint.Runner, r *LengthRule, block *hclext.Block, _ string, name string, _ string) {
	limit := r.Config.Length

	if len(name) > limit {
		message := fmt.Sprintf("'%s' is %d characters and should not be longer than %d.", name, len(name), limit)
		runner.EmitIssue(r, message, block.DefRange)
		logger.Debug(message)
	}
}

// NewLengthRule returns a new rule.
func NewLengthRule() *LengthRule {
	rule := &LengthRule{}
	rule.Config.Length = defaultLength
	return rule
}

// Enabled returns whether the rule is enabled by default
func (rule *LengthRule) Enabled() bool {
	return true
}

// Link returns the rule reference link
func (rule *LengthRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_length.md"
}

// Name returns the rule name.
func (rule *LengthRule) Name() string {
	return "eos_length"
}

// Severity returns the rule severity
func (rule *LengthRule) Severity() tflint.Severity {
	return tflint.WARNING
}
