// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const defaultCommentsBlocked = true
const defaultCommentsColumn = 80
const defaultCommentsJammed = true

// commentsRuleConfig represents the configuration for the CommentsRule.
type commentsRuleConfig struct {
	Block  bool `hclext:"block,optional"`
	Column int  `hclext:"column,optional"`
	Jammed bool `hclext:"jammed,optional"`
}

// CommentsRule checks for comment style.
type CommentsRule struct {
	tflint.DefaultRule
	Config commentsRuleConfig
}

// Check checks whether the rule conditions are met.
func (r *CommentsRule) Check(runner tflint.Runner) error {
	if err := runner.DecodeRuleConfig(r.Name(), &r.Config); err != nil {
		return err
	}

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		return nil
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	for name, file := range files {
		if err := r.checkComments(runner, name, file); err != nil {
			return err
		}
	}

	return nil
}

func (r *CommentsRule) checkComments(runner tflint.Runner, filename string, file *hcl.File) error {
	tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return diags
	}

	for _, token := range tokens {
		if token.Type != hclsyntax.TokenComment {
			continue
		}

		text := string(token.Bytes)

		// Check for a block comment if enabled.
		if r.Config.Block {
			if strings.HasPrefix(text, "/*") {
				message := "Block comments not allowed."
				runner.EmitIssue(r, message, token.Range)
				logger.Debug(message)
			}
		}

		// Check for jammed comments if enabled.
		if r.Config.Jammed {
			isJammed := false
			if strings.HasPrefix(text, "//") {
				if len(text) > 2 && text[2] != ' ' {
					isJammed = true
				}
			} else if strings.HasPrefix(text, "#") {
				if len(text) > 1 && text[1] != ' ' {
					isJammed = true
				}
			}

			if isJammed {
				trimmed := strings.TrimSpace(text)
				rns := []rune(trimmed)
				snippet := trimmed
				if len(rns) > 5 {
					snippet = string(rns[:5])
				}
				message := fmt.Sprintf("Comment is jammed ('%s ...').", snippet)
				runner.EmitIssue(r, message, token.Range)
				logger.Debug(message)
			}
		}

		// Check for comment extending beyond comment limit.
		if r.Config.Column > 0 {
			trimmedText := strings.TrimRight(text, "\r\n")
			end := token.Range.Start.Column + len(trimmedText) - 1

			if end > r.Config.Column {
				message := fmt.Sprintf("Comment extends beyond column %d to %d.", r.Config.Column, end)
				runner.EmitIssue(r, message, token.Range)
				logger.Debug(message)

			}
		}
	}

	return nil
}

// NewCommentsRule returns a new rule.
func NewCommentsRule() *CommentsRule {
	rule := &CommentsRule{}
	rule.Config = commentsRuleConfig{
		Block:  defaultCommentsBlocked,
		Column: defaultCommentsColumn,
		Jammed: defaultCommentsJammed,
	}

	return rule
}

// Enabled returns whether the rule is enabled by default.
func (r *CommentsRule) Enabled() bool {
	return true
}

// Link returns the rule link.
func (r *CommentsRule) Link() string {
	return "https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_comments.md"
}

// Name returns the rule name.
func (r *CommentsRule) Name() string {
	return "eos_comments"
}

// Severity returns the rule severity.
func (r *CommentsRule) Severity() tflint.Severity {
	return tflint.WARNING
}
