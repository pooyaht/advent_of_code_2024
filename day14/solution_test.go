package day14_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day14"
)

func TestSolvePartA(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	expectedAnswer := 12
	ans := day14.SolvePartA(strings.Split(input, "\n"), 100, 11, 7)
	if expectedAnswer != ans {
		t.Errorf("Expected %d got %d", expectedAnswer, ans)
	}
}
