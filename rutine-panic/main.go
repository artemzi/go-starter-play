package main

import (
	"errors"
	"log"
	"time"
)

type GoDoer func()

func Go(todo GoDoer) {
	go func() {
		defer func() { // first handle panic
			if err := recover(); err != nil {
				log.Printf("Panic appear: %s", err)
			}
		}()
		todo() // next call function
	}()
}

func message() {
	println("Inside gorutine")
	panic(errors.New("Oops!"))
}

func main() {
	Go(message) // call same as `go message()` but in safe
	println("Outside gorutine")
	time.Sleep(1 * time.Second) // just wait until gorutine finish job.
}
