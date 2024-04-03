#!/usr/bin/env bash

set -eu

# Get root of repository.
repo_root=$(git rev-parse --show-toplevel)

set -x

pushd "${repo_root}" > /dev/null
  bin/swagger generate client -f swagger_doc.json

  go mod tidy
popd > /dev/null
