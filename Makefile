GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
 
ifneq (,$(wildcard ./.env))
    include .env
	export $(shell sed 's/=.*//' .env)
endif
 
.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@echo "  ${YELLOW}help            ${RESET} Show this help message"
	@echo "  ${YELLOW}run             ${RESET} Run cli application"
	@echo "  ${YELLOW}build           ${RESET} Build application binary"
	@echo "  ${YELLOW}setup           ${RESET} Setup local environment"
	@echo "  ${YELLOW}check           ${RESET} Run tests, linters and tidy of the project"
	@echo "  ${YELLOW}test            ${RESET} Run tests only"
	@echo "  ${YELLOW}lint            ${RESET} Run linters via golangci-lint"
	@echo "  ${YELLOW}tidy            ${RESET} Run tidy for go module to remove unused dependencies" 

.PHONY: build
build:  
	OS="$(OS)" APP="cli" ./hacks/build.sh

.PHONY: build-docker
build-docker:  
	sudo chmod 666 /var/run/docker.sock
	docker build .
 
.PHONY: setup
setup:
	cp .env.example .env

.PHONY: check
check: %: tidy lint test

.PHONY: test
test:
	TEST_RUN_ARGS="$(TEST_RUN_ARGS)" TEST_DIR="$(TEST_DIR)" ./hacks/run-tests.sh

.PHONY: lint
lint:
	golangci-lint run --out-format=colored-line-number

.PHONY: fix
fix:
	golangci-lint run --fix  --out-format=colored-line-number

.PHONY: tidy
tidy:
	go mod tidy -v

.PHONY: upgrade
upgrade:
	go get -u ./... && go mod tidy -v
 
.PHONY: run
run:
	@OS="$(OS)" APP="cli" ./hacks/build.sh
	@./bin/cli ./cmd/cli/clouds.mp4 ./cmd/cli/dji.mp4 
