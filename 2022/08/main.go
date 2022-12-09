package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]
	edges := len(lines)*2 + len(lines[0])*2 - 4
	part1Answer := edges
	part2Answer := 0

	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[row])-1; col++ {
			num := getSize(lines[row][col])
			curScore, yes := checkAround(num, row, col, lines)
			if yes {
				part1Answer++
			}
			if curScore > part2Answer {
				part2Answer = curScore
			}
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func getSize(b byte) int {
	size, _ := strconv.Atoi(string(b))
	return size
}

func checkLeft(num, row, col int, lines []string) (int, bool) {
	score := 0
	visible := true
	for i := col - 1; i >= 0; i-- {
		if num <= getSize(lines[row][i]) {
			score++
			visible = false
			break
		} else {
			score++
		}
	}
	return score, visible
}

func checkRight(num, row, col int, lines []string) (int, bool) {
	score := 0
	visible := true
	for i := col + 1; i < len(lines[row]); i++ {
		if num <= getSize(lines[row][i]) {
			score++
			visible = false
			break
		} else {
			score++
		}
	}
	return score, visible
}

func checkDown(num, row, col int, lines []string) (int, bool) {
	score := 0
	visible := true
	for i := row + 1; i < len(lines); i++ {
		if num <= getSize(lines[i][col]) {
			score++
			visible = false
			break
		} else {
			score++
		}
	}
	return score, visible
}

func checkUp(num, row, col int, lines []string) (int, bool) {
	score := 0
	visible := true
	for i := row - 1; i >= 0; i-- {
		if num <= getSize(lines[i][col]) {
			score++
			visible = false
			break
		} else {
			score++
		}
	}
	return score, visible
}

func checkAround(num, row, col int, lines []string) (int, bool) {
	us, up := checkUp(num, row, col, lines)
	ds, down := checkDown(num, row, col, lines)
	ls, left := checkLeft(num, row, col, lines)
	rs, right := checkRight(num, row, col, lines)
	return ls * rs * ds * us, up || down || left || right
}
