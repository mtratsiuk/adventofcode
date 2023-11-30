#!/usr/bin/env bash

set -euo pipefail

. ./bin/env.sh

cd "$(dirname "$0")"/..
cd "./aoc-${AOC_YEAR}"


DAY_NUMBER=$1
DAY_NUMBER_PADDED=$(printf "%02d" $DAY_NUMBER)
BIN_NAME=$(ls -1 ./go/bin/ | grep "^${DAY_NUMBER_PADDED}_")

shift 1

echo "Building $BIN_NAME..."
time go build -o "./go/target/$BIN_NAME" "./go/bin/$BIN_NAME/$BIN_NAME.go" "$@"

echo "Running $BIN_NAME..."
time ./go/target/$BIN_NAME "$@"
