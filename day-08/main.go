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
	output := readInput(path)
	return findUnique(output)
}

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := []string{}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "|")
		output = append(output, strings.Fields(input[1])...)
	}
	return output
}

func findUnique(output []string) int {
	count := 0
	for _, s := range output {
		switch len(s) {
		case 2:
			count++
		case 3:
			count++
		case 4:
			count++
		case 7:
			count++
		}
	}
	return count
}

type entry struct {
	patterns []string
	output   []string
}

func partTwo(path string) int {
	entries := readInputV2(path)
	return sumAllEntries(entries)
}

func readInputV2(path string) []entry {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entries := []entry{}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "|")
		entry := entry{
			patterns: strings.Fields(input[0]),
			output:   strings.Fields(input[1]),
		}
		entries = append(entries, entry)
	}
	return entries
}

func alreadyIn(s string, b map[string]int) bool {
	for k := range b {
		if k == s {
			return true
		}
	}
	return false
}

func appearsIn(r rune, s string) bool {
	for _, l := range s {
		if l == r {
			return true
		}
	}
	return false
}

func numberOfCommonLetters(a, b string) int {
	cnt := 0
	for _, la := range a {
		if appearsIn(la, b) {
			cnt++
		}
	}
	return cnt
}

func reverseMap(m map[string]int) map[int]string {
	n := make(map[int]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func deduceEntry(patterns []string) map[string]int {
	p2d := map[string]int{}

	//Establish 1,4,7,8
	for _, s := range patterns {
		switch len(s) {
		case 2:
			p2d[s] = 1
		case 4:
			p2d[s] = 4
		case 3:
			p2d[s] = 7
		case 7:
			p2d[s] = 8
		}
	}

	d2p := reverseMap(p2d)

	//Establish 2,3,5
	for _, s := range patterns {
		if alreadyIn(s, p2d) {
			continue
		}
		if len(s) == 5 {
			if numberOfCommonLetters(s, d2p[1]) == 2 {
				p2d[s] = 3
			} else if numberOfCommonLetters(s, d2p[4]) == 3 {
				p2d[s] = 5
			} else {
				p2d[s] = 2
			}
		}
	}
	//Establish 0,6,9
	for _, s := range patterns {
		if alreadyIn(s, p2d) {
			continue
		}
		if len(s) == 6 {
			if numberOfCommonLetters(s, d2p[4]) == 4 {
				p2d[s] = 9
			} else if numberOfCommonLetters(s, d2p[1]) == 1 {
				p2d[s] = 6
			} else {
				p2d[s] = 0
			}
		}
	}
	return p2d
}

func equal(s1, s2 string) bool {
	if len(s1) == len(s2) {
		for _, l1 := range s1 {
			if !appearsIn(l1, s2) {
				return false
			}
		}
		return true
	}
	return false
}

func decodeOutput(m map[string]int, output []string) int {
	//need a way to compare patterns in my map with outputs, independent of order
	sum := ""
outer:
	for _, o := range output {
		for p := range m {
			if equal(o, p) {
				sum += strconv.Itoa(m[p])
				continue outer
			}
		}
	}
	sumInt, err := strconv.Atoi(sum)
	if err != nil {
		log.Fatal(err)
	}
	return sumInt
}

func sumEntry(entry entry) int {
	p2d := deduceEntry(entry.patterns)
	return decodeOutput(p2d, entry.output)
}

func sumAllEntries(entries []entry) int {
	sum := 0
	for _, entry := range entries {
		sum += sumEntry(entry)
	}
	return sum
}
