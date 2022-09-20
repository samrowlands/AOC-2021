package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	println(partOne("input.txt"))
	println(partTwo("input.txt"))
}

func partOne(path string) int {
	numbers, boards := readInput(path)
	winningBoards := findWinningBoards(numbers, boards)
	productOfUnmarked := winningBoards[0].unmarkedNumbers(numbers)

	return winningBoards[0].lastCalledNumber * productOfUnmarked
}

func partTwo(path string) int {
	numbers, boards := readInput(path)
	lastWinningBoard := lastWinningBoard(numbers, boards)
	productOfUnmarked := lastWinningBoard.unmarkedNumbers(numbers)

	return lastWinningBoard.lastCalledNumber * productOfUnmarked
}

type board struct {
	matrix           [][]int
	markedNumbers    []int
	lastCalledNumber int
	unmarked         []int
}

func readInput(path string) ([]int, []*board) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbers := strSliceToIntSlice(strings.Split(scanner.Text(), ","))

	boards := []*board{}
	isEmpty := false
	for scanner.Scan() {
		inputRow := strSliceToIntSlice(strings.Fields(scanner.Text()))
		if len(inputRow) == 0 {
			isEmpty = true
			continue
		}
		if len(inputRow) != 0 && isEmpty {
			currentMatrix := [][]int{}
			currentMatrix = append(currentMatrix, inputRow)

			for i := 0; i < 4; i++ {
				scanner.Scan()
				currentMatrix = append(currentMatrix, strSliceToIntSlice(strings.Fields(scanner.Text())))
			}
			boards = append(boards, &board{matrix: currentMatrix})
		}
	}
	return numbers, boards
}

func strSliceToIntSlice(slice []string) []int {
	numbers := []int{}
	for _, str := range slice {
		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func findWinningBoards(numbers []int, boards []*board) []*board {
	winningBoards := []*board{}
	for _, n := range numbers {
		for _, b := range boards {
			if b.numberExistsOnBoard(n) {
				b.markedNumbers = append(b.markedNumbers, n)
			}
			if len(b.markedNumbers) >= 5 {
				if b.hasWon() {
					b.lastCalledNumber = n
					winningBoards = append(winningBoards, b)
					return winningBoards
				}
			}
		}
	}
	return nil
}

func (b board) numberExistsOnBoard(n int) bool {
	for i := range b.matrix {
		for j := range b.matrix[i] {
			if b.matrix[i][j] == n {
				return true
			}
		}
	}
	return false
}

func (b board) hasWon() bool {
	if b.checkRows() || b.checkCols() {
		return true
	}
	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (b board) checkRows() bool {
	for _, row := range b.matrix {
		matchedNumberCount := 0
		for j := range row {
			if !contains(b.markedNumbers, row[j]) {
				break
			}
			matchedNumberCount++
			if matchedNumberCount == 5 {
				return true
			}
		}
	}
	return false
}

func (b board) checkCols() bool {
	for i := 0; i < 5; i++ {
		matchedNumberCount := 0
		for j := 0; j < 5; j++ {
			if !contains(b.markedNumbers, b.matrix[j][i]) {
				break
			}
			matchedNumberCount++
			if matchedNumberCount == 5 {
				return true
			}
		}
	}
	return false
}

func (b board) unmarkedNumbers(numbers []int) int {
	sum := 0
	for _, row := range b.matrix {
		for i := range row {
			if !contains(b.markedNumbers, row[i]) {
				sum += row[i]
			}
		}
	}
	return sum
}

func (b board) printMatrix() {
	for i := range b.matrix {
		for j := range b.matrix[i] {
			if j < 4 {
				fmt.Print(b.matrix[i][j], " ")
			}
			if j == 4 {
				println(b.matrix[i][j])
			}
		}
	}
}

func lastWinningBoard(numbers []int, boards []*board) *board {
	winningBoards := []*board{}
	for _, n := range numbers {
		for _, b := range boards {
			if containsBoard(winningBoards, b) {
				continue
			}
			if b.numberExistsOnBoard(n) {
				b.markedNumbers = append(b.markedNumbers, n)
			} else {
				b.unmarked = append(b.unmarked, n)
			}
			if len(b.markedNumbers) >= 5 {
				if b.hasWon() {
					b.lastCalledNumber = n
					winningBoards = append(winningBoards, b)
				}
			}
		}
	}
	return winningBoards[len(winningBoards)-1]
}

func containsBoard(s []*board, e *board) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
