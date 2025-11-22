# eos_length

Identify names longer than a configurable length (default 16 characters).

## Example

```hcl
resource "aws_instance" "very_long_instance_name" {
  # ...
}
```

```
$ tflint
1 issue(s) found:

Warning: 'very_long_instance_name' is 22 characters and should not be longer than 16 (eos_length)

  on config.tf line 1:
  1: resource "aws_instance" "very_long_instance_name" {

Reference: https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_length.md

```

## Why

Long names can make Terraform configurations harder to read and maintain. They can also cause issues with tools like `tfctl` or `terraform` by causing content to be pushed way past the right edge of the terminal. Keeping names concise encourages better naming practices and improves overall code quality.

## Configuration

The length limit can be customized using the `length` parameter in your `.tflint.hcl` configuration file. The default limit is 16 characters.

```hcl
rule "eos_length" {
  length = 20
  level = "warning"
}
```

## How To Fix

Rename the block to a shorter, more descriptive name.The rule can be ignored with -

```hcl
# tflint-ignore: eos_length
resource "aws_instance" "very_long_instance_name" {
  # ...
}
```