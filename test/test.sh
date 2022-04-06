#!/bin/bash

# GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go

# from the root
# chmod u+x ./test/test.sh
# ./test/test.sh

CURRENT_DIR=$PWD
GOMARVIN_V='v0.1.0'
GOMARVIN_CONFIG_BASE="gomarvin-${GOMARVIN_V}-"
GOMARVIN_CONFIG_DIR=${CURRENT_DIR}/examples/
BUILD_DIR=./test/build/
DANGEROUS_REGEN="true"

# Variable used in the config file name + the name of the
# generated dir ( "project_info": "name": "gin_ecommerce" )
EXAMPLES=(
  # 'gin'
  # 'gin_ecommerce'
  # 'fiber'
  'fiber_ecommerce'
  'echo_ecommerce'
)

# build binary to ./test/build/ and cd into the dir
GOOS=darwin GOARCH=arm64 go build -o ${BUILD_DIR}gomarvin main.go 
cd ${BUILD_DIR}


for example in "${EXAMPLES[@]}"; do

    GOMARVIN_CONFIG_FILE="${GOMARVIN_CONFIG_BASE}${example}.json"
    CONFIG_PATH="${GOMARVIN_CONFIG_DIR}${GOMARVIN_CONFIG_FILE}"

    # generate the project
    ./gomarvin -dangerous_regen=${DANGEROUS_REGEN} -config=${CONFIG_PATH}

    cd ${example}       # cd into the generated dir
    go mod tidy         # tidy things
    go mod download     # download dependencies at first
    # gofmt -s -w .       # format project
    code .              # open in vscdoe
    cd ..               # go back to build dir to run the binary again

done
