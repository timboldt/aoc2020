package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
// https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func part1(data []int) int {
	// This algorithm is O(N^2/2), and there are ways to solve this in O(N), but this solution is easier to read.
	for idx1, i := range data {
		for _, j := range data[idx1+1:] {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	panic("Oops!")
}

func part2(data []int) int {
	for idx1, i := range data {
		for idx2, j := range data[idx1+1:] {
			for _, k := range data[idx2+1:] {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	panic("Oops!")
}

func main() {
	// Assumes `go run aoc2020/day01` from the module-level directory.
	infile, err := os.Open("day01/input.txt")
	if err != nil {
		panic("Cannot find input file.")
	}
	data, err := ReadInts(infile)
	if err != nil {
		panic("Invalid input file.")
	}
	result1 := part1(data)
	fmt.Printf("Day 1 Part 1: %d\n", result1)
	result2 := part2(data)
	fmt.Printf("Day 1 Part 2: %d\n", result2)
}
