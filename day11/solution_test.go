package day11_test

import (
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day11"
)

func TestPartA(t *testing.T) {
	input := `125 17`

	expected_ans := 55312
	ans := day11.SolvePartA(input, 25)
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
