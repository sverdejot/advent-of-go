package main

import (
	"strings"
	"testing"
)

func TestTrebuchet1(t *testing.T) {
	r := strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	got := Trebuchet1(r)
	want := 142

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestTrebuchet2(t *testing.T) {
	r := strings.NewReader(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`)

	got := Trebuchet2(r)
	want := 281

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
