#!/usr/bin/env bash

set -eu

# Get root of repository.
repo_root=$(git rev-parse --show-toplevel)

set -x

pwd

pushd "${repo_root}" > /dev/null
  pwd
  bin/deps.sh github.com/go-swagger/go-swagger 
  ls bin
  swagger version

  rm -rf client models
  bin/swagger generate client -f swagger_doc.json

  go mod tidy
popd > /dev/null
