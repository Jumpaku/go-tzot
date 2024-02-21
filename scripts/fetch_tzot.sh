#!/bin/sh

set -eux

RELEASE_URL='https://api.github.com/repos/Jumpaku/tz-offset-transitions/releases/latest'
LATEST_TAG=$(curl "${RELEASE_URL}" | jq -r '.tag_name')

RAW_CONTENT_URL="https://github.com/Jumpaku/tz-offset-transitions/raw/${LATEST_TAG}"

OUTPUT_DIR='tz-offset-transitions'
mkdir -p "${OUTPUT_DIR}"
curl -L "${RAW_CONTENT_URL}/gen/version" > "${OUTPUT_DIR}/version"
curl -L "${RAW_CONTENT_URL}/gen/tzot.json" > "${OUTPUT_DIR}/tzot.json"
