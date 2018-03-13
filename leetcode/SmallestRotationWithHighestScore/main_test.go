package main

import (
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	inp := []int{2, 3, 1, 4, 0}
	move(&inp)

	if !reflect.DeepEqual(inp, []int{3, 1, 4, 0, 2}) {
		t.Errorf("Fail move case 1 return bad result: %v", inp)
	}

	move(&inp)

	if !reflect.DeepEqual(inp, []int{1, 4, 0, 2, 3}) {
		t.Errorf("Fail move case 2 return bad result: %v", inp)
	}
}

func TestBestRotation(t *testing.T) {
	inp := []int{2, 3, 1, 4, 0}
	res := bestRotation(inp)
	// Explanation:
	// Scores for each K are listed below:
	// K = 0,  A = [2,3,1,4,0],    score 2
	// K = 1,  A = [3,1,4,0,2],    score 3
	// K = 2,  A = [1,4,0,2,3],    score 3
	// K = 3,  A = [4,0,2,3,1],    score 4
	// K = 4,  A = [0,2,3,1,4],    score 3

	// For example, if we have [2, 4, 1, 3, 0], and we rotate by K = 2, it becomes [1, 3, 0, 2, 4].
	// This is worth 3 points because 1 > 0 [no points], 3 > 1 [no points], 0 <= 2 [one point],
	// 2 <= 3 [one point], 4 <= 4 [one point].

	if res != 3 {
		t.Errorf("Wrong result; got: %v, expected: 3.\n", res)
	}
}
