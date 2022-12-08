#!/usr/bin/env bash

set -eo pipefail

# get protoc executions
go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null

# get cosmos sdk from github
go get github.com/cosmos/cosmos-sdk@v0.45.9 2>/dev/null

echo "Generating gogo proto code"

proto_dirs=$(find ./modules -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -name '*.proto'); do
    if grep go_package $file &>/dev/null; then
      buf generate --template buf.gen.yaml $file
    fi
  done
done

