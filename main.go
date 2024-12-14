package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pooyaht/advent_of_code_2024/day1"
	"github.com/pooyaht/advent_of_code_2024/day10"
	"github.com/pooyaht/advent_of_code_2024/day11"
	"github.com/pooyaht/advent_of_code_2024/day12"
	"github.com/pooyaht/advent_of_code_2024/day13"
	"github.com/pooyaht/advent_of_code_2024/day2"
	"github.com/pooyaht/advent_of_code_2024/day3"
	"github.com/pooyaht/advent_of_code_2024/day4"
	"github.com/pooyaht/advent_of_code_2024/day5"
	"github.com/pooyaht/advent_of_code_2024/day6"
	"github.com/pooyaht/advent_of_code_2024/day7"
	"github.com/pooyaht/advent_of_code_2024/day8"
	"github.com/pooyaht/advent_of_code_2024/day9"
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
	case 2:
		if *part == "a" {
			result = day2.RunPartA()
		} else if *part == "b" {
			result = day2.RunPartB()
		}
	case 3:
		if *part == "a" {
			result = day3.RunPartA()
		}
		if *part == "b" {
			result = day3.RunPartB()
		}
	case 4:
		if *part == "a" {
			result = day4.RunPartA()
		}
		if *part == "b" {
			result = day4.RunPartB()
		}
	case 5:
		if *part == "a" {
			result = day5.RunPartA()
		}
		if *part == "b" {
			result = day5.RunPartB()
		}
	case 6:
		if *part == "a" {
			result = day6.RunPartA()
		}
		if *part == "b" {
			result = day6.RunPartB()
		}
	case 7:
		if *part == "a" {
			result = day7.RunPartA()
		}
		if *part == "b" {
			result = day7.RunPartB()
		}
	case 8:
		if *part == "a" {
			result = day8.RunPartA()
		}
		if *part == "b" {
			result = day8.RunPartB()
		}
	case 9:
		if *part == "a" {
			result = day9.RunPartA()
		}
		if *part == "b" {
			result = day9.RunPartB()
		}
	case 10:
		if *part == "a" {
			result = day10.RunPartA()
		}
		if *part == "b" {
			result = day10.RunPartB()
		}
	case 11:
		if *part == "a" {
			result = day11.RunPartA()
		}
		if *part == "b" {
			result = day11.RunPartB()
		}
	case 12:
		if *part == "a" {
			result = day12.RunPartA()
		}
		if *part == "b" {
			result = day12.RunPartB()
		}
	case 13:
		if *part == "a" {
			result = day13.RunPartA()
		}
		if *part == "b" {
			result = day13.RunPartB()
		}
	default:
		fmt.Printf("Day %d not implemented yet\n", *day)
		os.Exit(1)
	}

	fmt.Printf("Day %d Part %s Result: %v\n", *day, *part, result)
}
