#!/usr/bin/env bash

templ generate &&
go mod tidy &&
go build -o dist/
dist/authserver -h
