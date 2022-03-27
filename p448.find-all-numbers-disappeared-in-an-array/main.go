package main

func findDisappearedNumbers(nums []int) []int {
	extraN_1 := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i]%n == 0 {
			extraN_1 += n
			continue
		}
		nums[nums[i]%n] += n
	}
	result := []int{}
	for i := 1; i < n; i++ {
		if nums[i] <= n {
			result = append(result, i)
		}
	}
	if extraN_1 == 0 {
		result = append(result, n)
	}
	return result
}
