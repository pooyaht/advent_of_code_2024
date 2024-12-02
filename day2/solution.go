package day2

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() uint32 {
	reports := parseInput(input)
	return SolvePartA(reports)
}

func SolvePartA(reports [][]uint32) uint32 {
	var result uint32
	for _, report := range reports {
		allIncreasing := true
		allDecreasing := true
		adjacentCondition := true
		for j := 0; j < len(report)-1; j++ {
			if report[j] > report[j+1] {
				allIncreasing = false
			}
			if report[j] < report[j+1] {
				allDecreasing = false
			}
			if !allIncreasing && !allDecreasing {
				break
			}

			var diff uint32
			if report[j] > report[j+1] {
				diff = report[j] - report[j+1]
			} else {
				diff = report[j+1] - report[j]
			}

			if diff < 1 || diff > 3 {
				adjacentCondition = false
				break
			}
		}

		if (allIncreasing || allDecreasing) && adjacentCondition {
			result++
		}
	}
	return result
}

func RunPartB() uint32 {
	reports := parseInput(input)
	return SolvePartB(reports)
}

func SolvePartB(reports [][]uint32) uint32 {
	var result uint32

	for _, report := range reports {
		for j := 0; j < len(report); j++ {
			newReport := make([]uint32, 0, len(report)-1)
			newReport = append(newReport, report[:j]...)
			newReport = append(newReport, report[j+1:]...)
			resultAfterRemoval := SolvePartA([][]uint32{newReport})

			if resultAfterRemoval == 1 {
				result += 1
				break
			}
		}
	}

	return result
}

func parseInput(input string) [][]uint32 {
	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")
	var result [][]uint32

	for _, line := range lines {
		nums := strings.Fields(line)
		nums_int := make([]uint32, len(nums))
		for i, num := range nums {
			num_int, _ := strconv.ParseInt(num, 10, 32)
			nums_int[i] = uint32(num_int)
		}
		result = append(result, nums_int)
	}

	return result
}
