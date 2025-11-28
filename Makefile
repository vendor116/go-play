# Makefile

NAME := go-play
NAMESPACE := vendor116
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null | sed 's/^v//' || echo "dev")
IMAGE_TAG := $(NAMESPACE)/$(NAME):$(VERSION)

build:
	$(info Building executable binary file: $(VERSION)...)
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.version=$(VERSION) -s -w" -o bin/go-play ./cmd/go-play

build-docker:
	$(info Building Docker image...)
	docker build -f build/Dockerfile -t $(IMAGE_TAG) .

lint:
	$(info Linting golangci-lint...)
	go tool golangci-lint run ./...

fix:
	$(info Fix golangci-lint...)
	go tool golangci-lint run --fix ./...

generate:
	@echo "Generating OpenAPI models..."
	go tool oapi-codegen --config api/openapi/v1/models.cfg.yaml api/openapi/v1/openapi.yaml
	@echo "Generating OpenAPI server..."
	go tool oapi-codegen --config api/openapi/v1/server.cfg.yaml api/openapi/v1/openapi.yaml
	@echo "Generating OpenAPI client..."
	go tool oapi-codegen --config api/openapi/v1/client.cfg.yaml api/openapi/v1/openapi.yaml

GOLANGCI_LINT_VERSION := v2.6.2
OAPI_CODEGEN_VERSION := v2.5.1

install-tools:
	@echo "Installing oapi-codegen:$(OAPI_CODEGEN_VERSION) ..."
	go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)
	@echo "Installing golangci-lint version: $(GOLANGCI_LINT_VERSION) ..."
	go get -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.PHONY: \
	lint \
	build \
	build-docker \
	install-linter