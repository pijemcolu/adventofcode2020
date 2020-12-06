package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var dayFlag = flag.Int("d", 1, "advent of code day")
	var partFlag = flag.Int("p", 1, "advent of code day part")
	flag.Parse()

	inputFilePath := "./inputs/day" + fmt.Sprint(*dayFlag) + ".txt"
	var lines []string
	readLines(inputFilePath, &lines)

	solve(*dayFlag, *partFlag, &lines)
}

func solve(day int, part int, lines *[]string) {

	switch d := day; d {
	case 1:
		SolveDay1(part, lines)
	}
}

func readLines(filePath string, lines *[]string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Couldn't read file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		*lines = append(*lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
