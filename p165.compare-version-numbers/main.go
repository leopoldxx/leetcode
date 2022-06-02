package main

import (
	"strconv"
	"strings"
)

func convert(strs []string) []int {
	res := make([]int, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		n, _ := strconv.Atoi(strs[i])
		res = append(res, n)
	}
	return res
}

func pad0(nums *[]int, count int) {
	for i := 0; i < count; i++ {
		*nums = append(*nums, 0)
	}
}

func pad(nums1 *[]int, nums2 *[]int) {
	n1 := len(*nums1)
	n2 := len(*nums2)
	if n1 < n2 {
		pad0(nums1, n2-n1)
	} else if n1 > n2 {
		pad0(nums2, n1-n2)
	}
}

func compareVersion(version1 string, version2 string) int {
	version1nums := convert(strings.Split(version1, "."))
	version2nums := convert(strings.Split(version2, "."))
	pad(&version1nums, &version2nums)
	n := len(version1nums)
	for i := 0; i < n; i++ {
		if version1nums[i] < version2nums[i] {
			return -1
		} else if version1nums[i] > version2nums[i] {
			return 1
		}
		continue
	}
	return 0
}
