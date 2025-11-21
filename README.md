# TFLint Ruleset: Elements of Style

This is a custom TFLint ruleset that enforces both idiomatic conventions and opinionated styles for Terraform code, promoting "elements of style". Much like a style guide for writing, these rules are prescriptive by design. They reduce ambiguity and establish a baseline of readability often missing in Terraform projects.

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|eos_comments|Enforce simple comment styles.|WARNING|✔|[Link](docs/rules/eos_comments.md)|
|eos_hungarian|Disallow Hungarian notation in block labels.|WARNING|✔|[Link](docs/rules/eos_hungarian.md)|
|eos_length|Disallow block labels longer than configurable length (default 16).|WARNING|✔|[Link](docs/rules/eos_length.md)|
|eos_reminder|Disallow comments containing reminder tags.|WARNING|✔|[Link](docs/rules/eos_reminder.md)|
|eos_shout|Disallow all-uppercase block labels.|WARNING|✔|[Link](docs/rules/eos_shout.md)|
|eos_type_echo|Disallow type echoing in block labels.|WARNING|✔|[Link](docs/rules/eos_type_echo.md)|

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

  version = "0.3.9"
  source  = "github.com/staranto/tflint-ruleset-elements-of-style"
}
```
