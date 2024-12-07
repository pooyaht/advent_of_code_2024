package day7_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day7"
)

func TestPartA(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	expected_ans := uint64(3749)
	ans := day7.SolvePartA(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	expected_ans := uint64(11387)
	ans := day7.SolvePartB(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
