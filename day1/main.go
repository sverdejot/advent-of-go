package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Trebuchet2(f))
}

func Trebuchet2(r io.Reader) int {
	sc := bufio.NewScanner(r)

	var res int
	for sc.Scan() {
		line := sc.Text()

		n := readDigits(line)
		res += n
	}

	return res
}

var digits map[string]int = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
	"zero": 0,
}
func readDigits(line string) int {
	var first, last int
	for len(line) > 0 {
		for pattern, digit := range digits {
			if strings.HasPrefix(line, pattern) {
				if first == 0 {
					first = digit
				}
				last = digit
				continue
			}
			if num, err := strconv.Atoi(string(line[0])); err == nil {
				if first == 0 {
					first = num 
				}
				last = num
			}
		}
		line = line[1:]
	}
	return (first*10)+last
}

func Trebuchet1(r io.Reader) int {
	sc := bufio.NewScanner(r)

	var res int
	for sc.Scan() {
		line := sc.Text()

		n := getDigits(line)
		res += n
	}

	return res
}

func getDigits(code string) int {
	var first, last int
	for _, ch := range code {
		if n, err := strconv.Atoi(string(ch)); err == nil {
			if first == 0 {
				first = n
			}
			last = n
		}
	}
	return (first * 10) + last
}
