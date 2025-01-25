#!/usr/bin/env bash

NC='\033[0m'
YELLOW='\033[0;33m'
GREEN='\033[0;32m'
RED='\033[0;31m'

echo -e "${YELLOW}Generating templ pages$NC" &&
templ generate &&
go mod tidy &&

if [ "$1" = "--skiptests" ] || [ "$1" = "-st" ]; then
	echo "Skipping tests"
else 
	echo -e "${YELLOW}Running tests$NC"
	go test -count=1 ./...

	if [ $? -ne 0 ]; then
		echo -e "${RED}Tests failed. Fix issues first$NC"
		exit 1
	else
		echo
	fi
fi

echo -e "${YELLOW}Building go packages$NC"
go build -o dist/ -v &&
echo -e "${GREEN}All built$NC"

echo -e "${YELLOW}Starting server$NC"

if command -v sqlite3 &> /dev/null; then
	appId=$(sqlite3 -column dist/authserver_test.db "SELECT id FROM App;")
	echo "https://local.test?app=$appId&redirect=https://local.test"
fi

dist/authserver
