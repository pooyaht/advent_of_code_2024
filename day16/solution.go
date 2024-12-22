package day16

import (
	_ "embed"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n"))
}

func SolvePartA(input []string) int {
	var startPosition, endPosition [2]int

	for i, row := range input {
		for j, chr := range row {
			if chr == 'E' {
				endPosition = [2]int{i, j}
			}
			if chr == 'S' {
				startPosition = [2]int{i, j}
			}
		}
	}

	fromStart, _ := findPath(input, startPosition, endPosition)

	minCost := math.MaxInt32
	for i := 0; i < 4; i++ {
		if fromStart[endPosition[0]][endPosition[1]][i] < minCost {
			minCost = fromStart[endPosition[0]][endPosition[1]][i]
		}
	}

	minCost += minCost / 1000
	minCost += 1000

	return minCost
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(input []string) int {
	var startPosition, endPosition [2]int

	for i, row := range input {
		for j, chr := range row {
			if chr == 'E' {
				endPosition = [2]int{i, j}
			}
			if chr == 'S' {
				startPosition = [2]int{i, j}
			}
		}
	}

	fromStart, toEnd := findPath(input, startPosition, endPosition)

	minCost := math.MaxInt32
	for i := 0; i < 4; i++ {
		if fromStart[endPosition[0]][endPosition[1]][i] < minCost {
			minCost = fromStart[endPosition[0]][endPosition[1]][i]
		}
	}

	optimalTiles := make(map[[2]int]bool)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '#' {
				continue
			}

			for fromDir := 0; fromDir < 4; fromDir++ {
				costToHere := fromStart[i][j][fromDir]
				if costToHere == math.MaxInt32 {
					continue
				}

				for toDir := 0; toDir < 4; toDir++ {
					costFromHere := toEnd[i][j][toDir]
					if costFromHere == math.MaxInt32 {
						continue
					}

					turnCost := 0
					if int(math.Abs(float64(fromDir-toDir)))%2 == 1 {
						turnCost = 1000 - 1
					}

					totalCost := costToHere + turnCost + costFromHere
					if totalCost == minCost {
						optimalTiles[[2]int{i, j}] = true
						goto nextTile
					}
				}
			}
		nextTile:
		}
	}

	return len(optimalTiles)
}

func findPath(input []string, startPosition, endPosition [2]int) ([][][4]int, [][][4]int) {
	directions := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	fromStart := make([][][4]int, len(input))
	toEnd := make([][][4]int, len(input))
	for i := range input {
		fromStart[i] = make([][4]int, len(input[i]))
		toEnd[i] = make([][4]int, len(input[i]))
		for j := range input[i] {
			fromStart[i][j] = [4]int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}
			toEnd[i][j] = [4]int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}
		}
	}

	fromStart[startPosition[0]][startPosition[1]] = [4]int{0, 0, 0, 0}
	toEnd[endPosition[0]][endPosition[1]] = [4]int{0, 0, 0, 0}

	improved := true
	for improved {
		improved = false
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if input[i][j] == '#' {
					continue
				}
				for k := 0; k < 4; k++ {
					if fromStart[i][j][k] == math.MaxInt32 {
						continue
					}
					for l := 0; l < 4; l++ {
						newI, newJ := i+directions[l][0], j+directions[l][1]
						if newI < 0 || newI >= len(input) || newJ < 0 || newJ >= len(input[i]) || input[newI][newJ] == '#' {
							continue
						}

						newCost := fromStart[i][j][k]
						if int(math.Abs(float64(k-l)))%2 == 1 {
							newCost += 1000
						} else {
							newCost++
						}

						if newCost < fromStart[newI][newJ][l] {
							fromStart[newI][newJ][l] = newCost
							improved = true
						}
					}
				}
			}
		}

		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if input[i][j] == '#' {
					continue
				}
				for k := 0; k < 4; k++ {
					if toEnd[i][j][k] == math.MaxInt32 {
						continue
					}
					for l := 0; l < 4; l++ {
						newI, newJ := i-directions[l][0], j-directions[l][1]
						if newI < 0 || newI >= len(input) ||
							newJ < 0 || newJ >= len(input[i]) ||
							input[newI][newJ] == '#' {
							continue
						}

						newCost := toEnd[i][j][k]
						if int(math.Abs(float64(k-l)))%2 == 1 {
							newCost += 1000
						} else {
							newCost++
						}

						if newCost < toEnd[newI][newJ][l] {
							toEnd[newI][newJ][l] = newCost
							improved = true
						}
					}
				}
			}
		}
	}

	return fromStart, toEnd
}
