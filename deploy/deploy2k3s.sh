#!/bin/bash
DEPLOY_ACCESS_KEY=$1
DEPLOY_API_URL=$2
DEPLOY_VERSION=$3
DEPLOY_SERVER_NAME=web-screenshot-service
IMAGE_NAME=hkccr.ccs.tencentyun.com/debug.icu/$DEPLOY_SERVER_NAME

curl -X PUT \
    -H "content-type: application/json" \
    -H "Cookie: KuboardUsername=wertycn; KuboardAccessKey=${DEPLOY_ACCESS_KEY}" \
    -d '{"kind":"deployments","namespace":"debug-app","name":"${DEPLOY_SERVER_NAME}","images":{"'${IMAGE_NAME}'":"'${IMAGE_NAME}':'${DEPLOY_VERSION}'"}}' \
    "${DEPLOY_API_URL}"