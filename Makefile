.PHONY: clean build zip gen

build: ## Build the binary file: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html#golang-handler-naming
	GOOS=linux GOARCH=amd64 go build -o ./tmp/bootstrap ./main.go

zip: ## Zip the binary file
	zip -j ./tmp/function.zip ./tmp/bootstrap

gen: ## Generate zip file of the binary for AWS Lambda function
	make build
	make zip

