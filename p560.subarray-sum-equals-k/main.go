package main

func subarraySum(nums []int, k int) int {
	var (
		sumCount = map[int]int{}
		total    = 0
		sum      = 0
	)
	sumCount[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if count, exist := sumCount[sum-k]; exist {
			total += count
		}
		sumCount[sum]++
	}
	return total
}
