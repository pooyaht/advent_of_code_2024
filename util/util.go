package util

import "strings"

func Pad(input []string, padNum int, padSymbol string) []string {
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
