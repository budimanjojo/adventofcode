package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items        []int
	Operation    []string
	DivBy        int
	IfTrue       int
	IfFalse      int
	InspectCount int
}

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Answer := []int{}
	part2Answer := []int{}

	monkeys := genMonkeys(string(f))
	monkeys2 := genMonkeys(string(f))

	for i := 0; i < 20; i++ {
		doRound(monkeys, false)
	}
	for i := 0; i < 10000; i++ {
		doRound(monkeys2, true)
	}

	for i := range monkeys {
		part1Answer = append(part1Answer, monkeys[i].InspectCount)
	}
	for i := range monkeys2 {
		part2Answer = append(part2Answer, monkeys2[i].InspectCount)
	}

	sort.Ints(part1Answer)
	sort.Ints(part2Answer)
	fmt.Println(part1Answer[len(part1Answer)-1] * part1Answer[len(part1Answer)-2])
	fmt.Println(part2Answer[len(part2Answer)-1] * part2Answer[len(part2Answer)-2])
}

func genMonkeys(s string) []Monkey {
	monkey := Monkey{}
	monkeys := []Monkey{}
	ss := strings.Split(s, "\n\n")
	for i := range ss {
		data := strings.Split(ss[i], "\n")
		items := strings.Split(strings.TrimPrefix(data[1], "  Starting items: "), ", ")
		for i := range items {
			d, _ := strconv.Atoi(items[i])
			monkey.Items = append(monkey.Items, d)
		}
		db, _ := strconv.Atoi(strings.TrimPrefix(data[3], "  Test: divisible by "))
		ifTrue, _ := strconv.Atoi(string(data[4][len(data[4])-1]))
		ifFalse, _ := strconv.Atoi(string(data[5][len(data[5])-1]))
		monkey.Operation = strings.Split(strings.TrimPrefix(data[2], "  Operation: new = "), " ")
		monkey.DivBy = db
		monkey.IfTrue = ifTrue
		monkey.IfFalse = ifFalse
		monkey.InspectCount = 0
		monkeys = append(monkeys, monkey)
		monkey = Monkey{}
	}
	return monkeys
}

func doRound(monkeys []Monkey, isWorried bool) *[]Monkey {
	for i := range monkeys {
		for j := 0; j < len(monkeys[i].Items); j++ {
			wLevel := monkeys[i].Items[j]
			wLevel = monkeys[i].calculate(wLevel)
			wLevel = getWorryLevel(wLevel, monkeys, isWorried)
			if wLevel%monkeys[i].DivBy == 0 {
				monkeys[monkeys[i].IfTrue].Items = append(monkeys[monkeys[i].IfTrue].Items, wLevel)
			} else {
				monkeys[monkeys[i].IfFalse].Items = append(monkeys[monkeys[i].IfFalse].Items, wLevel)
			}
			monkeys[i].InspectCount++
		}
		monkeys[i].Items = []int{}
	}
	return &monkeys
}

func (m Monkey) calculate(i int) int {
	lhs, err := strconv.Atoi(m.Operation[0])
	if err != nil {
		lhs = i
	}
	rhs, err := strconv.Atoi(m.Operation[2])
	if err != nil {
		rhs = i
	}
	switch m.Operation[1] {
	case "+":
		return lhs + rhs
	case "*":
		return lhs * rhs
	}
	return 0
}

func getWorryLevel(i int, monkeys []Monkey, isWorried bool) int {
	div := 1
	if isWorried {
		for i := range monkeys {
			div *= monkeys[i].DivBy
		}
		return i % div
	} else {
		return i / 3
	}
}
