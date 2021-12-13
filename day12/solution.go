package main

import (
	f "fmt"
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

type cave struct {
	name  string
	size  string
	links map[string]*cave
}

func mapCaves(input []string) *cave {
	caveMap := map[string]*cave{}
	for _, row := range input {
		passage := strings.Split(row, "-")
		cave1 := passage[0]
		cave2 := passage[1]
		if _, ok := caveMap[cave1]; !ok {
			caveMap[cave1] = createCave(cave1)
		}
		if _, ok := caveMap[cave2]; !ok {
			caveMap[cave2] = createCave(cave2)
		}
		cave1Ref := caveMap[cave1]
		cave2Ref := caveMap[cave2]
		cave1Ref.links[cave2] = caveMap[cave2]
		cave2Ref.links[cave1] = caveMap[cave1]
	}
	return caveMap["start"]
}

func createCave(name string) *cave {
	size := "small"
	if name == strings.ToUpper(name) {
		size = "large"
	}
	return &cave{
		name:  name,
		size:  size,
		links: make(map[string]*cave),
	}
}

func isCaveAlreadyVisited(caveInQuestion *cave, path []string) bool {
	for _, visitedCave := range path {
		if visitedCave == caveInQuestion.name {
			return true
		}
	}
	return false
}

func printCave(caveToPrint *cave) {
	linkNames := make([]string, 0)
	for link := range caveToPrint.links {
		linkNames = append(linkNames, link)
	}
	sizeSymbol := ""
	if caveToPrint.size == "large" {
		sizeSymbol = "^"
	}
	f.Printf("[%s]%s --> %v\n", caveToPrint.name, sizeSymbol, linkNames)
}

func solution01(input []string) int {
	startCave := mapCaves(input)
	return traverseCaves01(startCave, []string{})
}

func traverseCaves01(currentCave *cave, path []string) int {
	path = append(path, currentCave.name)
	if currentCave.name == "end" {
		return 1
	}
	pathsToTake := 0
	for _, linkedCave := range currentCave.links {
		if linkedCave.size == "small" && isCaveAlreadyVisited(linkedCave, path) {
			continue
		}
		pathsToTake += traverseCaves01(linkedCave, path)
	}
	return pathsToTake
}

func solution02(input []string) int {
	startCave := mapCaves(input)
	return traverseCaves02(startCave, []string{})
}

func traverseCaves02(currentCave *cave, path []string) int {
	path = append(path, currentCave.name)
	if currentCave.name == "end" {
		return 1
	}
	pathsToTake := 0
	for _, linkedCave := range currentCave.links {
		if !canCaveBeVisited(linkedCave, path) {
			continue
		}
		pathsToTake += traverseCaves02(linkedCave, path)
	}
	return pathsToTake
}

func canCaveBeVisited(caveInQuestion *cave, path []string) bool {
	if caveInQuestion.size == "large" {
		return true
	}

	if caveInQuestion.size == "small" {
		if !isCaveAlreadyVisited(caveInQuestion, path) {
			return true
		}
		if caveInQuestion.name != "start" && caveInQuestion.name != "end" && !isAnySmallCaveVisitedTwiceBefore(path) {
			return true
		}
	}
	return false
}

func isAnySmallCaveVisitedTwiceBefore(path []string) bool {
	visitedMap := map[string]bool{}
	for _, cave := range path {
		if strings.ToLower(cave) == cave {
			if visitedMap[cave] == true {
				return true
			}
			visitedMap[cave] = true
		}
	}
	return false
}
