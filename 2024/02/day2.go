package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day2.txt")
	reports := strings.Split(string(input), "\n")

	// part 1
	count := 0
	for i := 0; i < len(reports); i++ {
		levelsStrings := strings.Fields(reports[i])
		levels := make([]int, len(levelsStrings))

		for i, level := range levelsStrings {
			levels[i], _ = strconv.Atoi(level)
		}
		if isReportSafe(levels) {
			count += 1
		}
	}
	fmt.Println(count) // part 1

	// part 2
	fixed_count := 0
	for i := 0; i < len(reports); i++ {
		levelsStrings := strings.Fields(reports[i])
		levels := make([]int, len(levelsStrings))

		for i, level := range levelsStrings {
			levels[i], _ = strconv.Atoi(level)
		}

		if isReportSafe(levels) {
			fixed_count += 1
		} else {
			for j := 0; j < len(levels); j++ {
				removed_current := make([]int, len(levels)-1)
				copy(removed_current[:j], levels[:j])
				copy(removed_current[j:], levels[j+1:])

				if isReportSafe(removed_current) {
					fixed_count += 1
					break
				}
			}
		}
	}
	fmt.Println(fixed_count) // part 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isReportSafe(report []int) bool {
	inc, dec := true, true
	for j := 0; j < len(report)-1; j++ {
		curr := report[j]
		next := report[j+1]

		diff := abs(curr - next)
		if diff < 1 || diff > 3 {
			return false
		}

		if curr < next {
			dec = false
		} else if curr > next {
			inc = false
		}
	}

	return inc || dec
}
