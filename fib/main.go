package main

import "fmt"

func main() {
	x := fib(10)
	fmt.Printf("Fibonacci digit #10 is %d\n", x)
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
