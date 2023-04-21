.POSIX:
.SUFFIXES:

GO?=go

-include .env
export

all: init lint vulnerabilities

init: # Downloads and verifies project dependencies and tooling.
	$(GO) install mvdan.cc/gofumpt@v0.5.0
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest

fmt: # Formats Go source files in this repository.
	gofumpt -e -extra -w .

lint: # Runs golangci-lint using the config at the root of the repository.
	golangci-lint run ./...

vulnerabilities: # Analyzes the codebase and looks for vulnerabilities affecting it.
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

test: # Runs unit tests.
	$(GO) test -cover -race -vet all -mod readonly ./...

test/coverage: # Generates a coverage profile and open it in a browser.
	$(GO) test -coverprofile cover.out
	$(GO) tool cover -html=cover.out

test/integration: # Runs integration tests.
ifndef $(and BUNNY_STORAGE_ZONE,BUNNY_READ_API_KEY,BUNNY_WRITE_API_KEY)
	$(error Missing required environment variables. Check test/README.md for more information.)
endif
	$(GO) test -cover -race -vet all -mod readonly ./test/integration

clean: # Cleans cache files from tests and deletes any build output.
	$(GO) clean -cache -fuzzcache -testcache ./...

.PHONY: all init fmt lint vulnerabilities test test/coverage test/integration clean
