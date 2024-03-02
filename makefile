build template:
	@echo "Building templates..."
	@go-bindata -o=internal/asset/asset.go -pkg=asset internal/template/...
	@echo "Done!"

local test:
	@echo "Running tests..."
	@go build -o go-dandelion-cli
	@mv go-dandelion-cli ../../.go/bin
	@echo "Done!"
