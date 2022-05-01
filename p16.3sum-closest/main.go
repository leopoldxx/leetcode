package main

import "sort"

func distant(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := 0
	closestDistant := 1000000
	for i := 0; i < len(nums)-2; i++ {
		x := i + 1
		y := len(nums) - 1
		for x < y {
			sum := nums[i] + nums[x] + nums[y]
			d := distant(sum, target)
			if d == 0 {
				return sum
			} else if closestDistant > d {
				closestSum = sum
				closestDistant = d
			}
			if sum > target {
				y--
			} else if sum < target {
				x++
			}
		}
	}
	return closestSum
}
