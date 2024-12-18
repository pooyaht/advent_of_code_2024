package day15

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n\n"))
}

func SolvePartA(input []string) int {
	board := make([][]rune, len(input[0]))
	for i, line := range strings.Split(input[0], "\n") {
		board[i] = []rune(line)
	}
	directions := input[1]

	x, y := -1, -1
	for i, row := range board {
		for j, cell := range row {
			if cell == '@' {
				x, y = i, j
				break
			}
		}
	}

	dirMap := map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}

	for _, dir := range directions {
		board, x, y = moveRobotA(board, x, y, dirMap[dir])
	}

	sum := 0
	for i, row := range board {
		for j, cell := range row {
			if cell == 'O' {
				sum += 100*i + j
			}
		}
	}

	return sum
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n\n"))
}

func SolvePartB(input []string) int {
	board := make([][]rune, len(input[0]))
	originalBoard := strings.Split(input[0], "\n")

	for i, line := range originalBoard {
		board[i] = make([]rune, len(line)*2)
		pos := 0
		for _, c := range line {
			switch c {
			case '#':
				board[i][pos] = '#'
				board[i][pos+1] = '#'
			case 'O':
				board[i][pos] = '['
				board[i][pos+1] = ']'
			case '.':
				board[i][pos] = '.'
				board[i][pos+1] = '.'
			case '@':
				board[i][pos] = '@'
				board[i][pos+1] = '.'
			}
			pos += 2
		}
	}

	x, y := -1, -1
	for i, row := range board {
		for j, cell := range row {
			if cell == '@' {
				x, y = i, j
				break
			}
		}
	}

	dirMap := map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}

	directions := input[1]
	for _, dir := range strings.Split(directions, "\n") {
		for _, chr := range dir {
			board, x, y = moveRobotB(board, x, y, dirMap[chr])
		}
	}

	var result int
	for i, row := range board {
		for j, chr := range row {
			if chr == '[' {
				result += 100*i + j
			}
		}
	}

	return result
}

func moveRobotA(board [][]rune, x, y int, direction [2]int) ([][]rune, int, int) {
	newX, newY := x+direction[0], y+direction[1]

	if newX < 0 || newX >= len(board) || newY < 0 || newY >= len(board[0]) || board[newX][newY] == '#' {
		return board, x, y
	}

	if board[newX][newY] == 'O' {
		var positions [][2]int
		curX, curY := newX, newY

		for {
			positions = append(positions, [2]int{curX, curY})
			nextX, nextY := curX+direction[0], curY+direction[1]

			if nextX < 0 || nextX >= len(board) || nextY < 0 || nextY >= len(board[0]) || board[nextX][nextY] == '#' {
				return board, x, y
			}

			if board[nextX][nextY] == 'O' {
				curX, curY = nextX, nextY
				continue
			}

			if board[nextX][nextY] == '.' {
				for i := len(positions) - 1; i >= 0; i-- {
					pos := positions[i]
					board[pos[0]+direction[0]][pos[1]+direction[1]] = 'O'
					board[pos[0]][pos[1]] = '.'
				}
				break
			}
		}
	}

	board[x][y] = '.'
	board[newX][newY] = '@'
	return board, newX, newY
}

func moveRobotB(board [][]rune, x, y int, direction [2]int) ([][]rune, int, int) {
	newX, newY := x+direction[0], y+direction[1]

	if newX < 0 || newX >= len(board) || newY < 0 || newY >= len(board[0]) || board[newX][newY] == '#' {
		return board, x, y
	}

	var ok bool
	var positions [][2]int
	if board[newX][newY] == '[' {
		ok, positions = canMovePartB(board, newX, newY, direction)
	} else if board[newX][newY] == ']' {
		ok, positions = canMovePartB(board, newX, newY-1, direction)
	}

	if !ok {
		if board[newX][newY] == '.' {
			board[x][y] = '.'
			board[newX][newY] = '@'
			return board, newX, newY
		}
		return board, x, y
	}

	for _, position := range positions {
		board[position[0]][position[1]] = '.'
		board[position[0]][position[1]+1] = '.'
	}
	for _, position := range positions {
		board[position[0]+direction[0]][position[1]+direction[1]] = '['
		board[position[0]+direction[0]][position[1]+direction[1]+1] = ']'
	}

	board[x][y] = '.'
	board[newX][newY] = '@'
	return board, newX, newY
}

func canMovePartB(board [][]rune, bracketOpenX, bracketOpenY int, direction [2]int) (bool, [][2]int) {
	positions := make([][2]int, 0)
	ok := false

	if direction[0] == 0 {
		curX, curY := bracketOpenX, bracketOpenY
		for {
			if curY < 0 || curY >= len(board[0]) || board[curX][curY] == '#' {
				return false, positions
			}
			if board[curX][curY] == '[' {
				positions = append(positions, [2]int{curX, curY})
			}
			if board[curX][curY] == '.' {
				return true, positions
			}
			curY += direction[1]
		}
	} else {
		visited := make(map[[2]int]bool)

		var findConnected func(int, int) bool
		findConnected = func(x, y int) bool {
			if x < 0 || x >= len(board) || y < 0 || y+1 >= len(board[0]) {
				return false
			}

			if board[x][y] == '#' || board[x][y+1] == '#' {
				return false
			}

			if visited[[2]int{x, y}] {
				return true
			}

			visited[[2]int{x, y}] = true
			positions = append(positions, [2]int{x, y})

			nextX := x + direction[0]
			if nextX < 0 || nextX >= len(board) {
				return false
			}

			if board[nextX][y] == '#' || board[nextX][y+1] == '#' {
				return false
			}

			if board[x+direction[0]][y] == '[' && !findConnected(x+direction[0], y) {
				return false
			}
			if y >= 1 && board[x+direction[0]][y-1] == '[' && !findConnected(x+direction[0], y-1) {
				return false
			}
			if y+1 < len(board[0]) && board[x+direction[0]][y+1] == '[' && !findConnected(x+direction[0], y+1) {
				return false
			}

			return true
		}

		ok = findConnected(bracketOpenX, bracketOpenY)
	}

	return ok, positions
}
