package main

func combineNums(nums []int, k int) [][]int {
	n := len(nums)
	if n == k {
		return [][]int{nums}
	} else if k == 1 {
		res := [][]int{}
		for i := 0; i < n; i++ {
			res = append(res, []int{nums[i]})
		}
		return res
	}
	res := [][]int{}
	for i := 0; i <= n-k; i++ {
		subRes := []int{nums[i]}
		for _, m := range combineNums(nums[i+1:], k-1) {
			res = append(res, append(subRes, m...))
		}
	}
	return res
}

func combine(n int, k int) [][]int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	if n == k {
		return [][]int{nums}
	} else if k == 1 {
		res := [][]int{}
		for i := 0; i < n; i++ {
			res = append(res, []int{nums[i]})
		}
		return res
	}
	res := [][]int{}
	for i := 0; i <= n-k; i++ {
		subRes := []int{nums[i]}
		for _, m := range combineNums(nums[i+1:], k-1) {
			res = append(res, append(subRes, m...))
		}
	}
	return res
}
