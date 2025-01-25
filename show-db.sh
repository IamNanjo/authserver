#!/usr/bin/env bash

echo "Apps"
sqlite3 -json ./dist/authserver_test.db "SELECT * FROM App;" | jq

echo "Domains"
sqlite3 -json ./dist/authserver_test.db "SELECT * FROM Domain;" | jq

echo "Users"
sqlite3 -json ./dist/authserver_test.db "SELECT * FROM User;" | jq

echo "App Managers"
sqlite3 -json ./dist/authserver_test.db "SELECT * FROM AppManager;" | jq
