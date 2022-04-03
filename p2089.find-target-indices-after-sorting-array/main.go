package main

func targetIndices(nums []int, target int) []int {
	bucket := [101]int{}
	for _, num := range nums {
		bucket[num]++
	}
	pre := 0
	for i := 0; i < target; i++ {
		pre += bucket[i]
	}
	res := make([]int, bucket[target])
	for i := 0; i < bucket[target]; i++ {
		res[i] = pre + i
	}
	return res

}
