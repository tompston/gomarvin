#!/bin/bash
# chmod +x ./test/generate-fetch.sh

go run main.go -dangerous-regen="true" \
    -config="./examples/tmp/gomarvin.json" \
    -fetch-only="true" generate