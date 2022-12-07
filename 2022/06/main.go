package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	fmt.Println(getFirstMarker(string(f), 4))
	fmt.Println(getFirstMarker(string(f), 14))
}

func getFirstMarker(s string, size int) int {
	for i := size; i < len(s); i++ {
		if isUnique(s[i-size:i]) {
			return i
		}
	}
	return 0
}

func isUnique(s string) bool {
	ss := strings.Split(s, "")
	set := make(map[string]struct{})
	for _, v := range ss {
		set[v] = struct{}{}
	}
	return len(set) == len(ss)
}
