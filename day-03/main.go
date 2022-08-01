package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}

func readInput(path string) []string {
	data := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func strBinToInt(binary string) int64 {
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func partOne(filepath string) {
	data := readInput(filepath)
	gammaRate := ""
	epsilonRate := ""
	for i := 0; i < len(data[0]); i++ {
		zeroCount := 0
		for _, v := range data {
			if v[i] == '0' {
				zeroCount++
			}
		}
		if zeroCount > len(data)/2 {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}
	println(strBinToInt(gammaRate) * strBinToInt(epsilonRate))
}

func partTwo(filepath string) {
	data := readInput(filepath)
	for i := 0; len(data) > 1; i++ {
		zeroCount := 0.0
		for _, v := range data {
			if v[i] == '0' {
				zeroCount++
			}
		}
		if zeroCount > float64(len(data))/2.0 {
			data = filter(data, '0', i)
		} else {
			data = filter(data, '1', i)
		}
	}
	oxygen := data[0]
	data = readInput(filepath)
	for i := 0; len(data) > 1; i++ {
		zeroCount := 0.0
		for _, v := range data {
			if v[i] == '0' {
				zeroCount++
			}
		}
		if zeroCount > float64(len(data))/2.0 {
			data = filter(data, '1', i)
		} else {
			data = filter(data, '0', i)
		}
	}
	co2 := data[0]
	println(strBinToInt(oxygen) * strBinToInt(co2))
}

func filter(array []string, number byte, charIndex int) []string {
	filtered := []string{}
	for _, v := range array {
		if v[charIndex] == number {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
