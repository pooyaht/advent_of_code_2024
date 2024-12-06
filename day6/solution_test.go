package day6_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day6"
)

func TestPartA(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected_ans := 41
	ans, _ := day6.SolvePartA(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected_ans := 6
	ans := day6.SolvePartB(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
