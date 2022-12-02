#!/usr/bin/env bash

set -eo pipefail

protoc \
  schema/thirdParty/proto/gogoproto/*.proto \
  schema/data/base/*.proto \
  schema/data/*.proto \
  schema/documents/base/*.proto \
  schema/ids/base/*.proto \
  schema/lists/base/*.proto \
  schema/properties/base/*.proto \
  schema/qualified/base/*.proto \
  schema/types/base/*.proto \
  modules/metas/internal/queries/meta/*.proto \
  modules/metas/internal/transactions/reveal/*.proto \
  -I. \
  --go_out=. \
  --go_opt=paths=source_relative
