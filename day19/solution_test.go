package day19_test

import (
	"testing"

	"github.com/pooyaht/advent_of_code_2024/day19"
)

func TestSolvePartA(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	expected := 6
	got := day19.SolvePartA(input)
	if expected != got {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestSolvePartB(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	expected := 16
	got := day19.SolvePartB(input)
	if expected != got {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
