package day1

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func parseLine(line string) (int, int, error) {
	parts := strings.Split(line, "   ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("not enough elems at line: %s", line)
	}
	l, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse num %s", parts[0])
	}
	r, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse num %s", parts[1])
	}
	return l, r, nil
}

func Parse(r io.Reader) ([]int, []int, error) {
	sc := bufio.NewScanner(r)

	// can't guess the initial size
	left := make([]int, 0)
	right := make([]int, 0)

	var lf, rg int
	var err error
	for sc.Scan() {
		rawLine := sc.Text()

		lf, rg, err = parseLine(rawLine)
		if err != nil {
			return nil, nil, err
		}

		left = append(left, lf)
		right = append(right, rg)
	}
	return left, right, err
}

func calculateDistance(l, r []int) int {
	var d int
	if len(l) != len(r) {
		return 0
	}
	for i := range len(l) {
		d += abs(l[i] - r[i])
	}
	return d
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func seqSortListInPlace(l, r []int) {
	slices.Sort(l)
	slices.Sort(r)
}

func parSortListInPlace(l, r []int) {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		slices.Sort(l)
	}()

	go func() {
		defer wg.Done()
		slices.Sort(r)
	}()

	wg.Wait()
}

func SimilarityScore(l, r []int) int {
	lmap := make(map[int]struct{}, len(l))
	for _, v := range l {
		lmap[v] = struct{}{}
	}
	rmap := make(map[int]int, len(l))
	for _, v := range r {
		if _, ok := lmap[v]; ok {
			rmap[v] += v
		}
	}
	var d int
	for _, v := range l {
		d += rmap[v]
	}
	return d
}

func TotalDistance(l, r []int) int {
	parSortListInPlace(l, r)
	
	return calculateDistance(l, r)
}
