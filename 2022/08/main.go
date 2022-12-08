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
	visible := edges
	score := 0

	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[row])-1; col++ {
			num := getSize(lines[row][col])
			if isVisible(num, row, col, lines) {
				visible++
			}
			if getScore(num, row, col, lines) > score {
				score = getScore(num, row, col, lines)
			}
		}
	}
	fmt.Println(visible)
	fmt.Println(score)
}

func parseTreeSlice(s string) []int {
	ss := strings.Split(s, "")
	si := []int{}
	for i := range ss {
		in, _ := strconv.Atoi(ss[i])
		si = append(si, in)
	}
	return si
}

func getSize(b byte) int {
	size, _ := strconv.Atoi(string(b))
	return size
}

func isVisibleLeft(num, row, col int, lines []string) bool {
	for i := col - 1; i >= 0; i-- {
		if num <= getSize(lines[row][i]) {
			return false
		}
	}
	return true
}

func isVisibleRight(num, row, col int, lines []string) bool {
	for i := col + 1; i < len(lines[row]); i++ {
		if num <= getSize(lines[row][i]) {
			return false
		}
	}
	return true
}

func isVisibleDown(num, row, col int, lines []string) bool {
	for i := row + 1; i < len(lines); i++ {
		if num <= getSize(lines[i][col]) {
			return false
		}
	}
	return true
}

func isVisibleTop(num, row, col int, lines []string) bool {
	for i := row - 1; i >= 0; i-- {
		if num <= getSize(lines[i][col]) {
			return false
		}
	}
	return true
}

func isVisible(num, row, col int, lines []string) bool {
	top := isVisibleTop(num, row, col, lines)
	bot := isVisibleDown(num, row, col, lines)
	left := isVisibleLeft(num, row, col, lines)
	right := isVisibleRight(num, row, col, lines)
	return top || bot || left || right
}

func getScore(num, row, col int, lines []string) int {
	var left, right, up, down int
	for i := col - 1; i >= 0; i-- {
		if num <= getSize(lines[row][i]) {
			left++
			break
		} else {
			left++
		}
	}
	for i := col + 1; i < len(lines[row]); i++ {
		if num <= getSize(lines[row][i]) {
			right++
			break
		} else {
			right++
		}
	}
	for i := row + 1; i < len(lines); i++ {
		if num <= getSize(lines[i][col]) {
			down++
			break
		} else {
			down++
		}
	}
	for i := row - 1; i >= 0; i-- {
		if num <= getSize(lines[i][col]) {
			up++
			break
		} else {
			up++
		}
	}
	if left == 0 { left++ }
	if right == 0 { right++ }
	if down == 0 { down++ }
	if up == 0 { up++ }
	return left * right * down * up
}
