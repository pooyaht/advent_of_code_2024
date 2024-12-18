package day15_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day15"
)

func TestSolvePartA(t *testing.T) {
	input := `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

	expectedAnswer := 2028
	ans := day15.SolvePartA(strings.Split(input, "\n\n"))
	if expectedAnswer != ans {
		t.Errorf("Expected %d got %d", expectedAnswer, ans)
	}
}

func TestSolvePartB(t *testing.T) {
	input := `#######
#...#.#
#.....#
#.#O@.#
#..O..#
#.....#
#######

<<<vv<<^^<<^^`

	expectedAnswer := 9021
	ans := day15.SolvePartB(strings.Split(input, "\n\n"))
	if expectedAnswer != ans {
		t.Errorf("Expected %d got %d", expectedAnswer, ans)
	}
}
