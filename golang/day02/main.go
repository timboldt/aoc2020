package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type password struct {
	n1 int
	n2 int
	ch rune
	pw string
}

func (p *password) isValid1() bool {
	var cnt int
	for _, ch := range p.pw {
		if ch == p.ch {
			cnt++
		}
	}
	return cnt >= p.n1 && cnt <= p.n2
}

func (p *password) isValid2() bool {
	var cnt int
	for idx, ch := range p.pw {
		if (idx == p.n1-1 || idx == p.n2-1) && ch != p.ch {
			cnt++
		}
	}
	return cnt == 1
}

func parsePassword(s string) (*password, error) {
	isFieldChar := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fields := strings.FieldsFunc(s, isFieldChar)
	n1, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return nil, err
	}
	n2, err := strconv.ParseInt(fields[1], 10, 32)
	if err != nil {
		return nil, err
	}
	var ch rune
	for _, r := range fields[2] {
		ch = r
		break
	}
	return &password{
		n1: int(n1),
		n2: int(n2),
		ch: ch,
		pw: fields[3],
	}, nil
}

func readPasswords(r io.Reader) ([]*password, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []*password
	for scanner.Scan() {
		x, err := parsePassword(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func part1(data []*password) int {
	var cnt int
	for _, p := range data {
		if p.isValid1() {
			cnt++
		}
	}
	return cnt
}

func part2(data []*password) int {
	var cnt int
	for _, p := range data {
		if p.isValid2() {
			cnt++
		}
	}
	return cnt
}

func main() {
	// Assumes `go run aoc2020/day02` from the module-level directory.
	infile, err := os.Open("day02/input.txt")
	if err != nil {
		panic("Cannot find input file.")
	}
	data, err := readPasswords(infile)
	if err != nil {
		panic("Invalid input file.")
	}
	result1 := part1(data)
	fmt.Printf("Day 1 Part 1: %d\n", result1)
	result2 := part2(data)
	fmt.Printf("Day 1 Part 2: %d\n", result2)
}
