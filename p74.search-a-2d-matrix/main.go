package main

func searchRow(matrix [][]int, target int) int {
	top, bottom := 0, len(matrix)-1
	for {
		if top > bottom {
			return bottom
		}
		mid := top + (bottom-top)/2
		if matrix[mid][0] == target {
			return mid
		} else if matrix[mid][0] > target {
			bottom = mid - 1
		} else {
			top = mid + 1
		}
	}
}

func searchCol(row []int, target int) bool {
	left, right := 0, len(row)-1
	for {
		if left > right {
			return false
		}
		mid := left + (right-left)/2
		if row[mid] == target {
			return true
		} else if row[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	row := searchRow(matrix, target)
	if row < 0 {
		return false
	}
	return searchCol(matrix[row], target)
}
