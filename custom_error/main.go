package main

import (
	"fmt"
	"log"
	"net/http"
)

// ResourceError just type for behavior testing
type ResourceError struct {
	Value error
}

func (c *ResourceError) Error() string {
	return fmt.Sprintf(
		"Resource error: %v\n",
		c.Value,
	)
}

func getURL() error {
	_, err := http.Get("some url") // wrong url
	if err != nil {
		return &ResourceError{Value: err}
	}
	return nil
}

func main() {
	err := getURL()
	if err != nil {
		switch err.(type) {
		case *ResourceError:
			err := err.(*ResourceError)
			log.Printf("resourse err: %v\n", err.Value.Error())
		default:
			log.Println("other error type")
		}
	}
}
