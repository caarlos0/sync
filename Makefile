SOURCE_FILES?=./...
TEST_PATTERN?=.

setup:
	go mod tidy
.PHONY: setup

build:
	go build
.PHONY: build

test:
	go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

ci: build test
.PHONY: ci

.DEFAULT_GOAL := ci
