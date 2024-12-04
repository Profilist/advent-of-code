package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Get input from day1.txt
	input, _ := os.ReadFile("day1.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var slice1 []int // left list
	var slice2 []int // right list

	freq2 := map[int]int{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		result := strings.Fields(line)
		val1, _ := strconv.Atoi(result[0]) // left number
		val2, _ := strconv.Atoi(result[1]) // right number

		slice1 = append(slice1, val1)
		slice2 = append(slice2, val2)

		freq2[val2]++
	}

	slices.Sort(slice1)
	slices.Sort(slice2)

	sum := 0.0
	similarity := 0
	for i := 0; i < len(slice1); i++ {
		sum += math.Abs(float64(slice1[i] - slice2[i]))
		similarity += slice1[i] * freq2[slice1[i]]
	}

	fmt.Println(int(sum))   // part 1
	fmt.Println(similarity) // part 2
}
