GO_BUILD_FLAGS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

MAIN_DIR=cmd/shipping_recommendations/main.go
USECASE_DIR=internal/usecase/shipping
LOGWRAPPER_DIR=pkg/logwrapper

.DEFAULT_GOAL := help

test: ## Execute all tests
	go test -v ./...

test-cover: ## Test with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build-mocks: ## Build mocks
	@go install github.com/vektra/mockery/v2@latest
	@~/go/bin/mockery --all --dir ${USECASE_DIR} --output ${USECASE_DIR}/mocks --case underscore
	@~/go/bin/mockery --all --dir ${LOGWRAPPER_DIR} --output ${LOGWRAPPER_DIR}/mocks --case underscore

lint: ## Lint the code with golangci-lint
	@~/go/bin/golangci-lint run

run: ## Run the application
	go run ${MAIN_DIR}

build: ## Build the application
	${GO_BUILD_FLAGS} go build -o bin/shipping_recommendations ${MAIN_DIR}

dep: ## Build dependencies
	go mod download

build-docs: ## Generate API documentation
	swag init -g ${MAIN_DIR} -o api

fmt-docs: ## Format comments for API documentation
	swag fmt -g ${MAIN_DIR}

clean: ## Clean the test and build artifacts
	rm -f coverage.out
	rm -f bin/shipping_recommendations

help: ## Show this help
	@egrep -h '\s##\s' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-30s %s\n", $$1, $$2}'
