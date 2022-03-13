package main

func productExceptSelf(nums []int) []int {
	prefixProduct := make([]int, len(nums))
	suffixProduct := make([]int, len(nums))
	product := make([]int, len(nums))
	currentProduct := 1

	for i := 1; i < len(nums); i++ {
		currentProduct *= nums[i-1]
		prefixProduct[i-1] = currentProduct
	}
	currentProduct = 1
	for i := len(nums) - 2; i >= 0; i-- {
		currentProduct *= nums[i+1]
		suffixProduct[i+1] = currentProduct
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			product[i] = suffixProduct[i+1]
		} else if i == len(nums)-1 {
			product[i] = prefixProduct[i-1]
		} else {

			product[i] = prefixProduct[i-1] * suffixProduct[i+1]
		}
	}
	return product
}
