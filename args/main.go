// print command line arguments in loop and with Join with execution time evaluation
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Command name: %s, command values: %v\n", os.Args[0], strings.Join(os.Args[1:], " "))
	// for i, v := range os.Args[1:] {
	// 	fmt.Println(i, v)
	// }
	elapsed := time.Since(start)
	log.Println("Execution time is >> ", elapsed)
}
