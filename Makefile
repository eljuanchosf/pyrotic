.PHONY: all

.DEFAULT_GOAL := help

COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
REPO := $(shell basename `git rev-parse --show-toplevel`)
DATE := $(shell date +%Y-%m-%d-%H-%M-%S)
APP_NAME := pyrotic

GO_BUILD_FLAGS=-ldflags="-X 'github.com/code-gorilla-au/pyrotic/internal/commands.version=dev-$(BRANCH)-$(COMMIT)'"

# commands
CMD_NAME ?= newCommand

test: ## Run unit tests
	go test --short -cover -failfast ./...

test_integration: build ## run integration test
	ENV=DEV ./pyrotic -p example/_templates generate fakr --meta foo=bar,bin=baz,enum_list=a long list

scan: ## run security scan
	gosec ./...

build: log scan ## build go files
	go build $(GO_BUILD_FLAGS) -o $(APP_NAME)

code_gen: ## generate code
	go generate

generate_cmd: build ## gernate new command
	./$(APP_NAME) generate cmd --name $(CMD_NAME)


# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)



#####################
#  Private Targets  #
#####################

log: # log env vars
	@echo "\n"
	@echo "COMMIT               $(COMMIT)"
	@echo "BRANCH               $(BRANCH)"
	@echo "APP_NAME             $(APP_NAME)"
	@echo "DATE                 $(DATE)"
	@echo "\n"
