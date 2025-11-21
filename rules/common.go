// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/staranto/tflint-ruleset-elements-of-style/terraform"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// BlockDef represents a block definition for schema generation.
type BlockDef struct {
	Typ     string
	Labels  []string
	Synonym string
}

// allLintableBlocks defines all block types and their label structures to
// check. The order is not important here, but we try to sync up with the order
// found in the *_test.tf sources.
var allLintableBlocks = []BlockDef{
	{Typ: "variable", Labels: []string{"name"}},
	// {Typ: "locals", Labels: []string{}},
	{Typ: "check", Labels: []string{"name"}},
	{Typ: "data", Labels: []string{"type", "name"}},
	{Typ: "ephemeral", Labels: []string{"type", "name"}},
	{Typ: "module", Labels: []string{"name"}},
	{Typ: "output", Labels: []string{"name"}},
	{Typ: "resource", Labels: []string{"type", "name"}},
}

// buildBlockSchemas generates a slice of BlockSchema from block definitions.
func buildBlockSchemas(defs []BlockDef) []hclext.BlockSchema {
	var blocks []hclext.BlockSchema
	for _, def := range defs {
		blocks = append(blocks, hclext.BlockSchema{
			Type:       def.Typ,
			LabelNames: def.Labels,
			Body:       &hclext.BodySchema{},
		})
	}
	return blocks
}

func normalizeBlock(block *hclext.Block, myBlocks []BlockDef) (string, string, string) {
	// logger.Debug(fmt.Sprintf("#### block=%v", block))

	var name string
	var typ string

	if len(block.Labels) == 2 {
		typ = block.Labels[0]
		name = block.Labels[1]
	} else if len(block.Labels) == 1 {
		typ = block.Type
		name = block.Labels[0]
	} else {
		typ = block.Type
		name = ""
	}

	synonym := ""
	for _, def := range myBlocks {
		if def.Typ == typ && def.Synonym != "" {
			synonym = def.Synonym
			break
		}
	}
	return typ, name, synonym
}

// CheckBlocksAndLocals iterates over blocks and locals and applies the check function.
func CheckBlocksAndLocals[T any](
	runner tflint.Runner,
	myBlocks []BlockDef,
	rule T,
	checkFunc func(tflint.Runner, T, *hclext.Block, string, string, string),
) error {
	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: buildBlockSchemas(myBlocks),
	}, nil)

	if err != nil {
		return err
	}
	logger.Debug(fmt.Sprintf("rule=%T body.len=%d", rule, len(body.Blocks)))

	// Check blocks.
	for _, block := range body.Blocks {
		typ, name, synonym := normalizeBlock(block, myBlocks)
		logger.Debug(fmt.Sprintf("typ=%s name=%s synonym=%s", typ, name, synonym))
		checkFunc(runner, rule, block, typ, name, synonym)
	}

	// Process locals. The local {} blocks are, apparently, not included by
	// GetModuleContent(), so we have to get them separately.
	locals, diags := getLocals(runner)
	if diags != nil {
		return diags
	}

	for name, local := range locals {
		checkFunc(runner, rule, &hclext.Block{DefRange: local.DefRange}, "local", name, "")
	}

	return nil
}

// getLocals is a helper function to get local {} blocks since GetModuleContent
// does not.
func getLocals(runner tflint.Runner) (map[string]*terraform.Local, hcl.Diagnostics) {
	myRunner := terraform.NewRunner(runner)
	locals, diags := myRunner.GetLocals()
	if diags.HasErrors() {
		return nil, diags
	}
	return locals, nil
}

// toSeverity converts a string level to a tflint.Severity.
func toSeverity(level string) tflint.Severity {
	switch strings.ToLower(level) {
	case "notice":
		return tflint.NOTICE
	case "warning":
		return tflint.WARNING
	}

	return tflint.ERROR
}
