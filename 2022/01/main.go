package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(f), "\n\n")

	var s int
	var sums []int
	for i := 0; i < len(input); i++ {
		arr := strings.Split(input[i], "\n")
		for I := range arr {
			d, _ := strconv.Atoi(arr[I])
			s += d
		}
		sums = append(sums, s)
		s = 0
	}
	sort.Ints(sums)
	fmt.Println(sums[len(sums)-1])
	fmt.Println(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3])
}
