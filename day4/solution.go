package day4

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n"))
}

func SolvePartA(input []string) int {
	input = pad(input, 4, ".")
	var result int
	for rowIndex, row := range input {
		for colIndex, val := range row {
			if val != 'X' {
				continue
			}

			if input[rowIndex+1][colIndex] == 'M' && input[rowIndex+2][colIndex] == 'A' && input[rowIndex+3][colIndex] == 'S' {
				result++
			}
			if input[rowIndex-1][colIndex] == 'M' && input[rowIndex-2][colIndex] == 'A' && input[rowIndex-3][colIndex] == 'S' {
				result++
			}

			if input[rowIndex][colIndex+1] == 'M' && input[rowIndex][colIndex+2] == 'A' && input[rowIndex][colIndex+3] == 'S' {
				result++
			}
			if input[rowIndex][colIndex-1] == 'M' && input[rowIndex][colIndex-2] == 'A' && input[rowIndex][colIndex-3] == 'S' {
				result++
			}

			if input[rowIndex+1][colIndex+1] == 'M' && input[rowIndex+2][colIndex+2] == 'A' && input[rowIndex+3][colIndex+3] == 'S' {
				result++
			}
			if input[rowIndex-1][colIndex-1] == 'M' && input[rowIndex-2][colIndex-2] == 'A' && input[rowIndex-3][colIndex-3] == 'S' {
				result++
			}

			if input[rowIndex+1][colIndex-1] == 'M' && input[rowIndex+2][colIndex-2] == 'A' && input[rowIndex+3][colIndex-3] == 'S' {
				result++
			}
			if input[rowIndex-1][colIndex+1] == 'M' && input[rowIndex-2][colIndex+2] == 'A' && input[rowIndex-3][colIndex+3] == 'S' {
				result++
			}
		}
	}

	return result
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(input []string) int {
	input = pad(input, 2, ".")
	var result int
	for rowIndex, row := range input {
		for colIndex, val := range row {
			if val != 'A' {
				continue
			}

			masPatterns := checkXPattern(input, rowIndex, colIndex)
			result += masPatterns
		}
	}
	return result
}

func checkXPattern(input []string, row, col int) int {
	patterns := 0
	diagonals := [][4][2]int{
		{{-1, -1}, {1, 1}, {-1, 1}, {1, -1}},
		{{1, -1}, {-1, 1}, {-1, -1}, {1, 1}},
	}

	for _, d := range diagonals {
		if (input[row+d[0][0]][col+d[0][1]] == 'M' &&
			input[row+d[1][0]][col+d[1][1]] == 'S') &&
			(input[row+d[2][0]][col+d[2][1]] == 'M' &&
				input[row+d[3][0]][col+d[3][1]] == 'S') {
			patterns++
		}
		if (input[row+d[0][0]][col+d[0][1]] == 'S' &&
			input[row+d[1][0]][col+d[1][1]] == 'M') &&
			(input[row+d[2][0]][col+d[2][1]] == 'S' &&
				input[row+d[3][0]][col+d[3][1]] == 'M') {
			patterns++
		}
	}

	return patterns
}

func pad(input []string, padNum int, padSymbol string) []string {
	originalHeight := len(input)
	originalWidth := len(input[0])

	newHeight := originalHeight + padNum*2
	newWidth := originalWidth + padNum*2
	grid := make([]string, newHeight)

	emptyRow := strings.Repeat(padSymbol, newWidth)
	for i := range grid {
		grid[i] = emptyRow
	}

	for i, row := range input {
		grid[i+padNum] = strings.Repeat(padSymbol, padNum) + row + strings.Repeat(padSymbol, padNum)
	}

	return grid

}
