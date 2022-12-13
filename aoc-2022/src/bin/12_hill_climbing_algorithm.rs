use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input =
        fs::read_to_string(pwd.join("data").join("12_hill_climbing_algorithm.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

type Grid = Vec<Vec<(i32, i32)>>;

struct State {
    grid: Vec<Vec<(i32, i32)>>,
    start_pos: (usize, usize),
    end_pos: (usize, usize),
}

impl State {
    fn from(input: &str) -> (Self, Vec<(usize, usize)>) {
        let mut grid: Grid = vec![];
        let mut a_pos: Vec<(usize, usize)> = vec![];
        let mut start_pos: (usize, usize) = (0, 0);
        let mut end_pos: (usize, usize) = (0, 0);

        for (y, line) in input.trim().lines().enumerate() {
            grid.push(vec![]);

            for (x, char) in line.chars().enumerate() {
                let height = match char {
                    'S' => {
                        start_pos = (x, y);
                        'a' as i32
                    }
                    'E' => {
                        end_pos = (x, y);
                        'z' as i32
                    }
                    'a' => {
                        a_pos.push((x, y));

                        'a' as i32
                    }
                    ch => ch as i32,
                };

                grid[y].push((height, i32::MAX));
            }
        }

        (
            Self {
                grid,
                start_pos,
                end_pos,
            },
            a_pos,
        )
    }

    fn get_distance(&self, (x, y): (usize, usize)) -> i32 {
        self.grid[y][x].1
    }
}

fn go(state: &mut State, (x, y): (usize, usize), distance: i32) {
    state.grid[y][x].1 = distance;

    let height = state.grid[y][x].0;

    let neighbours = [
        if x != 0 { Some((x - 1, y)) } else { None },
        Some((x + 1, y)),
        if y != 0 { Some((x, y - 1)) } else { None },
        Some((x, y + 1)),
    ]
    .into_iter()
    .filter_map(|x| x)
    .collect::<Vec<_>>();

    for (x, y) in neighbours {
        if let Some((h, d)) = state.grid.get(y).and_then(|c| c.get(x)) {
            if *h - height <= 1 && *d > distance + 1 {
                go(state, (x, y), distance + 1);
            }
        }
    }
}

fn solve_0(input: &str) -> i32 {
    let (mut state, _) = State::from(input);
    let start_pos = state.start_pos;

    go(&mut state, start_pos, 0);

    state.get_distance(state.end_pos)
}

fn solve_1(input: &str) -> i32 {
    let mut distances: Vec<i32> = vec![];
    let (state, a_pos) = State::from(input);

    for start_pos in a_pos {
        let mut run_state = State {
            grid: state.grid.clone(),
            start_pos,
            ..state
        };

        go(&mut run_state, start_pos, 0);

        distances.push(run_state.get_distance(run_state.end_pos));
    }

    distances.into_iter().min().unwrap()
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "\
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
";

        assert_eq!(super::solve_0(&input), 31);
    }

    #[test]
    fn solve_1() {
        let input = "\
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
";

        assert_eq!(super::solve_1(&input), 29);
    }
}
