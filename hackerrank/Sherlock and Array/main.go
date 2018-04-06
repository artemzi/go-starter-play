package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &t)
	data := make([][]int, 0, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(io, &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(io, &arr[j])
		}

		data = append(data, arr)
	}

	for _, arr := range data {
		fmt.Printf("%T: %v\n", arr, arr)
	}
}
