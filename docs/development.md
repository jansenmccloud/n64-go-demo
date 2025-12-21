# Development

- [Development](#development)
  - [Setup environment](#setup-environment)
  - [Build and run project](#build-and-run-project)

## Setup environment

* install go + IDE
* go install github.com/embeddedgo/dl/go1.24.5-embedded@latest
* go1.24.4-embedded download
* go install github.com/clktmr/n64/tools/n64go@latest (or v0.1.2)
* create go.env file im root
* export GOENV=go.env
  * optionally: for code completion start gopls in that environment (e.g. GOENV=go.env <editor>)
* go get github.com/clktmr/n64@latest (or v0.1.2)
* go mod init <moduleName>
* add main.go
* go mod tidy

## Build and run project

* have ares emulator in your path
* go build ./cmd/n64go
* go run ./cmd/n64go