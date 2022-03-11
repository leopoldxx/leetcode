package main

func maxArea(height []int) int {
	area, left, right := 0, 0, len(height)-1
	max, min := func(l, r int) int {
		if l < r {
			return r
		}
		return l
	}, func(l, r int) int {
		if l < r {
			return l
		}
		return r
	}
	for left < right {
		area = max(area, (right-left)*(min(height[left], height[right])))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
