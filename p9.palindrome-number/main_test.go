package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		num    int
		result bool
	}{
		{
			num:    121,
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, isPalindrome(tc.num))
	}
}
