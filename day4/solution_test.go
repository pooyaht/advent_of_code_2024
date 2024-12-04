package day4_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day4"
)

func TestPartA(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expected_ans := 18
	ans := day4.SolvePartA(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}

func TestPartB(t *testing.T) {
	input := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

	expected_ans := 9
	ans := day4.SolvePartB(strings.Split(input, "\n"))
	if ans != expected_ans {
		t.Errorf("Expected %d got %d", expected_ans, ans)
	}
}
