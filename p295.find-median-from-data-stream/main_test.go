package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	median := medianFinder.FindMedian()
	assert.Equal(t, float64(1.5), median)
	medianFinder.AddNum(3)
	median = medianFinder.FindMedian()
	assert.Equal(t, float64(2), median)
}
