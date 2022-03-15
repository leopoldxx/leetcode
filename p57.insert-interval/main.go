package main

func merge(left, right []int) ([]int, []int) {
	if left[0] <= right[0] && left[1] >= right[1] {
		return left, nil
	} else if right[0] <= left[0] && right[1] >= left[1] {
		return right, nil
	} else if left[0] <= right[0] && left[1] >= right[0] && left[1] <= right[1] {
		return []int{left[0], right[1]}, nil
	} else if right[0] <= left[0] && right[1] >= left[0] && right[1] <= left[1] {
		return []int{right[0], left[1]}, nil
	} else if left[1] < right[0] {
		return left, right
	} else {
		return right, left
	}

}
func appendLeft(result [][]int, left [][]int) [][]int {
	return append(result, left...)
}
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	for i := 0; i < len(intervals); i++ {
		a, b := merge(intervals[i], newInterval)

		if b == nil {
			newInterval = a
		} else {
			result = append(result, a)
			newInterval = b
			if intervals[i][0] == newInterval[0] {
				return appendLeft(result, intervals[i:])
			}
		}
	}
	return append(result, newInterval)

}
