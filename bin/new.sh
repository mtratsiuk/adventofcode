#!/usr/bin/env bash

set -euo pipefail

. ./bin/env.sh

cd "$(dirname "$0")"/..

AOC_YEAR=${AOC_YEAR:-2022}
AOC_TOKEN=${AOC_TOKEN}

DAY_NUMBER=$1
DAY_NUMBER_PADDED=$(printf "%02d" $DAY_NUMBER)
DAY_NAME=$(echo "$2" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')

INPUT_URL="https://adventofcode.com/${AOC_YEAR}/day/${DAY_NUMBER}/input"
FILE_NAME="${DAY_NUMBER_PADDED}_${DAY_NAME}"
INPUT_PATH="./aoc-${AOC_YEAR}/data/${FILE_NAME}.txt"
SOLUTION_PATH="./aoc-${AOC_YEAR}/src/bin/${FILE_NAME}.rs"


curl --cookie "session=${AOC_TOKEN}" $INPUT_URL > $INPUT_PATH

echo 'use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join('"\"${FILE_NAME}.txt\""')).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {

}

fn solve_1(input: &str) -> i32 {

}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "
";

        assert_eq!(super::solve_0(&input), 0);
    }

    #[test]
    fn solve_1() {
        let input = "
";

        assert_eq!(super::solve_1(&input), 0);
    }
}
' > $SOLUTION_PATH
