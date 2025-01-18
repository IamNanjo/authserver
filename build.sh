#!/bin/env bash

NC='\033[0m'
YELLOW='\033[0;33m'
GREEN='\033[0;32m'

echo -e "${YELLOW}Generating templ pages$NC"
templ generate &&
echo &&
go mod tidy

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

platforms=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")

echo -e "${YELLOW}Building project for ${#platforms[@]} platforms$NC"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	
    output_name="dist/authserver-$GOOS-$GOARCH"

	if [ $GOOS = "windows" ]; then
		output_name+=".exe"
	fi

	printf "Building $output_name - "

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name

	if [ $? -ne 0 ]; then
   		echo "failed"
		exit 1
	else
		echo -e "${GREEN}ok$NC"
	fi
done
