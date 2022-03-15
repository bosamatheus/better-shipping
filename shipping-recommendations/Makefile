USECASE_DIR=internal/usecase/recommendation

.DEFAULT_GOAL := help

test: ## Execute all tests
	go test ./...

test-cover: ## Test with coverage
	go test ./... -coverprofile=coverage.out

build-mocks: ## Build mocks
	@go install github.com/vektra/mockery/v2@latest
	@~/go/bin/mockery --all --dir $(USECASE_DIR) --output $(USECASE_DIR)/mocks --case underscore

run: ## Run the application
	go run cmd/shipping_recommendations/main.go

dep: ## Build dependencies
	go mod download

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-30s %s\n", $$1, $$2}'
