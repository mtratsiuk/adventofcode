use std::collections::HashMap;
use std::collections::VecDeque;
use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("05_supply_stacks.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> String {
    let mut crane = Crane::from(&input);

    crane.operate();
    crane.get_top_crates()
}

fn solve_1(input: &str) -> String {
    let mut crane = Crane::from(&input);

    crane.operate_9001();
    crane.get_top_crates()
}

#[derive(Debug)]
struct Move {
    from: usize,
    to: usize,
    count: usize,
}

#[derive(Debug)]
struct Crane {
    crate_stacks: HashMap<usize, VecDeque<char>>,
    operations: VecDeque<Move>,
}

impl Crane {
    fn from(input: &str) -> Self {
        let (stacks_str, operations_str) = input.split_once("\n\n").unwrap();
        let mut crate_stacks: HashMap<usize, VecDeque<char>> = HashMap::new();
        let mut operations: VecDeque<Move> = VecDeque::new();

        for line in stacks_str.lines().map(|l| l.chars().collect::<Vec<_>>()) {
            if !line.contains(&'[') {
                continue;
            }

            let mut stack_number = 1;
            let mut label_idx = 1;

            while let Some(label) = line.get(label_idx) {
                if !label.is_whitespace() {
                    crate_stacks
                        .entry(stack_number)
                        .and_modify(|stack| stack.push_back(*label))
                        .or_insert(VecDeque::from([*label]));
                }

                stack_number += 1;
                label_idx += 4;
            }
        }

        for line in operations_str.lines() {
            let values = line
                .split_ascii_whitespace()
                .filter_map(|x| x.parse::<usize>().ok());

            match values.collect::<Vec<_>>()[..] {
                [count, from, to] => operations.push_back(Move { count, from, to }),
                _ => panic!("Unexpected input format"),
            }
        }

        Self {
            crate_stacks,
            operations,
        }
    }

    fn operate(&mut self) {
        for Move { from, to, count } in &self.operations {
            for _ in 0..*count {
                let label = self
                    .crate_stacks
                    .get_mut(from)
                    .unwrap()
                    .pop_front()
                    .unwrap();

                self.crate_stacks.get_mut(to).unwrap().push_front(label);
            }
        }
    }

    fn operate_9001(&mut self) {
        for Move { from, to, count } in &self.operations {
            let labels = self
                .crate_stacks
                .get_mut(from)
                .unwrap()
                .drain(0..*count)
                .collect::<VecDeque<_>>();

            for label in labels.iter().rev() {
                self.crate_stacks.get_mut(to).unwrap().push_front(*label);
            }
        }
    }

    fn get_top_crates(&self) -> String {
        let mut top = vec![];

        for idx in 1..=self.crate_stacks.len() {
            top.push(self.crate_stacks.get(&idx).unwrap().get(0).unwrap());
        }

        String::from_iter(top)
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
";

        assert_eq!(super::solve_0(&input), "CMZ");
    }

    #[test]
    fn solve_1() {
        let input = "
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
";

        assert_eq!(super::solve_1(&input), "MCD");
    }
}
