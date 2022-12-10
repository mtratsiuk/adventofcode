use std::collections::HashSet;
use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("09_rope_bridge.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {
    solve_rope_len(input, 2)
}

fn solve_1(input: &str) -> i32 {
    solve_rope_len(input, 10)
}

fn solve_rope_len(input: &str, len: i32) -> i32 {
    let mut visited = HashSet::from([(0, 0)]);
    let mut rope = (0..len).map(|_| (0, 0)).collect::<Vec<_>>();

    for (dir_str, count_str) in input.trim().lines().map(|l| l.split_once(" ").unwrap()) {
        let count = count_str.parse::<i32>().unwrap();
        let (xm, ym) = match dir_str {
            "U" => (0, 1),
            "D" => (0, -1),
            "L" => (-1, 0),
            "R" => (1, 0),
            _ => panic!("Unexpected direction code {dir_str}"),
        };

        for _ in 0..count {
            let (mut xh, mut yh) = rope.first().unwrap();

            (xh, yh) = (xh + xm, yh + ym);
            rope[0] = (xh, yh);

            for idx in 1..rope.len() {
                let (xh, yh) = rope[idx - 1];
                let (mut xt, mut yt) = rope[idx];

                let dx = i32::abs(xh - xt);
                let dy = i32::abs(yh - yt);

                if dx + dy <= 1 || (dx + dy == 2 && dx * dy != 0) {
                    // overlapping or adjacent
                    // noop
                } else if dx + dy == 2 && dx * dy == 0 {
                    // same row or column
                    (xt, yt) = ((xt + xh) / 2, (yt + yh) / 2);
                } else if dx + dy >= 3 {
                    // different row and column, move diagonally
                    (xt, yt) = (
                        if dx == 1 { xh } else { (xt + xh) / 2 },
                        if dy == 1 { yh } else { (yt + yh) / 2 },
                    );
                }

                rope[idx] = (xt, yt);
            }

            visited.insert(*rope.last().unwrap());
        }
    }

    visited.len() as i32
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "\
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
";

        assert_eq!(super::solve_0(&input), 13);
    }

    #[test]
    fn solve_1() {
        let input = "\
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
";

        assert_eq!(super::solve_1(&input), 36);
    }
}
