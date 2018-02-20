package main

import (
	"log"
	"time"
)

func main() {
	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	log.Printf("%d wants the lock\n", id)
	lock <- true
	log.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond) // safe under lock
	log.Printf("%d is releasing the lock\n", id)
	<-lock // remove lock by reading from chan
}
