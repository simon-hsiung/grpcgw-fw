.DEFAULT_GOAL := help

COMMIT_SHA?=$(shell git rev-parse --short HEAD)
CGO_ENABLED?=0

# For Mac M2
# GOOS?=darwin
# GOARCH?=arm64
# For Linux
GOOS?=linux
GOARCH?=amd64
OUTBIN?=bin
GO ?= go

$(OUTBIN):
	mkdir -p ./${OUTBIN}

.PHONY: setup
setup: $(OUTBIN) ## setup go modules
	go mod tidy
	go fmt ./...

.PHONY: clean
clean: ## cleans the binary
	rm -rf ./${OUTBIN}
	go clean ./...

.PHONY: lint
lint: setup ## lint go
	go vet ./...
	golangci-lint run

.PHONY: test
test: ## runs go test the application
	CGO_ENABLED=${CGO_ENABLED} go test -v ./...

.PHONY: build-pb
build-pb: ## build the protobuf
	cd ./protocol && $(MAKE) build

.PHONY: build
build: setup ## build the application
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go build \
		-ldflags=" \
			-w -s \
			-X 'main.Commit=${COMMIT_SHA}' \
		" \
		-trimpath \
		./...

.PHONY: help
help: ## prints this help message
	@echo "Usage: \n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
