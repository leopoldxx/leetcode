package main

func distance(x1, y1, x2, y2 int) int {
	x := x1 - x2
	y := y1 - y2
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func nearestValidPoint(x int, y int, points [][]int) int {
	min := -1
	minDistance := -1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			d := distance(x, y, p[0], p[1])
			if min == -1 {
				min = i
				minDistance = d
			} else if d < minDistance {
				min = i
				minDistance = d
			}
		}
	}
	return min
}
