#!/bin/bash

# GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go
# newman run ./test/postman/gomarvin-tests.postman_collection.json

# from the root
# chmod u+x ./test/test.sh
# ./test/test.sh 0.9.0

# Check if an argument is provided
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

# Util function to check if running the generator again
# wont overwrite the initially generated file
# which should not be updated after the initial generation.
check_file_update() {
  local file_path="$1"
  local new_line="// this is appended by the test script"
  echo $new_line >>$file_path # append a new line to the file
  local file_pre_gen=$(cat $file_path)
  local file_pre_gen_no_ws=$(echo "$file_pre_gen" | tr -d '[:space:]')
  # run the generator again
  ./gomarvin -gut=true -config=${CONFIG_PATH} generate
  local file_post_gen=$(cat $file_path)
  local file_post_gen_no_ws=$(echo "$file_post_gen" | tr -d '[:space:]')
  if [ "$file_pre_gen_no_ws" == "$file_post_gen_no_ws" ]; then
    echo "The updated file contents match."
  else
    echo "The updated file contents do not match the contents in memory."
    exit 1
  fi
}

expected_version="$1"

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
  ./gomarvin -gut=true -config=${CONFIG_PATH} generate # generate the project

  PROJECT_PATH=${PWD}/${example}

  # --------------
  # Test block
  # --------------

  # Search for the version string in the output
  output=$(./gomarvin)
  version_string=$(echo "$output" | awk -F"Version: " '/Version: /{print $2}')
  if [ "$version_string" == "$expected_version" ]; then
    echo "Test passed. Version: $version_string"
  else
    echo "Test failed. Expected version: $expected_version, Found version: $version_string"
    exit 1
  fi

  # check if the generated files dont get overwritten
  check_file_update ${PWD}/${example}/pkg/app/app.go
  check_file_update ${PWD}/${example}/pkg/settings/settings.go
  check_file_update ${PWD}/${example}/cmd/api/main.go
  check_file_update ${PWD}/${example}/internal/api/v1/server/router.go

  # copy the ts client to test/build dir, so that it could be called from client.ts test file
  TS_CLIENT=${PWD}/${example}/client/gomarvin.ts
  cp ${TS_CLIENT} ./
  # copy the python client to test/build dir, so that it could be called from client.ts test file
  PY_CLIENT=${PWD}/${example}/client/gomarvin.py
  cp ${PY_CLIENT} ../gomarvin.py

  cd ${example}   # cd into the generated dir
  go mod tidy     # tidy things
  go mod download # download dependencies at first
  code .

  if [[ ${example} == *"with_modules"* ]]; then
    if [[ ${example} != gin_with_modules ]]; then
      echo " * Running fetch tests for ${example}"
      cd ${PROJECT_PATH}
      # go run main.go &
      # sleep 5   # wait for the server to start

      echo ${TS_CLIENT}

      # go run main.go &
      # kill -9 $(lsof -t -i:4444)
    fi
  fi

  cd .. # go back to build dir to run the binary again
done
