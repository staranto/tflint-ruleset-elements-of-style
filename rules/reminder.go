// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var defaultReminderTags = []string{
	"BUG",
	"FIXME",
	"HACK",
	"TODO",
}

var defaultReminderConfig = reminderRuleConfig{
	Tags:  defaultReminderTags,
	Level: "warning",
}

// reminderRuleConfig represents the configuration for the ReminderRule.
type reminderRuleConfig struct {
	Tags  []string `hclext:"tags,optional"`
	Level string   `hclext:"level,optional"`
}

// ReminderRule checks for reminders.
type ReminderRule struct {
	tflint.DefaultRule
	Config reminderRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *ReminderRule) Check(runner tflint.Runner) error {
	if err := runner.DecodeRuleConfig(r.Name(), &r.Config); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode rule config for %s: %v\n", r.Name(), err)
		return err
	}

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	for name, file := range files {
		if err := r.checkReminders(runner, name, file); err != nil {
			return err
		}
	}

	return nil
}

func (r *ReminderRule) checkReminders(runner tflint.Runner, filename string, file *hcl.File) error {
	tags := r.Config.Tags

	tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return diags
	}

	for _, token := range tokens {
		if token.Type != hclsyntax.TokenComment {
			continue
		}

		text := string(token.Bytes)
		tokens := strings.SplitAfterN(strings.ToUpper(text), " ", 2)
		if len(tokens) < 2 {
			continue
		}

		for _, t := range tags {
			if strings.HasSuffix(strings.TrimSpace(tokens[0]), t) || strings.HasPrefix(tokens[1], t) {
				message := fmt.Sprintf("'%s' has a reminder tag.", strings.TrimSpace(text))
				runner.EmitIssue(r, message, token.Range)
				logger.Debug(message)
			}
		}
	}

	return nil
}

// NewReminderRule returns a new rule.
func NewReminderRule() *ReminderRule {
	rule := &ReminderRule{}
	rule.Config = defaultReminderConfig
	return rule
}

// Enabled returns whether the rule is enabled by default.
func (r *ReminderRule) Enabled() bool {
	return true
}

// Link returns the rule reference link.
func (r *ReminderRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_reminder.md"
}

// Name returns the rule name.
func (r *ReminderRule) Name() string {
	return "eos_reminder"
}

// Severity returns the rule severity.
func (r *ReminderRule) Severity() tflint.Severity {
	return toSeverity(r.Config.Level)
}
