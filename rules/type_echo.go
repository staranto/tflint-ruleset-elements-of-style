package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Check checks whether ...
func (r *TypeEchoRule) Check(runner tflint.Runner) error {
	return r.walkModules(runner)
}

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
		runner.EmitIssueWithFix(
			r,
			fmt.Sprintf("The type \"%s\" is echoed in the label \"%s\"", typ, name),
			block.DefRange,
			func(f tflint.Fixer) error { return f.RemoveExtBlock(block) },
		)
	}
}

// TypeEchoRule checks whether ...
type TypeEchoRule struct {
	tflint.DefaultRule
}

// NewTypeEchoRule returns a new rule
func NewTypeEchoRule() *TypeEchoRule {
	return &TypeEchoRule{}
}

// Name returns the rule name
func (r *TypeEchoRule) Name() string {
	return "type_echo"
}

// Enabled returns whether the rule is enabled by default
func (r *TypeEchoRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TypeEchoRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TypeEchoRule) Link() string {
	return "https://www.google.com"
}
