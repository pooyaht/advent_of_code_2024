package day18_test

import (
	"strings"
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day18"
)

func TestSolvePartA(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1`

	expected := 22
	result := day18.SolvePartA(strings.Split(input, "\n"), 7)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSolvePartB(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`

	expected := "6,1"
	result := day18.SolvePartB(strings.Split(input, "\n"), 7)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
