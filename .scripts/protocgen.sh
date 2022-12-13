#!/usr/bin/env bash

set -xeo pipefail

# get protoc executions
go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null


# get cosmos sdk from github
go get github.com/cosmos/cosmos-sdk@v0.45.9 2>/dev/null

