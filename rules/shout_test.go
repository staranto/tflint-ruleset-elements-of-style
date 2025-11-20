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

var shoutDeep = flag.Bool("shoutDeep", false, "enable deep assert")
var shoutName = "SHOUT"

func TestShoutRule(t *testing.T) {
	flag.Parse()

	cases := []struct {
		Name    string
		Content string
		Want    helper.Issues
	}{
		{
			Name: "shouted_names",
			Content: func() string {
				content, _ := os.ReadFile("testdata/shout_test.tf")
				return string(content)
			}(),
			Want: helper.Issues{
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 7, Column: 1},
						End:      hcl.Pos{Line: 7, Column: 19},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 10, Column: 3},
						End:      hcl.Pos{Line: 10, Column: 12},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 13, Column: 1},
						End:      hcl.Pos{Line: 13, Column: 16},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 20, Column: 1},
						End:      hcl.Pos{Line: 20, Column: 38},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 22, Column: 1},
						End:      hcl.Pos{Line: 22, Column: 38},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 26, Column: 1},
						End:      hcl.Pos{Line: 26, Column: 16},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 30, Column: 1},
						End:      hcl.Pos{Line: 30, Column: 16},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: makeShoutMessage(shoutName),
					Range: hcl.Range{
						Filename: "shout_test.tf",
						Start:    hcl.Pos{Line: 35, Column: 1},
						End:      hcl.Pos{Line: 35, Column: 34},
					},
				},
			},
		},
	}

	for _, tc := range cases {

		// Run the tests and make sure the basic results are found...
		runner := helper.TestRunner(t, map[string]string{"shout_test.tf": tc.Content})
		rule := NewShoutRule()

		// ... no errors.
		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		// ... and the expected number of issues.
		if len(runner.Issues) != len(tc.Want) {
			t.Fatalf("Number of issues mismatch: got %d, want %d", len(runner.Issues), len(tc.Want))
		}

		t.Run(tc.Name, func(t *testing.T) {
			if *shoutDeep {
				helper.AssertIssues(t, tc.Want, runner.Issues)
			} else {
				helper.AssertIssuesWithoutRange(t, tc.Want, runner.Issues)
			}
		})
	}
}

func makeShoutMessage(name string) string {
	return fmt.Sprintf("'%s' should not be all uppercase.", name)
}
