#!/bin/bash
# chmod +x ./test/generate-project.sh

go run main.go -dangerous-regen=true \
    -config="./examples/v0.7.0/gomarvin-echo_with_modules.json" \
    generate

cd echo_with_modules
go mod tidy
go mod download
code .
cd ..
