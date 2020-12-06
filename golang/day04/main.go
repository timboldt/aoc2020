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
		"byr": regexp.MustCompile(`^19[2-9][0-9]|200[0-2]$`),
		"iyr": regexp.MustCompile(`^20(1[0-9]|20)$`),
		"eyr": regexp.MustCompile(`^20(2[0-9]|30)$`),
		"hgt": regexp.MustCompile(`^1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in$`),
		"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
		"ecl": regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`),
		"pid": regexp.MustCompile(`^[0-9]{9}$`),
		"cid": regexp.MustCompile(`.*`),
	}
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
		// fmt.Printf("%v\t%v:%v\n", rv.MatchString(p.kv[rk]), rk, p.kv[rk])
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
		if scanner.Text() == "" {
			if kv != nil {
				result = append(result, &passport{kv: kv})
				kv = nil
			}
			continue
		}
		if kv == nil {
			kv = make(map[string]string)
		}
		for _, f := range strings.Fields(scanner.Text()) {
			s := strings.Split(f, ":")
			if len(s) == 2 && validRegex[s[0]] != nil {
				if kv[s[0]] == "" {
					kv[s[0]] = s[1]
				} else {
					// Duplicate field
					kv[s[0]] = "****"
				}
			}
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
