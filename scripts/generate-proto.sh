#!/bin/bash

# Generate protobuf Go code with custom tags
echo "Generating protobuf Go code..."

# Generate standard protobuf code
protoc --go_out=. --go_opt=paths=source_relative internal/models/models.proto

# If you have a custom tag processor, you can add it here
# protoc --go-tag_out=. --go-tag_opt=paths=source_relative internal/models/models.proto

echo "Protobuf generation complete!" 