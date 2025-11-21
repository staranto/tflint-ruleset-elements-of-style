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

// typeEchoRuleConfig represents the configuration for the TypeEchoRule.
type typeEchoRuleConfig struct {
	Synonyms map[string][]string `hclext:"synonyms,optional"`
	Level    string              `hclext:"level,optional"`
}

var defaultTypeEchoConfig = typeEchoRuleConfig{
	Level: "warning",
}

// TypeEchoRule checks whether a block's type is echoed in its name.
type TypeEchoRule struct {
	tflint.DefaultRule
	Config typeEchoRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *TypeEchoRule) Check(runner tflint.Runner) error {
	if err := runner.DecodeRuleConfig(r.Name(), &r.Config); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode rule config for %s: %v\n", r.Name(), err)
		return err
	}

	return CheckBlocksAndLocals(runner, allLintableBlocks, r, checkForEcho)
}

// checkForEcho checks if the type is echoed in the name.
func checkForEcho(runner tflint.Runner,
	r *TypeEchoRule, block *hclext.Block,
	typ string, name string, synonym string) {

	// logger.Debug(fmt.Sprintf("checking for echo in type='%s' name='%s'", typ, name))
	echo := false

	lowerTyp := strings.ToLower(typ)   // aws_s3_bucket
	lowerName := strings.ToLower(name) // my_bucket
	splitName := strings.SplitSeq(lowerName, "_-")
	synonymText := ""

	for part := range strings.SplitSeq(lowerTyp, "_") {
		// logger.Debug(fmt.Sprintf("checking if '%s' contains part '%s'", lowerName, part))
		if strings.Contains(lowerName, part) {
			echo = true
			break
		}

		synonyms, _ := r.Config.Synonyms[part]
		if synonym != "" {
			synonyms = append(synonyms, synonym)
		}

		// Check synonyms
		for _, syn := range synonyms {
			for n := range splitName {
				// logger.Debug(fmt.Sprintf("checking if synonym '%s' matches name part '%v'", syn, n))
				if strings.Contains(n, syn) {
					echo = true
					synonymText = fmt.Sprintf(" (via synonym '%s')", syn)
					break
				}
			}

			if echo {
				break
			}
		}
	}

	// logger.Debug(fmt.Sprintf("echo=%v for type='%s' name='%s'", echo, typ, name))

	if echo {
		// logger.Debug(fmt.Sprintf("emiting issue for type='%s' name='%s'", typ, name))
		runner.EmitIssue(
			r,
			fmt.Sprintf("The type \"%s\" is echoed%s in the label \"%s\".", typ, synonymText, name),
			block.DefRange,
		)
	}
}

// NewTypeEchoRule returns a new rule.
func NewTypeEchoRule() *TypeEchoRule {
	rule := &TypeEchoRule{}
	rule.Config = defaultTypeEchoConfig
	return rule
}

// Enabled returns whether the rule is enabled by default.
func (r *TypeEchoRule) Enabled() bool {
	return true
}

// Link returns the rule reference link.
func (r *TypeEchoRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_type_echo.md"
}

// Name returns the rule name.
func (r *TypeEchoRule) Name() string {
	return "eos_type_echo"
}

// Severity returns the rule severity.
func (r *TypeEchoRule) Severity() tflint.Severity {
	return toSeverity(r.Config.Level)
}
