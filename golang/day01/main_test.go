package main

import (
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	input := strings.NewReader(`
		123
		456`)
	got, err := ReadInts(input)
	if err != nil {
		t.Error(err)
	}
	want := []int{123, 456}
	if len(got) != 2 || got[0] != 123 || got[1] != 456 {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	data := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	want := 514579
	if got := part1(data); got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	want := 241861950
	if got := part2(data); got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
