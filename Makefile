.PHONY: regenerate-docs start help stop cleanup unit-tests

SWAG=${HOME}/go/bin/swag
MOCKGEN=${HOME}/go/bin/mockgen
DESTINATION=mocks/mock_workflowrepository.go
PACKAGE=mocks
NTERFACE=WorkflowRepository

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

regenerate-docs: ## Regenerate the docks for this application
	@echo generating docs...
	$(SWAG) init

stop: ## Stop docker containers
	@docker-compose stop

start: ## Start docker containers and run main.go
	@echo starting app...
	docker-compose up -d
	go run main.go	 

unit-tests: ## Runs only unit tests 
	@echo running unit tests
	go test ./... -short

test: ## Runs all tests (unit and component)
	@echo running tests
	go test ./...

cleanup: ## Clean tests cache
	go clean -v -x -cache -testcache -modcache