package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	head := Position{X: 0, Y: 0}
	tail := Position{X: 0, Y: 0}
	visited := []Position{}
	visited = append(visited, tail)

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]
	for i := range lines {
		direction := string(lines[i][0])
		amount, _ := strconv.Atoi(string(lines[i][2:]))

		for i := 0; i < amount; i++ {
			head.MoveHead(direction)
			tail.MoveTail(direction, head)
			if !contains(visited, tail) {
				visited = append(visited, tail)
		}
		}
	}
	fmt.Println(len(visited))
}

func (head *Position) MoveHead(dir string) *Position {
	switch dir {
	case "L":
		head.X--
	case "R":
		head.X++
	case "U":
		head.Y++
	case "D":
		head.Y--
	}
	return head
}

func (tail *Position) MoveTail(dir string, head Position) *Position {
	if tail.X == head.X && tail.Y == head.Y {
		return tail
	}
	switch dir {
	case "L":
		if tail.Y != head.Y {
			if tail.X == head.X+2 {
				tail.X--
				tail.Y = head.Y
			}
		} else {
			tail.X = head.X+1
		}
	case "R":
		if tail.Y != head.Y {
			if tail.X == head.X-2 {
				tail.X++
				tail.Y = head.Y
			}
		} else {
			tail.X = head.X-1
		}
	case "U":
		if tail.X != head.X {
			if tail.Y == head.Y-2 {
				tail.X = head.X
				tail.Y++
			}
		} else {
			tail.Y = head.Y-1
		}
	case "D":
		if tail.X != head.X {
			if tail.Y == head.Y+2 {
				tail.X = head.X
				tail.Y--
			}
		} else {
			tail.Y = head.Y+1
		}

	}
	return tail
}

func contains(visited []Position, pos Position) bool {
	for _, v := range visited {
		if v == pos {
			return true
		}
	}
	return false
}
