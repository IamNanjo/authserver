SHELL := /usr/bin/env bash

# Paths
SCRIPT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
DIST_PATH := ${SCRIPT_PATH}/bin
TEST_DB_PATH := /tmp/authserver/authserver_test.db

# Build options
NAME := authserver
PLATFORMS := linux/amd64 linux/arm64 windows/amd64 windows/arm64 darwin/amd64 darwin/arm64
RELEASE_FLAGS := -ldflags "-s -w"
DEV_FLAGS := -v

# Output colors
NC := \033[0m
RED := '\033[0;31m'
YELLOW := \033[0;33m
GREEN := \033[0;32m

.PHONY: tidy templ clean test dev-setup dev-run dev release-build release

tidy:
	@echo -e "${YELLOW}Running go mod tidy${NC}"
	@go mod tidy
	@echo -e "${GREEN}Finished go mod tidy${NC}\n"

templ:
	@echo -e "${YELLOW}Generating templ pages${NC}"
	@go tool github.com/a-h/templ/cmd/templ generate
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
	@AUTHSERVER_DB="${TEST_DB_PATH}" go test -v -failfast -count=1 ./... || (echo -e "${RED}Tests failed. Fix issues first${NC}" && exit 1)
	@echo -e "${GREEN}Finished tests${NC}\n"

dev-setup: tidy clean sqlc test

dev-run:
	@echo -e "${YELLOW}Building development version of ${NAME}${NC}"
	@CGO_ENABLED=0 go build -o "${DIST_PATH}/${NAME}" ${DEV_FLAGS} ./cmd/server && echo -e "${GREEN}Finished building development version of ${NAME}${NC}\n"
	@echo -e "${YELLOW}Starting server${NC}"
	@if command -v sqlite3 &> /dev/null; then \
		appId=$$(sqlite3 -list ${TEST_DB_PATH} "SELECT id FROM App LIMIT 1;"); \
		echo "https://local.test?app=$$appId&redirect=https://local.test"; \
	fi
	@${DIST_PATH}/authserver

sqlc:
	@echo -e "${YELLOW}Generating sqlc functions${NC}"
	@go tool github.com/sqlc-dev/sqlc/cmd/sqlc compile
	@go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate
	@echo -e "${GREEN}Finished generating sqlc functions${NC}\n"

dev: sqlc
	@go tool github.com/bokwoon95/wgo -xfile _templ.go -xdir ${DIST_PATH} clear :: date :: echo :: make tidy templ dev-run :: sleep 1

release-build:
	@echo -e "${YELLOW}Building release version of ${NAME} for $(words ${PLATFORMS}) platforms${NC}"
	@for platform in ${PLATFORMS}; do \
		IFS='/' read -r GOOS GOARCH <<< "$$platform"; unset IFS; \
		output_name="${DIST_PATH}/${NAME}-$${GOOS}-$${GOARCH}"; \
		if [ "$$GOOS" = "windows" ]; then \
			output_name+=".exe"; \
		fi; \
		printf "Building $$output_name - "; \
		CGO_ENABLED=0 go build -o $$output_name ${RELEASE_FLAGS} ./cmd/server; \
		if [ $$? -ne 0 ]; then \
			echo -e "${RED}failed${NC}\n"; \
			exit 1; \
		else \
			echo -e "${GREEN}ok${NC}"; \
		fi \
	done
	@echo -e "${GREEN}Finished building release version of ${NAME}${NC}\n"

release: tidy clean test release-build
