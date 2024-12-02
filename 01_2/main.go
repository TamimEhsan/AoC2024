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

func calcSimilarity(left, right *[]int) int {
	/**
	calculate the difference between the pair of integers
	*/
	slices.Sort(*left)
	freq := make(map[int]int)

	similarity := 0
	for _, r := range *right {
		freq[r]++
	}

	for _, l := range *left {
		f, ok := freq[l]
		if !ok {
			continue
		}

		similarity += f * l
	}

	return similarity
}

func main() {
	// declare empty slice
	var left, right []int

	input(&left, &right)
	diff := calcSimilarity(&left, &right)

	fmt.Println(diff)
}
