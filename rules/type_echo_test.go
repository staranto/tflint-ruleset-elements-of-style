// Copyright (c) 2025 Steve Taranto <staranto@gmail.com>.
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"io/ioutil"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestTypeEchoRule(t *testing.T) {
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
					Rule:    NewTypeEchoRule(),
					Message: `The type "aws_instance" is echoed in the label "my_instance"`,
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 38},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: `The type "aws_instance" is echoed in the label "MY_INSTANCE"`,
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 7, Column: 1},
						End:      hcl.Pos{Line: 7, Column: 38},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: `The type "aws_instance" is echoed in the label "very_long_instance_name_that_exceeds_limit"`,
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 13, Column: 1},
						End:      hcl.Pos{Line: 13, Column: 69},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: `The type "aws_s3_bucket" is echoed in the label "my_bucket"`,
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 19, Column: 1},
						End:      hcl.Pos{Line: 19, Column: 37},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: `The type "check" is echoed in the label "health_check"`,
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 63, Column: 1},
						End:      hcl.Pos{Line: 63, Column: 21},
					},
				},
				{
					Rule:    NewTypeEchoRule(),
					Message: `The type "check" is echoed in the label "HEALTH_CHECK"`,
					Range: hcl.Range{
						Filename: "main.tf",
						Start:    hcl.Pos{Line: 74, Column: 1},
						End:      hcl.Pos{Line: 74, Column: 21},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tc.Content})

			rule := NewTypeEchoRule()

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
