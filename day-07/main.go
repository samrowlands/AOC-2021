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
	println(partTwo("input.txt"))
}

func partOne(path string) int {
	positions := readInput(path)
	return fuel(positions)
}

func partTwo(path string) int {
	positions := readInput(path)
	return fuel2(positions)
}

func readInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return strSliceToIntSlice(strings.Split(scanner.Text(), ","))
}

func strSliceToIntSlice(slice []string) []int {
	intSlice := []int{}
	for _, s := range slice {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, n)
	}
	return intSlice
}

func fuel(positions []int) int {
	n := largestInt(positions)
	fuel := 10000000000
	for i := 0; i < n; i++ {
		fc := fuelConsumption(i, positions)
		if fc < fuel {
			fuel = fc
		}
	}
	return fuel
}

func fuel2(positions []int) int {
	n := largestInt(positions)
	fuel := 10000000000
	for i := 0; i < n; i++ {
		fc := fuelConsumptionV2(i, positions)
		if fc < fuel {
			fuel = fc
		}
	}
	return fuel
}

func largestInt(positions []int) int {
	l := 0
	for _, p := range positions {
		if p > l {
			l = p
		}
	}
	return l
}

func fuelConsumption(i int, positions []int) int {
	sum := 0
	for _, p := range positions {
		sum += Abs(p - i)
	}
	return sum
}

func fuelConsumptionV2(x int, positions []int) int {
	sum := 0
	for _, p := range positions {
		increment := 0
		for i := 0; i < Abs(p-x); i++ {
			increment++
			sum += increment
		}
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
