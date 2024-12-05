package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day4.txt")
	rows := strings.Split(strings.TrimSpace(string(input)), "\n")

	var grid [][]rune
	for _, row := range rows {
		grid = append(grid, []rune(strings.TrimRight(row, "\r")))
	}

	fmt.Println(countXMAS(grid)) // part 1
}

// Part 1
func countXMAS(grid [][]rune) int {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	isValid := func(x, y int) bool {
		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
	}

	count := 0
	target := "XMAS"
	targetLen := len(target)

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == rune(target[0]) {
				for _, dir := range directions {
					dr, dc := dir[0], dir[1]
					found := true
					for k := 1; k < targetLen; k++ {
						nr, nc := r+k*dr, c+k*dc
						if !isValid(nr, nc) {
							found = false
							break
						}
						if grid[nr][nc] != rune(target[k]) {
							found = false
							break
						}
					}
					if found {
						count++
					}
				}
			}
		}
	}

	return count
}
