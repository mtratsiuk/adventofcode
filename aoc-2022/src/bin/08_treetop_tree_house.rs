use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("08_treetop_tree_house.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {
    let grid = get_tree_grid(input);
    let mut visible = 0;

    for (y, row) in grid.iter().enumerate() {
        for (x, _) in row.iter().enumerate() {
            let pos = (x, y);
            let is_visible = is_visible_from(&grid, pos, Direction::West)
                || is_visible_from(&grid, pos, Direction::East)
                || is_visible_from(&grid, pos, Direction::North)
                || is_visible_from(&grid, pos, Direction::South);

            visible += is_visible as i32;
        }
    }

    visible
}

fn solve_1(input: &str) -> i32 {
    let grid = get_tree_grid(input);
    let mut max_score = 0;

    for (y, row) in grid.iter().enumerate() {
        for (x, _) in row.iter().enumerate() {
            let pos = (x, y);
            let score = get_visible_count(&grid, pos, Direction::West)
                * get_visible_count(&grid, pos, Direction::East)
                * get_visible_count(&grid, pos, Direction::North)
                * get_visible_count(&grid, pos, Direction::South);

            max_score = max_score.max(score);
        }
    }

    max_score
}

type Grid = Vec<Vec<u8>>;
type Position = (usize, usize);

enum Direction {
    West,
    East,
    North,
    South,
}

fn is_visible_from(grid: &Grid, (x, y): Position, direction: Direction) -> bool {
    let height = grid.len();
    let width = grid.first().unwrap().len();
    let tree_size = grid[y][x];

    match direction {
        Direction::West => x == 0 || grid[y][0..x].iter().all(|h| *h < tree_size),
        Direction::East => x == width || grid[y][x + 1..].iter().all(|h| *h < tree_size),
        Direction::North => y == 0 || grid[0..y].iter().all(|row| row[x] < tree_size),
        Direction::South => y == height || grid[y + 1..].iter().all(|row| row[x] < tree_size),
    }
}

fn get_visible_count(grid: &Grid, (x, y): Position, direction: Direction) -> i32 {
    let height = grid.len();
    let width = grid.first().unwrap().len();
    let tree_size = grid[y][x];

    match direction {
        Direction::West => {
            if x == 0 {
                0
            } else {
                grid[y][0..x]
                    .iter()
                    .rev()
                    .fold((0, false), |acc @ (count, done), size| {
                        if done {
                            acc
                        } else {
                            (count + 1, *size >= tree_size)
                        }
                    })
                    .0
            }
        }
        Direction::East => {
            if x == width {
                0
            } else {
                grid[y][x + 1..]
                    .iter()
                    .fold((0, false), |acc @ (count, done), size| {
                        if done {
                            acc
                        } else {
                            (count + 1, *size >= tree_size)
                        }
                    })
                    .0
            }
        }
        Direction::North => {
            if y == 0 {
                0
            } else {
                grid[0..y]
                    .iter()
                    .rev()
                    .fold((0, false), |acc @ (count, done), row| {
                        if done {
                            acc
                        } else {
                            (count + 1, row[x] >= tree_size)
                        }
                    })
                    .0
            }
        }
        Direction::South => {
            if y == height {
                0
            } else {
                grid[y + 1..]
                    .iter()
                    .fold((0, false), |acc @ (count, done), row| {
                        if done {
                            acc
                        } else {
                            (count + 1, row[x] >= tree_size)
                        }
                    })
                    .0
            }
        }
    }
}

fn get_tree_grid(input: &str) -> Grid {
    input
        .trim()
        .lines()
        .map(|l| {
            l.chars()
                .map(|c| c.to_string().parse::<u8>().unwrap())
                .collect::<_>()
        })
        .collect::<_>()
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "\
30373
25512
65332
33549
35390
";

        assert_eq!(super::solve_0(&input), 21);
    }

    #[test]
    fn solve_1() {
        let input = "\
30373
25512
65332
33549
35390
";

        assert_eq!(super::solve_1(&input), 8);
    }
}
