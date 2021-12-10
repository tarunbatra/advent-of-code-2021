package main

import (
	f "fmt"
	"strconv"

	"tarunbatra.com/aoc2021/utils"
)

func main() {
	input := utils.GetInput("./input.txt")
	result1 := solution01(input)
	f.Println(result1)
	result2 := solution02(input)
	f.Println(result2)
}

func solution01(input []string) int {
	increasingDepth := 0
	for i := 1; i < len(input); i++ {
		num1, _ := strconv.Atoi(input[i-1])
		num2, _ := strconv.Atoi(input[i])
		if num1 < num2 {
			increasingDepth++
		}
	}
	return increasingDepth
}

func solution02(input []string) int {
	increasingDepth := 0
	for i := 2; i < len(input)-1; i++ {
		num1, _ := strconv.Atoi(input[i-2])
		num2, _ := strconv.Atoi(input[i-1])
		num3, _ := strconv.Atoi(input[i])
		sumA := num1 + num2 + num3

		num4, _ := strconv.Atoi(input[i+1])
		sumB := num2 + num3 + num4
		if sumA < sumB {
			increasingDepth++
		}
	}
	return increasingDepth
}
