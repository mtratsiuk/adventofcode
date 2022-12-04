use std::collections::HashMap;
use std::collections::HashSet;
use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input =
        fs::read_to_string(pwd.join("data").join("03_rucksack_reorganization.txt")).unwrap();

    let prios_map = build_priorities_map();
    let sum_0 = solve_0(&input, &prios_map);
    let sum_1 = solve_1(&input, &prios_map);

    println!("sum_0: {sum_0}");
    println!("sum_1: {sum_1}");
}

pub fn build_priorities_map() -> HashMap<char, i32> {
    let mut cur_prio = 1;

    ('a'..='z')
        .chain('A'..='Z')
        .fold(HashMap::new(), |mut prios, char| {
            prios.insert(char, cur_prio);
            cur_prio += 1;
            prios
        })
}

pub fn solve_0(input: &str, prios_map: &HashMap<char, i32>) -> i32 {
    let mut sum = 0;

    for line in input.trim().lines() {
        let buckets = line.split_at(line.len() / 2);
        let left = buckets.0.chars().collect::<HashSet<_>>();
        let right = buckets.1.chars().collect::<HashSet<_>>();
        let intersection = left.intersection(&right).next().unwrap();

        sum += prios_map.get(intersection).unwrap();
    }

    sum
}

pub fn solve_1(input: &str, prios_map: &HashMap<char, i32>) -> i32 {
    let mut sum = 0;

    for group in input
        .trim()
        .lines()
        .map(|line| line.chars().collect::<HashSet<_>>())
        .collect::<Vec<HashSet<_>>>()
        .chunks(3)
    {
        match group {
            [first, second, third] => {
                let right_intersection = second
                    .intersection(&third)
                    .cloned()
                    .collect::<HashSet<_>>();

                let intersection = first.intersection(&right_intersection).next().unwrap();

                sum += prios_map.get(intersection).unwrap();
            }
            _ => panic!("Unexpected group size"),
        }
    }

    sum
}

#[cfg(test)]
mod tests {
    #[test]
    fn sum_0() {
        let input = "
abcddxyz
abcAAdef
";
        let prios = super::build_priorities_map();

        assert_eq!(super::solve_0(&input, &prios), 4 + 27);
    }

    #[test]
    fn sum_1() {
        let input = "
abcA
eAgh
igAl
mZop
Zrst
uvwZ
";
        let prios = super::build_priorities_map();

        assert_eq!(super::solve_1(&input, &prios), 27 + 52);
    }
}
