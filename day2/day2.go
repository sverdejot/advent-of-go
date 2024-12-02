package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Parse(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)

	var safe int
	for sc.Scan() {
		report, err := parseLine(sc.Text())
		if err != nil {
			return 0, err
		}

		if isSafe(report) {
			safe++
		}
	}
	return safe, nil
}

func parseLine(l string) ([]int, error) {
	n := utf8.RuneCountInString(l)
	if n == 0 {
		return nil, nil
	}

	res := make([]int, 0, n)
	var v int
	var err error
	for _, elem := range strings.Split(l, " ") {
		v, err = strconv.Atoi(elem)
		if err != nil {
			return res, err
		}
		res = append(res, v)
	}
	return res, nil
}

func isSafeWithTolerance(report []int, tolerance int) bool {
	// guess this should be the rule
	if len(report) < 2 {
		return true
	}

	// 'set' only marks whether 'inc' is the default value or not
	// as soon as the first comparison is done we can trust the value set at 'inc' 
	// which btw means 'increasing'
	var inc, set bool
	last := report[0]
	for i := 1; i < len(report); i++ {
		if last < report[i] {
			if set && inc {
				// was increasing but now is lower
				return false
			}
			inc = false
		} else {
			// was decreasing but now is higher
			if set && !inc {
				return false
			}
			inc = true
		}
		if last == report[i] {
			return false
		}
		set = true
		if t := abs(last - report[i]); t < 1 || t > 3 {
			return false
		}
		
		last = report[i]
	}
	return true
}


func isSafe(report []int) bool {
	return isSafeWithTolerance(report, 0)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
