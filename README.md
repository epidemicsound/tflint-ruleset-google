# TFLint Ruleset for terraform-provider-google
[![Build Status](https://github.com/epidemicsound/tflint-ruleset-google/workflows/build/badge.svg?branch=main)](https://github.com/epidemicsound/tflint-ruleset-google/actions)
[![GitHub release](https://img.shields.io/github/release/epidemicsound/tflint-ruleset-google.svg)](https://github.com/epidemicsound/tflint-ruleset-google/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for Terraform Google Cloud Platform provider (fork of https://github.com/terraform-linters/setup-tflint)

## Requirements

- TFLint v0.42+
- Go v1.22

## Installation

You can install the plugin by adding a config to `.tflint.hcl` and running `tflint --init`:

```hcl
plugin "google" {
    enabled = true
    version = "0.31.1"
    source  = "github.com/terraform-linters/tflint-ruleset-google"
}
```

For more configuration about the plugin, see [Plugin Configuration](docs/configuration.md).

## Rules

100+ rules are available. See the [documentation](docs/rules/README.md).

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

Note that if you install the plugin with `make install`, you must omit the `version` and `source` attributes in `.tflint.hcl`:

```hcl
plugin "google" {
    enabled = true
}
```

## Add a new rule

If you are interested in adding a new rule to this ruleset, you can use the generator. Run the following command:

```
$ go run ./rules/generator
```

Follow the instructions to edit the generated files and open a new pull request.

## Create a new release

1. Add your changes to a PR (include a commit with the version bump into the PR), get review, merge to main.
2. Wait for a successful completion of the build on main.
3. Fetch the latest state of main to your local machine, create a tag and push it to remote
   (i.e. `git checkout main && git pull && git tag v0.31.2 && git push origin v0.31.2`).
4. Wait for the [release](https://github.com/epidemicsound/tflint-ruleset-google/actions/workflows/release.yml) workflow to finish.
5. Go to [releases](https://github.com/epidemicsound/tflint-ruleset-google/releases), your new release will be in the `draft` state.
6. Click edit -> add the release notes (or auto-generate them) -> Publish release.
