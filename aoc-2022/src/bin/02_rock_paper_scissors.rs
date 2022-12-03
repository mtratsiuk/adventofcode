use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("02_rock_paper_scissors.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {
    get_code_pairs(&input)
        .map(|(them, us)| (Shape::from(&them).unwrap(), Shape::from(&us).unwrap()))
        .fold(0, |mut sum, (them, us)| {
            sum += us.play(&them).score();
            sum += us.score();
            sum
        })
}

fn solve_1(input: &str) -> i32 {
    get_code_pairs(&input)
        .map(|(them, outcome)| {
            (
                Shape::from(&them).unwrap(),
                Outcome::from(&outcome).unwrap(),
            )
        })
        .fold(0, |mut sum, (them, outcome)| {
            let us = Shape::from_outcome(&outcome, &them);
            sum += outcome.score();
            sum += us.score();
            sum
        })
}

fn get_code_pairs(input: &str) -> impl Iterator<Item = (char, char)> + '_ {
    input
        .trim()
        .split("\n")
        .map(|line| match line.chars().collect::<Vec<char>>()[..] {
            [left, _, right] => (left, right),
            _ => panic!("Unexpected input format"),
        })
}

#[derive(Debug, PartialEq, Eq)]
enum Shape {
    Rock,
    Paper,
    Scissors,
}

#[derive(Debug, PartialEq)]
enum Outcome {
    Win,
    Lose,
    Draw,
}

impl Shape {
    fn from(code: &char) -> Result<Self, String> {
        match code {
            'A' => Ok(Shape::Rock),
            'B' => Ok(Shape::Paper),
            'C' => Ok(Shape::Scissors),
            'X' => Ok(Shape::Rock),
            'Y' => Ok(Shape::Paper),
            'Z' => Ok(Shape::Scissors),
            _ => Err(format!("Unexpected code {code}").to_string()),
        }
    }

    fn from_outcome(outcome: &Outcome, other: &Self) -> Self {
        for guess in [Shape::Rock, Shape::Paper, Shape::Scissors] {
            if guess.play(&other) == *outcome {
                return guess;
            }
        }

        panic!("Unreachable")
    }

    fn play(&self, other: &Self) -> Outcome {
        match self {
            Shape::Rock => match other {
                Shape::Rock => Outcome::Draw,
                Shape::Paper => Outcome::Lose,
                Shape::Scissors => Outcome::Win,
            },
            Shape::Paper => match other {
                Shape::Rock => Outcome::Win,
                Shape::Paper => Outcome::Draw,
                Shape::Scissors => Outcome::Lose,
            },
            Shape::Scissors => match other {
                Shape::Rock => Outcome::Lose,
                Shape::Paper => Outcome::Win,
                Shape::Scissors => Outcome::Draw,
            },
        }
    }

    fn score(&self) -> i32 {
        match self {
            Shape::Rock => 1,
            Shape::Paper => 2,
            Shape::Scissors => 3,
        }
    }
}

impl Outcome {
    fn from(code: &char) -> Result<Self, String> {
        match code {
            'X' => Ok(Outcome::Lose),
            'Y' => Ok(Outcome::Draw),
            'Z' => Ok(Outcome::Win),
            _ => Err(format!("Unexpected code {code}").to_string()),
        }
    }

    fn score(&self) -> i32 {
        match self {
            Outcome::Lose => 0,
            Outcome::Draw => 3,
            Outcome::Win => 6,
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "
A Y
B X
C Z
";

        assert_eq!(super::solve_0(&input), 15);
    }

    #[test]
    fn solve_1() {
        let input = "
A Y
B X
C Z
";

        assert_eq!(super::solve_1(&input), 12);
    }
}
