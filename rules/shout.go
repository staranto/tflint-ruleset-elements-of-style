// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"unicode"

	"github.com/staranto/tflint-ruleset-elements-of-style/terraform"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ShoutRule checks whether a block's type is shouted in its name.
type ShoutRule struct {
	tflint.DefaultRule
}

// Check checks whether the rule conditions are met.
func (r *ShoutRule) Check(runner tflint.Runner) error {
	myBlocks := []BlockDef{
		{Typ: "data", Labels: []string{"type", "name"}},
		{Typ: "resource", Labels: []string{"type", "name"}},
		{Typ: "check", Labels: []string{"name"}},
		{Typ: "output", Labels: []string{"name"}},
	}

	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: buildBlockSchemas(myBlocks),
	}, nil)

	if err != nil {
		return err
	}

	// Process data blocks
	for _, block := range body.Blocks {
		typ, name, synonym := normalizeBlock(block, myBlocks)
		checkForShout(runner, r, block, typ, name, synonym)
	}

	// Wrap the runner to access custom methods
	myRunner := terraform.NewRunner(runner)
	locals, diags := myRunner.GetLocals()
	if diags.HasErrors() {
		return diags
	}

	for name, local := range locals {
		logger.Debug(fmt.Sprintf("#### SHOUT local name='%s' value='%v'", name, local))
		checkForShout(runner, r, &hclext.Block{DefRange: local.DefRange}, "local", name, "")
	}

	return nil
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
		runner.EmitIssue(r, fmt.Sprintf("'%s' should not be all uppercase", name),
			block.DefRange)
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

// shoutRuleConfig represents the configuration for the ShoutRule.
type shoutRuleConfig struct {
	// Ignore provider prefix
}
