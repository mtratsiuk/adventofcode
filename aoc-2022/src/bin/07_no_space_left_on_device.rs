use std::collections::HashMap;
use std::collections::VecDeque;
use std::env;
use std::fs;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input =
        fs::read_to_string(pwd.join("data").join("07_no_space_left_on_device.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> i32 {
    let dirs_map = get_dir_size_map(input);

    dirs_map.values().filter(|x| **x <= 100_000).sum()
}

fn solve_1(input: &str) -> i32 {
    let dirs_map = get_dir_size_map(input);

    let total_space = 70000000;
    let required_space = 30000000;
    let occupied_space = *dirs_map.get("/").unwrap();
    let free_space = total_space - occupied_space;
    let min_space_to_free = required_space - free_space;

    *dirs_map
        .values()
        .filter(|x| **x >= min_space_to_free)
        .min()
        .unwrap()
}

type DirSize = i32;

fn get_dir_size_map(input: &str) -> HashMap<String, DirSize> {
    let mut dirs_map: HashMap<String, DirSize> = HashMap::new();
    let mut pwd: VecDeque<String> = VecDeque::new();

    for line in input.trim().lines() {
        match line.split_ascii_whitespace().collect::<Vec<&str>>()[..] {
            ["$", "cd", dir] => {
                match dir {
                    ".." => {
                        pwd.pop_back().unwrap();
                    }
                    dir_name => {
                        pwd.push_back(
                            pwd.back()
                                .map(|path| format!("{path}/{dir_name}"))
                                .or_else(|| Some(dir_name.to_string()))
                                .unwrap()
                                .replace("//", "/"),
                        );
                    }
                };
            }
            ["$", "ls"] => (),
            ["dir", _] => (),
            [size, _] => {
                for dir in pwd.iter().rev() {
                    dirs_map
                        .entry(dir.to_string())
                        .and_modify(|s| *s += size.parse::<i32>().unwrap())
                        .or_insert(size.parse::<i32>().unwrap());
                }
            }
            _ => panic!("Unexpected input format"),
        };
    }

    dirs_map
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "\
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
";

        assert_eq!(super::solve_0(&input), 95437);
    }

    #[test]
    fn solve_1() {
        let input = "\
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
";

        assert_eq!(super::solve_1(&input), 24933642);
    }
}
