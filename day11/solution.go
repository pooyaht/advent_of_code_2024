package day11

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(input, 25)
}

func SolvePartA(input string, blinks int) int {
	stones := make([]int, 0, len(input))
	for _, num := range strings.Split(input, " ") {
		numInt, _ := strconv.Atoi(num)
		stones = append(stones, numInt)
	}

	newStones := make([]int, 0, len(stones))
	for blinks > 0 {
		blinks--

		for _, stone := range stones {
			stoneNumDigits := len(strconv.Itoa(stone))

			if stone == 0 {
				newStones = append(newStones, stone+1)
			} else if stoneNumDigits%2 == 0 {
				newStones = append(newStones, stone/int(math.Pow10(stoneNumDigits/2)))
				newStones = append(newStones, stone%int(math.Pow10(stoneNumDigits/2)))
			} else {
				newStones = append(newStones, stone*2024)
			}
		}

		stones = newStones
		newStones = make([]int, 0, len(stones))
	}

	return len(stones)
}

func RunPartB() int {
	return SolvePartB(input, 75)
}

func SolvePartB(input string, blinks int) int {
	currentCounts := make(map[int]int)

	for _, num := range strings.Split(input, " ") {
		numInt, _ := strconv.Atoi(num)
		currentCounts[numInt]++
	}

	for blink := 0; blink < blinks; blink++ {
		nextCounts := make(map[int]int)

		for num, count := range currentCounts {
			if num == 0 {
				nextCounts[1] += count
			} else {
				digits := len(strconv.Itoa(num))
				if digits%2 == 0 {
					left := num / int(math.Pow10(digits/2))
					right := num % int(math.Pow10(digits/2))
					nextCounts[left] += count
					nextCounts[right] += count
				} else {
					nextCounts[num*2024] += count
				}
			}
		}

		currentCounts = nextCounts
	}

	totalStones := 0
	for _, count := range currentCounts {
		totalStones += count
	}

	return totalStones
}
