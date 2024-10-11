.PHONY: build zip gen generate

# 使用例: make build DIR=hello ENV=dev
# 命名に関して: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html#golang-handler-naming
build: ## Build the binary file
	@if [ -z "$(DIR)" ]; then echo "Please provide DIR as an argument. Example: make build DIR=folder_name"; exit 1; fi
	rm -rf tmp/$(DIR)/*
	mkdir -p tmp/$(DIR)
	GOOS=linux go build -ldflags="-X 'env.env=${ENV}' -s -w" -trimpath -o tmp/$(DIR)/bootstrap cmd/$(DIR)/main.go


# 使用例: make zip DIR=hello
zip: ## Zip the binary file
	@if [ -z "$(DIR)" ]; then echo "Please provide DIR as an argument. Example: make zip DIR=folder_name"; exit 1; fi
	zip -j tmp/$(DIR)/function.zip tmp/$(DIR)/bootstrap


# 使用例: make gen DIR=hello [ENV=dev]
gen: ## Generate zip file of the binary for AWS Lambda function
	@if [ -z "$(DIR)" ]; then echo "Please provide DIR as an argument. Example: make gen DIR=folder_name"; exit 1; fi
	make build DIR=$(DIR) ENV=$(ENV)
	make zip DIR=$(DIR)


generate: ## Run go generate
	go generate ./...

