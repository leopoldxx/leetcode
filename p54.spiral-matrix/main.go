package main

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

type direct func() direct

type result struct {
	offset int
	xspan  int
	yspan  int
	matrix [][]int
	maxLen int
	data   []int
}

func (r *result) init() direct {
	r.maxLen = r.xspan * r.yspan
	r.data = make([]int, 0, r.maxLen)
	return r.right
}

func (r *result) right() direct {
	x := r.offset
	y := r.offset
	maxy := r.yspan - r.offset
	for y < maxy {
		r.data = append(r.data, r.matrix[x][y])
		y++
	}
	if len(r.data) == r.maxLen {
		return nil
	}
	return r.down
}
func (r *result) down() direct {
	x := r.offset + 1
	y := r.yspan - 1 - r.offset
	maxx := r.xspan - r.offset
	for x < maxx {
		r.data = append(r.data, r.matrix[x][y])
		x++
	}
	if len(r.data) == r.maxLen {
		return nil
	}
	return r.left

}
func (r *result) left() direct {
	x := r.xspan - r.offset - 1
	y := r.yspan - r.offset - 2
	miny := r.offset
	for y >= miny {
		r.data = append(r.data, r.matrix[x][y])
		y--
	}
	if len(r.data) == r.maxLen {
		return nil
	}
	return r.up

}
func (r *result) up() direct {
	x := r.xspan - r.offset - 2
	y := r.offset
	minx := r.offset
	for x > minx {
		r.data = append(r.data, r.matrix[x][y])
		x--
	}
	if len(r.data) == r.maxLen {
		return nil
	}
	r.offset++
	return r.right
}

func spiralOrder(matrix [][]int) []int {
	xspan, yspan := len(matrix), len(matrix[0])
	r := &result{
		xspan:  xspan,
		yspan:  yspan,
		matrix: matrix,
	}

	direct := r.init()
	for direct != nil {
		direct = direct()
	}
	return r.data
}
