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

	head := newPosition()
	tail := newPosition()
	part1Answer := []Position{newPosition()}
	part2Answer := []Position{newPosition()}
	rope := make([]Position, 9)
	for i := range rope {
		rope[i] = newPosition()
	}

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]
	for i := range lines {
		direction := string(lines[i][0])
		amount, _ := strconv.Atoi(string(lines[i][2:]))

		for i := 0; i < amount; i++ {
			head.moveHead(direction)
			tail.moveTail(head)
			if !tail.contains(part1Answer) {
				part1Answer = append(part1Answer, tail)
			}
			rope[0].moveTail(head)
			for k := 1; k < len(rope); k++ {
				rope[k].moveTail(rope[k-1])
				if !rope[len(rope)-1].contains(part2Answer) {
					part2Answer = append(part2Answer, rope[len(rope)-1])
				}
			}
		}
	}

	fmt.Println(len(part1Answer))
	fmt.Println(len(part2Answer))
}

func (head *Position) moveHead(dir string) *Position {
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

func (tail *Position) moveTail(head Position) *Position {
	if tail.X == head.X && tail.Y == head.Y {
		return tail
	}
	if head.X-tail.X == -2 {
		tail.X = head.X + 1
		tail.Y = head.Y
	}
	if head.X-tail.X == 2 {
		tail.X = head.X - 1
		tail.Y = head.Y
	}
	if head.Y-tail.Y == 2 {
		tail.X = head.X
		tail.Y = head.Y - 1
	}
	if head.Y-tail.Y == -2 {
		tail.X = head.X
		tail.Y = head.Y + 1
	}
	return tail
}

func (pos Position) contains(visited []Position) bool {
	for _, v := range visited {
		if v == pos {
			return true
		}
	}
	return false
}

func newPosition() Position {
	return Position{X: 0, Y: 0}
}
