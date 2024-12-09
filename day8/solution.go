package day8

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
	antennaToLocation := make(map[byte][][2]int, len(input))

	for i, row := range input {
		for j, chr := range row {
			if chr == '.' {
				continue
			}
			if _, ok := antennaToLocation[byte(chr)]; !ok {
				antennaToLocation[byte(chr)] = make([][2]int, 0, len(row))
			}
			antennaToLocation[byte(chr)] = append(antennaToLocation[byte(chr)], [2]int{i, j})
		}
	}

	result := make(map[[2]int]bool, 0)
	for _, locations := range antennaToLocation {
		for i, locationA := range locations {
			for _, locationB := range locations[i+1:] {
				diff := [2]int{locationB[0] - locationA[0], locationB[1] - locationA[1]}
				antinodeA := [2]int{locationA[0] - diff[0], locationA[1] - diff[1]}
				if antinodeA[0] >= 0 && antinodeA[0] < len(input) && antinodeA[1] >= 0 && antinodeA[1] < len(input[0]) {
					result[antinodeA] = true
				}
				antinodeB := [2]int{locationB[0] + diff[0], locationB[1] + diff[1]}
				if antinodeB[0] >= 0 && antinodeB[0] < len(input) && antinodeB[1] >= 0 && antinodeB[1] < len(input[0]) {
					result[antinodeB] = true
				}
			}
		}
	}

	return len(result)
}

func RunPartB() int {
	return SolvePartB(strings.Split(input, "\n"))
}

func SolvePartB(input []string) int {
	antennaToLocation := make(map[byte][][2]int, len(input))

	for i, row := range input {
		for j, chr := range row {
			if chr == '.' {
				continue
			}
			if _, ok := antennaToLocation[byte(chr)]; !ok {
				antennaToLocation[byte(chr)] = make([][2]int, 0, len(row))
			}
			antennaToLocation[byte(chr)] = append(antennaToLocation[byte(chr)], [2]int{i, j})
		}
	}

	result := make(map[[2]int]bool, 0)
	for _, locations := range antennaToLocation {
		for i, locationA := range locations {
			for _, locationB := range locations[i+1:] {
				diff := [2]int{locationB[0] - locationA[0], locationB[1] - locationA[1]}
				i := 0

				for {
					antinodeA := [2]int{locationA[0] - i*diff[0], locationA[1] - i*diff[1]}
					if antinodeA[0] >= 0 && antinodeA[0] < len(input) && antinodeA[1] >= 0 && antinodeA[1] < len(input[0]) {
						result[antinodeA] = true
					} else {
						break
					}
					i++
				}

				i = 0
				for {
					antinodeB := [2]int{locationB[0] + i*diff[0], locationB[1] + i*diff[1]}
					if antinodeB[0] >= 0 && antinodeB[0] < len(input) && antinodeB[1] >= 0 && antinodeB[1] < len(input[0]) {
						result[antinodeB] = true
					} else {
						break
					}
					i++
				}
			}
		}
	}

	return len(result)
}
