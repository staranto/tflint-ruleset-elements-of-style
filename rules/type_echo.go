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

// TypeEchoRule checks whether a block's type is echoed in its name.
type TypeEchoRule struct {
	tflint.DefaultRule
}

// Check checks whether the rule conditions are met.
func (r *TypeEchoRule) Check(runner tflint.Runner) error {
	return r.walkModules(runner)
}

// Enabled returns whether the rule is enabled by default
func (r *TypeEchoRule) Enabled() bool {
	return true
}

// Link returns the rule reference link
func (r *TypeEchoRule) Link() string {
	return "https://www.example.com/blah"
}

// Name returns the rule name.
func (r *TypeEchoRule) Name() string {
	return "type_echo"
}

// Severity returns the rule severity
func (r *TypeEchoRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// NewTypeEchoRule returns a new rule.
func NewTypeEchoRule() *TypeEchoRule {
	return &TypeEchoRule{}
}

// typeEchoRuleConfig represents the configuration for the TypeEchoRule.
type typeEchoRuleConfig struct {
	// Ignore provider prefix
}

// walkModules walks through the modules and checks for type echoes.
func (r *TypeEchoRule) walkModules(runner tflint.Runner) error {
	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "data",
				LabelNames: []string{"type", "name"},
				Body:       &hclext.BodySchema{},
			},
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body:       &hclext.BodySchema{},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	blocks := body.Blocks.ByType()

	// Process data blocks
	for _, block := range blocks["data"] {
		checkForEcho(runner, r, block, block.Labels[0], block.Labels[1])
	}

	// Process resource blocks
	for _, block := range blocks["resource"] {
		checkForEcho(runner, r, block, block.Labels[0], block.Labels[1])
	}

	return nil
}

// checkForEcho checks if the type is echoed in the name.
func checkForEcho(runner tflint.Runner, r *TypeEchoRule, block *hclext.Block, typ string, name string) {
	echo := false

	for part := range strings.SplitSeq(typ, "_") {
		logger.Debug(fmt.Sprintf("checking if '%s' contains part '%s'", name, part))
		if strings.Contains(name, part) {
			echo = true
			break
		}
	}

	logger.Debug(fmt.Sprintf("echo=%v for type='%s' name='%s'", echo, typ, name))

	if echo {
		logger.Debug(fmt.Sprintf("emiting issue for type='%s' name='%s'", typ, name))
		runner.EmitIssue(
			r,
			fmt.Sprintf("The type \"%s\" is echoed in the label \"%s\"", typ, name),
			block.DefRange,
		)
	}
}
