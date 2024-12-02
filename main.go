package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sverdejot/advent-of-go/day1"
	"github.com/sverdejot/advent-of-go/day2"
)

func main() {
	// DAY 1
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	l, r, err := day1.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	d := day1.TotalDistance(l, r)
	fmt.Printf("Day 1: total distance between lists: %d\n", d)
	sc := day1.SimilarityScore(l, r)
	fmt.Printf("Day 1: similarity score between lists: %d\n", sc)

	// DAY 2
	f2, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	sf, err := day2.Parse(f2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 2: safe records %d\n", sf)
}
