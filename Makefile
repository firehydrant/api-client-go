.DEFAULT_GOAL := generate

SWAGGER_VERSION ?= v0.30.3
GO_VERSION ?= 1.19.3-alpine

generate: dependencies
	@rm -rf ./client
	@rm -rf ./models
	@curl https://api.firehydrant.io/v1/swagger_doc | jq . > swagger_doc.json
	@docker run -v $(shell pwd):/go/src/github.com/firehydrant/api-client-go quay.io/goswagger/swagger:$(SWAGGER_VERSION) generate client -f /go/src/github.com/firehydrant/api-client-go/swagger_doc.json -t /go/src/github.com/firehydrant/api-client-go

dependencies:
	@docker run -w/go/src/github.com/firehydrant/api-client-go -v $(shell pwd):/go/src/github.com/firehydrant/api-client-go golang:$(GO_VERSION) go get ./...

tidy:
	@docker run -w/go/src/github.com/firehydrant/api-client-go -v $(shell pwd):/go/src/github.com/firehydrant/api-client-go golang:$(GO_VERSION) go mod tidy

release: tidy generate
