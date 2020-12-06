package main

import (
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	input := strings.NewReader(`
		123
		456`)
	ints, err := ReadInts(input)
	if err != nil {
		t.Error(err)
	}
	if len(ints) != 2 || ints[0] != 123 || ints[1] != 456 {
		t.Errorf("want: [123,456], got: %v", ints)
	}
}
