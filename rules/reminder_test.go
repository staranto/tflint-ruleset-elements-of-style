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

var reminderDeep = flag.Bool("reminderDeep", false, "enable deep assert")

func TestReminderRule(t *testing.T) {
	flag.Parse()

	cases := []struct {
		Name    string
		Content string
		Want    helper.Issues
	}{
		{
			Name: "reminders",
			Content: func() string {
				content, _ := os.ReadFile("testdata/reminder_test.tf")
				return string(content)
			}(),
			Want: helper.Issues{
				{
					Rule:    NewReminderRule(),
					Message: makeReminderMessage("// TODO Reminder found."),
					Range: hcl.Range{
						Filename: "reminder_test.tf",
						Start:    hcl.Pos{Line: 4, Column: 1},
						End:      hcl.Pos{Line: 4, Column: 23},
					},
				},
				{
					Rule:    NewReminderRule(),
					Message: makeReminderMessage("# TODO Reminder found."),
					Range: hcl.Range{
						Filename: "reminder_test.tf",
						Start:    hcl.Pos{Line: 5, Column: 1},
						End:      hcl.Pos{Line: 5, Column: 22},
					},
				},
				{
					Rule:    NewReminderRule(),
					Message: makeReminderMessage("# TODO Reminder found."),
					Range: hcl.Range{
						Filename: "reminder_test.tf",
						Start:    hcl.Pos{Line: 8, Column: 11},
						End:      hcl.Pos{Line: 8, Column: 32},
					},
				},
				{
					Rule:    NewReminderRule(),
					Message: makeReminderMessage("# FIXME Reminder found."),
					Range: hcl.Range{
						Filename: "reminder_test.tf",
						Start:    hcl.Pos{Line: 14, Column: 37},
						End:      hcl.Pos{Line: 14, Column: 59},
					},
				},
			},
		},
	}

	for _, tc := range cases {

		// Run the tests and make sure the basic results are found...
		runner := helper.TestRunner(t, map[string]string{"reminder_test.tf": tc.Content})
		rule := NewReminderRule()

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
			if *reminderDeep {
				helper.AssertIssues(t, tc.Want, runner.Issues)
			} else {
				helper.AssertIssuesWithoutRange(t, tc.Want, runner.Issues)
			}
		})
	}
}

func makeReminderMessage(comment string) string {
	return fmt.Sprintf("'%s' has a reminder tag.", comment)
}
