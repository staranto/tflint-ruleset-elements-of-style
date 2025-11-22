// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var defaultHungarianTags = []string{
	"str",
	"int", "num",
	"bool",
	"list", "lst", "set", "map", "arr", "array",
}

var defaultHungarianConfig = hungarianRuleConfig{
	Tags:  defaultHungarianTags,
	Level: "warning",
}

// hungarianRuleConfig represents the configuration for the HungarianRule.
type hungarianRuleConfig struct {
	Tags  []string `hclext:"tags,optional"`
	Level string   `hclext:"level,optional"`
}

// HungarianRule checks whether a block's type is echoed in its name.
type HungarianRule struct {
	tflint.DefaultRule
	Config hungarianRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *HungarianRule) Check(runner tflint.Runner) error {
	if err := runner.DecodeRuleConfig(r.Name(), &r.Config); err != nil {
		return err
	}

	return CheckBlocksAndLocals(runner, allLintableBlocks, r, checkForHungarian)
}

// checkForHungarian checks if the name uses Hungarian notation.
func checkForHungarian(runner tflint.Runner, r *HungarianRule, block *hclext.Block, typ string, name string, _ string) {
	tags := r.Config.Tags

	for _, t := range tags {
		if strings.HasPrefix(name, t) || strings.HasSuffix(name, t) || strings.Contains(name, "_"+t) {
			if err := runner.EmitIssue(r, fmt.Sprintf("'%s' uses Hungarian notation with '%s'.", name, t),
				block.DefRange); err != nil {
				logger.Error(err.Error())
			}
			return
		}
	}
}

// NewHungarianRule returns a new rule.
func NewHungarianRule() *HungarianRule {
	rule := &HungarianRule{}
	rule.Config = defaultHungarianConfig
	return rule
}

// Enabled returns whether the rule is enabled by default.
func (r *HungarianRule) Enabled() bool {
	return true
}

// Link returns the rule reference link.
func (r *HungarianRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_hungarian.md"
}

// Name returns the rule name.
func (r *HungarianRule) Name() string {
	return "eos_hungarian"
}

// Severity returns the rule severity.
func (r *HungarianRule) Severity() tflint.Severity {
	return toSeverity(r.Config.Level)
}
