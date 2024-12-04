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

func isValid(grid *[]string, x, y int) bool {
	/**
	check if the position is valid
	*/
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) {
		return false
	}
	return true
}

func checkMAS(grid *[]string, x, y, dir int) bool {
	/*
		check if the current position is the middle of a mas or sam in the direction dir
	*/
	dirx := [][]int{{-1, 0, 1}, {-1, 0, 1}}
	diry := [][]int{{-1, 0, 1}, {1, 0, -1}}
	MAS := []string{"MAS", "SAM"}

	for i := 0; i < 2; i++ {
		match := true
		for j := 0; j < 3; j++ {
			if (*grid)[x+dirx[dir][j]][y+diry[dir][j]] != MAS[i][j] {
				match = false
			}
		}
		if match {
			return true
		}
	}
	return false
}

func checkXMAS(grid *[]string, x, y int) bool {
	/**
	check if the current position is the middle of a x-mas
	*/
	dirx := []int{-1, 0, 1, -1, 0, 1}
	diry := []int{-1, 0, 1, 1, 0, -1}
	for i := 0; i < 6; i++ {
		if !isValid(grid, x+dirx[i], y+diry[i]) {
			return false
		}
	}

	return checkMAS(grid, x, y, 0) && checkMAS(grid, x, y, 1)
}

func countXMAS(grid *[]string) int {
	/**
	count the number of XMAS around the position x, y
	*/
	count := 0
	for x := 0; x < len(*grid); x++ {
		for y := 0; y < len((*grid)[0]); y++ {
			if checkXMAS(grid, x, y) {
				count++
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
