package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	if len(nums) < 4 {
		return nil
	}
	res := make([][]int, 0, 100)
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			x := j + 1
			y := len(nums) - 1
			for x < y {
				s := nums[i] + nums[j] + nums[x] + nums[y]
				if s < target {
					x++
					continue
				} else if s > target {
					y--
					continue
				}
				res = append(res, []int{nums[i], nums[j], nums[x], nums[y]})
				x++
				y--
				for x < y && nums[x] == nums[x-1] {
					x++
				}
				for x < y && nums[y] == nums[y+1] {
					y--
				}
			}
		}
	}
	return res
}
