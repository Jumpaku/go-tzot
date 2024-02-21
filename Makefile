.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Builds
	sh scripts/fetch_tzot.sh
	go generate -v ./cmd/tzot/...
	go build ./...


.PHONY: examples
examples: ## Builds examples
	go generate -v ./examples/...