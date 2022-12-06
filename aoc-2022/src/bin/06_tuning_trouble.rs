use std::collections::HashSet;
use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("06_tuning_trouble.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {
    let line = input.lines().nth(0).unwrap();
    let marker = line
        .char_indices()
        .position(|(i, _)| line[i.saturating_sub(3)..=i].chars().collect::<HashSet<char>>().len() == 4)
        .unwrap();

    (marker + 1) as i32
}

fn solve_1(input: &str) -> i32 {
    let line = input.lines().nth(0).unwrap();
    let marker = line
        .char_indices()
        .position(|(i, _)| line[i.saturating_sub(13)..=i].chars().collect::<HashSet<char>>().len() == 14)
        .unwrap();

    (marker + 1) as i32
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "mjqjpqmgbljsphdztnvjfqwrcgsmlb";

        assert_eq!(super::solve_0(&input), 7);
    }

    #[test]
    fn solve_0_1() {
        let input = "bvwbjplbgvbhsrlpgdmjqwftvncz";

        assert_eq!(super::solve_0(&input), 5);
    }

    #[test]
    fn solve_1() {
        let input = "mjqjpqmgbljsphdztnvjfqwrcgsmlb";

        assert_eq!(super::solve_1(&input), 19);
    }

    #[test]
    fn solve_1_1() {
        let input = "bvwbjplbgvbhsrlpgdmjqwftvncz";

        assert_eq!(super::solve_1(&input), 23);
    }
}
