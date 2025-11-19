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

var hungarianDeep = flag.Bool("hungarianDeep", false, "enable deep assert")

func TestHungarianRule(t *testing.T) {
	flag.Parse()

	cases := []struct {
		Name    string
		Content string
		Want    helper.Issues
	}{
		{
			Name: "hungarian_names",
			Content: func() string {
				content, _ := os.ReadFile("testdata/hungarian_test.tf")
				return string(content)
			}(),
			Want: helper.Issues{
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("str_hung", "str"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 7, Column: 1},
						End:      hcl.Pos{Line: 7, Column: 22},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("hung_int", "int"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 10, Column: 3},
						End:      hcl.Pos{Line: 10, Column: 11},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("hung_bool_check", "bool"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 13, Column: 1},
						End:      hcl.Pos{Line: 13, Column: 24},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("map_hung", "map"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 20, Column: 1},
						End:      hcl.Pos{Line: 20, Column: 41},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("hung_lst", "lst"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 22, Column: 1},
						End:      hcl.Pos{Line: 22, Column: 41},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("hung_set_mod", "set"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 26, Column: 1},
						End:      hcl.Pos{Line: 26, Column: 23},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("num_hung", "num"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 30, Column: 1},
						End:      hcl.Pos{Line: 30, Column: 19},
					},
				},
				{
					Rule:    NewHungarianRule(),
					Message: makeHungarianMessage("str_hung", "str"),
					Range: hcl.Range{
						Filename: "hungarian_test.tf",
						Start:    hcl.Pos{Line: 35, Column: 1},
						End:      hcl.Pos{Line: 35, Column: 37},
					},
				},
			},
		},
	}

	for _, tc := range cases {

		// Run the tests and make sure the basic results are found...
		runner := helper.TestRunner(t, map[string]string{"hungarian_test.tf": tc.Content})
		rule := NewHungarianRule()

		// ... no errors.
		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		// ... and the expected number of issues.
		if len(runner.Issues) != len(tc.Want) {
			t.Fatalf("Number of issues mismatch: got %d, want %d", len(runner.Issues), len(tc.Want))
		}

		t.Run(tc.Name, func(t *testing.T) {
			if *hungarianDeep {
				helper.AssertIssues(t, tc.Want, runner.Issues)
			} else {
				helper.AssertIssuesWithoutRange(t, tc.Want, runner.Issues)
			}
		})
	}
}

func makeHungarianMessage(name string, key string) string {
	return fmt.Sprintf("'%s' uses Hungarian notation with '%s'", name, key)
}
