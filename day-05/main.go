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
	lines := readInput(path, false)
	return overlaps(lines)
}

func partTwo(path string) int {
	lines := readInput(path, true)
	return overlaps(lines)
}

func readInput(path string, diagonals bool) []line {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []line{}

	for scanner.Scan() {
		x1, y1, x2, y2 := coordinates(scanner.Text())
		if lineIsValid(x1, y1, x2, y2, diagonals) {
			l := line{
				x1: x1,
				y1: y1,
				x2: x2,
				y2: y2,
			}
			l.addCoordinates(diagonals)
			lines = append(lines, l)
		}
	}
	return lines
}

type line struct {
	x1          int
	y1          int
	x2          int
	y2          int
	coordinates []coordinate
}

type coordinate struct {
	x     int
	y     int
	count int
}

func coordinates(text string) (int, int, int, int) {
	xy1 := strings.Split(strings.Fields(text)[0], ",")
	xy2 := strings.Split(strings.Fields(text)[2], ",")

	x1, _ := strconv.Atoi(xy1[0])
	y1, _ := strconv.Atoi(xy1[1])

	x2, _ := strconv.Atoi(xy2[0])
	y2, _ := strconv.Atoi(xy2[1])

	return x1, y1, x2, y2
}

func lineIsValid(x1, y1, x2, y2 int, diagonals bool) bool {
	if x1 == x2 || y1 == y2 {
		return true
	}
	if diagonals && (Abs(y2-y1) == Abs(x2-x1)) {
		return true
	}
	return false
}

//A distinct grid (as our benchmark)

func (l *line) addCoordinates(diagonals bool) {
	c := []coordinate{}
	//if vertical
	if l.x1 == l.x2 {
		if l.y1 == l.y2+1 {
			c = append(c, coordinate{x: l.x1, y: l.y1})
			c = append(c, coordinate{x: l.x2, y: l.y2})
		}
		if l.y1 > l.y2+1 {
			c = append(c, coordinate{x: l.x1, y: l.y1})
			c = append(c, coordinate{x: l.x2, y: l.y2})
			for i := l.y2 + 1; i < l.y1; i++ {
				c = append(c, coordinate{x: l.x1, y: i})
			}
		}
		if l.y2 > l.y1+1 {
			c = append(c, coordinate{x: l.x1, y: l.y1})
			c = append(c, coordinate{x: l.x2, y: l.y2})
			for i := l.y1 + 1; i < l.y2; i++ {
				c = append(c, coordinate{x: l.x1, y: i})
			}
		}
	}

	//if horizontal
	if l.y1 == l.y2 {
		if l.x1 == l.x2+1 {
			c = append(c, coordinate{x: l.x1, y: l.y1})
			c = append(c, coordinate{x: l.x2, y: l.y2})
		}
		if l.x1 > l.x2+1 {
			c = append(c, coordinate{x: l.x1, y: l.y1})
			c = append(c, coordinate{x: l.x2, y: l.y2})
			for i := l.x2 + 1; i < l.x1; i++ {
				c = append(c, coordinate{x: i, y: l.y1})
			}
		}
		if l.x2 > l.x1+1 {
			c = append(c, coordinate{x: l.x1, y: l.y1})
			c = append(c, coordinate{x: l.x2, y: l.y2})
			for i := l.x1 + 1; i < l.x2; i++ {
				c = append(c, coordinate{x: i, y: l.y1})
			}
		}
	}
	if diagonals {
		//if diagonal
		if Abs(l.y2-l.y1) == Abs(l.x2-l.x1) {

			//need to think this through better:
			//if negative gradient
			if (l.y2-l.y1)/(l.x2-l.x1) == -1 {
				if l.x1 < l.x2 {
					for i := l.x1; i <= l.x2; i++ {
						count := i - l.x1
						c = append(c, coordinate{x: i, y: l.y1 - count})
					}
				} else {
					for i := l.x2; i <= l.x1; i++ {
						count := i - l.x2
						c = append(c, coordinate{x: i, y: l.y2 - count})
					}
				}
			}
			//if positive gradient
			if (l.y2-l.y1)/(l.x2-l.x1) == 1 {
				if l.x1 < l.x2 {
					for i := l.x1; i <= l.x2; i++ {
						count := i - l.x1
						c = append(c, coordinate{x: i, y: l.y1 + count})
					}
				} else {
					for i := l.x2; i <= l.x1; i++ {
						count := i - l.x2
						c = append(c, coordinate{x: i, y: l.y2 + count})
					}
				}
			}
		}
	}

	l.coordinates = c
}

func overlaps(lines []line) int {
	distinctXY := []*coordinate{}
	for _, l := range lines {
		for _, c := range l.coordinates {
			coord := c
			if !contains(distinctXY, &coord) {
				coord.count = 1
				distinctXY = append(distinctXY, &coord)
			} else {
				distinctXY = incrementCount(distinctXY, &coord)
			}
		}
	}
	return countOverlaps(distinctXY)
}

func contains(cs []*coordinate, c *coordinate) bool {
	for _, a := range cs {
		if a.x == c.x && a.y == c.y {
			return true
		}
	}
	return false
}

func countOverlaps(cs []*coordinate) int {
	count := 0
	for _, c := range cs {
		if c.count >= 2 {
			count++
		}
	}
	return count
}

func incrementCount(cs []*coordinate, c *coordinate) []*coordinate {
	for i, a := range cs {
		if a.x == c.x && a.y == c.y {
			cs[i].count++
		}
	}
	return cs
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
