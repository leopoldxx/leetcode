package main

func threeSum(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		for left := i + 1; left < len(nums)-1; left++ {
			for right := left + 1; right < len(nums); right++ {
				s := nums[i] + nums[left] + nums[right]
				if s == target {
					count++
				}
			}
		}
	}
	return count
}

func countQuadruplets(nums []int) int {
	count := 0
	for j := len(nums) - 1; j > 2; j-- {
		count += threeSum(nums[:j], nums[j])
	}
	return count
}
