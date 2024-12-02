package day1

import (
	"math/rand"
	"strings"
	"testing"
)

const maxElems = 1000

func BenchmarkSortListInPlace(t *testing.B) {
	t.Run("sequential", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l, r := rand.Perm(maxElems), rand.Perm(maxElems)
			seqSortListInPlace(l, r)
		}
	})
	t.Run("parallel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l, r := rand.Perm(maxElems), rand.Perm(maxElems)
			parSortListInPlace(l, r)
		}
	})
}

func Test_parseLine(t *testing.T) {
	tests := []struct{
		Name	string
		Input	string
		L, R int
		Error bool
	}{
		{
			Name: "empty line", Input: "", Error: true,
		},
		{
			Name: "valid line", Input: "1   3", L: 1, R: 3,
		},
		{
			Name: "contains letter", Input: "1   d", Error: true,
		},
		{
			Name: "not enough sep", Input: "1  2", Error: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			l, r, err := parseLine(tt.Input)

			if !tt.Error && err != nil{
				t.Fatalf("found: %v", err)
			} else {
				return
			}
			
			if l != tt.L || r != tt.R {
				t.Fatalf("got %d, %d but want %d, %d", l, r, tt.L, tt.R)
			}
		})
	}
}

func TestSimilarityScore(t *testing.T) {
	t.Run("should return 31", func(t *testing.T) {
		const want = 31

		l := []int{3,4,2,1,3,3}
		r := []int{4,3,5,3,9,3}

		d := SimilarityScore(l, r)
		if d != want {
			t.Fatalf("got %d want %d", want, d)
		}
	})
}

func TestParseInput(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		given := strings.NewReader("")

		right, left, err := Parse(given)

		assertNoError(t, err)
		assertLength(t, 0, right, left)
	})

	t.Run("single line", func(t *testing.T) {
		given := strings.NewReader(`1   2`)

		right, left, err := Parse(given)

		assertNoError(t, err)
		assertLength(t, 1, right, left)
		assertElems(t, right, 1)
		assertElems(t, left, 2)
	})

	t.Run("invalid single-line", func(t *testing.T) {
		given := strings.NewReader(`1   d`)

		_, _, err := Parse(given)

		assertError(t, err)
	})

	t.Run("multi-line", func(t *testing.T) {
		given := strings.NewReader(`1   2
3   4
7   8
124   8963`)

		right, left, err := Parse(given)

		assertNoError(t, err)
		assertLength(t, 4, right, left)
		assertElems(t, right, 1, 3, 7, 124)
		assertElems(t, left, 2, 4, 8, 8963)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got %+v want nil", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("got nil want %+v", err)
	}
}

func assertLength(t *testing.T, length int, lists ...[]int) {
	t.Helper()
	for _, l := range lists {
		if len(l) != length {
			t.Fatalf("list %v has not the expected length: got %d want %d", l, len(l), length)
		}
	}
}

func assertElems(t *testing.T, l []int, elems ...int) {
	t.Helper()
	for idx, v := range elems {
		if l[idx] != v {
			t.Fatalf("not same elems at idx %d for list %v: got %d want %d", idx, l, v, l[idx])
		}
	}
}
