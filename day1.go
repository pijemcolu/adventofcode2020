package main

import (
	"fmt"
	"strconv"
)

// SolveDay1 solves the puzzle for day1
func SolveDay1(part int, lines *[]string) {
	switch p := part; p {
	case 1:
		solveDay1Part1(lines)
	case 2:
		solveDay1Part2(lines)
	}
}

func solveDay1Part1(lines *[]string) {
	solution := 0

	for i, lineI := range *lines {
		valueI, _ := strconv.Atoi(lineI)
		for _, lineJ := range (*lines)[i+1:] {
			valueJ, _ := strconv.Atoi(lineJ)

			if valueI+valueJ == 2020 {
				solution = valueI * valueJ
				break
			}
		}
	}
	fmt.Println(solution)
}

func solveDay1Part2(lines *[]string) {
	solution := 0

	for i, lineI := range *lines {
		valueI, _ := strconv.Atoi(lineI)
		for j, lineJ := range (*lines)[i+1:] {
			valueJ, _ := strconv.Atoi(lineJ)
			for _, lineK := range (*lines)[j+1:] {
				valueK, _ := strconv.Atoi(lineK)
				if valueI+valueJ+valueK == 2020 {
					solution = valueI * valueJ * valueK
					break
				}
			}
		}
	}
	fmt.Println(solution)
}
