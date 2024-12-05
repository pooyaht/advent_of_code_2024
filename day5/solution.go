package day5

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(input)
}

func SolvePartA(input string) int {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	sequences := parseSequences(parts[1])
	var result int

	for _, seq := range sequences {
		isValid := true
	mainLoop:
		for i, before := range seq {
			for _, after := range seq[i+1:] {
				if rules[after][before] {
					isValid = false
					break mainLoop
				}
			}
		}
		if isValid {
			result += seq[len(seq)/2]
		}
	}
	return result
}

func RunPartB() int {
	return SolvePartB(input)
}

func SolvePartB(input string) int {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	sequences := parseSequences(parts[1])
	var result int

	for _, seq := range sequences {
		isValid := true
	checkLoop:
		for i, before := range seq {
			for _, after := range seq[i+1:] {
				if rules[after][before] {
					isValid = false
					break checkLoop
				}
			}
		}

		if isValid {
			continue
		}

		n := len(seq)
		for i := 0; i < n-1; i++ {
			swapped := false

			for j := 0; j < n-i-1; j++ {
				if rules[seq[j]][seq[j+1]] {
					seq[j], seq[j+1] = seq[j+1], seq[j]
					swapped = true
				}
			}

			if !swapped {
				break
			}
		}
		result += seq[len(seq)/2]
	}

	return result
}

func parseRules(data string) map[int]map[int]bool {
	rules := make(map[int]map[int]bool)
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])

		if rules[before] == nil {
			rules[before] = make(map[int]bool)
		}

		rules[before][after] = true
	}

	return rules
}

func parseSequences(data string) [][]int {
	var sequences [][]int
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		numStrs := strings.Split(line, ",")
		sequence := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			sequence[i] = num
		}
		sequences = append(sequences, sequence)
	}

	return sequences
}
