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
	score := 0
	luckyNumbers, bingoBoards := processInput(input)
game:
	for count, luckyNumber := range luckyNumbers {
		// f.Printf("Lucky number: %d\n", luckyNumber)
		for _, bingoBoard := range bingoBoards {
			for row, bingoRow := range bingoBoard {
				for col, bingoNum := range bingoRow {
					if bingoNum.num == luckyNumber && bingoNum.marked == false {
						bingoRow[col].marked = true
						// printBoard(bingoBoard)
						if count >= 5 {
							if checkIfBingo(bingoBoard, row, col) {
								score = calculateScore(bingoBoard, luckyNumber)
								break game
							}
						}
					}
				}
			}
		}
	}
	return score
}

func solution02(input []string) int {
	score := 0
	luckyNumbers, bingoBoards := processInput(input)
	winningBoards := make([]bool, len(bingoBoards))
	winningBoardIndex := 0
game:
	for count, luckyNumber := range luckyNumbers {
		// f.Printf("Lucky number: %d\n", luckyNumber)
		for boardIndex, bingoBoard := range bingoBoards {
			if winningBoards[boardIndex] {
				continue
			}
			for row, bingoRow := range bingoBoard {
				for col, bingoNum := range bingoRow {
					if bingoNum.num == luckyNumber && bingoNum.marked == false {
						bingoRow[col].marked = true
						// printBoard(bingoBoard)
						if count >= 5 {
							if checkIfBingo(bingoBoard, row, col) {
								score = calculateScore(bingoBoard, luckyNumber)
								winningBoards[boardIndex] = true
								winningBoardIndex++
								if winningBoardIndex == len(bingoBoards) {
									break game
								}
							}
						}
					}
				}
			}
		}
	}
	return score
}

type bingoCell struct {
	num    int
	marked bool
}

func checkIfBingo(bingoBoard [][]bingoCell, row int, col int) bool {
	for _, row := range bingoBoard {
		markedCellsInRows := 0
		for _, cell := range row {
			if cell.marked == true {
				markedCellsInRows++
			}
		}
		if markedCellsInRows == len(row) {
			return true
		}
	}
	for i := 0; i < len(bingoBoard); i++ {
		markedCellsInColumns := 0
		for j := 0; j < len(bingoBoard[i]); j++ {
			cell := bingoBoard[j][i]
			if cell.marked == true {
				markedCellsInColumns++
			}
		}
		if markedCellsInColumns == len(bingoBoard) {
			return true
		}
	}
	return false
}

func calculateScore(bingoBoard [][]bingoCell, luckyNumber int) int {
	sumOfUnmarkedNums := 0
	for i := 0; i < len(bingoBoard); i++ {
		for j := 0; j < len(bingoBoard[i]); j++ {
			cell := bingoBoard[i][j]
			if cell.marked == false {
				sumOfUnmarkedNums += cell.num
			}
		}
	}
	return sumOfUnmarkedNums * luckyNumber
}

func processInput(input []string) ([]int, [][][]bingoCell) {
	luckyNumStr := strings.Split(input[0], ",")
	luckyNumbers := make([]int, len(luckyNumStr))
	for i, str := range luckyNumStr {
		luckyNumbers[i], _ = strconv.Atoi((str))
	}
	bingoBoards := make([][][]bingoCell, 0)
	currentBoard := make([][]bingoCell, 5)
	currentBoardIndex := 0
	for i := 2; i <= len(input); i++ {
		if i >= len(input) || input[i] == "" {
			bingoBoards = append(bingoBoards, currentBoard)
			currentBoard = make([][]bingoCell, 5)
			currentBoardIndex = 0
			continue
		}
		splitRow := strings.Split(input[i], " ")
		currentRow := make([]bingoCell, 0)
		for _, str := range splitRow {
			if str == "" {
				continue
			}
			num, _ := strconv.Atoi(str)
			currentRow = append(currentRow, bingoCell{num, false})
		}
		currentBoard[currentBoardIndex] = currentRow
		currentBoardIndex++
	}

	return luckyNumbers, bingoBoards
}

func printBoard(bingoBoard [][]bingoCell) {
	f.Println("_____________________________________")
	for i := 0; i < len(bingoBoard); i++ {
		for j := 0; j < len(bingoBoard[i]); j++ {
			cell := bingoBoard[i][j]
			mark := ""
			if cell.marked {
				mark = "❌"
			}
			f.Printf("%d%v\t", cell.num, mark)
		}
		f.Println()
	}
	f.Println("-------------------------------------")
}
