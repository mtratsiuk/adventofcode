use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("04_camp_cleanup.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {
    get_range_pairs(&input)
        .iter()
        .fold(0, |count, (left, right)| {
            if left.contains(&right) || right.contains(&left) {
                count + 1
            } else {
                count
            }
        })
}

fn solve_1(input: &str) -> i32 {
    get_range_pairs(&input)
        .iter()
        .fold(0, |count, (left, right)| {
            if left.overlaps(&right) {
                count + 1
            } else {
                count
            }
        })
}

fn get_range_pairs(input: &str) -> Vec<(Range, Range)> {
    input
        .trim()
        .lines()
        .flat_map(|line| line.split(",").flat_map(|range| range.split("-")))
        .collect::<Vec<_>>()
        .chunks(4)
        .map(|chunk| match chunk {
            [a, b, c, d] => (
                Range(a.parse::<_>().unwrap(), b.parse::<_>().unwrap()),
                Range(c.parse::<_>().unwrap(), d.parse::<_>().unwrap()),
            ),
            _ => panic!("Unexpected input format"),
        })
        .collect::<Vec<_>>()
}

#[derive(Debug)]
struct Range(i32, i32);

impl Range {
    fn contains_point(&self, point: i32) -> bool {
        point >= self.0 && point <= self.1
    }

    fn contains(&self, other: &Self) -> bool {
        self.contains_point(other.0) && self.contains_point(other.1)
    }

    fn overlaps(&self, other: &Self) -> bool {
        self.contains_point(other.0)
            || self.contains_point(other.1)
            || other.contains_point(self.0)
            || other.contains_point(self.1)
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
";

        assert_eq!(super::solve_0(&input), 2);
    }

    #[test]
    fn solve_1() {
        let input = "
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
4-8,8-10
";

        assert_eq!(super::solve_1(&input), 4);
    }
}
