use aoc_runner_derive::{aoc, aoc_generator};
use std::convert::Infallible;

fn num_trees(rows: &[String], col_step: usize, row_step: usize) -> u64 {
    let mut col: usize = 0;
    let mut row: usize = 0;
    let mut trees = 0;
    for rowtxt in rows {
        if row == 0 {
            let ch = rowtxt.chars().nth(col).unwrap();
            if ch == '#' {
                trees = trees + 1;
            }
            // println!("{}", &rowtxt[0..col+1]);
            col = (col + col_step) % rowtxt.len();
        }
        row = (row + 1) % row_step;
    }
    trees
}

#[aoc_generator(day3)]
fn parse_input(input: &str) -> Result<Vec<String>, Infallible> {
    input.lines().map(|l| l.parse()).collect()
}

#[aoc(day3, part1)]
fn part1(rows: &[String]) -> u64 {
    num_trees(rows, 3, 1)
}

#[aoc(day3, part2)]
fn part2(rows: &[String]) -> u64 {
    num_trees(rows, 1, 1)
        * num_trees(rows, 3, 1)
        * num_trees(rows, 5, 1)
        * num_trees(rows, 7, 1)
        * num_trees(rows, 1, 2)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn num_trees_test() {
        let data = &[
            "..##.......".to_string(),
            "#...#...#..".to_string(),
            ".#....#..#.".to_string(),
            "..#.#...#.#".to_string(),
            ".#...##..#.".to_string(),
            "..#.##.....".to_string(),
            ".#.#.#....#".to_string(),
            ".#........#".to_string(),
            "#.##...#...".to_string(),
            "#...##....#".to_string(),
            ".#..#...#.#".to_string(),
        ];
        assert_eq!(num_trees(data, 1, 1), 2);
        assert_eq!(num_trees(data, 3, 1), 7);
        assert_eq!(num_trees(data, 5, 1), 3);
        assert_eq!(num_trees(data, 7, 1), 4);
        assert_eq!(num_trees(data, 1, 2), 2);
    }
}
