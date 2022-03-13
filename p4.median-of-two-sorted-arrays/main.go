package main

func findMedianInSingleSortedArrays(nums []int) float64 {
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
	return float64(nums[len(nums)/2])
}

func median(nums1 []int, nums2 []int, p, p1, p2 int) float64 {
	isOdd := (len(nums1)+len(nums2))%2 == 1
	if isOdd {
		return float64(nums1[p])
	}
	if p1 >= len(nums1) {
		return float64(nums1[p]+nums2[p2]) / 2
	} else if p2 >= len(nums2) {
		return float64(nums1[p]+nums1[p1]) / 2
	}
	if nums1[p1] < nums2[p2] {
		return float64(nums1[p]+nums1[p1]) / 2
	}
	return float64(nums1[p]+nums2[p2]) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 == 0 && len2 == 0 {
		return 0
	} else if len1 == 0 {
		return findMedianInSingleSortedArrays(nums2)
	} else if len2 == 0 {
		return findMedianInSingleSortedArrays(nums1)
	}

	k := (len1 + len2 + 1) / 2

	left := 0
	right := len1
	for left < right {
		index1 := left + (right-left)/2
		index2 := k - (index1 + 1) - 1

		if index2+1 < 0 {
			right = index1
			continue
		} else if index2+1 == 0 && nums1[index1] <= nums2[0] {
			return median(nums1, nums2, index1, index1+1, index2+1)
		} else if index2+1 == 0 && nums1[index1] > nums2[0] {
			right = index1
			continue
		} else if index2+1 > len2 {
			left = index1 + 1
			continue
		} else if index2+1 == len2 && nums1[index1] >= nums2[index2] {
			return median(nums1, nums2, index1, index1+1, index2+1)
		} else if index2+1 == len2 && nums1[index1] < nums2[index2] {
			left = index1 + 1
			continue
		} else if nums1[index1] >= nums2[index2] && nums1[index1] <= nums2[index2+1] {
			return median(nums1, nums2, index1, index1+1, index2+1)
		} else if nums1[index1] >= nums2[index2] {
			right = index1
		} else {
			left = index1 + 1
		}
	}
	return findMedianSortedArrays(nums2, nums1)
}
