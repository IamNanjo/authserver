#!/usr/bin/env bash

NC='\033[0m'
YELLOW='\033[0;33m'
GREEN='\033[0;32m'

echo -e "${YELLOW}Generating templ pages$NC" &&
templ generate &&
go mod tidy &&
echo -e "${YELLOW}Building go packages$NC" &&
go build -o dist/ -v &&
echo -e "${GREEN}All built$NC" &&
echo -e "${YELLOW}Starting server$NC" &&
dist/authserver
