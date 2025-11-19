// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ReminderRule checks for reminders.
type ReminderRule struct {
	tflint.DefaultRule
}

// Check checks whether the rule conditions are met.
func (r *ReminderRule) Check(runner tflint.Runner) error {
	return CheckBlocksAndLocals(runner, allLintableBlocks, r, checkReminder)
}

// checkReminder checks for reminders.
func checkReminder(runner tflint.Runner, r *ReminderRule, block *hclext.Block, typ string, name string, synonym string) {
	// return null for now
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
	return tflint.WARNING
}

// NewReminderRule returns a new rule.
func NewReminderRule() *ReminderRule {
	return &ReminderRule{}
}
