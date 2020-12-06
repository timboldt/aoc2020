package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type point struct {
	row int
	col int
}

type terrain struct {
	rows  int
	cols  int
	trees map[point]bool
}

func parseTerrain(r io.Reader) (terrain, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := terrain{trees: make(map[point]bool)}
	var row int
	for scanner.Scan() {
		if result.cols == 0 {
			result.cols = len(scanner.Text())
		}
		for col, ch := range scanner.Text() {
			if ch == '#' {
				result.trees[point{row: row, col: col}] = true
			}
		}
		row++
	}
	result.rows = row
	return result, scanner.Err()
}

func countTrees(data terrain, rowStep int, colStep int) int {
	var trees int
	var col int
	for row := 0; row < data.rows; row = row + rowStep {
		if data.trees[point{row: row, col: col}] {
			trees++
		}
		col = (col + colStep) % data.cols
	}
	return trees
}

func part1(data terrain) int {
	return countTrees(data, 1, 3)
}

func part2(data terrain) int {
	return countTrees(data, 1, 1) *
		countTrees(data, 1, 3) *
		countTrees(data, 1, 5) *
		countTrees(data, 1, 7) *
		countTrees(data, 2, 1)
}

func main() {
	// Assumes `go run aoc2020/day03` from the module-level directory.
	infile, err := os.Open("day03/input.txt")
	if err != nil {
		panic("Cannot find input file.")
	}
	data, err := parseTerrain(infile)
	if err != nil {
		panic("Invalid input file.")
	}
	result1 := part1(data)
	fmt.Printf("Day 3 Part 1: %d\n", result1)
	result2 := part2(data)
	fmt.Printf("Day 3 Part 2: %d\n", result2)
}
