// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
)

// BlockDef represents a block definition for schema generation.
type BlockDef struct {
	Typ     string
	Labels  []string
	Synonym string
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
	logger.Debug(fmt.Sprintf("#### block=%v", block))

	var name string
	var typ string

	if len(block.Labels) == 2 {
		typ = block.Labels[0]
		name = block.Labels[1]
	} else {
		typ = block.Type
		name = block.Labels[0]
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
