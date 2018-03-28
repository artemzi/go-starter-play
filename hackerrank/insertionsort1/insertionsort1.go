package main

import (
	"bufio"
	"fmt"
	"os"
)

func insertionsort1(data *[]string) string {
	var result string
	// for range *data {
	// 	result += fmt.Sprintf("%v\n", strings.Join(*data, " "))
	// }
	return result
}

func main() {
	var l int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &l)

	data := make([]string, l)
	for i := 0; i < l; i++ {
		fmt.Fscan(io, &data[i])
	}

	fmt.Printf("%v", insertionsort1(&data))
}
