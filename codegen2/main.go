//go:generate ./queue MyInt
package main

import "log"

type MyInt int

func main() {
	var one, two, three MyInt = 1, 2, 3
	q := NewMyIntQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	log.Printf("First value: %d\n", q.Remove())
}

// use it with:
// go build gen/* && ./generator main.go MyInt
