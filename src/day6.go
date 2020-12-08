package main

import (
	"fmt"
	"strconv"
)

// SolveDay6 solves the puzzle for day6
func SolveDay6(part int, lines *[]string) {
	switch p := part; p {
	case 1:
		solveDay6Part1(lines)
	case 2:
		solveDay6Part2(lines)
	}
}

func solveDay6Part1(lines *[]string) {
	solution := 0

	*lines = append(*lines, "") // append an empty line as the differentiator between groups
	var uniqueGroupAnswers []byte
	for _, line := range *lines {

		if line != "" {
			for i := 0; i < len(line); i++ {
				personAnswer := line[i]
				exists := false
				for _, groupAnswer := range uniqueGroupAnswers {
					if groupAnswer == personAnswer {
						exists = true
					}
				}
				if exists == false {
					uniqueGroupAnswers = append(uniqueGroupAnswers, personAnswer)
				}
			}
		} else {
			solution = solution + len(uniqueGroupAnswers)
			uniqueGroupAnswers = []byte{}
		}
	}

	fmt.Println(solution)
}

func solveDay6Part2(lines *[]string) {
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
