package main

import (
	"reflect"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	score := calculateScore([]int{1, 3, 0, 2, 4})

	if score != 3 {
		t.Errorf("Wrong score; got: %v, expected: 3.\n", score)
	}
}

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
		t.Errorf("Wrong result case 1; got: %v, expected: 3.\n", res)
	}

	res = bestRotation([]int{0, 3, 1, 3, 1})
	if res != 0 {
		t.Errorf("Wrong result case 2; got: %v, expected: 3.\n", res)
	}

	res = bestRotation([]int{6, 2, 8, 3, 5, 2, 4, 3, 7, 6})
	if res != 9 {
		t.Errorf("Wrong result case 3; got: %v, expected: 3.\n", res)
	}

	res = bestRotation([]int{33, 17, 66, 78, 44, 82, 60, 85, 93, 17, 2, 68, 80, 29, 30, 16, 27, 97, 75, 45, 52, 51, 59, 27, 26, 60, 26, 0, 23, 49, 65, 68, 45, 18, 97, 64, 99, 4, 96, 85, 12, 50, 95, 96, 76, 84, 90, 56, 66, 74, 7, 99, 24, 45, 40, 20, 57, 88, 89, 87, 77, 68, 55, 65, 56, 54, 48, 47, 89, 22, 91, 39, 49, 62, 89, 93, 3, 44, 9, 9, 89, 31, 41, 12, 58, 53, 31, 66, 83, 6, 93, 20, 10, 68, 57, 72, 65, 60, 87, 75})
	if res != 63 {
		t.Errorf("Wrong result case 4; got: %v, expected: 3.\n", res)
	}
}
