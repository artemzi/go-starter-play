package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Post is a post data
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	cli := &http.Client{Timeout: 3 * time.Second}
	r, err := cli.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("Error in getting json: %v", err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error in reading response: %v", err)
	}

	data := []Post{}
	json.Unmarshal(body, &data)

	fmt.Printf("%T\n userId: %v\n id: %v\n title: %v\n body: %v\n",
		data[0], data[0].UserID, data[0].ID, data[0].Title, data[0].Body)
}
