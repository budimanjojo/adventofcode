package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Answer := 0
	part2Answer := []int{}

	input := strings.Split(string(f), "\n")
	dirList := []string{}

	sizeMap := map[string]int{}

	for i := range input {
		switch processString(input[i]) {
		case "goRoot":
			dirList = []string{"/"}
		case "goUp":
			dirList = dirList[:len(dirList)-1]
		case "goTo":
			dirList = append(dirList, strings.Split(input[i], " ")[2])
		case "fInfo":
			for j := len(dirList); j != 0; j-- {
				d := strings.Join(dirList[:j], "/")
				if strings.HasPrefix(d, "//") {
					d = strings.TrimPrefix(d, "/")
				}
				sizeMap[d] += getSize(input[i])
			}
		}
	}

	used := 70000000-sizeMap["/"]
	needed := 30000000

	keys := []string{}
	for k := range sizeMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		if sizeMap[v] <= 100000 {
			part1Answer += sizeMap[v]
		}
		if sizeMap[v] >= needed-used {
			part2Answer = append(part2Answer, sizeMap[v])
		}
	}
	sort.Ints(part2Answer)

	fmt.Println(part1Answer)
	fmt.Println(part2Answer[0])
}

func processString(s string) string {
	if goRoot, _ := regexp.Match("\\$ cd /$", []byte(s)); goRoot {
		return "goRoot"
	}
	if goUp, _ := regexp.Match("\\$ cd \\.\\.", []byte(s)); goUp {
		return "goUp"
	}
	if goTo, _ := regexp.Match("\\$ cd .+", []byte(s)); goTo {
		return "goTo"
	}
	if fInfo, _ := regexp.Match("[0-9]+ .+", []byte(s)); fInfo {
		return "fInfo"
	}
	return ""
}

func getSize(s string) int {
	ss := strings.Split(s, " ")
	v, _ := strconv.Atoi(ss[0])
	return v
}
