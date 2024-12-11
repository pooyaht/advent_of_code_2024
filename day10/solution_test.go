package day10_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day10"
)

func TestPartA(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	expected_ans := 36
	ans := day10.SolvePartA(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	expected_ans := 81
	ans := day10.SolvePartB(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
