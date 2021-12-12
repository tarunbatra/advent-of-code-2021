package main

import (
	f "fmt"
	"strconv"
	"strings"
	"time"

	"tarunbatra.com/aoc2021/utils"
)

func main() {
	input := utils.GetInput("./input.txt")
	t1 := time.Now()
	result1 := solution01(input)
	f.Printf("Part 1: %d (in %v)\n", result1, time.Since((t1)))
	t2 := time.Now()
	result2 := solution02(input)
	f.Printf("Part 2: %d (in %v)\n", result2, time.Since((t2)))
}

func solution01(input []string) int {
	gamma := ""
	epsilon := ""
	countOfOnes := make([]int, len(input[0]))
	for _, row := range input {
		binaryData := strings.Split(row, "")
		for i, col := range binaryData {
			if col == "1" {
				countOfOnes[i]++
			}
		}
	}
	for _, count := range countOfOnes {
		if count > len(input)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaValue, _ := strconv.ParseInt(gamma, 2, 0)
	epsilonValue, _ := strconv.ParseInt(epsilon, 2, 0)
	return int(gammaValue) * int(epsilonValue)
}

func solution02(input []string) int {
	binaryO2Value := bitCriteria(input, 0, false)[0]
	binaryCO2Value := bitCriteria(input, 0, true)[0]
	o2Value, _ := strconv.ParseInt(binaryO2Value, 2, 0)
	co2Value, _ := strconv.ParseInt(binaryCO2Value, 2, 0)
	return int(o2Value * co2Value)
}

func bitCriteria(list []string, bitPosition int, selectLeast bool) []string {
	values := map[string][]string{"1": {}, "0": {}}
	if len(list) == 1 {
		return list
	}
	count := map[string]int{"1": 0, "0": 0}
	for _, binaryStr := range list {
		bit := string(binaryStr[bitPosition])
		values[bit] = append(values[bit], binaryStr)
		count[bit]++
	}
	if selectLeast {
		if count["0"] > count["1"] {
			return bitCriteria(values["1"], bitPosition+1, selectLeast)
		} else {
			return bitCriteria(values["0"], bitPosition+1, selectLeast)
		}
	}
	if count["1"] >= count["0"] {
		return bitCriteria(values["1"], bitPosition+1, selectLeast)
	} else {
		return bitCriteria(values["0"], bitPosition+1, selectLeast)
	}
}
