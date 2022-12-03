#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")"/..

for project in $(ls | grep '^aoc-'); do
  cd $project

  cargo check
  cargo test

  cd ..
done
