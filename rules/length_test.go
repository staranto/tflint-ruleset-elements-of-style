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

var deep = flag.Bool("deep", false, "enable deep assert")
var lengthName = "zakpxy_very_long_name"

func TestLengthRule(t *testing.T) {
	flag.Parse()

	cases := []struct {
		Name    string
		Content string
		Want    helper.Issues
	}{
		{
			Name: "long_names",
			Content: func() string {
				content, _ := os.ReadFile("testdata/length_test.tf")
				return string(content)
			}(),
			Want: helper.Issues{
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 7, Column: 1},
						End:      hcl.Pos{Line: 7, Column: 33},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 10, Column: 3},
						End:      hcl.Pos{Line: 10, Column: 28},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 13, Column: 1},
						End:      hcl.Pos{Line: 13, Column: 30},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 20, Column: 1},
						End:      hcl.Pos{Line: 20, Column: 55},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 22, Column: 1},
						End:      hcl.Pos{Line: 22, Column: 52},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 26, Column: 1},
						End:      hcl.Pos{Line: 26, Column: 31},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 30, Column: 1},
						End:      hcl.Pos{Line: 30, Column: 31},
					},
				},
				{
					Rule:    NewLengthRule(),
					Message: makeLengthMessage(lengthName),
					Range: hcl.Range{
						Filename: "length_test.tf",
						Start:    hcl.Pos{Line: 34, Column: 1},
						End:      hcl.Pos{Line: 34, Column: 48},
					},
				},
			},
		},
	}

	for _, tc := range cases {

		// Run the tests and make sure the basic results are found...
		runner := helper.TestRunner(t, map[string]string{"length_test.tf": tc.Content})
		rule := NewLengthRule()

		// ... no errors.
		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		// ... and the expected number of issues.
		if len(runner.Issues) != len(tc.Want) {
			t.Fatalf("Number of issues mismatch: got %d, want %d", len(runner.Issues), len(tc.Want))
		}

		t.Run(tc.Name, func(t *testing.T) {
			if *deep {
				helper.AssertIssues(t, tc.Want, runner.Issues)
			} else {
				helper.AssertIssuesWithoutRange(t, tc.Want, runner.Issues)
			}
		})
	}
}

func makeLengthMessage(name string) string {
	return fmt.Sprintf("'%s' is %d characters and should not be longer than %d", name, len(name), 16)
}
