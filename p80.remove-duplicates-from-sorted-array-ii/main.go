package main

func removeDuplicates(nums []int) int {
	fast, slow := 1, 0

	num := nums[0]
	count := 1

	for fast < len(nums) {
		if nums[fast] == num {
			count++
			fast++
			continue
		}
		if count > 2 {
			count = 2
		}
		for i := 0; i < count; i++ {
			nums[slow] = num
			slow++
		}

		num = nums[fast]
		count = 1
		fast++
	}
	if count > 2 {
		count = 2
	}
	for i := 0; i < count; i++ {
		nums[slow] = num
		slow++
	}
	return slow
}
