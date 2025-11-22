# TFLint Ruleset: Elements of Style

This is a custom TFLint ruleset that enforces both idiomatic conventions and opinionated styles for Terraform code, promoting "elements of style". Much like a style guide for writing, these rules are prescriptive by design. They reduce ambiguity and establish a baseline of readability often missing in Terraform projects.

## Rules

|Name|Description|Link|
| --- | --- | --- |
|eos_comments|Identify non-standard comment styles.|[Link](docs/rules/eos_comments.md)|
|eos_hungarian|Identify Hungarian notation in names.|[Link](docs/rules/eos_hungarian.md)|
|eos_length|Identify names longer than configurable length (default 16).|[Link](docs/rules/eos_length.md)|
|eos_reminder|Identify comments containing reminder tags.|[Link](docs/rules/eos_reminder.md)|
|eos_shout|Identify all-uppercase names.|[Link](docs/rules/eos_shout.md)|
|eos_type_echo|Identify type echoing in names.|[Link](docs/rules/eos_type_echo.md)|

## Installation

### Pre-built binary

1. Download the zip file for your platform from [Release](https://github.com/staranto/tflint-ruleset-elements-of-style/releases/latest).

2. Unzip it to your `${HOME}/.tflint.d/plugins` folder.

### Building the plugin from source

Building from source requires Go 1.25+.

Clone the repository locally and then build the binary:

```
$ make
```

Install the plugin binary with the following:

```
$ make install
```

## Requirements

- TFLint v0.46+

## Configuration

The plugin can be enabled with `tflint --init` after declaring the plugin in `.tflint.hcl` as follows:

```hcl
plugin "elements-of-style" {
  enabled = true

  version = "0.3.14"
  source  = "github.com/staranto/tflint-ruleset-elements-of-style"
}
```
