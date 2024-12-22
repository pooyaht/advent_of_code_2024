package day16_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day16"
)

func TestSolvePartA(t *testing.T) {
	input := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

	expectedAnswer := 7035
	ans := day16.SolvePartA(strings.Split(input, "\n"))
	if expectedAnswer != ans {
		t.Errorf("Expected %d got %d", expectedAnswer, ans)
	}
}

func TestSolvePartB(t *testing.T) {
	input := `###############
#.......#.....#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#..E..#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

	expectedAnswer := 5
	ans := day16.SolvePartB(strings.Split(input, "\n"))
	if expectedAnswer != ans {
		t.Errorf("Expected %d got %d", expectedAnswer, ans)
	}
}
