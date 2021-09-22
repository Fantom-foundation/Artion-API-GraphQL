# --------------------------------------------------------------------------
# Makefile for the Artion API GraphQL Server
#
# v0.1 (2020/03/09)  - Initial version, base API server build.
# (c) Fantom Foundation, 2020
# --------------------------------------------------------------------------

# project related vars
PROJECT := $(shell basename "$(PWD)")

# go related vars
GO_BASE := $(shell pwd)
GO_BIN := $(CURDIR)/build

# compile time variables will be injected into the app
APP_VERSION := 1.1.0
BUILD_DATE := $(shell date)
BUILD_COMPILER := $(shell go version)
BUILD_COMMIT := $(shell git show --format="%H" --no-patch)
BUILD_COMMIT_TIME := $(shell git show --format="%cD" --no-patch)

build/artionapi: internal/graphql/schema/gen/schema.graphql
	go build \
	-gcflags="all=-N -l" \
	-ldflags="-X 'artion-api-graphql/cmd/artionapi/build.Version=$(APP_VERSION)' -X 'artion-api-graphql/cmd/artionapi/build.Time=$(BUILD_DATE)' -X 'artion-api-graphql/cmd/artionapi/build.Compiler=$(BUILD_COMPILER)' -X 'artion-api-graphql/cmd/artionapi/build.Commit=$(BUILD_COMMIT)' -X 'artion-api-graphql/cmd/artionapi/build.CommitTime=$(BUILD_COMMIT_TIME)'" \
	-o build/artionapi \
	./cmd/artionapi

test:
	go test ./...

internal/graphql/schema/gen/schema.graphql: internal/graphql/definition
	@bash tools/make_graphql_bundle.sh $@ $<

internal/repository/rpc/contracts/FantomMarketplace.go: internal/repository/rpc/contracts/abi/FantomMarketplace.json
	abigen --type FantomMarketplace --pkg contracts --abi $< --out $@

.PHONY: build/artionapi help test
all: help
help: Makefile
	@echo
	@echo "Choose a make command in "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
