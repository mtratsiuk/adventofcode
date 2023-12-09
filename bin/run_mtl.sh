#!/usr/bin/env bash

set -euo pipefail

. ./bin/env.sh

COMPILER_PATH=~/Git/mt-lang/src/cli/compile-file.ts

cd "$(dirname "$0")"/..

echo "Running..."
time cat "$@" | deno run $COMPILER_PATH | deno run --allow-env --allow-read --no-check -q -
