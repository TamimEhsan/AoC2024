package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(grid *[]string) {
	/**
	take input from file
	*/
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		*grid = append(*grid, line)
	}
}

func checkXMAS(grid *[]string, x, y, dir int) bool {
	/**
	extract the XMAS from the position x, y, and direction dir
	*/
	dirx := []int{0, 0, 1, -1, 1, 1, -1, -1}
	diry := []int{1, -1, 0, 0, 1, -1, 1, -1}

	XMAS := "XMAS"

	// check if the position is valid
	for i := 0; i < 4; i++ {
		if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) {
			return false
		}
		if (*grid)[x][y] != XMAS[i] {
			return false
		}
		x += dirx[dir]
		y += diry[dir]
	}

	return true
}

func countXMAS(grid *[]string) int {
	/**
	count the number of XMAS around the position x, y
	*/
	count := 0
	for x := 0; x < len(*grid); x++ {
		for y := 0; y < len((*grid)[0]); y++ {
			for dir := 0; dir < 8; dir++ {
				if checkXMAS(grid, x, y, dir) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	// declare empty slice
	var grid []string

	// take input from file
	input(&grid)

	// calculate the result
	res := countXMAS(&grid)

	fmt.Println(res)
}
