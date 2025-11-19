// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"flag"
	"fmt"
	"testing"

	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

var typeEchoDeep = flag.Bool("typeEchoDeep", false, "enable deep assert")

func TestTypeEchoRule(t *testing.T) {
	flag.Parse()

	cases := []struct {
		Name    string
		Content string
		Want    helper.Issues
	}{
		{
			Name: "echoed_names",
			Content: func() string {
				content, _ := os.ReadFile("testdata/type_echo_test.tf")
				return string(content)
			}(),
			Want: helper.Issues{
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("variable", "variable_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 7, Column: 1},
						End:      hcl.Pos{Line: 7, Column: 25},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("check", "check_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 13, Column: 1},
						End:      hcl.Pos{Line: 13, Column: 19},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("aws_caller_identity", "caller_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 20, Column: 1},
						End:      hcl.Pos{Line: 20, Column: 41},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("random_password", "password_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 22, Column: 1},
						End:      hcl.Pos{Line: 22, Column: 44},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("module", "module_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 26, Column: 1},
						End:      hcl.Pos{Line: 26, Column: 21},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("output", "output_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 30, Column: 1},
						End:      hcl.Pos{Line: 30, Column: 21},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("aws_instance", "instance_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 35, Column: 1},
						End:      hcl.Pos{Line: 35, Column: 40},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: makeTypeEchoMessage("local", "local_echo"),
					Range: hcl.Range{
						Filename: "type_echo_test.tf",
						Start:    hcl.Pos{Line: 10, Column: 3},
						End:      hcl.Pos{Line: 10, Column: 17},
					},
				},
			},
		},
	}

	for _, tc := range cases {

		// Run the tests and make sure the basic results are found...
		runner := helper.TestRunner(t, map[string]string{"type_echo_test.tf": tc.Content})
		rule := NewTypeEchoRule()

		// ... no errors.
		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		// ... and the expected number of issues.
		if len(runner.Issues) != len(tc.Want) {
			t.Logf("Expected %d issues, got %d", len(tc.Want), len(runner.Issues))
			for i, issue := range runner.Issues {
				t.Logf("Issue %d: %s at %s", i, issue.Message, issue.Range)
			}
			t.Fatalf("Number of issues mismatch: got %d, want %d", len(runner.Issues), len(tc.Want))
		}

		t.Run(tc.Name, func(t *testing.T) {
			if *typeEchoDeep {
				helper.AssertIssues(t, tc.Want, runner.Issues)
			} else {
				helper.AssertIssuesWithoutRange(t, tc.Want, runner.Issues)
			}
		})
	}
}

func makeTypeEchoMessage(typ string, name string) string {
	return fmt.Sprintf("The type \"%s\" is echoed in the label \"%s\"", typ, name)
}
