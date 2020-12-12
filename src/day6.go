package main

import (
	"fmt"
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

	groupAnswersMap := map[string]int{}
	personCount := 0
	*lines = append(*lines, "") // append an empty line as the differentiator between groups
	for _, personAnswers := range *lines {

		if personAnswers != "" {
			personCount++
			for i := 0; i < len(personAnswers); i++ {
				singleAnswer := personAnswers[i]
				groupAnswersMap[string(singleAnswer)] = groupAnswersMap[string(singleAnswer)] + 1
			}
			fmt.Println(groupAnswersMap)
		} else {
			fmt.Println("eog")
			for _, v := range groupAnswersMap {
				if v == personCount {
					solution++
				}
			}

			groupAnswersMap = map[string]int{}
			personCount = 0
		}
	}

	fmt.Println(solution)

}
