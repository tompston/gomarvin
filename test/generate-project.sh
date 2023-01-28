#!/bin/bash
# chmod +x ./test/generate-project.sh

go run main.go -dangerous-regen="true" \
    -config="./examples/v0.6.0/gomarvin-echo_with_modules.json" \
    generate

cd chi_with_modules
go mod tidy
go mod download
code .
cd ..