package day19

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(input)
}

func SolvePartA(input string) int {
	parts := strings.Split(input, "\n\n")
	patterns := strings.Split(strings.TrimSpace(parts[0]), ", ")
	designs := strings.Split(strings.TrimSpace(parts[1]), "\n")

	var total int
	memo := make(map[string]int)
	for _, design := range designs {
		ways := countWaysToMakeDesign(design, patterns, memo)
		if ways > 0 {
			total++
		}
	}
	return total
}

func RunPartB() int {
	return SolvePartB(input)
}

func SolvePartB(input string) int {
	parts := strings.Split(input, "\n\n")
	patterns := strings.Split(strings.TrimSpace(parts[0]), ", ")
	designs := strings.Split(strings.TrimSpace(parts[1]), "\n")

	var total int
	for _, design := range designs {
		ways := countWaysToMakeDesign(design, patterns, make(map[string]int))
		total += ways
	}
	return total
}

func countWaysToMakeDesign(design string, patterns []string, memo map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	if count, exists := memo[design]; exists {
		return count
	}

	total := 0
	for _, pattern := range patterns {
		if len(pattern) <= len(design) && design[:len(pattern)] == pattern {
			total += countWaysToMakeDesign(design[len(pattern):], patterns, memo)
		}
	}

	memo[design] = total
	return total
}
