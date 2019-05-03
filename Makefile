##
# This Makefile exists only for developer's convenience. It does not do actual building.
##

.DEFAULT_GOAL=help

.PHONY: help test test-cover

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run tests
	go test -cover ./...

test-cover: ## Run tests with file-by-file coverage
	go test -coverprofile cover.out ./...; go tool cover -func cover.out; rm cover.out

cover: test-cover ## Alias of `test-cover`
