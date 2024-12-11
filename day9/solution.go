package day9

import (
	_ "embed"
	"slices"
	"strconv"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(input)
}

func SolvePartA(input string) int {
	indexToID := make(map[int]int)
	freeIndexes := make([]int, 0, len(input))
	occupiedIndexes := make([]int, 0, len(input))
	id := 0
	index := 0

	for i, chr := range input {
		num, _ := strconv.Atoi(string(chr))

		if i%2 == 0 {
			for j := 0; j < num; j++ {
				indexToID[index] = id
				occupiedIndexes = append(occupiedIndexes, index)
				index++
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				freeIndexes = append(freeIndexes, index)
				index++
			}
		}
	}

	left := 0
	right := len(occupiedIndexes) - 1

	for right > left && left < len(freeIndexes) && right >= 0 {
		if freeIndexes[left] > occupiedIndexes[right] {
			break
		}
		indexToID[freeIndexes[left]] = indexToID[occupiedIndexes[right]]
		delete(indexToID, occupiedIndexes[right])
		left++
		right--
	}

	var checksum int
	for index, id := range indexToID {
		checksum += index * id
	}

	return checksum
}

func RunPartB() int {
	return SolvePartB(input)
}

func SolvePartB(input string) int {
	indexToID := make(map[[2]int]int)
	freeIndexes := make([][2]int, 0, len(input))
	occupiedIndexes := make([][2]int, 0, len(input))
	id := 0
	index := 0

	for i, chr := range input {
		num, _ := strconv.Atoi(string(chr))

		if i%2 == 0 {
			occupiedIndexes = append(occupiedIndexes, [2]int{index, index + num - 1})
			indexToID[[2]int{index, index + num - 1}] = id
			index += num
			id++
		} else {
			freeIndexes = append(freeIndexes, [2]int{index, index + num - 1})
			index += num
		}
	}

	right := len(occupiedIndexes) - 1

	for right >= 0 {
		left := 0
		for left < len(freeIndexes) {
			if freeIndexes[left][0] > occupiedIndexes[right][0] {
				break
			}

			freeBlockSize := freeIndexes[left][1] - freeIndexes[left][0]
			occupiedBlockSize := occupiedIndexes[right][1] - occupiedIndexes[right][0]

			if freeBlockSize >= occupiedBlockSize {
				indexToID[[2]int{freeIndexes[left][0], freeIndexes[left][0] + occupiedBlockSize}] = indexToID[occupiedIndexes[right]]
				delete(indexToID, occupiedIndexes[right])
				freeRemaning := freeBlockSize - occupiedBlockSize
				if freeRemaning > 0 {
					freeIndexes[left] = [2]int{freeIndexes[left][0] + occupiedBlockSize + 1, freeIndexes[left][0] + occupiedBlockSize + freeRemaning}
				} else {
					freeIndexes = slices.Delete(freeIndexes, left, left+1)
				}
				break
			}
			left++
		}

		right--
	}

	var checksum int
	for index, id := range indexToID {
		for i := index[0]; i <= index[1]; i++ {
			checksum += i * id
		}
	}

	return checksum
}
