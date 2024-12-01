package day1_test

import (
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day1"
)

func TestPartA(t *testing.T) {
	input := [][]uint32{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	expected_ans := uint32(11)
	ans := day1.SolvePartA(input[0], input[1])
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := [][]uint32{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	expected_ans := uint32(31)
	ans := day1.SolvePartB(input[0], input[1])
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
