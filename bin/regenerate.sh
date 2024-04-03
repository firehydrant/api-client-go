#!/usr/bin/env bash

set -eu

# Get root of repository.
repo_root=$(git rev-parse --show-toplevel)

pushd "${repo_root}" > /dev/null
  bin/deps.sh github.com/go-swagger/go-swagger 

  rm -rf client models
  bin/swagger generate client -f swagger_doc.json

  go mod tidy
popd > /dev/null
