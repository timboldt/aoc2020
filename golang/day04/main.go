package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type passport struct {
	kv map[string]string
}

var (
	validRegex map[string]*regexp.Regexp
)

func init() {
	validRegex = map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`^(19[2-9][0-9])|(200[0-2])$`),
		"iyr": regexp.MustCompile(`^(201[0-9])|2020$`),
		"eyr": regexp.MustCompile(`^(202[0-9])|2030$`),
		"hgt": regexp.MustCompile(`^(((1[5-8][0-9])|(19[0-3]))cm)|((59|(6[0-9])|(7[0-6]))in)$`),
		"hcl": regexp.MustCompile(`^\#[0-9a-f]{6}$`),
		"ecl": regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`),
		"pid": regexp.MustCompile(`^[0-9]{9}$`),
	}
	// 	byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
}

func (p *passport) isValid1() bool {
	return p.kv["byr"] != "" &&
		p.kv["iyr"] != "" &&
		p.kv["eyr"] != "" &&
		p.kv["hgt"] != "" &&
		p.kv["hcl"] != "" &&
		p.kv["ecl"] != "" &&
		p.kv["pid"] != ""
}

func (p *passport) isValid2() bool {
	for rk, rv := range validRegex {
		//fmt.Printf("%v:%v\t\t%v\n", rk, p.kv[rk], rv.MatchString(p.kv[rk]))
		if !rv.MatchString(p.kv[rk]) {
			return false
		}
	}
	return true
}

func parseInput(r io.Reader) ([]*passport, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []*passport
	var kv map[string]string
	for scanner.Scan() {
		if scanner.Text() == "" && kv != nil {
			result = append(result, &passport{kv: kv})
			kv = nil
			continue
		}
		if kv == nil {
			kv = make(map[string]string)
		}
		for _, f := range strings.Fields(scanner.Text()) {
			s := strings.Split(f, ":")
			kv[s[0]] = s[1]
		}
	}
	if kv != nil {
		result = append(result, &passport{kv: kv})
		kv = nil
	}
	return result, scanner.Err()
}

func part1(data []*passport) int {
	var cnt int
	for _, p := range data {
		if p.isValid1() {
			cnt++
		}
	}
	return cnt
}

func part2(data []*passport) int {
	var cnt int
	for _, p := range data {
		if p.isValid2() {
			cnt++
		}
	}
	return cnt
}

func main() {
	// Assumes `go run aoc2020/day04` from the module-level directory.
	infile, err := os.Open("day04/input.txt")
	if err != nil {
		panic("Cannot find input file.")
	}
	data, err := parseInput(infile)
	if err != nil {
		panic("Invalid input file.")
	}
	result1 := part1(data)
	fmt.Printf("Day 4 Part 1: %d\n", result1)
	result2 := part2(data)
	fmt.Printf("Day 4 Part 2: %d\n", result2)
}
