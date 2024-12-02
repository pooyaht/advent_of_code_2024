package day2_test

import (
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day2"
)

func TestPartA(t *testing.T) {
	input := [][]uint32{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	expected_ans := uint32(2)
	ans := day2.SolvePartA(input)
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := [][]uint32{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	expected_ans := uint32(4)
	ans := day2.SolvePartB(input)
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
