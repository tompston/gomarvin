#!/bin/bash
# chmod +x ./test/generate-project.sh

go run main.go -dangerous-regen="true" \
    -config="./examples/v0.5.0/gomarvin-chi_with_modules.json" \
    generate

cd chi_with_modules
go mod tidy
go mod download
code .
cd ..