package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionsort1(t *testing.T) {
	data := []string{"2", "4", "6", "8", "3"}
	expected := "2 4 6 8 8\n2 4 6 6 8\n2 4 4 6 8\n2 3 4 6 8\n"
	output := insertionsort1(data, len(data))

	assert.Equal(t, expected, output, "Wrong output")
}
