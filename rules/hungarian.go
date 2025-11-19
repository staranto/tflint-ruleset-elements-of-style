// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"os"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// hungarianRuleConfig represents the configuration for the HungarianRule.
type hungarianRuleConfig struct {
	Defaults []string `hclext:"defaults,optional"`
	More     []string `hclext:"more,optional"`
}

// HungarianRule checks whether a block's type is echoed in its name.
type HungarianRule struct {
	tflint.DefaultRule
	Config hungarianRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *HungarianRule) Check(runner tflint.Runner) error {

	config := hungarianRuleConfig{}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode rule config for %s: %v\n", r.Name(), err)
		return err
	}
	r.Config = config

	return CheckBlocksAndLocals(runner, allLintableBlocks, r, checkForHungarian)
}

// checkForHungarian checks if the name uses Hungarian notation.
func checkForHungarian(runner tflint.Runner, r *HungarianRule, block *hclext.Block, typ string, name string, _ string) {
	keys := r.Config.Defaults
	if len(keys) == 0 {
		keys = []string{
			"str",
			"int", "num",
			"bool",
			"list", "lst", "set", "map", "arr"} // default
	}

	keys = append(keys, r.Config.More...)

	for _, k := range keys {
		if strings.HasPrefix(name, k) || strings.HasSuffix(name, k) || strings.Contains(name, "_"+k) {
			runner.EmitIssue(r, fmt.Sprintf("'%s' uses Hungarian notation with '%s'", name, k),
				block.DefRange)
			return
		}
	}
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
	return tflint.WARNING
}

// NewHungarianRule returns a new rule.
func NewHungarianRule() *HungarianRule {
	return &HungarianRule{}
}
