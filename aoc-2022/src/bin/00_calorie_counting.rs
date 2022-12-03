use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("00_calorie_counting.txt")).unwrap();

    let max = solve_0(&input);
    let max_3 = solve_1(&input);

    println!("max: {max}");
    println!("max_3: {max_3}");
}

pub fn solve_0(input: &str) -> i32 {
    let mut max = 0;
    let mut current = 0;

    for x in input.split("\n") {
        if x.is_empty() {
            max = max.max(current);
            current = 0;
        } else {
            let num_value = i32::from_str_radix(x, 10).unwrap();
            current += num_value;
        }
    }

    return max.max(current);
}

pub fn solve_0_with_fold(input: &str) -> i32 {
    let (max, last) = input.split("\n").fold((0, 0), |(max, cur), val| {
        if val.is_empty() {
            (max.max(cur), 0)
        } else {
            let num_val = i32::from_str_radix(val, 10).unwrap();
            (max, cur + num_val)
        }
    });

    max.max(last)
}

pub fn solve_1(input: &str) -> i32 {
    let mut sums = input
        .split("\n")
        .fold((vec![], 0), |(mut sums, cur), val| {
            if val.is_empty() {
                sums.push(cur);
                (sums, 0)
            } else {
                let num_val = i32::from_str_radix(val, 10).unwrap();
                (sums, cur + num_val)
            }
        })
        .0;

    sums.sort();

    sums.iter().rev().take(3).sum()
}

#[cfg(test)]
mod tests {
    #[test]
    fn find_max_3() {
        let input = "
1
2
3

10
20

100
200

40
50

200
300

1
2

300
400
";
        assert_eq!(super::solve_1(input), 1500);
    }

    #[test]
    fn find_total_of_single() {
        assert_eq!(super::solve_0("10\n20\n30"), 60);
        assert_eq!(super::solve_0_with_fold("10\n20\n30"), 60);
    }

    #[test]
    fn find_total_of_two() {
        let input = "
10
20
30

100
200
";
        assert_eq!(super::solve_0(input), 300);
        assert_eq!(super::solve_0_with_fold(input), 300);
    }

    #[test]
    fn should_support_singles() {
        assert_eq!(
            super::solve_0(
                "
10
20
30

9000

60
70
800
"
            ),
            9000
        );
    }

    #[test]
    fn first_should_be_max() {
        assert_eq!(
            super::solve_0(
                "
100
200
300

10
20
30

40
50
60
"
            ),
            600
        );
    }
}
