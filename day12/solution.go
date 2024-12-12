package day12

import (
	_ "embed"
	"strings"

	"github.com/pooyaht/advent_of_code_2024/util"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n"))
}

func SolvePartA(grid []string) int {
	grid = util.Pad(grid, 1, ".")

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	var result int
	for i, row := range grid {
		for j, chr := range row {
			if chr == '.' {
				continue
			}
			result += fencingPriceA(grid, visited, i, j)
		}
	}

	return result
}

func fencingPriceA(grid []string, visited [][]bool, row, col int) int {
	area := 0
	perimeter := 0
	char := grid[row][col]

	var dfs func(row, column int)
	dfs = func(row, column int) {
		currentChar := grid[row][column]
		if currentChar != char || visited[row][column] {
			return
		}

		visited[row][column] = true

		area++
		for _, direction := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			newRow := row + direction[0]
			newColumn := column + direction[1]
			if grid[newRow][newColumn] != char {
				perimeter++
			}
			dfs(newRow, newColumn)
		}
	}

	dfs(row, col)

	return area * perimeter
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(grid []string) int {
	grid = util.Pad(grid, 1, ".")
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	var result int
	for i, row := range grid {
		for j, chr := range row {
			if chr == '.' || visited[i][j] {
				continue
			}
			result += fencingPriceB(grid, visited, i, j)
		}
	}
	return result
}

func countCorners(grid []string, row, col int, char byte) int {
	corners := 0
	if grid[row-1][col] != char && grid[row][col-1] != char {
		corners++
	}

	if grid[row-1][col] != char && grid[row][col+1] != char {
		corners++
	}

	if grid[row+1][col] != char && grid[row][col-1] != char {
		corners++
	}

	if grid[row+1][col] != char && grid[row][col+1] != char {
		corners++
	}

	if grid[row][col-1] == char && grid[row+1][col] == char && grid[row+1][col-1] != char {
		corners++
	}

	if grid[row][col-1] == char && grid[row-1][col] == char && grid[row-1][col-1] != char {
		corners++
	}

	if grid[row][col+1] == char && grid[row+1][col] == char && grid[row+1][col+1] != char {
		corners++
	}

	if grid[row-1][col] == char && grid[row][col+1] == char && grid[row-1][col+1] != char {
		corners++
	}

	return corners
}

func fencingPriceB(grid []string, visited [][]bool, row, col int) int {
	area := 0
	sides := 0
	char := grid[row][col]

	var dfs func(row, column int)
	dfs = func(row, column int) {
		if grid[row][column] != char || visited[row][column] {
			return
		}

		visited[row][column] = true
		area++
		sides += countCorners(grid, row, column, char)

		for _, dir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			dfs(row+dir[0], column+dir[1])
		}
	}

	dfs(row, col)

	return area * sides
}
