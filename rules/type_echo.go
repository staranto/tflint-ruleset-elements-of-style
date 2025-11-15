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
	Config typeEchoRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *TypeEchoRule) Check(runner tflint.Runner) error {

	config := typeEchoRuleConfig{}

	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	r.Config = config

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
			{
				Type:       "check",
				LabelNames: []string{"name"},
				Body:       &hclext.BodySchema{},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	// Process data blocks
	for _, block := range body.Blocks {
		logger.Debug(fmt.Sprintf("#### block=%v", block))

		var name string
		var typ string

		if block.Type == "check" {
			typ = "check"
			name = block.Labels[0]
		} else {
			typ = block.Labels[0]
			name = block.Labels[1]
		}

		checkForEcho(runner, r, block, typ, name)
	}

	return nil
}

// Enabled returns whether the rule is enabled by default
func (r *TypeEchoRule) Enabled() bool {
	return true
}

// Link returns the rule reference link
func (r *TypeEchoRule) Link() string {
	return "https://www.example.com/type_echo"
}

// Name returns the rule name.
func (r *TypeEchoRule) Name() string {
	return "eos_type_echo"
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
	Synonyms map[string][]string `hclext:"synonyms,optional"`
}

// checkForEcho checks if the type is echoed in the name.
func checkForEcho(runner tflint.Runner, r *TypeEchoRule, block *hclext.Block, typ string, name string) {
	echo := false
	synonymText := ""

	lowerTyp := strings.ToLower(typ)   // aws_s3_bucket
	lowerName := strings.ToLower(name) // my_bucket
	splitName := strings.SplitSeq(lowerName, "_-")

	for part := range strings.SplitSeq(lowerTyp, "_") {
		logger.Debug(fmt.Sprintf("checking if '%s' contains part '%s'", lowerName, part))
		if strings.Contains(lowerName, part) {
			echo = true
			break
		}

		// Check synonyms
		if synonyms, ok := r.Config.Synonyms[part]; ok {
			for _, syn := range synonyms {
				for n := range splitName {
					logger.Debug(fmt.Sprintf("checking if synonym '%s' matches name part '%v'", syn, n))
					if strings.Contains(n, syn) {
						echo = true
						synonymText = fmt.Sprintf(" (via synonym '%s')", syn)
						break
					}
				}
			}

			if echo {
				break
			}
		}
	}

	logger.Debug(fmt.Sprintf("echo=%v for type='%s' name='%s'", echo, typ, name))

	if echo {
		logger.Debug(fmt.Sprintf("emiting issue for type='%s' name='%s'", typ, name))
		runner.EmitIssue(
			r,
			fmt.Sprintf("The type \"%s\" is echoed %s in the label \"%s\"", typ, synonymText, name),
			block.DefRange,
		)
	}
}
