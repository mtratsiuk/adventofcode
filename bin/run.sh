#!/usr/bin/env bash

set -euo pipefail

. ./bin/env.sh

cd "$(dirname "$0")"/..
cd "./aoc-${AOC_YEAR}"


DAY_NUMBER=$1
DAY_NUMBER_PADDED=$(printf "%02d" $DAY_NUMBER)
BIN_NAME=$(ls -1 ./src/bin/ | grep "^${DAY_NUMBER_PADDED}_" | sed 's/\.rs//g')

echo "Running ${BIN_NAME}..."
cargo run --bin $BIN_NAME
