package main

import (
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	input := strings.NewReader("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc")
	got, err := readPasswords(input)
	if err != nil {
		t.Error(err)
	}
	want := []*password{
		&password{n1: 1, n2: 3, ch: 'a', pw: "abcde"},
		&password{n1: 1, n2: 3, ch: 'b', pw: "cdefg"},
		&password{n1: 2, n2: 9, ch: 'c', pw: "ccccccccc"},
	}
	for i, v := range got {
		if *v != *want[i] {
			t.Errorf("got: %v, want: %v", *v, *want[i])
		}
	}
	if len(got) != len(want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	data := []*password{
		&password{n1: 1, n2: 3, ch: 'a', pw: "abcde"},
		&password{n1: 1, n2: 3, ch: 'b', pw: "cdefg"},
		&password{n1: 2, n2: 9, ch: 'c', pw: "ccccccccc"},
	}
	if !data[0].isValid1() {
		t.Errorf("expected first password to be valid")
	}
	if data[1].isValid1() {
		t.Errorf("expected second password to be invalid")
	}
	if !data[2].isValid1() {
		t.Errorf("expected third password to be valid")
	}
	want := 2
	if got := part1(data); got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := []*password{
		&password{n1: 1, n2: 3, ch: 'a', pw: "abcde"},
		&password{n1: 1, n2: 3, ch: 'b', pw: "cdefg"},
		&password{n1: 2, n2: 9, ch: 'c', pw: "ccccccccc"},
	}
	if !data[0].isValid2() {
		t.Errorf("expected first password to be valid")
	}
	if data[1].isValid2() {
		t.Errorf("expected second password to be invalid")
	}
	if data[2].isValid2() {
		t.Errorf("expected third password to be invalid")
	}
	want := 1
	if got := part2(data); got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
