package day1

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() uint32 {
	listA, listB := parseInput(input)
	return SolvePartA(listA, listB)
}

func SolvePartA(listA, listB []uint32) uint32 {
	slices.Sort(listA)
	slices.Sort(listB)
	var result uint32

	for i, num := range listA {
		if num >= listB[i] {
			result = result + num - listB[i]
		} else {
			result = result + listB[i] - num
		}
	}
	return uint32(result)
}

func RunPartB() uint32 {
	listA, listB := parseInput(input)
	return SolvePartB(listA, listB)
}

func SolvePartB(listA, listB []uint32) uint32 {
	var result uint32
	counter := make(map[uint32]uint32, len(listB))

	for _, num := range listB {
		counter[num]++
	}

	for _, num := range listA {
		result += counter[num] * num
	}

	return result
}

func parseInput(input string) (listA, listB []uint32) {
	input = strings.TrimSpace(input)
	n := len(strings.Split(input, "\n"))
	listA = make([]uint32, n)
	listB = make([]uint32, n)

	lines := strings.Split(input, "\n")
	for _, row := range lines {
		fields := strings.Fields(row)
		num1, _ := strconv.ParseInt(fields[0], 10, 32)
		num2, _ := strconv.ParseInt(fields[1], 10, 32)

		listA = append(listA, uint32(num1))
		listB = append(listB, uint32(num2))
	}

	return listA, listB
}
