#!/bin/bash

# from the root
# chmod u+x ./test/frontend.sh
# ./test/run-frontend.sh


# https://learning.postman.com/docs/running-collections/using-newman-cli/command-line-integration-with-newman/
# npm install -g newman
newman run ./test/postman/gomarvin-tests.postman_collection.json


## -- start frontend app to test fetch functions
# CURRENT_DIR=$PWD
# FRONTEND_DIR=${CURRENT_DIR}/test/fetch/

# cd ${FRONTEND_DIR}
# npm i
# npm run dev