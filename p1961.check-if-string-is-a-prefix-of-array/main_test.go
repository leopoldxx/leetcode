package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.True(t, isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"}))
}
