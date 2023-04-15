#!/bin/bash

# GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go
# newman run ./test/postman/gomarvin-tests.postman_collection.json

# from the root
# chmod u+x ./test/test.sh
# ./test/test.sh

CURRENT_DIR=$PWD
GOMARVIN_V='v0.9.0'
GOMARVIN_CONFIG_BASE="gomarvin-"
GOMARVIN_CONFIG_DIR=${CURRENT_DIR}/examples/${GOMARVIN_V}/
BUILD_DIR=./test/build/
DANGEROUS_REGEN=true
GUT=true

# Variable used in the config file name + the name of the
# generated dir ( "project_info": "name": "gin_ecommerce" )
EXAMPLES=(
  # 'gin'
  # 'fiber'
  # 'echo'
  # 'chi'
  # 'gin_with_modules'
  # 'echo_with_modules'
  # 'chi_with_modules'
  'fiber_with_modules'
)

# delete the build dir before generating the project
rm -rf ${BUILD_DIR}
# create a fresh build dir
mkdir ${BUILD_DIR}

# build binary to ./test/build/ and cd into the dir
GOOS=darwin GOARCH=arm64 go build -o ${BUILD_DIR}gomarvin main.go
cd ${BUILD_DIR}

for example in "${EXAMPLES[@]}"; do

  GOMARVIN_CONFIG_FILE="${GOMARVIN_CONFIG_BASE}${example}.json"
  CONFIG_PATH="${GOMARVIN_CONFIG_DIR}${GOMARVIN_CONFIG_FILE}"

  # generate the project
  ./gomarvin -dangerous-regen=${DANGEROUS_REGEN} -gut=true -config=${CONFIG_PATH} generate

  # copy the ts client to test/build dir, so that it could be called from client.ts test file
  TS_CLIENT=${PWD}/${example}/client/gomarvin.gen.ts
  cp ${TS_CLIENT} ./

  # copy the python client to test/build dir, so that it could be called from client.ts test file
  PY_CLIENT=${PWD}/${example}/client/gomarvin.gen.py
  cp ${PY_CLIENT} ../gomarvin.py

  cd ${example}   # cd into the generated dir
  go mod tidy     # tidy things
  go mod download # download dependencies at first
  code .

  # run postman tests on servers that hold the testable endpoints
  # if [[ ${example} == *"with_modules"* ]]; then
  #     echo "--- Running postman tests for ${example}"
  #     # nohup go run main.go &
  #     # newman run ${CURRENT_DIR}/test/postman/gomarvin-tests.postman_collection.json
  #     # kill -9 $(lsof -t -i:4444)
  #     # code .                  # open in vscdoe
  # fi

  cd .. # go back to build dir to run the binary again

done
