package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int) string {
	input := bufio.NewScanner(f)
	var duplicates string
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			duplicates = f.Name()
		}
	}
	// here we ignore all errors from input.Err()
	return duplicates
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var duplicates []string
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err: %v\n", err)
				continue
			}

			d := countLines(f, counts)
			duplicates = append(duplicates, d)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println("duplicates found in: ", strings.Join(duplicates, " "))
}
