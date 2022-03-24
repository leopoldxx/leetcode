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
	cases := []struct {
		head   *ListNode
		result *ListNode
	}{
		{
			head:   buildList(1, 2, 3, 4),
			result: buildList(1, 4, 2, 3),
		},
		{
			head:   buildList(1, 2, 3, 4, 5),
			result: buildList(1, 5, 2, 4, 3),
		},
	}
	for _, tc := range cases {
		reorderList(tc.head)
		assert.Equal(t, tc.result, tc.head)
	}

}
