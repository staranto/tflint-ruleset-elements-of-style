// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"io/ioutil"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestShoutRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "main",
			Content: func() string {
				content, _ := ioutil.ReadFile("testdata/main.tf")
				return string(content)
			}(),
			Expected: helper.Issues{
				{
					Rule:    NewShoutRule(),
					Message: "'MY_INSTANCE' should not be all uppercase",
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 6, Column: 1},
						End:      hcl.Pos{Line: 6, Column: 38},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: "'UBUNTU' should not be all uppercase",
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 29, Column: 1},
						End:      hcl.Pos{Line: 29, Column: 24},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: "'INSTANCE_ID' should not be all uppercase",
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 44, Column: 1},
						End:      hcl.Pos{Line: 44, Column: 21},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: "'COMMON_TAGS' should not be all uppercase",
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 55, Column: 3},
						End:      hcl.Pos{Line: 57, Column: 4},
					},
				},
				{
					Rule:    NewShoutRule(),
					Message: "'HEALTH_CHECK' should not be all uppercase",
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 71, Column: 1},
						End:      hcl.Pos{Line: 71, Column: 21},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tc.Content})

			rule := NewShoutRule()

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
