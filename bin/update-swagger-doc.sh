#!/usr/bin/env bash

set -eo pipefail

# If API_ENDPOINT is undefined, use the default value for the FireHydrant API
API_ENDPOINT=${API_ENDPOINT:-https://api.firehydrant.io/v1/swagger_doc}

curl "${API_ENDPOINT}" | jq . > swagger_doc.json
