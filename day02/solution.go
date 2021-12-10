package main

import (
	f "fmt"
	"strconv"
	"strings"

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
	x := 0
	y := 0
	for _, cmd := range input {
		subcmd := strings.Split(cmd, " ")
		direction := subcmd[0]
		unit, _ := strconv.Atoi(subcmd[1])
		switch direction {
		case "forward":
			x += unit
		case "up":
			y -= unit
		case "down":
			y += unit
		}
	}
	return x * y
}

func solution02(input []string) int {
	x := 0
	y := 0
	aim := 0
	for _, cmd := range input {
		subcmd := strings.Split(cmd, " ")
		direction := subcmd[0]
		unit, _ := strconv.Atoi(subcmd[1])
		switch direction {
		case "forward":
			x += unit
			y += aim * unit
		case "up":
			aim -= unit
		case "down":
			aim += unit
		}
	}
	return x * y
}
