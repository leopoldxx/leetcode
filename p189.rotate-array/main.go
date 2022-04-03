package main

func rotate2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j := 0, len(nums)-1
	for {
		if i == j || i > j {
			return
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	k = k % len(nums)
	rotate2(nums[:len(nums)-k])
	rotate2(nums[len(nums)-k:])
	rotate2(nums)
}
