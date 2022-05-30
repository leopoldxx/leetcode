package main

func singleNumber(nums []int) int {
	var res int32
	var tmp [32]int32
	for k := 0; k < 32; k++ {
		for i := 0; i < len(nums); i++ {
			v := int32(nums[i])
			tmp[k] += v >> k & 1
			tmp[k] %= 3
		}
		res |= int32(tmp[k] << k)
	}
	return int(res)
}
