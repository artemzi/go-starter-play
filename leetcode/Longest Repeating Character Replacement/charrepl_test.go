package charrepl

import (
	"runtime"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				t.Errorf("Panic: %v\n", r)
				// debug.PrintStack()
			}
		}
	}()

	s := "ABAB"
	k := 2
	expected := 4

	if r := characterReplacement(s, k); r != expected {
		t.Errorf("Result for %s is wrong. got: %d, expected: %d\n", s, r, expected)
	}

	s = "AABABBA"
	k = 1

	if r := characterReplacement(s, k); r != expected {
		t.Errorf("Result for %s is wrong. got: %d, expected: %d\n", s, r, expected)
	}
}
