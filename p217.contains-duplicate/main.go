package main

func containsDuplicate(nums []int) bool {
	cache := map[int]struct{}{}
	for _, num := range nums {
		if _, exists := cache[num]; exists {
			return true
		}
		cache[num] = struct{}{}
	}
	return false
}
