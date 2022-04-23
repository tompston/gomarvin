#!/bin/bash

# GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go
# newman run ./test/postman/gomarvin-tests.postman_collection.json

# from the root
# chmod u+x ./test/test.sh
# ./test/test.sh

CURRENT_DIR=$PWD
GOMARVIN_V='v0.3.0'
GOMARVIN_CONFIG_BASE="gomarvin-"
GOMARVIN_CONFIG_DIR=${CURRENT_DIR}/examples/${GOMARVIN_V}/
BUILD_DIR=./test/build/
DANGEROUS_REGEN="true"

# Variable used in the config file name + the name of the
# generated dir ( "project_info": "name": "gin_ecommerce" )
EXAMPLES=(
  'gin'
  'gin_with_modules'
  'fiber'
  'fiber_with_modules'
  'echo'
  'echo_with_modules'
  'chi'
  'chi_with_modules'
)

# build binary to ./test/build/ and cd into the dir
GOOS=darwin GOARCH=arm64 go build -o ${BUILD_DIR}gomarvin main.go 
cd ${BUILD_DIR}

for example in "${EXAMPLES[@]}"; do

    GOMARVIN_CONFIG_FILE="${GOMARVIN_CONFIG_BASE}${example}.json"
    CONFIG_PATH="${GOMARVIN_CONFIG_DIR}${GOMARVIN_CONFIG_FILE}"

    # generate the project
    ./gomarvin -dangerous_regen=${DANGEROUS_REGEN} -config=${CONFIG_PATH}

    cd ${example}           # cd into the generated dir
    go mod tidy             # tidy things
    go mod download         # download dependencies at first
    code .
    

    # run postman tests on servers that hold the testable endpoints
    # TODO : figure out how to echo only summary
    if [[ ${example} == *"with_modules"* ]]; then

        echo "------- Running postman tests for ${example} !"

        
        # nohup go run main.go &  
        # newman run ${CURRENT_DIR}/test/postman/gomarvin-tests.postman_collection.json
        # kill -9 $(lsof -t -i:4444)
        # code .                  # open in vscdoe
    
    fi
    
    cd ..               # go back to build dir to run the binary again

done