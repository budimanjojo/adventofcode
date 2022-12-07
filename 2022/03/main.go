package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	partOneResult := 0
	partTwoResult := 0
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	for _, line := range lines {
		count := len(line)
		if count > 0 {
			firstCompart := strings.Split(line[0:count/2], "")
			secondCompart := strings.Split(line[count/2:count], "")
			dup := getDuplicates(firstCompart, secondCompart)
			value, err := stringtoNumber(dup[0])
			if err != nil {
				panic(err)
			}
			partOneResult += value
		}
	}

	groups := make([][]string, len(lines)/3)
	k := 0
	for i := range groups {
		groups[i] = make([]string, 3)
		for g := range groups[i] {
			groups[i][g] = lines[k]
			k++
		}
	}

	for k := range groups {
		s1 := strings.Split(groups[k][0], "")
		s2 := strings.Split(groups[k][1], "")
		s3 := strings.Split(groups[k][2], "")
		dup1 := getDuplicates(s1, s2)
		dup2 := getDuplicates(dup1, s3)
		value, err := stringtoNumber(dup2[0])
		if err != nil {
			panic(err)
		}

		partTwoResult += value
	}

	fmt.Println(partOneResult)
	fmt.Println(partTwoResult)
}

func stringtoNumber(s string) (int, error) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	char := strings.Split(chars, "")
	for i := range char {
		if s == char[i] {
			i++
			return i, nil
		}
	}
	return 0, fmt.Errorf("Error converting string to number!")
}

func getDuplicates(a, b []string) []string {
	var got []string
	for k, s := range a {
		if contains(b, s) {
			got = append(got, a[k])
		}
	}
	return got
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}
