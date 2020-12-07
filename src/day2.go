package main

import (
	"fmt"
	"strconv"
	"strings"
)

// SolveDay2 solves the puzzle for day1
func SolveDay2(part int, lines *[]string) {
	switch p := part; p {
	case 1:
		solveDay2Part1(lines)
	case 2:
		solveDay2Part2(lines)
	}
}

func solveDay2Part1(lines *[]string) {
	solution := 0

	for _, line := range *lines {
		split := strings.Split(line, " ")

		min, _ := strconv.Atoi(strings.Split(split[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(split[0], "-")[1])
		character := string(split[1][0])
		password := split[2]

		charactercount := strings.Count(password, character)
		validPassword := charactercount >= min && charactercount <= max
		if validPassword {
			solution++
		}
	}

	fmt.Println(solution)
}

func solveDay2Part2(lines *[]string) {
	solution := 0

	for _, line := range *lines {
		split := strings.Split(line, " ")

		pos1, _ := strconv.Atoi(strings.Split(split[0], "-")[0])
		pos2, _ := strconv.Atoi(strings.Split(split[0], "-")[1])
		character := string(split[1][0])
		password := split[2]

		bothCharsSame := string(password[pos1-1]) == string(password[pos2-1])
		if bothCharsSame {
			continue
		}

		pos1Valid := string(password[pos1-1]) == character
		if pos1Valid {
			solution++
			continue
		}

		pos2Valid := string(password[pos2-1]) == character
		if pos2Valid {
			solution++
		}
	}

	fmt.Println(solution)
}
