use aoc_runner_derive::{aoc, aoc_generator};
use regex::Regex;
use std::collections::HashMap;
use std::collections::HashSet;
use std::convert::Infallible;

fn is_valid_passport(raw: &str, validate_values: bool) -> bool {
    lazy_static! {
        static ref RE_FIELD_SEPERATOR: Regex = Regex::new(r"(\s+)").unwrap();
        static ref RE_VALID: HashMap<&'static str, Regex> = {
            let mut map = HashMap::new();
            map.insert("byr", Regex::new(r"^(19[2-9][0-9]|200[0-2])$").unwrap());
            map.insert("iyr", Regex::new(r"^20(1[0-9]|20)$").unwrap());
            map.insert("eyr", Regex::new(r"^20(2[0-9]|30)$").unwrap());
            map.insert("hgt", Regex::new(r"^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$").unwrap());
            map.insert("hcl", Regex::new(r"^#[0-9a-f]{6}$").unwrap());
            map.insert("ecl", Regex::new(r"^(amb|blu|brn|gry|grn|hzl|oth)$").unwrap());
            map.insert("pid", Regex::new(r"^[0-9]{9}$").unwrap());
            map.insert("cid", Regex::new(r".*").unwrap());
            map
        };
    }
    let expected: HashSet<&str> = vec!["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
        .into_iter()
        .collect();
    let mut set = HashSet::new();
    for kv in RE_FIELD_SEPERATOR.split(raw) {
        let parts: Vec<&str> = kv.split(":").collect();
        if validate_values {
            if !RE_VALID.get(parts[0]).unwrap().is_match(parts[1]) {
                return false;
            }
        }
        set.insert(parts[0]);
    }
    set.is_superset(&expected)
}

#[aoc_generator(day4)]
fn parse_input(input: &str) -> Result<Vec<String>, Infallible> {
    input.split("\n\n").map(|l| l.parse()).collect()
}

#[aoc(day4, part1)]
fn part1(passports: &[String]) -> i32 {
    let mut good = 0;
    for passport in passports {
        if is_valid_passport(passport, false) {
            good = good + 1;
        }
    }
    good
}

#[aoc(day4, part2)]
fn part2(passports: &[String]) -> i32 {
    let mut good = 0;
    for passport in passports {
        if is_valid_passport(passport, true) {
            good = good + 1;
        }
    }
    good
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn valid_passport_test() {
        assert_eq!(
            is_valid_passport(
                "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
        byr:1937 iyr:2017 cid:147 hgt:183cm"
            ),
            true
        );
        assert_eq!(
            is_valid_passport(
                "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
                hcl:#cfa07d byr:1929"
            ),
            false
        );
        assert_eq!(
            is_valid_passport(
                "hcl:#ae17e1 iyr:2013
                eyr:2024
                ecl:brn pid:760753108 byr:1931
                hgt:179cm"
            ),
            true
        );
        assert_eq!(
            is_valid_passport(
                "hcl:#cfa07d eyr:2025 pid:166559648
                iyr:2011 ecl:brn hgt:59in"
            ),
            false
        );
    }
}
