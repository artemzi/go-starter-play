package main

import (
	"fmt"
	"strings"
)

func insertionsort1(data []string, n int) string {
	var result string
	i := n - 1
	j := i
	x := data[i]

	for j > 0 && data[j-1] > x {
		data[j] = data[j-1]
		j--
		result += fmt.Sprintf("%v\n", strings.Join(data, " "))
	}
	data[j] = x
	result += fmt.Sprintf("%v\n", strings.Join(data, " "))
	return result
}

func main() {
	var l int
	fmt.Scan(&l)

	data := make([]string, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Printf("%v", insertionsort1(data, l))
}
