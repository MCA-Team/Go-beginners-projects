.PHONY: help init
.DEFAULT_GOAL = help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


init: ## Install external dependencies to run this project
	@echo "\nInstalation of external dependencies...\n"
	@go get github.com/fxtlabs/primes