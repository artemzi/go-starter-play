package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var digit, length int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &digit)
	fmt.Fscan(io, &length)

	data := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(io, &data[i])
	}

	for i := range data {
		if data[i] == digit {
			fmt.Fprintln(os.Stdout, i)
			break
		}
	}
}
