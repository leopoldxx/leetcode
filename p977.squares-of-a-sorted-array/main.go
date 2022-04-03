package main

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}
	res := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	k := len(nums) - 1
	for {
		if i > j {
			break
		}
		ii := nums[i] * nums[i]
		jj := nums[j] * nums[j]
		if ii >= jj {
			res[k] = ii
			i++
		} else {
			res[k] = jj
			j--
		}
		k--
	}
	return res
}
