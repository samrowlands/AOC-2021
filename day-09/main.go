package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	println(partOne("input.txt"))
	println(partTwo("input.txt"))
}

func partOne(path string) int {
	matrix := readInput(path)
	colLen := len(matrix[0])
	rowLen := len(matrix)
	lowPoints, _ := computeLowPoints(matrix, colLen, rowLen)
	return sumRiskLevels(lowPoints)
}

func partTwo(path string) int {
	matrix := readInput(path)
	colLen := len(matrix[0])
	rowLen := len(matrix)

	_, lowPointsReference := computeLowPoints(matrix, colLen, rowLen)

	basinLengths := []int{}
	for _, lowPointReference := range lowPointsReference {
		basinLengths = append(basinLengths, bfs(matrix, lowPointReference, colLen, rowLen))
	}

	sort.Ints(basinLengths)
	return basinLengths[len(basinLengths)-1] * basinLengths[len(basinLengths)-2] * basinLengths[len(basinLengths)-3]
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

func computeLowPoints(matrix matrix, colLen int, rowLen int) ([]int, []gridReference) {
	lowPoints := []int{}
	lowPointsReference := []gridReference{}

	for rowIndex := range matrix {
		for colIndex, value := range matrix[rowIndex] {
			if matrix.isLowPoint(rowIndex, colIndex, colLen, rowLen) {
				lowPoints = append(lowPoints, value)
				lowPointsReference = append(lowPointsReference, gridReference{rowIndex: rowIndex, columnIndex: colIndex})
			}
		}
	}
	return lowPoints, lowPointsReference
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

//ok so we need to find the starting coords for the DFS
//i.e. the coords of the low points
//do DFS
//store basin in slice of basins?

// execute a dfs per basis

// need the row, col and the grid?
// func dfs(n int) {
//keep track of visited nodes
//we to store a node struct with coords and value in visited nodes (probs don't even need to store value)

//one option: compute the adjacency list for the starting low point?
//could be tricky
//second option: use breadth-first search
//a dfs up to the 9 boundary, how?
// }

//lets do bfs bc we don't have to build an adjaceny list this way:

func bfs(matrix matrix, lowPoint gridReference, colLen int, rowLen int) int {
	queue := []gridReference{lowPoint}
	visited := []gridReference{}
	//initialise a 'queue' i.e. a slice and populate with low point grid reference.

	//while the queue is not empty
	for len(queue) != 0 {
		//get the first reference in the queue store it and remove it from the queue
		cell := queue[0]
		queue = queue[1:]

		//if this gridReference has been visited continue the loop
		if hasVisited(visited, cell) {
			continue
		}
		visited = append(visited, cell)

		//for each neighbour of the gridReference.
		for _, neighbour := range neighbours(cell, colLen, rowLen) {
			row := neighbour.rowIndex
			col := neighbour.columnIndex
			//check its not a 9 and check it hasn't already been visited
			//if these are true add it to the queue
			//add the reference to visited (another slice of gridReferences)
			if matrix[row][col] != 9 && !hasVisited(visited, neighbour) {
				queue = append(queue, neighbour)
			}
		}
	}
	return len(visited)
}

type gridReference struct {
	rowIndex    int
	columnIndex int
}

func hasVisited(visited []gridReference, cell gridReference) bool {
	for _, ref := range visited {
		if ref.columnIndex == cell.columnIndex && ref.rowIndex == cell.rowIndex {
			return true
		}
	}
	return false
}

//neighbours func, compute slice of grid references

func neighbours(cell gridReference, colLen int, rowLen int) []gridReference {
	deltas := []gridReference{
		{rowIndex: 1, columnIndex: 0},
		{rowIndex: -1, columnIndex: 0},
		{rowIndex: 0, columnIndex: 1},
		{rowIndex: 0, columnIndex: -1},
	}
	neighbourList := []gridReference{}

	for _, g := range deltas {
		c := cell.columnIndex + g.columnIndex
		r := cell.rowIndex + g.rowIndex
		if 0 <= c && c < colLen && 0 <= r && r < rowLen {
			neighbourList = append(neighbourList, gridReference{columnIndex: c, rowIndex: r})
		}
	}
	return neighbourList
}
