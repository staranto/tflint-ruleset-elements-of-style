// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// LengthRule checks whether a block's type is shouted in its name.
type LengthRule struct {
	tflint.DefaultRule
	Config lengthRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *LengthRule) Check(runner tflint.Runner) error {

	config := lengthRuleConfig{}
	config.Length = 16

	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	r.Config = config

	myBlocks := []BlockDef{
		{Typ: "data", Labels: []string{"type", "name"}},
		{Typ: "resource", Labels: []string{"type", "name"}},
		{Typ: "check", Labels: []string{"name"}},
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
		checkForLength(runner, r, block, typ, name, synonym)
	}

	return nil
}

// checkForLength checks if the type is shouted in the name.
func checkForLength(runner tflint.Runner, r *LengthRule, block *hclext.Block, _ string, name string, _ string) {
	limit := r.Config.Length

	if len(name) > limit {
		runner.EmitIssue(r, fmt.Sprintf("'%s' is %d characters and should not be longer than %d", name, len(name), limit),
			block.DefRange)
	}
}

// Enabled returns whether the rule is enabled by default
func (r *LengthRule) Enabled() bool {
	return true
}

// Link returns the rule reference link
func (r *LengthRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_length.md"
}

// Name returns the rule name.
func (r *LengthRule) Name() string {
	return "eos_length"
}

// Severity returns the rule severity
func (r *LengthRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// NewLengthRule returns a new rule.
func NewLengthRule() *LengthRule {
	return &LengthRule{}
}

// lengthRuleConfig represents the configuration for the LengthRule.
type lengthRuleConfig struct {
	Length int `hclext:"length,optional"`
}
