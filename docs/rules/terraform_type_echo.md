# terraform_type_echo

Disallow type echoing in block labels.

## Example

```hcl
resource "aws_s3_bucket" "logging-bucket" {
  # ...
}

```

```
$ tflint
1 issue(s) found:

Warning: [Fixable] The type "aws_s3_bucket" is echoed in the label "logging-bucket" (type_echo)

  on config.tf line 1:
  1: resource "aws_s3_bucket" "logging-bucket" {

Reference: https://github.com/terraform-linters/tflint-ruleset-terraform/blob/v0.1.0/docs/rules/terraform_unused_declarations.md

```

## Why

Type echoing (aka. type jittering or [Hungarian Notation](https://en.wikipedia.org/wiki/Hungarian_notation)) is considered a bad practice when writing Terraform.  In *all* cases, the Terraform and OpenTofu tooling displays the type (`aws_s3_bucket`) immediately adjacent to the label, or name, (`logging-bucket`) of the occurence.

In the HCL language itself, the syntax is -

```hcl
resource "aws_s3_bucket" "logging-bucket" {
  # ...
}
```
not -

```hcl
resource "aws_s3_bucket"
# A whle bunch of comments describing
# what this resource is about
  "logging-bucket" {
  # ...
}
```

When listing the contents of a state file (with `terraform state list` or `tfctl sq`), or executing a `plan/apply`, the output is -

```
aws_s3_bucket.logging-bucket
```

In *all* cases, you would "jitter" as you pronounced this - "aws S3 bucket logging bucket". Or, even more pronounced - "bucket logging bucket".  Neither "flows" as well as simply saying "S3 bucket logging".

Since HCL is a verbose language this can also quickly spin out of control if you were to write something like -

```hcl
resource "aws_security_group" "primary_security_group" {
  # ...
}

resource "aws_vpc_security_group_ingress_rule" "sg_rule_ingress" {
  security_group_id = aws_security_group.primary_security_group.id
  # ...
}
```

It's much more readable and, thus, maintanable to write -

```hcl
resource "aws_security_group" "primary" {
  # ...
}

resource "aws_vpc_security_group_ingress_rule" "ingress" {
  security_group_id = aws_security_group.primary.id
  # ...
}
```


## How To Fix

Remove the declaration. For `variable` and `data`, remove the entire block. For a `local` value, remove the attribute from the `locals` block.

While data sources should generally not have side effects, take greater care when removing them. For example, removing `data "http"` will cause Terraform to no longer perform an HTTP `GET` request during each plan. If a data source is being used for side effects, add an annotation to ignore it:

```tf
# tflint-ignore: terraform_unused_declarations
data "http" "example" {
  url = "https://checkpoint-api.hashicorp.com/v1/check/terraform"
}
```
