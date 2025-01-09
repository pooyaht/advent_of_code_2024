package day18

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n")[:1024], 71)
}

func SolvePartA(memorySpace []string, n int) int {
	type queueItem struct {
		pos  [2]int
		path [][2]int
	}
	queue := make([]queueItem, 0, len(memorySpace))
	visited := make(map[[2]int]bool)

	grid := makeGrid(memorySpace, n)

	queue = append(queue, queueItem{
		pos:  [2]int{0, 0},
		path: make([][2]int, 0),
	})
	visited[[2]int{0, 0}] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		x, y := item.pos[0], item.pos[1]
		if x == n-1 && y == n-1 {
			return len(item.path)
		}

		for _, dir := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] == '#' {
				continue
			}
			if visited[[2]int{nx, ny}] {
				continue
			}
			queue = append(queue, queueItem{
				pos:  [2]int{nx, ny},
				path: append(item.path, [2]int{nx, ny}),
			})
			visited[[2]int{nx, ny}] = true
		}
	}

	return -1
}

func RunPartB() string {
	return SolvePartB(strings.Split(input, "\n"), 71)
}

func SolvePartB(memorySpace []string, n int) string {
	left, right := 0, len(memorySpace)
	ans := ""
	for left < right {
		mid := (left + right) / 2
		res := SolvePartA(memorySpace[:mid], n)
		if res == -1 {
			ans = memorySpace[mid-1]
			right = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}

func makeGrid(memorySpace []string, n int) [][]rune {
	grid := make([][]rune, n)

	for i := range n {
		grid[i] = make([]rune, n)
		for j := range n {
			grid[i][j] = '.'
		}
	}

	for _, line := range memorySpace {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		grid[y][x] = '#'
	}

	return grid
}
