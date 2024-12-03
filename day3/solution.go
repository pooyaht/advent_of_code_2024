package day3

import (
	_ "embed"
	"strconv"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(input)
}

func SolvePartA(input string) int {
	mulParser := keywordParser("mul")
	leftParanParser := charParser('(')
	rightParanParser := charParser(')')
	semiColonParser := charParser(',')
	numParser := numberParser()
	position := 0
	var result int

	for {
		if position >= len(input) {
			break
		}

		mulResult := mulParser(input, position)
		position = mulResult.position
		if !mulResult.success {
			position++
			continue
		}

		leftParanResult := leftParanParser(input, position)
		position = leftParanResult.position
		if !leftParanResult.success {
			position++
			continue
		}

		numResult1 := numParser(input, position)
		position = numResult1.position
		if !numResult1.success {
			continue
		}

		semiColonResult := semiColonParser(input, position)
		position = semiColonResult.position
		if !semiColonResult.success {
			position++
			continue
		}

		numResult2 := numParser(input, position)
		position = numResult2.position
		if !numResult2.success {
			continue
		}

		rightParanParser := rightParanParser(input, position)
		position = rightParanParser.position
		if !rightParanParser.success {
			position++
			continue
		}

		num1, _ := strconv.ParseInt(numResult1.value, 10, 64)
		num2, _ := strconv.ParseInt(numResult2.value, 10, 64)

		result += int(num1) * int(num2)
	}

	return result
}

func RunPartB() int {
	return SolvePartB(input)
}

func SolvePartB(input string) int {
	mulParser := keywordParser("mul")
	doParser := keywordParser("do()")
	dontParser := keywordParser("don't()")
	leftParanParser := charParser('(')
	rightParanParser := charParser(')')
	semiColonParser := charParser(',')
	numParser := numberParser()

	position := 0
	var result int

	for {
		if position >= len(input) {
			break
		}

		dontResult := dontParser(input, position)
		position = dontResult.position
		if dontResult.success {
			for {
				if position >= len(input) {
					break
				}

				doResult := doParser(input, position)
				position = doResult.position
				if doResult.success {
					break
				} else {
					position++
				}
			}
		}

		mulResult := mulParser(input, position)
		position = mulResult.position
		if !mulResult.success {
			position++
			continue
		}

		leftParanResult := leftParanParser(input, position)
		position = leftParanResult.position
		if !leftParanResult.success {
			position++
			continue
		}

		numResult1 := numParser(input, position)
		position = numResult1.position
		if !numResult1.success {
			continue
		}

		semiColonResult := semiColonParser(input, position)
		position = semiColonResult.position
		if !semiColonResult.success {
			position++
			continue
		}

		numResult2 := numParser(input, position)
		position = numResult2.position
		if !numResult2.success {
			continue
		}

		rightParanParser := rightParanParser(input, position)
		position = rightParanParser.position
		if !rightParanParser.success {
			position++
			continue
		}

		num1, _ := strconv.ParseInt(numResult1.value, 10, 64)
		num2, _ := strconv.ParseInt(numResult2.value, 10, 64)

		result += int(num1) * int(num2)
	}

	return result
}

type Parser = func(input string, position int) parseResult

type parseResult struct {
	success  bool
	position int
	value    string
}

func charParser(char byte) Parser {
	return func(input string, position int) parseResult {
		if position >= len(input) {
			return parseResult{
				success:  false,
				position: position,
				value:    "",
			}
		}

		if input[position] == char {
			return parseResult{
				success:  true,
				position: position + 1,
				value:    string(char),
			}
		} else {
			return parseResult{
				success:  false,
				position: position,
				value:    "",
			}
		}
	}
}

func keywordParser(keyword string) Parser {
	return func(input string, position int) parseResult {
		if position+len(keyword) >= len(input) {
			return parseResult{
				success:  false,
				position: position,
				value:    "",
			}
		}

		if input[position:position+len(keyword)] == keyword {
			return parseResult{
				success:  true,
				position: position + len(keyword),
				value:    keyword,
			}
		} else {
			return parseResult{
				success:  false,
				position: position,
				value:    "",
			}
		}
	}
}

func numberParser() Parser {
	isDigit := func(ch byte) bool {
		return ch >= '0' && ch <= '9'
	}

	return func(input string, position int) parseResult {
		if position >= len(input) {
			return parseResult{
				success:  false,
				position: position,
				value:    "",
			}
		}

		if !isDigit(input[position]) {
			return parseResult{
				success:  false,
				position: position,
				value:    "",
			}
		}

		startPos := position
		for position < len(input) && isDigit(input[position]) {
			position++
		}

		number := input[startPos:position]

		return parseResult{
			success:  true,
			position: position,
			value:    number,
		}
	}
}
