package main

func compare(left, right []int) int {
	if left[1] < right[0] {
		return 1
	} else if left[0] <= right[0] && left[1] <= right[1] {
		return 2
	} else if left[0] <= right[0] && left[1] > right[1] {
		return 3
	} else if left[0] > right[0] && left[1] <= right[1] {
		return 4
	} else if left[0] > right[0] && left[0] <= right[1] && left[1] > right[1] {
		return 5
	} else {
		return 6
	}
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	for {
		if len(firstList) == 0 || len(secondList) == 0 {
			break
		}
		A := firstList[0]
		B := secondList[0]
		ret := compare(A, B)
		switch ret {
		case 1:
			firstList = firstList[1:]
		case 2:
			res = append(res, []int{B[0], A[1]})
			firstList = firstList[1:]
		case 3:
			res = append(res, []int{B[0], B[1]})
			secondList = secondList[1:]
		case 4:
			res = append(res, []int{A[0], A[1]})
			firstList = firstList[1:]
		case 5:
			res = append(res, []int{A[0], B[1]})
			secondList = secondList[1:]
		case 6:
			secondList = secondList[1:]
		}

	}
	return res
}
