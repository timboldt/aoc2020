use aoc_runner_derive::{aoc, aoc_generator};
use std::num::ParseIntError;

#[aoc_generator(day1)]
fn parse_input(input: &str) -> Result<Vec<i32>, ParseIntError> {
    input.lines().map(|l| l.parse()).collect()
}

#[aoc(day1, part1)]
fn part1(entries: &[i32]) -> i32 {
    for v1 in entries {
        for v2 in entries {
            if v1 + v2 == 2020 {
                return v1 * v2;
            }
        }
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_example() {
        assert_eq!(part1(&[1721, 979, 366, 299, 675, 1456]), 514579);
    }
}
