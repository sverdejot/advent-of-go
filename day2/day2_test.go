package day2

import (
	"slices"
	"testing"
)

func Test_parseLine(t *testing.T) {
	tests := []struct{
		Name	string
		Input	string
		Want	[]int
		Error	bool
	}{
		{
			Name: "empty line",
			Input: "",
		},
		{
			Name: "single digit",
			Input: "9",
			Want: []int{9},
		},
		{
			Name: "invalid character",
			Input: "d",
			Error: true,
		},
		{
			Name: "valid report",
			Input: "1 9 3 2",
			Want: []int{1, 9, 3, 2},
		},
		{
			Name: "invalid multi-character",
			Input: "1 3 d 9 2",
			Error: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := parseLine(tt.Input)

			if err != nil {
				if tt.Error {
					return
				}
				t.Fatalf("oops: %v", err)
			}

			// leaking a impl. detail: returning a nil slice
			// handle with care
			if !slices.Equal(got, tt.Want) {
				t.Fatalf("slices %v, %v are not equal", got, tt.Want)
			}
		})
	}
}

func Test_isSafe(t *testing.T) {
	tests := []struct{
		Name	string
		Input	[]int
		Want	bool
	}{
		{
			Name: "empty report",
			Want: true,
		},
		{
			Name: "single elem",
			Input: []int{1},
			Want: true,
		},
		{
			Name: "all increasing",
			Input: []int{1, 3, 6},
			Want: true,
		},
		{
			Name: "safe",
			Input: []int{7, 6, 4, 2, 1},
			Want: true,
		},
		{
			Name: "unsafe",
			Input: []int{1, 2, 7, 8, 9},
			Want: false,
		},
		{
			Name: "unsafe",
			Input: []int{9, 7, 6, 2, 1},
			Want: false,
		},
		{
			Name: "unsafe",
			Input: []int{1, 3, 2, 4, 5},
			Want: false,
		},
		{
			Name: "unsafe",
			Input: []int{8, 6, 4, 4, 1},
			Want: false,
		},
		{
			Name: "safe",
			Input: []int{1, 3, 6, 7, 9},
			Want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := isSafe(tt.Input)

			if got != tt.Want {
				t.Fatalf("got %v want %v for %v", got, tt.Want, tt.Input)
			}
		})
	}
}
