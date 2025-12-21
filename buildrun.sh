#!/bin/bash

export GOENV=go.env

echo "... building"
go build ./cmd/n64go

echo "... running"
go run ./cmd/n64go