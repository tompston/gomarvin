#!/bin/bash

# from the root
# chmod u+x ./test/run-frontend.sh
# ./test/run-frontend.sh

CURRENT_DIR=$PWD
GOMARVIN_V='v0.1.0'
FRONTEND_DIR=${CURRENT_DIR}/test/fetch/

cd ${FRONTEND_DIR}
npm i
npm run dev
