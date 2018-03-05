// solution for https://tour.golang.org/moretypes/23
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount return map with world => count
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, val := range strings.Fields(s) {
		m[val]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
