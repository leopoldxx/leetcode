package main

func computeCount(length int) int {
	res := 1
	for i := 1; i <= length; i++ {
		res *= i
	}
	return res
}

func getPermutaionBySlice(nums []byte, k int) []byte {
	if len(nums) == 1 {
		return nums
	}
	c := computeCount(len(nums[1:]))

	d := (k - 1) / c
	k = (k - 1) % c

	char := nums[d]

	if d == 0 {
		nums = nums[1:]
	} else if d == len(nums)-1 {
		nums = nums[:d]
	} else {
		nums = append(nums[:d], nums[d+1:]...)
	}

	return append([]byte{char}, getPermutaionBySlice(nums, k+1)...)
}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		if k == 1 {
			return "12"
		} else {
			return "21"
		}
	}
	nums := make([]byte, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = byte('0' + i)
	}
	return string(getPermutaionBySlice(nums, k))
}
