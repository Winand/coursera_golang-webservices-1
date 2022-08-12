package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello there my dear friend say hello to everyone No smoking No loitering No no"
	s = strings.ToLower(s)
	counts := map[string]int{}
	for _, w := range strings.Fields(s) {
		counts[w] += 1 // also ++
	}
	fmt.Println(counts)
}
