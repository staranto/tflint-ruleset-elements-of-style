# TFLint Ruleset: Elements of Style

This is a custom TFLint ruleset that enforces idiomatic conventions and style guidelines for Terraform configurations, promoting "elements of style" in infrastructure as code.

## Requirements

- TFLint v0.46+
- Go v1.25

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "elements-of-style" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/staranto/tflint-ruleset-elements-of-style"
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|eos_type_echo|Disallow type echoing in block labels|WARNING|✔|[Link](docs/rules/eos_type_echo.md)|
|eos_length|Disallow block labels longer than configurable length (default 16)|WARNING|✔|[Link](docs/rules/eos_length.md)|
|eos_shout|Disallow all-uppercase block labels|WARNING|✔|[Link](docs/rules/eos_shout.md)|

## Installation

1. Download the zip file for your platform from [Release](https://github.com/staranto/tflint-ruleset-elements-of-style/releases/latest).

2. Unzip it to your `${HOME}/.tflint.d/plugins` folder

## Building the plugin from source

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

### Configuring tflint

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "elements-of-style" {
  enabled = true
}
EOS
$ tflint
```
