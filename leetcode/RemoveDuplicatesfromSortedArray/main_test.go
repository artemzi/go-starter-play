package main

import "testing"

func BenchmarkRemoveDuplicates(b *testing.B) {
	RemoveDuplicates([]int{1, 1, 2})
}
