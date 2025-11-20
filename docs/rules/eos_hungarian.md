# eos_hungarian

Disallow Hungarian notation in block labels.

## Example

```hcl
resource "aws_instance" "str_instance" {
  # ...
}
```

```
$ tflint
1 issue(s) found:

Warning: 'str_instance' uses Hungarian notation with 'str' (eos_hungarian)

  on config.tf line 1:
  1: resource "aws_instance" "str_instance" {

Reference: https://github.com/staranto/tflint-ruleset-elements-of-style/blob/main/docs/rules/eos_hungarian.md

```

## Why

Hungarian notation (encoding type information in variable names) is generally considered redundant in strongly typed languages or declarative configurations like Terraform where the type is often evident from the context (e.g., `resource "aws_instance"` clearly defines an instance). Avoiding it leads to cleaner and more readable code.

## Configuration

The list of disallowed prefixes/suffixes can be customized.

By default, the following strings are checked: `str`, `int`, `num`, `bool`, `list`, `lst`, `set`, `map`, `arr`, `array`.

You can override the defaults completely:

```hcl
rule "eos_hungarian" {
  defaults = ["foo", "bar"]
}
```

Or you can append to the defaults:

```hcl
rule "eos_hungarian" {
  more = ["foo", "bar"]
}
```

## How To Fix

Rename the block to remove the Hungarian notation.

```hcl
resource "aws_instance" "web" {
  # ...
}
```

The rule can be ignored with:

```hcl
# tflint-ignore: eos_hungarian
resource "aws_instance" "str_instance" {
  # ...
}
```
