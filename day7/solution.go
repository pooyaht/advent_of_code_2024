package day7

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() uint64 {
	return SolvePartA(strings.Split(input, "\n"))
}

func SolvePartA(input []string) uint64 {
	var result uint64

	for _, row := range input {
		parts := strings.Split(row, ":")
		target, _ := strconv.ParseUint(parts[0], 10, 64)

		parts[1] = strings.TrimSpace(parts[1])
		numberStrings := strings.Split(parts[1], " ")
		numbers := make([]uint64, 0, len(numberStrings))

		for _, numStr := range numberStrings {
			num, _ := strconv.ParseUint(numStr, 10, 64)
			numbers = append(numbers, num)
		}

		matchedTarget := possibleMatch(target, numbers[1:], []string{"*", "+"}, numbers[0])
		if matchedTarget != nil {
			result += *matchedTarget
		}
	}
	return result
}

func RunPartB() uint64 {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(input []string) uint64 {
	var result uint64

	for _, row := range input {
		parts := strings.Split(row, ":")
		target, _ := strconv.ParseUint(parts[0], 10, 64)

		parts[1] = strings.TrimSpace(parts[1])
		numberStrings := strings.Split(parts[1], " ")
		numbers := make([]uint64, 0, len(numberStrings))

		for _, numStr := range numberStrings {
			num, _ := strconv.ParseUint(numStr, 10, 64)
			numbers = append(numbers, num)
		}

		matchedTarget := possibleMatch(target, numbers[1:], []string{"*", "+", "||"}, numbers[0])
		if matchedTarget != nil {
			result += *matchedTarget
		}
	}
	return result
}

func possibleMatch(target uint64, numbers []uint64, operators []string, currentValue uint64) *uint64 {
	if target == currentValue && len(numbers) == 0 {
		return &target
	}
	if currentValue > target || len(numbers) == 0 {
		return nil
	}

	var result *uint64
	for _, operator := range operators {
		switch operator {
		case "+":
			result = possibleMatch(target, numbers[1:], operators, currentValue+numbers[0])
		case "*":
			result = possibleMatch(target, numbers[1:], operators, currentValue*numbers[0])
		case "||":
			currentValue, _ = strconv.ParseUint(fmt.Sprintf("%d%d", currentValue, numbers[0]), 10, 64)
			result = possibleMatch(target, numbers[1:], operators, currentValue)
		}
		if result != nil {
			return result
		}
	}
	return nil
}
