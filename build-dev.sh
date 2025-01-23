#!/usr/bin/env bash

NC='\033[0m'
YELLOW='\033[0;33m'
GREEN='\033[0;32m'

echo -e "${YELLOW}Generating templ pages$NC" &&
templ generate &&
go mod tidy &&

if [ "$1" = "--skiptests" ] || [ "$1" = "-st" ]; then
	echo "Skipping tests"
else 
	echo -e "${YELLOW}Running tests$NC"
	go test ./...

	if [ $? -ne 0 ]; then
		echo "Tests failed. Fix issues first"
		exit 1
	else
		echo
	fi
fi

echo -e "${YELLOW}Building go packages$NC" &&
go build -o dist/ -v &&
echo -e "${GREEN}All built$NC" &&
echo -e "${YELLOW}Starting server$NC" &&
dist/authserver
