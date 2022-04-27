package main

func getRow(rowIndex int) []int {
	length := rowIndex + 1
	nums := make([]int, length)

	nums[0] = 1

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			nums[j] = nums[j] + nums[j-1]
		}
	}
	return nums
}
