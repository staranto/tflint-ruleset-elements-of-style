# eos_shout

Disallow all-uppercase block labels.

## Example

```hcl
resource "aws_instance" "MY_INSTANCE" {
  # ...
}

variable "MY_VAR" {
  # ...
}

locals {
  MY_LOCAL = "value"
}
```

```
$ tflint
3 issue(s) found:

Warning: 'MY_INSTANCE' should not be all uppercase (eos_shout)

  on config.tf line 1:
  1: resource "aws_instance" "MY_INSTANCE" {

Warning: 'MY_VAR' should not be all uppercase (eos_shout)

  on config.tf line 5:
  5: variable "MY_VAR" {

Warning: 'MY_LOCAL' should not be all uppercase (eos_shout)

  on config.tf line 9:
  9: locals {

Reference: https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_shout.md

```

## Why

All-uppercase names (shouting) can be harder to read and may imply a significance, such as constants or macros, that doesn't exist. Using snake_case, mixedCase, or lowercase names improves readability and aligns with common naming conventions.

## Configuration

This rule has no configurable parameters.

## How To Fix

Rename the block to use snake_case, mixedCase or lowercase. The rule can be ignored with -

```hcl
# tflint-ignore: eos_shout
resource "aws_instance" "MY_INSTANCE" {
  # ...
}
```