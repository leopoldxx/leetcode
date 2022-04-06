package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	buildList := func(data ...int) *ListNode {
		var list *ListNode
		for idx := len(data) - 1; idx >= 0; idx-- {
			list = &ListNode{Val: data[idx], Next: list}
		}
		return list
	}

	assert.Equal(t,
		buildList(4, 5, 6),
		middleNode(buildList(1, 2, 3, 4, 5, 6)))

}
