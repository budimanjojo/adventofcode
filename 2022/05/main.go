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
	
	input := strings.Split(string(f), "\n\n")
	moves := strings.Split(string(input[1]), "\n")
	data := getStackData(input[0])
	data2 := make(map[int][]string)
	for k, v := range data {
		newData := make([]string, len(v))
		copy(newData, v)
		data2[k] = newData
	}

	for line := range moves {
		if len(moves[line]) > 0 {
			s := strings.Split(moves[line], " ")
			ss := []string{s[1], s[3], s[5]}
			move := make([]int, len(ss))
			for k := range ss {
				move[k], _ = strconv.Atoi(ss[k])
			}

			data = moveBoxes(move, data)
			data2 = moveBoxesTogether(move, data2)
		}
	}

	partOneAnswer := []string{}
	partTwoAnswer := []string{}

	for i := 1; i < len(data)+1; i++ {
		partOneAnswer = append(partOneAnswer, data[i][len(data[i])-1])
		partTwoAnswer = append(partTwoAnswer, data2[i][len(data2[i])-1])
	}
	fmt.Println(strings.Join(partOneAnswer, ""))
	fmt.Println(strings.Join(partTwoAnswer, ""))
}

func getStackData(s string) map[int][]string {
	ss := strings.Split(s, "\n")
	d := map[int][]string{}
	for i := len(ss)-2; i > -1; i-- {
		line := ss[i][1:]
		for j, k := 0, 0; k < len(line); j, k = j+1, k+4 {
			if string(line[k]) != " " {
				d[j+1] = append(d[j+1], string(line[k]))
			}
		}
	}
	return d
}

func moveBoxes(s []int, data map[int][]string) map[int][]string {
	toRemove := s[0]
	for i := 0; i < toRemove; i++ {
		data[s[2]] = append(data[s[2]], data[s[1]][len(data[s[1]])-1])
		data[s[1]] = removeLastIndex(data[s[1]])
	}
	return data
}

func moveBoxesTogether(s []int, data map[int][]string) map[int][]string {
	toRemove := s[0]
	for i := 0; i < toRemove; i++ {
		data[s[2]] = append(data[s[2]], data[s[1]][len(data[s[1]])-toRemove+i])
	}
	for i := 0; i < toRemove; i++ {
		data[s[1]] = removeLastIndex(data[s[1]])
	}
	return data
}

func removeLastIndex(s []string) []string {
	return s[:len(s)-1]
}
