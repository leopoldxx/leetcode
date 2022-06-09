package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	trie := Constructor()
	trie.Insert("hello world")
	assert.True(t, trie.StartsWith("hello"))
}
