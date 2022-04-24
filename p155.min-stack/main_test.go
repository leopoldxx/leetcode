package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	stack := Constructor()
	stack.Push(2147483646)
	stack.Push(2147483646)
	stack.Push(2147483647)
	assert.Equal(t, 2147483647, stack.Top())
	stack.Pop()
	assert.Equal(t, 2147483646, stack.GetMin())
	stack.Pop()
	assert.Equal(t, 2147483646, stack.GetMin())
	stack.Pop()
	stack.Push(2147483647)
	assert.Equal(t, 2147483647, stack.Top())
	assert.Equal(t, 2147483647, stack.GetMin())

}
