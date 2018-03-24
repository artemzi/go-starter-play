package coinchange

import (
	"testing"
)

func TestCoinChangeNormalResult(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	expected := 3 // (11 = 5 + 5 + 1)

	if r := coinChange(coins, amount); r != expected {
		t.Errorf("Result for %d with amount %d is wrong. got: %d, expected: %d\n",
			coins, amount, r, expected)
	}
}

func TestCoinChangeNilResult(t *testing.T) {
	coins := []int{2}
	amount := 3
	expected := -1

	if r := coinChange(coins, amount); r != expected {
		t.Errorf("Result for %d with amount %d is wrong. got: %d, expected: %d\n",
			coins, amount, r, expected)
	}
}
