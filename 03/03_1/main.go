package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func input(muls *[]string) {
	/**
	take input from file, the inputs are list a pair of integer
	*/
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pattern := "mul\\([0-9]+,[0-9]+\\)"
	r, _ := regexp.Compile(pattern)
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllString(line, -1)
		*muls = append(*muls, match...)
	}
}

func extractNumbers(s string) (int, int) {
	/**
	extract two numbers from the string
	*/
	r, _ := regexp.Compile("[0-9]+")
	nums := r.FindAllString(s, -1)
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return num1, num2
}

func calcResult(muls *[]string) int {
	/**
	calculate the result of the multiplication
	*/
	result := 0
	for _, mul := range *muls {
		num1, num2 := extractNumbers(mul)
		result += num1 * num2
	}
	return result
}
func main() {
	// declare empty slice
	var muls []string

	// take input from file
	input(&muls)

	// calculate the result
	res := calcResult(&muls)

	fmt.Println(res)
}
