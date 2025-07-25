.PHONY: proto server desktop tag clean help version build install dpkg deps test clean doc

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD) doc
BDIR=build
$(shell mkdir -p $(BDIR))
BINARY_NAME=hxe
VERSION=$(shell grep -e 'VERSION = ".*"' internal/version.go | cut -d= -f2 | sed  s/[[:space:]]*\"//g)

.PHONY: help version build install dpkg deps test clean doc



help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'


server: ## Building hxe backend server
	@echo "Building hxe backend server..."


desktop: ## Building hxe desktop app
	@echo "Building hxe desktop app..."
	cd desktop && npm run build



clean: ## Cleaning build artifacts...
	@echo "Cleaning build artifacts..."
 
tag: ## Tagging the current commit
	@echo "Tagging the current commit..."
	git tag -a v$(VERSION) -m "Update to version $(VERSION)"
	git push origin v$(VERSION)

proto: ## Generating protobuf code
	@echo "Generating protobuf code..."
	@protoc --go_out=. --go_opt=paths=source_relative internal/modules/services/models/models.proto

banner: ## Updating banner
	@echo "Updating banner..."
	./scripts/banner.sh

version: ## Updating version
	@echo "Updating version..."
	./scripts/version.sh

# #!/bin/bash

# # Generate protobuf Go code with custom tags
# echo "Generating protobuf Go code..."

# # Generate standard protobuf code
# protoc --go_out=. --go_opt=paths=source_relative internal/models/models.proto

# # If you have a custom tag processor, you can add it here
# # protoc --go-tag_out=. --go-tag_opt=paths=source_relative internal/models/models.proto

# echo "Protobuf generation complete!" 

