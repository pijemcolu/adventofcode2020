package main

import (
	"fmt"
)

// SolveDay3 solves the puzzle for day 3
func SolveDay3(part int, lines *[]string) {
	switch p := part; p {
	case 1:
		solveDay3Part1(lines)
	case 2:
		solveDay3Part2(lines)
	}
}

func solveDay3Part1(lines *[]string) {
	solution := 0

	solution = solveDay3SlopeRun(lines, 1, 3)

	fmt.Println(solution)
}

func solveDay3Part2(lines *[]string) {
	solution := 0

	down1right1 := solveDay3SlopeRun(lines, 1, 1)
	down1right3 := solveDay3SlopeRun(lines, 1, 3)
	down1right5 := solveDay3SlopeRun(lines, 1, 5)
	down1right7 := solveDay3SlopeRun(lines, 1, 7)
	down2right1 := solveDay3SlopeRun(lines, 2, 1)
	solution = down1right1 * down1right3 * down1right5 * down1right7 * down2right1

	fmt.Println(solution)
}

func solveDay3SlopeRun(lines *[]string, down int, right int) int {
	solution := 0

	for i := 0; i < len(*lines); i += down {
		line := (*lines)[i]

		treePos := (i / down * right) % len(line)
		character := string(line[treePos])

		if character == "#" {
			solution++
		}
	}

	return solution
}
