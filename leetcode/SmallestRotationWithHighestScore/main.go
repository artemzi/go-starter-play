// https://leetcode.com/contest/weekly-contest-75/problems/smallest-rotation-with-highest-score/
package main

import (
	"fmt"
)

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
	var (
		result, tmp int
	)
	shifts := make(map[int]int)
	for i := range A {
		shifts[i] = calculateScore(A)
		move(&A)
	}

	for i := 0; i < len(shifts); i++ {
		if tmp == 0 {
			tmp = shifts[i]
		} else {
			if shifts[i] == tmp {
				continue
			}
			if shifts[i] > tmp {
				tmp = shifts[i]
				result = i
			}
		}
	}

	return result
}

func main() {
	fmt.Print("Run with: go test -v\n")
}
