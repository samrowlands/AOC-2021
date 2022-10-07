package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	println(partOne("input.txt"))
}

func partOne(path string) int {
	matrix := readInput(path)
	lowPoints := computeLowPoints(matrix)
	return sumRiskLevels(lowPoints)
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func readInput(path string) matrix {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matrix := matrix{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strSlice := strings.Split(scanner.Text(), "")
		intSlice := []int{}
		for _, s := range strSlice {
			integer, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			intSlice = append(intSlice, integer)
		}
		matrix = append(matrix, intSlice)
	}
	return matrix
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// func printMatrix(matrix matrix) {
// 	for i := range matrix {
// 		for _, j := range matrix[i] {
// 			print(j)
// 		}
// 		println()
// 	}
// }

func computeLowPoints(matrix matrix) []int {
	lowPoints := []int{}

	colLen := len(matrix[0])
	rowLen := len(matrix)
	for rowIndex := range matrix {
		for colIndex, value := range matrix[rowIndex] {
			if matrix.isLowPoint(rowIndex, colIndex, colLen, rowLen) {
				lowPoints = append(lowPoints, value)
			}
		}
	}
	return lowPoints
}

func isEdge(rowIndex int, colIndex int, rowLen int, colLen int) bool {
	if rowIndex == 0 || rowIndex == (colLen-1) || colIndex == 0 || colIndex == (rowLen-1) {
		return true
	}
	return false
}

func isCorner(rowIndex int, colIndex int, rowLen int, colLen int) bool {
	if rowIndex == 0 && (colIndex == 0 || colIndex == (rowLen-1)) || rowIndex == colLen-1 && (colIndex == 0 || colIndex == (rowLen-1)) {
		return true
	}
	return false
}

type matrix [][]int

func (m matrix) isLowPoint(rowIndex int, colIndex int, rowLen int, colLen int) bool {
	current := m[rowIndex][colIndex]
	var left int
	var right int
	var below int
	var above int
	if colIndex > 0 {
		left = m[rowIndex][colIndex-1]
	}
	if rowIndex < colLen-1 {
		below = m[rowIndex+1][colIndex]
	}
	if rowIndex > 0 {
		above = m[rowIndex-1][colIndex]
	}
	if colIndex < rowLen-1 {
		right = m[rowIndex][colIndex+1]
	}
	if isCorner(rowIndex, colIndex, rowLen, colLen) {
		if rowIndex == 0 && colIndex == 0 {
			if right > current && below > current {
				return true
			}
		}
		if rowIndex == 0 && colIndex == rowLen-1 {
			if left > current && below > current {
				return true
			}
		}
		if rowIndex == colLen-1 && colIndex == 0 {
			if right > current && above > current {
				return true
			}
		}
		if rowIndex == colLen-1 && colIndex == rowLen-1 {
			if above > current && left > current {
				return true
			}
		}
	} else if isEdge(rowIndex, colIndex, rowLen, colLen) {
		if rowIndex == 0 {
			if left > current && right > current && below > current {
				return true
			}
		}
		if rowIndex == colLen-1 {
			if left > current && right > current && above > current {
				return true
			}
		}
		if colIndex == 0 {
			if above > current && right > current && below > current {
				return true
			}
		}
		if colIndex == rowLen-1 {
			if above > current && left > current && below > current {
				return true
			}
		}
	} else {
		if left > current && right > current && above > current && below > current {
			return true
		}
	}
	return false
}

func sumRiskLevels(lowPoints []int) int {
	sum := 0
	for _, n := range lowPoints {
		sum += n + 1
	}
	return sum
}

//
