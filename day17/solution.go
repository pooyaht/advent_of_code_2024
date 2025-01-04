package day17

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() string {
	return SolvePartA(strings.Split(input, "\n\n"))
}

func SolvePartA(input []string) string {
	registers, program := parseInput(input)

	output := make([]int, 0)

	pc := 0
	for pc < len(program) {
		instruction := program[pc]
		operand := program[pc+1]
		instructions := newInstructions()

		if instruction == 3 {
			next := instructions[instruction](registers, operand)
			if next != nil {
				pc = *next
				continue
			}
		} else if instruction == 5 {
			output = append(output, *instructions[instruction](registers, operand))
		} else {
			instructions[instruction](registers, operand)
		}

		pc += 2
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output)), ","), "[]")
}

func RunPartB() string {
	return SolvePartB(strings.Split(input, "\n\n"))
}

func SolvePartB(input []string) string {
	return ""
}

func parseInput(input []string) (map[string]int, []int) {
	registers := make(map[string]int)
	var program []int

	for _, line := range strings.Split(input[0], "\n") {
		parts := strings.Split(line, " ")
		registers[parts[1][:1]], _ = strconv.Atoi(parts[2])
	}

	for _, line := range strings.Split(strings.Split(input[1], "Program: ")[1], ",") {
		num, _ := strconv.Atoi(line)
		program = append(program, num)
	}

	return registers, program
}

func newInstructions() map[int]func(map[string]int, int) *int {
	handleComboOperator := func(registers map[string]int, operand int) int {
		if operand == 4 {
			operand = registers["A"]
		} else if operand == 5 {
			operand = registers["B"]
		} else if operand == 6 {
			operand = registers["C"]
		} else if operand == 7 {
			panic("operand 7")
		}

		return operand
	}

	instructions := map[int]func(map[string]int, int) *int{
		0: func(registers map[string]int, operand int) *int {
			//adv
			operand = handleComboOperator(registers, operand)
			numerator, ok := registers["A"]
			if !ok {
				return nil
			}

			denominator := math.Pow(2.0, float64(operand))
			registers["A"] = int(float64(numerator) / denominator)
			return nil
		},
		1: func(registers map[string]int, operand int) *int {
			//bxl
			registers["B"] = registers["B"] ^ operand
			return nil
		},
		2: func(registers map[string]int, operand int) *int {
			//bst
			operand = handleComboOperator(registers, operand)
			registers["B"] = operand % 8
			return nil
		},
		3: func(registers map[string]int, operand int) *int {
			//jnz
			if registers["A"] == 0 {
				return nil
			}
			return &operand
		},
		4: func(registers map[string]int, operand int) *int {
			//bxc
			registers["B"] = registers["B"] ^ registers["C"]
			return nil
		},
		5: func(registers map[string]int, operand int) *int {
			//out
			operand = handleComboOperator(registers, operand)
			num := operand % 8
			return &num
		},
		6: func(registers map[string]int, operand int) *int {
			//bdv
			operand = handleComboOperator(registers, operand)
			numerator, ok := registers["A"]
			if !ok {
				return nil
			}

			denominator := math.Pow(2.0, float64(operand))
			registers["B"] = int(float64(numerator) / denominator)
			return nil
		},
		7: func(registers map[string]int, operand int) *int {
			//cdv
			operand = handleComboOperator(registers, operand)
			numerator, ok := registers["A"]
			if !ok {
				return nil
			}

			denominator := math.Pow(2.0, float64(operand))
			registers["C"] = int(float64(numerator) / denominator)
			return nil
		},
	}

	return instructions
}
