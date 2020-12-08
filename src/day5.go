package main

import (
	"fmt"
	"sort"
)

// SolveDay5 solves the puzzle for day5
func SolveDay5(part int, lines *[]string) {
	switch p := part; p {
	case 1:
		solveDay5Part1(lines)
	case 2:
		solveDay5Part2(lines)
	}
}

func solveDay5Part1(lines *[]string) {
	maxSeatID := 0

	for _, line := range *lines {
		rowCode := line[:7]
		columnCode := line[7:]

		rowBase := 64
		row := 0
		for _, char := range rowCode {
			if char == 'B' {
				row = row + rowBase
			}
			rowBase = rowBase / 2
		}

		columnBase := 4
		column := 0
		for _, char := range columnCode {
			if char == 'R' {
				column = column + columnBase
			}
			columnBase = columnBase / 2
		}

		seatID := row*8 + column
		if seatID >= maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println(maxSeatID)
}

func solveDay5Part2(lines *[]string) {
	solution := 0

	var seatIDs []int
	for _, line := range *lines {
		rowCode := line[:7]
		columnCode := line[7:]

		rowBase := 64
		row := 0
		for _, char := range rowCode {
			if char == 'B' {
				row = row + rowBase
			}
			rowBase = rowBase / 2
		}

		columnBase := 4
		column := 0
		for _, char := range columnCode {
			if char == 'R' {
				column = column + columnBase
			}
			columnBase = columnBase / 2
		}

		seatID := row*8 + column
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)
	for i := 0; i < len(seatIDs)-1; i++ {
		seat1 := seatIDs[i]
		seat2 := seatIDs[i+1]

		if seat1+2 == seat2 {
			solution = seat1 + 1
		}
	}

	fmt.Println(solution)
}
