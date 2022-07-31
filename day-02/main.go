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
	partOne("input.txt")
	partTwo("input.txt")
}

type instruction struct {
	Direction string
	Magnitude int
}

func readInput(path string) []instruction {
	data := []instruction{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		//split the line to achieve
		data = append(data, instruction{
			Direction: s[0],
			Magnitude: toInt(s[1]),
		})
	}
	return data
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func partOne(path string) {
	instructions := readInput(path)
	horizontal := 0
	depth := 0
	for _, v := range instructions {
		switch v.Direction {
		case "forward":
			horizontal += v.Magnitude
		case "down":
			depth += v.Magnitude
		case "up":
			depth -= v.Magnitude
		}
	}
	fmt.Println(horizontal * depth)
}

func partTwo(path string) {
	instructions := readInput(path)
	horizontal := 0
	depth := 0
	aim := 0
	for _, v := range instructions {
		switch v.Direction {
		case "forward":
			horizontal += v.Magnitude
			depth += aim * v.Magnitude
		case "down":
			aim += v.Magnitude
		case "up":
			aim -= v.Magnitude
		}
	}
	fmt.Println(horizontal * depth)
}
