#!/usr/bin/make -f

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(HOME)/go/bin
SIMAPP = ./simulation/make
GOLANG_PROTOBUF_VERSION=1.28.0
GOGO_PROTOBUF_VERSION=1.3.2
GRPC_GATEWAY_VERSION=1.16.0

#install all dependencies for buf
install-buf-dependencies:
	@go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION}
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION}
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v${GRPC_GATEWAY_VERSION}
	git clone https://github.com/regen-network/protobuf.git; \
        cd protobuf; \
        go mod download; \
        make install
	git clone https://github.com/regen-network/cosmos-proto.git; \
        cd cosmos-proto/protoc-gen-gocosmos; \
        go install .
	go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos
	go get github.com/cosmos/cosmos-sdk@v0.45.9
	rm -rf cosmos-proto protobuf

export GO111MODULE = on

all: build test lintci

# The below include contains the tools and runsim targets.
include simulation/make/Makefile

########################################
### Build

build: go.sum
	@go build -mod=readonly ./...
.PHONY: build

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
.PHONY: go-mod-cache

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@go mod tidy

distclean:
	rm -rf \
    gitian-build-darwin/ \
    gitian-build-linux/ \
    gitian-build-windows/ \
    .gitian-builder-cache/
.PHONY: distclean

### Testing
TEST_TARGET ?= ./...
TEST_ARGS ?= -timeout 12m -race -coverprofile=./coverage.out -covermode=atomic -v

test: test-unit
test-all: test-unit test-ledger-mock test-race test-cover

test-ledger-mock:
	@go test -mod=readonly `go list github.com/cosmos/cosmos-sdk/crypto` -tags='cgo ledger test_ledger_mock'

test-ledger: test-ledger-mock
	@go test -mod=readonly -v `go list github.com/cosmos/cosmos-sdk/crypto` -tags='cgo ledger'

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' $(PACKAGES_NOSIMULATION) $(TEST_TARGET) $(TEST_ARGS)

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race $(PACKAGES_NOSIMULATION)

.PHONY: test test-all test-ledger-mock test-ledger test-unit test-race

run-simulations: test-sim-custom-genesis-fast test-sim-nondeterminism test-sim-import-export test-sim-after-import test-sim-custom-genesis-multi-seed test-sim-multi-seed-long test-sim-multi-seed-short test-sim-benchmark-invariants

test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@go test -mod=readonly $(SIMAPP) -run=TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

test-sim-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.assetNode/config/genesis.json will be used."
	@go test -mod=readonly $(SIMAPP) -run=TestFullAppSimulation -Genesis=${HOME}/.assetNode/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

test-sim-full-application:
	@echo "Running full app..."
	@go test -mod=readonly $(SIMAPP) -run=TestFullAppSimulation \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 50 5 TestAppImportExport


test-sim-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 50 5 TestAppSimulationAfterImport

test-sim-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.assetNode/config/genesis.json will be used."
	@$(BINDIR)/runsim -Genesis=${HOME}/.assetNode/config/genesis.json -SimAppPkg=$(SIMAPP) 400 5 TestFullAppSimulation

test-sim-multi-seed-long: runsim
	@echo "Running long multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 500 50 TestFullAppSimulation

test-sim-multi-seed-short: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 50 10 TestFullAppSimulation

test-sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly $(SIMAPP) -benchmem -bench=BenchmarkInvariants -run=^$ \
	-Enabled=true -NumBlocks=1000 -BlockSize=200 \
	-Period=1 -Commit=true -Seed=57 -v -timeout 24h

.PHONY: \
test-sim-full-application \
test-sim-nondeterminism \
test-sim-custom-genesis-fast \
test-sim-import-export \
test-sim-after-import \
test-sim-custom-genesis-multi-seed \
test-sim-multi-seed-short \
test-sim-multi-seed-long \
test-sim-benchmark-invariants

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true

test-sim-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkFullAppSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h

test-sim-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkFullAppSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out

.PHONY: test-sim-profile test-sim-benchmark

test-cover:
	@export VERSION=$(VERSION); bash -x tests/test_cover.sh
.PHONY: test-cover

benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_NOSIMULATION)
.PHONY: benchmark

lintci:
	golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 --config .golangci.yaml
.PHONY: lintci

lintci-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
.PHONY: lintci-install

lintci-remove:
	rm $(GOPATH)/bin/golangci-lint
.PHONY: lintci-remove

lintci-update: lintci-remove lintci-install
.PHONY: lintci-update

goimports:
	goimports -local="github.com/AssetMantle/modules" -w .
