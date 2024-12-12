package day12_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day12"
)

func TestPartA(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	expected_ans := 1930
	ans := day12.SolvePartA(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

	expected_ans := 368
	ans := day12.SolvePartB(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
