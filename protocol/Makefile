.DEFAULT_GOAL := build

DOCKER?=docker
BUF_IMAGE := bufbuild/buf:latest
BUF_CMD := $(DOCKER) run --rm -v "${PWD}":/src -w /src $(BUF_IMAGE) 

.PHONY: init
init: ## run buf mod update, format and lint to initialize
	$(DOCKER) pull $(BUF_IMAGE)
	$(BUF_CMD) dep update
	$(BUF_CMD) format -w
	$(BUF_CMD) lint

.PHONY: clean
clean: ## clean generated files
	find ./pb -type f -name "*.go" -delete 

.PHONY: build
build: init clean ## generate go file, doc from proto by using buf
	$(BUF_CMD) generate
	go mod tidy

.PHONY: test
test: build ## run test
	go test -v ./...

.PHONY: help
help: ## prints this help message
	@echo "Usage: \n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
