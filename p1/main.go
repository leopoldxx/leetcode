package main

func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	for idx, num := range nums {
		if idxPre, exists := cache[target-num]; exists {
			return []int{idxPre, idx}
		}
		cache[num] = idx
	}
	return nil
}
