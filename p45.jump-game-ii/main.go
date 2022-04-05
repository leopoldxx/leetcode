package main

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
func jump(nums []int) int {
	steps := make([]int, len(nums))
	steps[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if steps[i+j] == 0 {
				steps[i+j] = steps[i] + 1
			} else {
				steps[i+j] = min(steps[i+j], steps[i]+1)
			}
		}
	}
	return steps[len(nums)-1]
}
