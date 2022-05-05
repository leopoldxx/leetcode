package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "9534330", largestNumber([]int{3, 30, 34, 5, 9}))
}
