package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1Answer := 0
	part2Answer := 0
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	for k := range lines {
		if len(lines[k]) > 0 {
			group := strings.Split(lines[k], ",")
			g1 := strings.Split(group[0], "-")
			g2 := strings.Split(group[1], "-")
			var g1Int []int
			var g2Int []int
			for k := range g1 {
				he, _ := strconv.Atoi(g1[k])
				g1Int = append(g1Int, he)
			}
			for k := range g2 {
				he, _ := strconv.Atoi(g2[k])
				g2Int = append(g2Int, he)
			}
			slice1 := makeRange(g1Int[0], g1Int[1])
			slice2 := makeRange(g2Int[0], g2Int[1])
			if len(slice1) < len(slice2) {
				if fullyContain(slice1, slice2) {
					part1Answer++
				}
				if partialContain(slice1, slice2) {
					part2Answer++
				}
			} else {
				if fullyContain(slice2, slice1) {
					part1Answer++
				}
				if partialContain(slice1, slice2) {
					part2Answer++
				}
			}
		}
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func fullyContain(small, big []int) bool {
	smallStr := "(^|[^0-9])" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(small)), ", "), "[]") + "($|[^0-9])"
	bigStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(big)), ", "), "[]")
	if matched, _ := regexp.Match(smallStr, []byte(bigStr)); matched {
		return true
	}
	return false
}

func partialContain(small, big []int) bool {
	bigStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(big)), ", "), "[]")
	for k := range small {
		s := "(^|[^0-9])" + fmt.Sprint(small[k]) + "($|[^0-9])"
		if matched, _ := regexp.Match(s, []byte(bigStr)); matched {
			return true
		}
	}
	return false
}
