package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pooyaht/advent_of_code_2024/day1"
)

func main() {
	day := flag.Int("day", 1, "Day number (1-25)")
	part := flag.String("part", "a", "Part (a/b)")
	flag.Parse()

	if *day < 1 || *day > 25 {
		log.Fatal("Day must be between 1 and 25")
	}

	var result any
	switch *day {
	case 1:
		if *part == "a" {
			result = day1.RunPartA()
		} else if *part == "b" {
			result = day1.RunPartB()
		}
	default:
		fmt.Printf("Day %d not implemented yet\n", *day)
		os.Exit(1)
	}

	fmt.Printf("Day %d Part %s Result: %v\n", *day, *part, result)
}
