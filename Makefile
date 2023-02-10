# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

########################################################################################################################
### VARIABLES
########################################################################################################################
MAIN_DIR := src

# Image URL to use all building/pushing image targets
VERSION ?= latest
IMG ?= cluster-autoscaler-status-exporter:v${VERSION}
PLATFORM ?= linux/amd64
# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.23

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all
all: test build docker-build

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

########################################################################################################################
### DEVELOPMENT COMMANDS
########################################################################################################################
.PHONY: fmt
fmt: ## Run go fmt against code.
	@cd src && go fmt .

.PHONY: vet
vet: ## Run go vet against code.
	@cd src && go vet .

.PHONY: test
test: fmt vet ## Run tests.

run:
	cd src/ && go run .

########################################################################################################################
### BUILD COMMANDS
########################################################################################################################
build:
	cd src/ && go mod tidy && go build -o ./../bin/case .

.PHONY: docker-build
docker-build: test
	docker buildx build --platform ${PLATFORM} -t ${IMG} .

.PHONY: docker-push
docker-push:
	docker push ${IMG}

########################################################################################################################
### CLEAN COMMANDS
########################################################################################################################
clean:
	rm -r ./tmp
