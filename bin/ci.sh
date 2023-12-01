#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")"/..

for project in $(ls | grep '^aoc-'); do
  cd $project

  if [ -f ./Cargo.toml ]; then
    cargo check
    cargo test
  fi

  go test ./...
  not_formatted=$(gofmt -l .)

  if [ ! -z "$not_formatted" ]; then
    echo "Following files are not formatted:"
    echo "$not_formatted"
    echo "Please run 'go fmt'"
    exit 1
  fi

  cd ..
done
