.POSIX:
.SUFFIXES:

-include .env
export

GO=go
RM = rm

all: init fmt lint vulnerabilities test

init: # Downloads and verifies project dependencies and tooling.
	$(GO) install mvdan.cc/gofumpt@latest
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest

fmt: # Formats Go source files in this repository.
	gofumpt -e -extra -w .

lint: # Runs golangci-lint using the config at the root of the repository.
	golangci-lint run ./...

vulnerabilities: # Analyzes the codebase and looks for vulnerabilities affecting it.
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

test: # Runs unit tests.
	$(GO) test -short -cover -race -vet all -mod readonly ./...

test/integration: # Runs integration tests.
ifndef $(and BUNNY_STORAGE_ZONE,BUNNY_READ_API_KEY,BUNNY_WRITE_API_KEY)
	$(error Missing required environment variables. Check test/README.md for more information.)
endif
	$(GO) test -cover -race -vet all -mod readonly ./tests/integration

test/coverage: # Generates a coverage profile and open it in a browser.
	$(GO) test -short -coverprofile cover.out ./...
	$(GO) tool cover -html=cover.out

clean: # Cleans cache files from tests and deletes any build output.
	$(RM) -f cover.out

.PHONY: all init fmt lint vulnerabilities test test/coverage clean
