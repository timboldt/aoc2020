use aoc_runner_derive::{aoc, aoc_generator};
use std::num::ParseIntError;
use regex::Regex;

#[derive(Debug)]
struct Password {
    n1: usize,
    n2: usize,
    c: char,
    p: String,
}

fn parse_line(line: &str) -> Result<Password, ParseIntError> {
    lazy_static! {
        static ref RE: Regex = Regex::new(
            r"^(\d{1,2})-(\d{1,2}) ([a-z]): ([a-z]+)$"
        )
        .unwrap();
    }
    let caps = RE.captures(line).unwrap();

    Ok(Password {
        n1: caps[1].parse().unwrap(),
        n2: caps[2].parse().unwrap(),
        c: caps[3].parse().unwrap(),
        p: caps[4].to_string(),
    })
}

#[aoc_generator(day2)]
fn parse_input(input: &str) -> Result<Vec<Password>, ParseIntError> {
    input.lines().map(|l| parse_line(l)).collect()
}

#[aoc(day2, part1)]
fn part1(passwords: &[Password]) -> i32 {
    let mut good = 0;
    for pwd in passwords {
        let cnt = pwd.p.chars().filter(|ch| *ch == pwd.c).count();
        if cnt >= pwd.n1 && cnt <= pwd.n2 {
            good = good + 1;
        }
    }
    good
}

// #[aoc(day2, part2)]
// fn part2(entries: &[i32]) -> i32 {
// }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_example() {
        assert_eq!(
            part1(&[
                Password {
                    n1: 1,
                    n2: 3,
                    c: 'a',
                    p: "abcde"
                },
                Password {
                    n1: 1,
                    n2: 3,
                    c: 'b',
                    p: "cdefg"
                },
                Password {
                    n1: 2,
                    n2: 9,
                    c: 'c',
                    p: "ccccccccc"
                },
            ]),
            2
        );
    }
}
