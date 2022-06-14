package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	res := []string{}
	n := len(nums)

	var left = 0
	i := 0
	for i < n {
		left = nums[i]
		pre := left
		j := i + 1
		for j < n {
			if pre+1 != nums[j] {
				break
			}
			pre = nums[j]
			j++
		}
		if j == i+1 {
			res = append(res, fmt.Sprintf("%d", left))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", left, pre))
		}
		i = j
	}
	return res
}
