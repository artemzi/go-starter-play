// https://leetcode.com/contest/weekly-contest-75/problems/smallest-rotation-with-highest-score/
package main

import "fmt"

func move(A *[]int) {
	(*A) = append((*A)[1:], (*A)[0])
}

func calculateScore(A []int) int {
	score := 0
	for i, v := range A {
		if v <= i {
			score++
			continue
		}
	}
	return score
}

func bestRotation(A []int) int {
	return 1
}

func main() {
	fmt.Print("Run with: go test -v\n")
}
