// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"flag"
	"testing"

	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

var commentsDeep = flag.Bool("commentsDeep", false, "enable deep assert")

func TestCommentsRule(t *testing.T) {
	flag.Parse()

	cases := []struct {
		Name    string
		Content string
		Want    helper.Issues
	}{
		{
			Name: "comments",
			Content: func() string {
				content, _ := os.ReadFile("testdata/comments_test.tf")
				return string(content)
			}(),
			Want: helper.Issues{
				{
					Rule:    NewCommentsRule(),
					Message: "Comment is jammed ('#Bad  ...').",
					Range: hcl.Range{
						Filename: "comments_test.tf",
						Start:    hcl.Pos{Line: 4, Column: 1},
						End:      hcl.Pos{Line: 4, Column: 22},
					},
				},
				{
					Rule:    NewCommentsRule(),
					Message: "Comment is jammed ('//Bad ...').",
					Range: hcl.Range{
						Filename: "comments_test.tf",
						Start:    hcl.Pos{Line: 5, Column: 1},
						End:      hcl.Pos{Line: 5, Column: 23},
					},
				},
				{
					Rule:    NewCommentsRule(),
					Message: "Comment extends beyond column 80 to 126.",
					Range: hcl.Range{
						Filename: "comments_test.tf",
						Start:    hcl.Pos{Line: 7, Column: 1},
						End:      hcl.Pos{Line: 7, Column: 114},
					},
				},
				{
					Rule:    NewCommentsRule(),
					Message: "Block comments not allowed.",
					Range: hcl.Range{
						Filename: "comments_test.tf",
						Start:    hcl.Pos{Line: 9, Column: 1},
						End:      hcl.Pos{Line: 12, Column: 3},
					},
				},
				{
					Rule:    NewCommentsRule(),
					Message: "Comment extends beyond column 80 to 106.",
					Range: hcl.Range{
						Filename: "comments_test.tf",
						Start:    hcl.Pos{Line: 15, Column: 3},
						End:      hcl.Pos{Line: 15, Column: 108},
					},
				},
			},
		},
	}

	for _, tc := range cases {

		// Run the tests and make sure the basic results are found...
		runner := helper.TestRunner(t, map[string]string{"comments_test.tf": tc.Content})
		rule := NewCommentsRule()

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
			if *commentsDeep {
				helper.AssertIssues(t, tc.Want, runner.Issues)
			} else {
				helper.AssertIssuesWithoutRange(t, tc.Want, runner.Issues)
			}
		})
	}
}
