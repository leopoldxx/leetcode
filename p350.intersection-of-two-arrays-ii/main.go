package main

func intersect(nums1 []int, nums2 []int) []int {
	cache := [1001]int{}
	for _, num := range nums1 {
		cache[num]++
	}
	res := []int{}
	for _, num := range nums2 {
		if v := cache[num]; v > 0 {
			res = append(res, num)
			cache[num]--
		}
	}
	return res
}
