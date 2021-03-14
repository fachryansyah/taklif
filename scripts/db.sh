#!/bin/bash
echo "creating local DB for development.."
if [ ! -d ~/.taklif ]; then
  mkdir -p ~/.taklif;
fi
rm ~/.taklif/data.db
touch ~/.taklif/data.db
echo "local DB created!\n"
go run ${PWD}/databases/migration/main.go