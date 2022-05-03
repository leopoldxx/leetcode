package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 88, hashpow(3))
	assert.Equal(t, 4, hash("abr"))
	assert.Equal(t, 30, hashnew(4, 88, 'a', 'a'))
	assert.Equal(t, 1, strStr("abra", "bra"))
	assert.Equal(t, 2, strStr("hello", "ll"))
}
