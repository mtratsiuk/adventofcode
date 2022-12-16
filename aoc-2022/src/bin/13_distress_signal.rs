use std::cmp::Ordering;
use std::env;
use std::fs;
use std::iter::Peekable;
use std::str::Chars;

fn main() {
    let pwd = env::current_dir().unwrap();
    let input = fs::read_to_string(pwd.join("data").join("13_distress_signal.txt")).unwrap();

    let result_0 = solve_0(&input);
    let result_1 = solve_1(&input);

    println!("result_0: {result_0}");
    println!("result_1: {result_1}");
}

fn solve_0(input: &str) -> usize {
    input
        .split("\n\n")
        .map(
            |pair| match &pair.lines().map(Packet::parse).collect::<Vec<_>>()[..] {
                [left, right] => left.cmp(right),
                _ => panic!("Unexpected input format"),
            },
        )
        .enumerate()
        .filter(|(_, cmp)| *cmp == Ordering::Less)
        .map(|(idx, _)| idx + 1)
        .sum::<usize>()
}

fn solve_1(input: &str) -> usize {
    let p2 = Packet::parse("[[2]]");
    let p6 = Packet::parse("[[6]]");

    let mut packets = input
        .lines()
        .filter(|l| !l.is_empty())
        .map(Packet::parse)
        .chain(vec![p2.clone(), p6.clone()])
        .collect::<Vec<_>>();

    packets.sort();

    packets.into_iter().enumerate().fold(1, |key, (idx, pkt)| {
        if pkt == p2 || pkt == p6 {
            key * (idx + 1)
        } else {
            key
        }
    })
}

#[derive(Debug, Eq, Clone)]
enum Packet {
    Int(u32),
    List(Vec<Packet>),
}

impl Packet {
    fn parse(line: &str) -> Self {
        let mut parser = PacketParser {
            chars: line.chars().peekable(),
        };

        parser.parse().unwrap()
    }
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        match (self, other) {
            (Packet::Int(l), Packet::Int(r)) => l.cmp(r),
            (Packet::List(l), Packet::List(r)) => {
                for (i, lv) in l.iter().enumerate() {
                    if let Some(rv) = r.get(i) {
                        match lv.cmp(rv) {
                            r @ (Ordering::Less | Ordering::Greater) => {
                                return r;
                            }
                            _ => (),
                        }
                    } else {
                        return Ordering::Greater;
                    }
                }

                l.len().cmp(&r.len())
            }
            (Packet::Int(lv), r @ Packet::List(_)) => Packet::List(vec![Packet::Int(*lv)]).cmp(r),
            (l @ Packet::List(_), Packet::Int(rv)) => l.cmp(&Packet::List(vec![Packet::Int(*rv)])),
        }
    }
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl PartialEq for Packet {
    fn eq(&self, other: &Self) -> bool {
        self.cmp(other) == Ordering::Equal
    }
}

struct PacketParser<'a> {
    chars: Peekable<Chars<'a>>,
}

impl<'a> PacketParser<'a> {
    fn parse(&mut self) -> Option<Packet> {
        match self.chars.peek() {
            Some('[') => self.parse_list(),
            Some(('0'..='9')) => self.parse_int(),
            _ => None,
        }
    }

    fn parse_list(&mut self) -> Option<Packet> {
        assert_eq!(self.chars.next().unwrap(), '[');

        let mut list: Vec<Packet> = vec![];

        loop {
            if let Some(ch) = self.chars.peek() {
                match *ch {
                    ']' => {
                        self.chars.next();
                        break;
                    }
                    ',' => {
                        self.chars.next();
                    }
                    _ => {
                        if let Some(packet) = self.parse() {
                            list.push(packet);
                        } else {
                            return None;
                        }
                    }
                }
            } else {
                return None;
            }
        }

        return Some(Packet::List(list));
    }

    fn parse_int(&mut self) -> Option<Packet> {
        let mut value = self.chars.next().and_then(|x| x.to_digit(10)).unwrap();

        loop {
            let peek = self.chars.peek().cloned();

            match peek {
                Some(',' | ']') | None => break,
                Some(ch @ ('0'..='9')) => {
                    self.chars.next();

                    value = value * 10 + ch.to_digit(10)?;
                }
                _ => {
                    return None;
                }
            }
        }

        Some(Packet::Int(value))
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_0() {
        let input = "\
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
";

        assert_eq!(super::solve_0(&input), 13);
    }

    #[test]
    fn solve_1() {
        let input = "\
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
";

        assert_eq!(super::solve_1(&input), 140);
    }
}
