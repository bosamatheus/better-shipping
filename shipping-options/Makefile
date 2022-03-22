.DEFAULT_GOAL := help

run: ## Run the application
	cargo run src/main.rs

help: ## Show this help
	@egrep -h '\s##\s' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-30s %s\n", $$1, $$2}'
