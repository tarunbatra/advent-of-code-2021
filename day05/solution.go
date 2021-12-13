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
	coordSeenCountMap := map[string]int{}
	for _, row := range input {
		start, end := parseLineSegment(row)
		lineCoords := getLineCoords1(start, end)
		for _, lineCoord := range lineCoords {
			coordKey := f.Sprint(lineCoord[0], ",", lineCoord[1])
			if _, ok := coordSeenCountMap[coordKey]; !ok {
				coordSeenCountMap[coordKey] = 0
			}
			coordSeenCountMap[coordKey] += 1
		}
	}
	coordsSeenTwice := 0
	for _, seenCount := range coordSeenCountMap {
		if seenCount > 1 {
			coordsSeenTwice += 1
		}
	}
	return coordsSeenTwice
}

func solution02(input []string) int {
	coordSeenCountMap := map[string]int{}
	for _, row := range input {
		start, end := parseLineSegment(row)
		lineCoords := getLineCoords2(start, end)
		for _, lineCoord := range lineCoords {
			coordKey := f.Sprint(lineCoord[0], ",", lineCoord[1])
			if _, ok := coordSeenCountMap[coordKey]; !ok {
				coordSeenCountMap[coordKey] = 0
			}
			coordSeenCountMap[coordKey] += 1
		}
	}
	coordsSeenTwice := 0
	for _, seenCount := range coordSeenCountMap {
		if seenCount > 1 {
			coordsSeenTwice += 1
		}
	}
	return coordsSeenTwice
}

func parseLineSegment(row string) ([]int, []int) {
	segments := strings.Split(row, " -> ")
	startCoord := parseStrCoord(segments[0])
	endCoord := parseStrCoord(segments[1])
	return startCoord, endCoord
}

func parseStrCoord(str string) []int {
	segment := strings.Split(str, ",")
	x, _ := strconv.Atoi(segment[0])
	y, _ := strconv.Atoi(segment[1])
	return []int{x, y}
}

func getLineCoords1(start []int, end []int) [][]int {

	movingAxis := 0
	stableAxis := 1
	lineCoords := [][]int{}
	if start[0] != end[0] && start[1] != end[1] {
		return lineCoords
	}
	if start[1] != end[1] {
		movingAxis = 1
		stableAxis = 0
	}
	if start[movingAxis] < end[movingAxis] {
		for i := 0; i <= end[movingAxis]-start[movingAxis]; i++ {
			newCoord := make([]int, 2)
			newCoord[stableAxis] = start[stableAxis]
			newCoord[movingAxis] = start[movingAxis] + i
			lineCoords = append(lineCoords, newCoord)
		}
	}
	if start[movingAxis] > end[movingAxis] {
		for i := 0; i <= start[movingAxis]-end[movingAxis]; i++ {
			newCoord := make([]int, 2)
			newCoord[stableAxis] = end[stableAxis]
			newCoord[movingAxis] = end[movingAxis] + i
			lineCoords = append(lineCoords, newCoord)
		}
	}
	return lineCoords
}

func getLineCoords2(start []int, end []int) [][]int {

	movingAxis := 0
	stableAxis := 1
	lineCoords := [][]int{}
	if start[1] != end[1] {
		movingAxis = 1
		stableAxis = 0
	}

	diff := start[movingAxis] - end[movingAxis]
	if diff < 0 {
		diff = end[movingAxis] - start[movingAxis]
	}

	for i := 0; i <= diff; i++ {
		newCoord := make([]int, 2)
		if start[movingAxis] < end[movingAxis] {
			newCoord[movingAxis] = start[movingAxis] + i
		} else if start[movingAxis] > end[movingAxis] {
			newCoord[movingAxis] = start[movingAxis] - i
		} else {
			newCoord[movingAxis] = start[movingAxis]
		}
		newCoord[stableAxis] = start[stableAxis]

		if start[stableAxis] < end[stableAxis] {
			newCoord[stableAxis] = start[stableAxis] + i
		} else if start[stableAxis] > end[stableAxis] {
			newCoord[stableAxis] = start[stableAxis] - i
		} else {
			newCoord[stableAxis] = start[stableAxis]
		}
		lineCoords = append(lineCoords, newCoord)
	}

	return lineCoords
}
