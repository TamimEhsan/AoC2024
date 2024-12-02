package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func input(left, right *[]int) {
	/**
	take input from file, the inputs are list a pair of integer
	*/
	file, _ := os.Open("input.txt")
	defer file.Close()
	in := bufio.NewReader(file)

	var a, b int
	for {
		_, err := fmt.Fscan(in, &a, &b)
		if err != nil {
			break
		}
		*left = append(*left, a)
		*right = append(*right, b)
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcDiff(left, right *[]int) int {
	/**
	calculate the difference between the pair of integers
	*/
	slices.Sort(*left)
	slices.Sort(*right)

	diff := 0
	for i := 0; i < len(*left); i++ {
		diff += abs((*left)[i] - (*right)[i])
	}

	return diff
}

func main() {
	// declare empty slice
	var left, right []int

	input(&left, &right)
	diff := calcDiff(&left, &right)

	fmt.Println(diff)
}
