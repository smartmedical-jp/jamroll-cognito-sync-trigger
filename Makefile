.PHONY: build zip gen generate

# 命名に関して: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html#golang-handler-naming
# 使用例: make build DIR=hello
build: ## Build the binary file
	@if [ -z "$(DIR)" ]; then echo "Please provide DIR as an argument. Example: make build DIR=folder_name"; exit 1; fi
	rm -rf ./tmp/$(DIR)/*
	mkdir -p ./tmp/$(DIR)
	go build -o ./tmp/$(DIR)/bootstrap ./cmd/$(DIR)/main.go

# 使用例: make zip DIR=hello
zip: ## Zip the binary file
	@if [ -z "$(DIR)" ]; then echo "Please provide DIR as an argument. Example: make zip DIR=folder_name"; exit 1; fi
	zip -j ./tmp/$(DIR)/function.zip ./tmp/$(DIR)/bootstrap

# 使用例: make gen DIR=hello
gen: ## Generate zip file of the binary for AWS Lambda function
	@if [ -z "$(DIR)" ]; then echo "Please provide DIR as an argument. Example: make gen DIR=folder_name"; exit 1; fi
	make build DIR=$(DIR)
	make zip DIR=$(DIR)

generate: ## Run go generate
	go generate ./...

