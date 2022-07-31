package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}

func partOne(filePath string) {
	depths := readInput(filePath)
	count := 0
	for i := range depths {
		if i == 0 {
			continue
		}
		if depths[i] > depths[i-1] {
			count++
		}
	}
	fmt.Println(count)
}

func partTwo(filePath string) {
	depths := readInput(filePath)
	count := 0
	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			count++
		}
	}
	fmt.Println(count)
}

func readInput(filePath string) []int {
	var depths []int
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depths = append(depths, toInt(scanner.Text()))
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return depths
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
