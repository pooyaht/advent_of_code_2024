package day10

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
	trailHeads := make([][2]int, 0, len(input))
	for i, row := range input {
		for j, chr := range row {
			if chr == '0' {
				trailHeads = append(trailHeads, [2]int{i, j})
			}
		}
	}

	result := countTrailPaths(input, trailHeads, false)

	return result
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(input []string) int {
	trailHeads := make([][2]int, 0, len(input))
	for i, row := range input {
		for j, chr := range row {
			if chr == '0' {
				trailHeads = append(trailHeads, [2]int{i, j})
			}
		}
	}

	result := countTrailPaths(input, trailHeads, true)

	return result
}

func countTrailPaths(grid []string, trailHeads [][2]int, allowRevisit bool) int {
	var result int
	rows, cols := len(grid), len(grid[0])

	for _, trailHead := range trailHeads {
		var dfs func(row, column int, visited [][]bool)
		dfs = func(row, column int, visited [][]bool) {
			if grid[row][column] == '9' {
				result++
				if !allowRevisit {
					visited[row][column] = true
				}
				return
			}

			if !allowRevisit {
				visited[row][column] = true
			}

			for _, direction := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				newRow := row + direction[0]
				newColumn := column + direction[1]

				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < cols &&
					grid[newRow][newColumn] == grid[row][column]+1 &&
					(allowRevisit || !visited[newRow][newColumn]) {
					if !allowRevisit {
						dfs(newRow, newColumn, visited)
					} else {
						dfs(newRow, newColumn, nil)
					}
				}
			}

			if !allowRevisit {
				visited[row][column] = false
			}
		}

		if !allowRevisit {
			visited := make([][]bool, rows)
			for i := range visited {
				visited[i] = make([]bool, cols)
			}
			dfs(trailHead[0], trailHead[1], visited)
		} else {
			dfs(trailHead[0], trailHead[1], nil)
		}
	}

	return result
}
