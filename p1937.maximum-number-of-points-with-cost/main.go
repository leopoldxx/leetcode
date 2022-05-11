package main

func split(points []int, intoLR []int, intoRL []int) {
	for i := 0; i < len(points); i++ {
		intoLR[i] = points[i] + i
		intoRL[i] = points[i] - i
	}
}

func maxLeftToRight(points []int, preSumLeft []int) {
	max := preSumLeft[0]
	for i := 0; i < len(points); i++ {
		if max < preSumLeft[i] {
			max = preSumLeft[i]
		}
		preSumLeft[i] = max + points[i] - i
	}
}

func maxRightToLeft(points []int, preSumRight []int) {
	n := len(points)
	max := points[n-1]
	for i := n - 1; i >= 0; i-- {
		if max < preSumRight[i] {
			max = preSumRight[i]
		}
		preSumRight[i] = max + points[i] + i
	}
}

func mergeMax(points []int, preSumLeft []int, preSumRight []int) {
	n := len(points)
	for i := 0; i < n; i++ {
		if preSumLeft[i] < preSumRight[i] {
			points[i] = preSumRight[i]
		} else {
			points[i] = preSumLeft[i]
		}
	}
}

func maxPoints(points [][]int) int64 {
	n := len(points)
	m := len(points[0])

	preSumLeft := make([]int, m)
	preSumRight := make([]int, m)

	split(points[0], preSumLeft, preSumRight)

	for i := 1; i < n; i++ {
		maxLeftToRight(points[i], preSumLeft)
		maxRightToLeft(points[i], preSumRight)
		mergeMax(points[i], preSumLeft, preSumRight)
		split(points[i], preSumLeft, preSumRight)
	}

	max := points[n-1][0]
	for i := 0; i < m; i++ {
		if max < points[n-1][i] {
			max = points[n-1][i]
		}
	}
	return int64(max)

}
