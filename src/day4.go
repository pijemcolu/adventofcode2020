package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	validFieldKeys         = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	minValidFieldsRequired = 7
)

// SolveDay4 solves the puzzle for day 4
func SolveDay4(part int, lines *[]string) {
	switch p := part; p {
	case 1:
		solveDay4Part1(lines)
	case 2:
		solveDay4Part2(lines)
	}
}

func solveDay4Part1(lines *[]string) {
	solution := 0

	passports := parsePassports(lines)
	for _, passport := range passports {

		validFields := 0
		for _, field := range passport.Fields {

			contains := contains(validFieldKeys, field.Key)
			if contains {
				validFields++
			}
		}
		if validFields >= minValidFieldsRequired {
			solution++
		}
	}

	fmt.Println(solution)
}

func solveDay4Part2(lines *[]string) {
	solution := 0

	passports := parsePassports(lines)
	for _, passport := range passports {
		fmt.Println(passport)

		validFields := 0
		for _, field := range passport.Fields {
			valid := validateField(field)
			if valid {
				validFields++
			}
		}
		if validFields >= minValidFieldsRequired {
			solution++
		}
	}

	fmt.Println(solution)
}

func parsePassports(lines *[]string) []Passport {
	var passports []Passport

	var passport Passport
	*lines = append(*lines, "") // append an empty line as the differentiator between passports
	for _, line := range *lines {

		kvs := strings.Split(line, " ")
		var fields []Field

		if line != "" {
			for _, kv := range kvs {

				kvSplit := strings.Split(kv, ":")
				k := kvSplit[0]
				v := kvSplit[1]

				field := Field{
					Key:   k,
					Value: v,
				}
				fields = append(fields, field)
			}

			passport.Fields = append(passport.Fields, fields...)
		} else {
			passports = append(passports, passport)
			passport = Passport{}
		}
	}
	return passports
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if element == e {
			return true
		}
	}
	return false
}

func validateField(field Field) bool {

	switch key := field.Key; key {
	case "byr":
		byr, _ := strconv.Atoi(field.Value)
		valid := byr >= 1920 && byr <= 2002
		fmt.Println(field.Value, valid)
		return valid
	case "iyr":
		iyr, _ := strconv.Atoi(field.Value)
		valid := iyr >= 2010 && iyr <= 2020
		fmt.Println(field.Value, valid)
		return valid
	case "eyr":
		eyr, _ := strconv.Atoi(field.Value)
		valid := eyr >= 2020 && eyr <= 2030
		fmt.Println(field.Value, valid)
		return valid
	case "hgt":
		valid := false
		if strings.HasSuffix(field.Value, "cm") {
			hgtString := field.Value[:len(field.Value)-2]
			hgt, _ := strconv.Atoi(hgtString)
			valid = hgt >= 150 && hgt <= 193
		}
		if strings.HasSuffix(field.Value, "in") {
			hgtString := field.Value[:len(field.Value)-2]
			hgt, _ := strconv.Atoi(hgtString)
			valid = hgt >= 59 && hgt <= 76
		}
		fmt.Println(field.Value, valid)
		return valid
	case "hcl":
		valid, _ := regexp.MatchString("^[#](.{6}$)[a-f0-9]*$", field.Value)
		fmt.Println(field.Value, valid)
		return valid
	case "ecl":
		valid := (field.Value == "amb" ||
			field.Value == "blu" ||
			field.Value == "brn" ||
			field.Value == "gry" ||
			field.Value == "grn" ||
			field.Value == "hzl" ||
			field.Value == "oth")
		fmt.Println(field.Value, valid)
		return valid
	case "pid":
		valid, _ := regexp.MatchString("^(.{9}$)[0-9]*$", field.Value)
		fmt.Println(field.Value, valid)
		return valid
	default:
		return false
	}
}

// Passport is a slice of Fields
type Passport struct {
	Fields []Field
}

// Field is a key value pair
type Field struct {
	Key   string
	Value string
}
