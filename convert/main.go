package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inp := bufio.NewScanner(os.Stdin)
	fmt.Println("Input integer for binary conversion:")
	for inp.Scan() {
		i, err := strconv.Atoi(inp.Text())
		if err != nil {
			log.Println("Wrong integer value in input")
			os.Exit(2)
		}
		fmt.Printf("Bin value of input is: %08b\n", i)
	}
}
