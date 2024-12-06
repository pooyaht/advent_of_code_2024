package day6

import (
	_ "embed"
	"strings"

	"github.com/pooyaht/advent_of_code_2024/util"
)

//go:embed input.txt
var input string

func RunPartA() int {
	res, _ := SolvePartA(strings.Split(input, "\n"))
	return res
}

func SolvePartA(grid []string) (int, []string) {
	guardStartRow := -1
	guardStartColumn := -1
	grid = util.Pad(grid, 1, "*")

	for i, row := range grid {
		col := strings.Index(row, "^")
		if col != -1 {
			guardStartColumn = col
			guardStartRow = i
			break
		}
	}

	if guardStartColumn == -1 || guardStartRow == -1 {
		return 0, grid
	}

	var result int
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	directionIndex := 0

	for {
		direction := directions[directionIndex]
		newGuardRow := guardStartRow + direction[0]
		newGuardColumn := guardStartColumn + direction[1]

		if grid[newGuardRow][newGuardColumn] == '#' {
			directionIndex = (directionIndex + 1) % len(directions)
			continue
		}
		if grid[newGuardRow][newGuardColumn] == '*' {
			break
		}

		guardStartRow = newGuardRow
		guardStartColumn = newGuardColumn
		if grid[newGuardRow][newGuardColumn] == 'X' {
			continue
		} else {
			result++
			grid[newGuardRow] = grid[newGuardRow][:newGuardColumn] + "X" + grid[newGuardRow][newGuardColumn+1:]
		}
	}

	return result, grid
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(grid []string) int {
	guardStartRow := -1
	guardStartColumn := -1
	for i, row := range grid {
		col := strings.Index(row, "^")
		if col != -1 {
			guardStartColumn = col
			guardStartRow = i
			break
		}
	}

	if guardStartColumn == -1 || guardStartRow == -1 {
		return 0
	}
	guardStartColumn++
	guardStartRow++

	_, grid = SolvePartA(grid)
	grid[guardStartRow] = grid[guardStartRow][:guardStartColumn] + "^" + grid[guardStartRow][guardStartColumn+1:]

	var visitedPositions [][2]int
	for i, row := range grid {
		for j, val := range row {
			if val == 'X' {
				visitedPositions = append(visitedPositions, [2]int{i, j})
			}
		}
	}

	var result int
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for _, position := range visitedPositions {
		row, col := position[0], position[1]
		gridCopy := make([]string, len(grid))
		copy(gridCopy, grid)

		gridCopy[row] = gridCopy[row][:col] + "#" + gridCopy[row][col+1:]

		guardRow, guardCol := guardStartRow, guardStartColumn
		directionIndex := 0
		loopFound := false
		count := 0

		for {
			count++
			tempDirection := directions[directionIndex]
			newRow := guardRow + tempDirection[0]
			newCol := guardCol + tempDirection[1]

			if gridCopy[newRow][newCol] == '#' {
				directionIndex = (directionIndex + 1) % len(directions)
				continue
			}

			if gridCopy[newRow][newCol] == '*' {
				break
			}

			if count > 9999 {
				loopFound = true
				break
			}

			guardRow, guardCol = newRow, newCol
		}

		if loopFound {
			result++
		}
	}

	return result
}
