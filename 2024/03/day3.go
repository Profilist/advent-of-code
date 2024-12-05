package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("day3.txt")

	regex := regexp.MustCompile(`mul\(\d+,\d+\)`) // get all mul(a,b) strings
	matches := regex.FindAllString(string(input), -1)
	sum := 0

	regex2 := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`) // get all mul(a,b), do, and don't strings
	matches2 := regex2.FindAllString(string(input), -1)
	sum_filtered := 0

	// Part 1 calculations
	for _, match := range matches {
		regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // get a and b from mul(a,b)
		numbers := regex.FindAllStringSubmatch(match, -1)
		num1, _ := strconv.Atoi(numbers[0][1])
		num2, _ := strconv.Atoi(numbers[0][2])
		sum += num1 * num2
	}

	// Part 2 calculations
	final_matches := []string{}
	do := true
	// Filter out don't() mul(a, b) strings
	for _, match := range matches2 {
		if match == "don't()" {
			do = false
		}
		if do && match != "do()" {
			final_matches = append(final_matches, match)
		}
		if match == "do()" {
			do = true
		}
	}

	for _, match := range final_matches {
		regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // get a and b from mul(a,b)
		numbers := regex.FindAllStringSubmatch(match, -1)
		num1, _ := strconv.Atoi(numbers[0][1])
		num2, _ := strconv.Atoi(numbers[0][2])
		sum_filtered += num1 * num2
	}

	fmt.Println(sum)          // part 1
	fmt.Println(sum_filtered) // part 2
}
