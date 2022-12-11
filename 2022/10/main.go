package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cycles struct {
	Count []int
}

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]
	cycle := new(Cycles)
	cycle.Count = make([]int, 1)
	cycle.Count[0] = 1
	part1Answer := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "noop" {
			cycle.Count = append(cycle.Count, cycle.Count[len(cycle.Count)-1])
		} else {
			number, _ := strconv.Atoi(lines[i][5:])
			cycle.Count = append(cycle.Count, cycle.Count[len(cycle.Count)-1])
			cycle.Count = append(cycle.Count, cycle.Count[len(cycle.Count)-1]+number)
		}
	}

	for i := 20; i <= 220; i = i + 40 {
		part1Answer += cycle.Count[i-1] * i
	}

	fmt.Println(part1Answer)
	cycle.drawScreen()
}

func (c Cycles) drawScreen() {
	crt := map[int][]string{}
	column := 0
	pixel := 0
	for i := 0; i < len(c.Count); i++ {
		if abs(c.Count[i]-pixel) <=1 {
			crt[column] = append(crt[column], "")
		} else {
			crt[column] = append(crt[column], " ")
		}
		pixel++
		if pixel == 40 {
			column++
			pixel = 0
		}
	}
	for i := 0; i < 6; i++ {
		println(strings.Join(crt[i], ""))
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
