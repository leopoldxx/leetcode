package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	data := []byte("aabbccc")
	res := compress(data)
	assert.Equal(t, 6, res)
	assert.Equal(t, []byte("a2b2c3"), data[:res])
}
