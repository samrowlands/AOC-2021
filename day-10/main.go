package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func calculateSyntaxErrorScore(input []string) int {
	sMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	bMap := map[rune]rune{
		'}': '{',
		']': '[',
		'>': '<',
		')': '(',
	}
	open := []rune{}
	// close := []rune{}
	score := 0

OUTER:
	for _, str := range input {
		for _, b := range str {
			//if open bracket, append to open
			if b == '{' || b == '[' || b == '(' || b == '<' {
				open = append(open, b)
			} else {
				//if close bracket doesn't equal latest open then add to score
				if bMap[b] != open[len(open)-1] {
					score += sMap[b]
					continue OUTER
					//if it does, delete latest from open
				} else if bMap[b] == open[len(open)-1] {
					open = open[:len(open)-1]
				}
			}
		}
	}
	return score
}

func autoCompleteScore(input []string) int {
	bMap := map[rune]rune{
		'}': '{',
		']': '[',
		'>': '<',
		')': '(',
	}
	sMap := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	//discard all corrupted lines
	incomplete := []string{}

	//construct incomplete list
OUTER:
	for _, str := range input {
		open := []rune{}
		for _, b := range str {
			//for each character in current string,
			//if bracket is open append to open
			if b == '{' || b == '[' || b == '(' || b == '<' {
				open = append(open, b)
			} else {
				//if the current bracket is closed and doesn't equal latest, then continue
				if bMap[b] != open[len(open)-1] {
					continue OUTER
				} else if bMap[b] == open[len(open)-1] {
					open = open[:len(open)-1]
				}
			}
		}
		incomplete = append(incomplete, str)
	}

	//
	scores := []int{}
	for _, str := range incomplete {
		open := []rune{}
		for _, b := range str {
			if b == '{' || b == '[' || b == '(' || b == '<' {
				open = append(open, b)
			} else {
				if bMap[b] == open[len(open)-1] {
					open = open[:len(open)-1] //pop
				} else {
					log.Fatal("corrupt string")
				}
			}
		}
		score := 0

		for i := len(open) - 1; i > -1; i-- {
			score = (score * 5) + sMap[open[i]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func partOne(path string) int {
	input := readInput(path)
	return calculateSyntaxErrorScore(input)
}

func partTwo(path string) int {
	input := readInput(path)
	return autoCompleteScore(input)
}

func main() {
	println(partOne("input.txt"))
	println(partTwo("input.txt"))
}

func reverseSlice(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
