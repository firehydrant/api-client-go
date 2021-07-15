generate:
	@echo ""
	@echo "###################################################################################"
	@echo "# Generating API client!                                                          #"
	@echo "###################################################################################"
	@echo ""
	@echo ""
	@rm -rf ./client
	@rm -rf ./models
	@curl -o swagger_doc.json https://api.firehydrant.io/v1/swagger_doc
	@docker run -v $(shell pwd):/go/src/github.com/firehydrant/api-client-go quay.io/goswagger/swagger generate client -f /go/src/github.com/firehydrant/api-client-go/swagger_doc.json -t /go/src/github.com/firehydrant/api-client-go
	@rm swagger_doc.json
	@echo ""
	@echo "⚠️ ⚠️ ⚠️  the 'go get' suggestion above will not work if this repository is not located at that path -- you can run 'make gen-dependencies' instead."
	@echo ""

gen-dependencies:
	@echo ""
	@echo "###################################################################################"
	@echo "# 'go get'ing Swagger Dependencies!                                               #"
	@echo "###################################################################################"
	@echo ""
	go get -u $(shell pwd)/...