VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

build: ## Build the binary
	go build ldflags "-X main.version=$(VERSION)" -o polyenv .

install: ## Install to GOPATH/bin
	go install ldflags "-X main.version=$(VERSION)" .

test: ## Run tests (unit only)
	go test -v ./...

test-network: ## Run tests including network (hits tio.run)
	POLYENV_NETWORK_TESTS=1 go test -v ./...

vet: ## Run go vet
	go vet ./...

lint: vet ## Run staticcheck (includes vet)
	@command -v staticcheck >/dev/null 2>&1 || { echo "install staticcheck: go install honnef.co/go/tools/cmd/staticcheck@latest"; exit 1; }
	staticcheck ./...

clean: ## Remove build artifacts
	rm -f polyenv

help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*##' $(MAKEFILE_LIST) | sort | \
		awk -F ':.*## ' '{printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: build install test test-network vet lint clean help
.DEFAULT_GOAL := help
