package day14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n"), 100, 101, 103)
}

func SolvePartA(input []string, seconds, width, height int) int {
	positions, velocities := parseInput(input)
	var q1, q2, q3, q4 int

	for i := range positions {
		positions[i][0] = (((positions[i][0] + seconds*velocities[i][0]) % width) + width) % width
		positions[i][1] = (((positions[i][1] + seconds*velocities[i][1]) % height) + height) % height

		if positions[i][0] >= 0 && positions[i][0] < width/2 && positions[i][1] >= 0 && positions[i][1] < height/2 {
			q1++
			continue
		}
		if positions[i][0] >= 0 && positions[i][0] < width/2 && positions[i][1] >= height/2+1 && positions[i][1] < height {
			q3++
			continue
		}
		if positions[i][0] >= width/2+1 && positions[i][0] < width && positions[i][1] >= 0 && positions[i][1] < height/2 {
			q2++
			continue
		}
		if positions[i][0] >= width/2+1 && positions[i][0] < width && positions[i][1] >= height/2+1 && positions[i][1] < height {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"), 101, 103)
}

func SolvePartB(input []string, width, height int) int {
	return -1
}

func parseInput(input []string) ([][2]int, [][2]int) {
	positions := make([][2]int, 0, len(input))
	velocities := make([][2]int, 0, len(input))

	for _, line := range input {
		parts := strings.Split(line, " ")
		var x, y, vx, vy int

		fmt.Sscanf(parts[0][2:], "%d,%d", &x, &y)
		positions = append(positions, [2]int{x, y})

		fmt.Sscanf(parts[1][2:], "%d,%d", &vx, &vy)
		velocities = append(velocities, [2]int{vx, vy})
	}

	return positions, velocities
}
