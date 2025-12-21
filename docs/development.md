# Development

- [Development](#development)
  - [Initial setup environment](#initial-setup-environment)
  - [Build and run n64go project](#build-and-run-n64go-project)

## Initial setup environment

* have a go dev environment
* have ares emulator installed and its executable in your path
* `go install github.com/embeddedgo/dl/go1.24.4-embedded@latest`
* `go1.24.4-embedded download`
* `go install github.com/clktmr/n64/tools/n64go@v0.1.2`
* for code completion in your IDE start gopls in the embedded-go environment (GOENV=go.env)
  * e.g. in VScode settings: `"go.alternateTools": { "go": "pathTo/go.env"}`

## Build and run n64go project

> **Info:** Building and running your project requires to have the embedded-go toolchain in your GOENV

To build your go project into a runnable N64 ROM execute the following from project root directory: 
* `./buildrun.sh`
