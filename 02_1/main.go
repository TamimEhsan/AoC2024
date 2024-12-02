package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func input(matrix *[][]int) {
	/**
	take input from file, the inputs are list a pair of integer
	*/
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		rowStr := strings.Split(line, " ")
		rowInt := make([]int, len(rowStr))
		for i, str := range rowStr {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			rowInt[i] = num
		}
		*matrix = append(*matrix, rowInt)
	}

}

func checkSorted(row *[]int) bool {

	sortedAsc := slices.IsSorted(*row)
	slices.Reverse(*row)
	sortedDesc := slices.IsSorted(*row)

	return sortedAsc || sortedDesc
}

func checkDiff(row *[]int) bool {

	for i := 1; i < len(*row); i++ {
		diff := abs((*row)[i] - (*row)[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func safetyCheck(matrix *[][]int) int {
	safe := 0
	for _, row := range *matrix {
		if checkSorted(&row) && checkDiff(&row) {
			safe++
		}
	}
	return safe
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// declare empty slice
	var matrix [][]int

	input(&matrix)

	safe := safetyCheck(&matrix)

	fmt.Println(safe)
}
