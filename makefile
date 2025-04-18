SHELL := /usr/bin/env bash

# Paths
SCRIPT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
DIST_PATH := ${SCRIPT_PATH}/dist
TEST_DB_PATH := ${DIST_PATH}/authserver_test.db

# Build options
NAME := authserver
PLATFORMS := linux/amd64 linux/arm64 windows/amd64 windows/arm64 darwin/amd64 darwin/arm64
RELEASE_FLAGS := -ldflags "-s -w"
DEV_FLAGS :=

# Output colors
NC := \033[0m
RED := '\033[0;31m'
YELLOW := \033[0;33m
GREEN := \033[0;32m

.PHONY: tidy templ clean test dev release

tidy:
	@echo -e "${YELLOW}Running go mod tidy${NC}"
	@go mod tidy
	@echo -e "${GREEN}Finished go mod tidy${NC}\n"

templ:
	@echo -e "${YELLOW}Generating templ pages${NC}"
	@go tool templ generate
	@echo -e "${GREEN}Finished generating templ pages${NC}\n"

clean:
	@echo -e "${YELLOW}Cleaning test DB and binaries${NC}"
	@rm -fv "${TEST_DB_PATH}" "${TEST_DB_PATH}-shm" "${TEST_DB_PATH}-wal" "${DIST_PATH}/${NAME}"
	@for platform in ${PLATFORMS}; do \
		IFS='/' read -r GOOS GOARCH <<< "$$platform"; unset IFS; \
		filename="${DIST_PATH}/${NAME}-$${GOOS}-$${GOARCH}"; \
		if [ "$$GOOS" = "windows" ]; then \
			filename+=".exe"; \
		fi; \
		rm -fv "$$filename"; \
	done
	@echo -e "${GREEN}Files cleaned${NC}\n"

test:
	@echo -e "${YELLOW}Running tests${NC}"
	@AUTHSERVER_DB="${TEST_DB_PATH}" go test -count=1 ./... || (echo -e "${RED}Tests failed. Fix issues first${NC}" && exit 1)
	@echo -e "${GREEN}Finished tests${NC}\n"

dev:
	@echo -e "${YELLOW}Building development version of ${NAME}${NC}"
	@go build -o "dist/${NAME}" -v ${DEV_FLAGS} && echo -e "${GREEN}Finished building development version of ${NAME}${NC}\n"
	@echo -e "${YELLOW}Starting server${NC}"
	@if command -v sqlite3 &> /dev/null; then \
		appId=$$(sqlite3 -column dist/authserver_test.db "SELECT id FROM App;"); \
		echo "https://local.test?app=$$appId&redirect=https://local.test"; \
	fi
	@./dist/authserver

release:
	@echo -e "${YELLOW}Building release version of ${NAME} for $(words ${PLATFORMS}) platforms${NC}"
	@for platform in ${PLATFORMS}; do \
		IFS='/' read -r GOOS GOARCH <<< "$$platform"; unset IFS; \
		output_name="${DIST_PATH}/${NAME}-$${GOOS}-$${GOARCH}"; \
		if [ "$$GOOS" = "windows" ]; then \
			output_name+=".exe"; \
		fi; \
		printf "Building $$output_name - "; \
		go build -o $$output_name ${RELEASE_FLAGS}; \
		if [ $$? -ne 0 ]; then \
			echo -e "${RED}failed${NC}\n"; \
			exit 1; \
		else \
			echo -e "${GREEN}ok${NC}"; \
		fi \
	done
	@echo -e "${GREEN}Finished building release version of ${NAME}${NC}\n"
