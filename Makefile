.PHONY: server desktop clean


server: ## Building hxe backend server
	@echo "Building hxe backend server..."


desktop: ## Building hxe desktop app
	@echo "Building hxe desktop app..."


clean: ## Cleaning build artifacts...
	@echo "Cleaning build artifacts..."
 
tag: ## Tagging the current commit
	@echo "Tagging the current commit..."
	git tag -a v0.1.0 -m "Initial release"
	git push origin v0.1.0



# protoc --go_out=. --go_opt=paths=source_relative internal/models/models.proto

# #!/bin/bash

# # Generate protobuf Go code with custom tags
# echo "Generating protobuf Go code..."

# # Generate standard protobuf code
# protoc --go_out=. --go_opt=paths=source_relative internal/models/models.proto

# # If you have a custom tag processor, you can add it here
# # protoc --go-tag_out=. --go-tag_opt=paths=source_relative internal/models/models.proto

# echo "Protobuf generation complete!" 

