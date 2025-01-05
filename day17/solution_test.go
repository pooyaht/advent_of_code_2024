package day17_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day17"
)

func TestSolvePartA(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	expected := "4,2,5,6,7,7,7,7,3,1,0"
	result := day17.SolvePartA(strings.Split(input, "\n\n"))
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSolvePartB(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

	expected := "117440"
	result := day17.SolvePartB(strings.Split(input, "\n\n"))
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
