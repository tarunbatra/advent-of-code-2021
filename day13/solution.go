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
	f.Printf("Part 2: \n%s \n(in %v)\n", result2, time.Since((t2)))
}

func solution01(input []string) int {
	points, folds := parseInput(input)
	newPoints := points
	for i := 0; i < 1; i++ {
		newPoints = getPointsAfterFolding(newPoints, folds[i])
	}
	return len(newPoints)
}

func solution02(input []string) string {
	points, folds := parseInput(input)
	newPoints := points
	for i := 0; i < len(folds); i++ {
		newPoints = getPointsAfterFolding(newPoints, folds[i])
	}
	return printPaper(newPoints)
}

type point struct {
	x int
	y int
}

func (p *point) str() string {
	return f.Sprint(p.x, ",", p.y)
}

type foldingInstruction struct {
	axis  string
	point int
}

func parseInput(input []string) (map[string]point, []foldingInstruction) {
	foldInstruction := false
	points := map[string]point{}
	folds := []foldingInstruction{}
	for _, row := range input {
		if row == "" {
			foldInstruction = true
			continue
		}

		if foldInstruction {
			instructionStr := strings.TrimPrefix(row, "fold along ")
			instruction := strings.Split(instructionStr, "=")
			axis := instruction[0]
			point, _ := strconv.Atoi(instruction[1])
			newFold := foldingInstruction{axis, point}
			folds = append(folds, newFold)
		} else {
			pointStr := strings.Split(row, ",")
			x, _ := strconv.Atoi(string(pointStr[0]))
			y, _ := strconv.Atoi(string(pointStr[1]))
			p := point{x, y}
			points[p.str()] = p
		}
	}
	return points, folds
}

func getPointsAfterFolding(points map[string]point, fold foldingInstruction) map[string]point {
	newPoints := map[string]point{}
	for _, p := range points {
		newPoint := foldPoint(p, fold)
		newPoints[newPoint.str()] = newPoint
	}
	return newPoints
}

func foldPoint(p point, fold foldingInstruction) point {
	if fold.axis == "y" && p.y > fold.point {
		newY := p.y - 2*(p.y-fold.point)
		return point{p.x, newY}
	}
	if fold.axis == "x" && p.x > fold.point {
		newX := p.x - 2*(p.x-fold.point)
		return point{newX, p.y}
	}
	return p
}

func printPaper(points map[string]point) string {
	leftBound := 99999
	rightBound := 0
	topBound := 99999
	bottomBound := 0
	for _, p := range points {
		if p.x < leftBound {
			leftBound = p.x
		}
		if p.x > rightBound {
			rightBound = p.x
		}
		if p.y > bottomBound {
			bottomBound = p.y
		}
		if p.y < topBound {
			topBound = p.y
		}
	}

	paper := ""
	for i := topBound; i <= bottomBound; i++ {
		for j := leftBound; j <= rightBound; j++ {
			pointRef := f.Sprint(j, ",", i)
			if _, ok := points[pointRef]; ok {
				paper += "#"
			} else {
				paper += " "
			}
		}
		paper += "\n"
	}
	return paper
}
