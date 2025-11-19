# THINK This is fragile af. Order dependent.  See comment at bottom.
# TODO Figure out how to use .tflint.hcl in a test.

# #########
# Tests that will emit issues.

variable "zakpxy_very_long_name" {}

locals {
  zakpxy_very_long_name = 1
}

check "zakpxy_very_long_name" {
  assert {
    condition     = local.zakpxy_very_long_name > 0
    error_message = "Must be > 0."
  }
}

data "aws_get_caller_identity" "zakpxy_very_long_name" {}

ephemeral "random_password" "zakpxy_very_long_name" {
  length = 8
}

module "zakpxy_very_long_name" {
  source = "./modules/"
}

output "zakpxy_very_long_name" {
  value = local.zakpxy_very_long_name
}

// tflint-ignore: eos_length
resource "aws_instance" "zakpxy_very_long_name" {
  ami = "ami-12345678"
}

# #########
# Tests that will not emit issues.

variable "short" {}

locals {
  short = 1
}

resource "aws_instance" "short" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
}


# Input for length_test.go. Each of the blocks in the first ("emit issues")
# section has a matching test case in length_test.go. What's more, the order of
# the blocks here must be synced with the order of the test cases. This is super
# fragile, but I don't have the will to think of a better way right now.  The
# blocks align with common.go/allLintableBlocks.
