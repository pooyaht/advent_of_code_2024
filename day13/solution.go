package day13

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func RunPartA() int {
	return SolvePartA(strings.Split(input, "\n\n"))
}

func SolvePartA(input []string) int {
	machines, _ := parseInput(input)
	var result int

	for _, machine := range machines {
		min_tokens := 0
		for nA := 0; nA <= 100; nA++ {
			for nB := 0; nB <= 100; nB++ {
				if (uint64(nA*machine.ButtonA.X+nB*machine.ButtonB.X) == machine.Prize.X) && (uint64(nA*machine.ButtonA.Y+nB*machine.ButtonB.Y) == machine.Prize.Y) {
					if min_tokens == 0 || 3*nA+nB < min_tokens {
						min_tokens = 3*nA + nB
					}
				}
			}
		}
		result += min_tokens
	}

	return result
}

func RunPartB() uint64 {
	return SolvePartB(strings.Split(input, "\n\n"))
}

func SolvePartB(input []string) uint64 {
	machines, _ := parseInput(input)
	var result uint64

	for _, machine := range machines {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000

		nA, nB, ok := solveSystem(machine)
		if !ok {
			continue
		}

		result += nA*3 + nB
	}

	return result
}

type Button struct {
	X int
	Y int
}

type Prize struct {
	X uint64
	Y uint64
}

type Machine struct {
	ButtonA Button
	ButtonB Button
	Prize   Prize
}

func parseInput(machines []string) ([]Machine, error) {
	result := make([]Machine, 0, len(machines))

	for i, machineStr := range machines {
		machine, err := parseMachine(machineStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing machine %d: %v", i+1, err)
		}
		result = append(result, machine)
	}

	return result, nil
}

func parseMachine(input string) (Machine, error) {
	buttonAPattern := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBPattern := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizePattern := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	aMatches := buttonAPattern.FindStringSubmatch(input)
	if len(aMatches) != 3 {
		return Machine{}, fmt.Errorf("invalid Button A format")
	}
	aX, _ := strconv.Atoi(aMatches[1])
	aY, _ := strconv.Atoi(aMatches[2])

	bMatches := buttonBPattern.FindStringSubmatch(input)
	if len(bMatches) != 3 {
		return Machine{}, fmt.Errorf("invalid Button B format")
	}
	bX, _ := strconv.Atoi(bMatches[1])
	bY, _ := strconv.Atoi(bMatches[2])

	pMatches := prizePattern.FindStringSubmatch(input)
	if len(pMatches) != 3 {
		return Machine{}, fmt.Errorf("invalid Prize format")
	}
	pX, _ := strconv.ParseUint(pMatches[1], 10, 64)
	pY, _ := strconv.ParseUint(pMatches[2], 10, 64)

	return Machine{
		ButtonA: Button{X: aX, Y: aY},
		ButtonB: Button{X: bX, Y: bY},
		Prize:   Prize{X: pX, Y: pY},
	}, nil
}

func solveSystem(m Machine) (nA uint64, nB uint64, possible bool) {
	aX := m.ButtonA.X
	aY := m.ButtonA.Y
	bX := m.ButtonB.X
	bY := m.ButtonB.Y
	pX := int(m.Prize.X)
	pY := int(m.Prize.Y)

	denom := aY*bX - aX*bY
	if denom == 0 {
		return 0, 0, false
	}

	num := pY*bX - pX*bY
	if num%denom != 0 {
		return 0, 0, false
	}

	nAInt := num / denom
	if nAInt < 0 {
		return 0, 0, false
	}

	num = pX - nAInt*aX
	if num%bX != 0 {
		return 0, 0, false
	}

	nBInt := num / bX
	if nBInt < 0 {
		return 0, 0, false
	}

	if nAInt*aX+nBInt*bX != pX ||
		nAInt*aY+nBInt*bY != pY {
		return 0, 0, false
	}

	return uint64(nAInt), uint64(nBInt), true
}
