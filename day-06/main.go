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
	fish := readInput(path)
	return computeFish(fish, 80)
}

func partTwo(path string) int {
	fish := readInput(path)

	fishMap := make(map[int]int)
	for _, f := range fish {
		fishMap[f]++
	}
	return computeFishV3(fishMap, 256)
}

func readInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := strings.Split(scanner.Text(), ",")

	return strSliceToIntSlice(s)
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

func computeFish(fish []int, loops int) int {
	for i := 0; i < loops; i++ {
		z := zeroCount(fish)
		fish = reduceTimer(fish)
		fish = appendEights(z, fish)
	}
	return len(fish)
}

func zeroCount(fish []int) int {
	count := 0
	for _, f := range fish {
		if f == 0 {
			count++
		}
	}
	return count
}

func appendEights(zeroCount int, fish []int) []int {
	for zeroCount > 0 {
		fish = append(fish, 8)
		zeroCount--
	}
	return fish
}

func reduceTimer(fish []int) []int {
	for i := range fish {
		if fish[i] == 0 {
			fish[i] = 6
		} else {
			fish[i]--
		}
	}
	return fish
}

func computeFishV2(fish []int, loops int) int {
	for i := 0; i < loops; i++ {
		zeros := 0
		n := len(fish)
		for i := 0; i < n; i++ {
			if fish[i] == 0 {
				zeros++
				fish[i] = 6
				fish = append(fish, 8)
			} else {
				fish[i]--
			}
		}
	}
	return len(fish)
}

func computeFishV3(fishMap map[int]int, loops int) int {
	//one needs to be a map from previous day
	//one needs to be a new map
	currentMap := fishMap
	for i := 0; i < loops; i++ {
		newMap := make(map[int]int)
		for k := range currentMap {
			if k == 0 {
				newMap[6] += currentMap[k]
				newMap[8] += currentMap[k]
			}
			if k > 0 {
				newMap[k-1] += currentMap[k]
			}
		}
		currentMap = newMap
	}
	sum := 0
	for _, c := range currentMap {
		sum += c
	}
	return sum
}
