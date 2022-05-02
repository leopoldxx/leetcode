package main

func decompressRLElist(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i = i + 2 {
		freq := nums[i]
		val := nums[i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}
	return res
}
