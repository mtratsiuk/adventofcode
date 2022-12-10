use std::collections::HashSet;
use std::env;
use std::fmt::Display;
use std::fs;

const SCREEN_WIDTH: usize = 40;
const SCREEN_HEIGHT: usize = 6;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("10_cathode-ray_tube.txt")).unwrap();

    let result_0 = solve_0(&input);
    solve_1(&input);

    println!("result_0: {result_0}");
}

fn solve_0(input: &str) -> i32 {
    let device = Device::from(input);
    let target_cycles = (20..=220).step_by(40).collect::<HashSet<_>>();

    device
        .enumerate()
        .filter_map(|(cycle, state)| {
            let tick_number = cycle as i32 + 1;

            if target_cycles.contains(&tick_number) {
                Some(state.x * tick_number)
            } else {
                None
            }
        })
        .sum()
}

fn solve_1(input: &str) -> () {
    let mut device = Device::from(input);

    device
        .by_ref()
        .take(SCREEN_WIDTH * SCREEN_HEIGHT)
        .fold((), |_, _| ());

    println!("{device}");
}

#[derive(Debug)]
struct Device {
    ops: Vec<String>,
    op_ptr: usize,
    op_cycles: usize,
    cycle: usize,
    state: DeviceState,
    bitmap: Vec<Vec<bool>>,
}

#[derive(Debug, Clone, Copy)]
struct DeviceState {
    x: i32,
}

impl Device {
    fn from(input: &str) -> Self {
        Self {
            ops: input
                .trim()
                .lines()
                .map(|l| l.to_string())
                .collect::<Vec<_>>(),
            op_ptr: 0,
            op_cycles: 0,
            cycle: 0,
            state: DeviceState { x: 1 },
            bitmap: (0..SCREEN_HEIGHT)
                .map(|_| vec![false; SCREEN_WIDTH])
                .collect::<_>(),
        }
    }
}

impl Iterator for Device {
    type Item = DeviceState;

    fn next(&mut self) -> Option<Self::Item> {
        let x = self.cycle % SCREEN_WIDTH;
        let y = self.cycle / SCREEN_WIDTH;

        if (self.state.x - 1..=self.state.x + 1).contains(&(x as i32)) {
            self.bitmap[y][x] = true;
        }

        self.cycle += 1;

        if self.op_ptr >= self.ops.len() {
            return None;
        }

        match self.ops[self.op_ptr]
            .split_ascii_whitespace()
            .collect::<Vec<_>>()[..]
        {
            ["noop"] => {
                self.op_ptr += 1;

                Some(self.state)
            }
            ["addx", val] => {
                self.op_cycles += 1;

                if self.op_cycles < 2 {
                    Some(self.state)
                } else {
                    let cur_state = self.state;

                    self.state.x += val.parse::<i32>().unwrap();
                    self.op_ptr += 1;
                    self.op_cycles = 0;

                    Some(cur_state)
                }
            }
            _ => panic!("Unexpected op code"),
        }
    }
}

impl Display for Device {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for line in &self.bitmap {
            for pixel in line {
                let repr = if *pixel { "#" } else { "." };

                write!(f, "{repr}")?;
            }

            writeln!(f)?;
        }

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "\
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
";

        assert_eq!(super::solve_0(&input), 13140);
    }

    #[test]
    fn solve_1() {
        let input = "\
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
";

        assert_eq!(super::solve_1(&input), ());
    }
}
