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
	part1Answer := 0
	cycles := genCycles(lines)

	for i := 20; i <= 220; i = i + 40 {
		part1Answer += cycles.Count[i-1] * i
	}
	fmt.Println(part1Answer)

	cycles.drawScreen()
}

func genCycles(ss []string) *Cycles {
	c := new(Cycles)
	c.Count = make([]int, 1)
	c.Count[0] = 1

	for i := range ss {
		if ss[i] == "noop" {
			c.Count = append(c.Count, c.Count[len(c.Count)-1])
		} else {
			number, _ := strconv.Atoi(ss[i][5:])
			c.Count = append(c.Count, c.Count[len(c.Count)-1])
			c.Count = append(c.Count, c.Count[len(c.Count)-1]+number)
		}
	}
	return c
}

func (c *Cycles) drawScreen() {
	for k, v := range c.Count[:len(c.Count)] {
		spritePos := k % 40
		if spritePos == 0 && k != 0 && k != len(c.Count)-1{
			fmt.Printf("\n")
		}
		if v-spritePos >= -1 && v-spritePos <= 1 {
			fmt.Printf("")
		} else {
			fmt.Printf(" ")
		}
	}
}
