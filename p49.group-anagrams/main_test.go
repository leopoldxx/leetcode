package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	// cases := []struct {
	// 	strs   []string
	// 	result [][]string
	// }{

	// 	{
	// 		strs:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
	// 		result: [][]string{{"eat", "ate", "tea"}, {"bat"}, {"nat", "tan"}},
	// 	},
	// }
	// for _, tc := range cases {
	// 	result := groupAnagrams(tc.strs)
	// 	r, _ := yaml.Marshal(tc.result)
	// 	b, _ := yaml.Marshal(result)
	// 	assert.Equal(t, string(r), string(b))
	// }
	groupAnagrams(nil)
}
