#! /bin/bash

num=$1

mkdir day$num

cd day$num

touch solution.go
touch input.txt

echo "package main

import (
	f \"fmt\"
	\"time\"

	\"tarunbatra.com/aoc2021/utils\"
)

func main() {
	input := utils.GetInput(\"./input.txt\")
	t1 := time.Now()
	result1 := solution01(input)
	f.Printf(\"Part 1: %d (in %v)\n\", result1, time.Since((t1)))
	t2 := time.Now()
	result2 := solution02(input)
	f.Printf(\"Part 2: %d (in %v)\n\", result2, time.Since((t2)))
}

func solution01(input []string) int {
	return 0
}

func solution02(input []string) int {
	return 0
}
" >> solution.go