package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, true, isMatch("aab", "a*b"))
	assert.Equal(t, true, isMatch("aab", "a.b"))
	assert.Equal(t, false, isMatch("aab", "ab"))
}
