package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var dayFlag = flag.Int("d", 1, "advent of code day")
	var partFlag = flag.Int("p", 1, "advent of code day part")
	flag.Parse()

	executablePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	inputFilePath := executablePath + "/inputs/day" + fmt.Sprint(*dayFlag) + ".txt"
	var lines []string
	readLines(inputFilePath, &lines)

	solve(*dayFlag, *partFlag, &lines)
}

func solve(day int, part int, lines *[]string) {

	switch d := day; d {
	case 1:
		SolveDay1(part, lines)
	case 2:
		SolveDay2(part, lines)
	case 3:
		SolveDay3(part, lines)
	}
}

func readLines(filePath string, lines *[]string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Couldn't read input file from path: %v. Error: %v \n", filePath, err)
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
