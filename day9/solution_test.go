package day9_test

import (
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day9"
)

func TestPartA(t *testing.T) {
	input := `2333133121414131402`

	expected_ans := 1928
	ans := day9.SolvePartA(input)
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := `2333133121414131402`

	expected_ans := 2858
	ans := day9.SolvePartB(input)
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
