package main

func trap(height []int) int {
	left, leftMaxHeight := 0, height[0]
	right, rightMaxHeight := len(height)-1, height[len(height)-1]
	capacity := 0

	for left < right {
		if leftMaxHeight <= rightMaxHeight {
			left++
			if leftMaxHeight <= height[left] {
				leftMaxHeight = height[left]
			} else {
				capacity += (leftMaxHeight - height[left]) //diff
			}
		} else {
			right--
			if rightMaxHeight <= height[right] {
				rightMaxHeight = height[right]
			} else {
				capacity += (rightMaxHeight - height[right])
			}
		}
	}

	return capacity
}
