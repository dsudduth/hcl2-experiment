# HCL2 Experiment

[HCL2](https://github.com/hashicorp/hcl2) is the second iteration of the HashiCorp Configuration Language. It is a powerful DSL that serves as an alternative to YAML, TOML, INI, JSON, etc. This experiment attempts to evaluate HCL2 as a possible config language for Golang applications.

## Why HCL

HCL has an incredibly dynamic nature allowing users to perform tasks such as interpolation and recursion. In more advanced implementations, configuration files can be implemented as modules; making it shareable and reusable. The syntax itself is [declarative](https://en.wikipedia.org/wiki/Declarative_programming) and excellent choice when clarity is a must. While this list is not exhaustive, it does provide a small peek into why you may want to consider it as a config language for your project.

It's also important to note that HCL2 is now fully compatible with JSON. Conversion is now possible, further lending to its usefulness.

## Running this experiment

First you'll need to install the latest version of Golang if you do not already have it. Next, install the necessary dependencies.

```bash
go mod download
```

With the dependencies installed, the code can simply be run using the `go run` command.

```bash
go run main.go
```

If successful, a JSON representation of the object should be printed.

## Where can I read more

Documentation for HCL2 is a bit sparse at the moment. The following resources can help you get started.

* [HCL2 Docs](https://github.com/hashicorp/hcl2/tree/master/guide) - I have not found this published anywhere, but it's a great starting point. It contains an intro, examples, and a few implementation details.
* [Terraform Repository](https://github.com/hashicorp/terraform) - Official HashiCorp project and perhaps the most complex implementation of HCL, taking full advantage of HCL's capabilities.
* [Consul Repository](https://github.com/hashicorp/consul/) - Official HashiCorp project. While not as complex as Terraform, it still demonstrates an advanced implementation of the config language.
* [Golicense Repository](https://github.com/mitchellh/golicense) - This tool was created by Mitchell Hashimoto and it contains a simple example of HCL2 in action.
* [Terragrunt Repository](https://github.com/gruntwork-io/terragrunt) - Terragrunt is a popular wrapper for Terraform. It has an excellent implementation of HCL2.
* [TF Lint Repository](https://github.com/wata727/tflint) - This is a third-party application written as a linter for Terraform. It leverages HCL2 as its config language.
