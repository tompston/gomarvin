#!/bin/bash
# chmod +x ./test/generate-fetch.sh


go run main.go -dangerous-regen=true \
    -config="./examples/v0.7.0/gomarvin-chi_with_modules.json" \
    -fetch-only=true generate

go run main.go -dangerous-regen=true \
    -config="./examples/v0.7.0/gomarvin-chi_with_modules.json" \
    -fetch-only=true -gut=true generate