package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		// it's guaranteed that the solution must be exist.
		if left == right {
			return []int{left + 1, right + 1}
		}
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
}
