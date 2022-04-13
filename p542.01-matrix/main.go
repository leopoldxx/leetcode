package main

type queue struct {
	data [][]int
}

func (q *queue) push(c []int) {
	q.data = append(q.data, c)
}
func (q *queue) pop() []int {
	if len(q.data) == 0 {
		return nil
	}
	top := q.data[0]
	q.data = q.data[1:]
	return top
}

func setDistance(x, y, d int, mat [][]int) []int {
	if x < 0 || x == len(mat) || y < 0 || y == len(mat[0]) || mat[x][y] >= 0 {
		return nil
	}
	mat[x][y] = d
	return []int{x, y}
}

func updateMatrix(mat [][]int) [][]int {
	if mat == nil {
		return nil
	}
	res := make([][]int, len(mat))
	que := &queue{}

	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
				que.push([]int{i, j})
			} else {
				res[i][j] = -1
			}
		}
	}
	pushIfNotNil := func(c []int) {
		if c != nil {
			que.push(c)
		}
	}
	top := que.pop()
	for top != nil {
		x := top[0]
		y := top[1]
		d := res[x][y]
		pushIfNotNil(setDistance(x-1, y, d+1, res))
		pushIfNotNil(setDistance(x+1, y, d+1, res))
		pushIfNotNil(setDistance(x, y-1, d+1, res))
		pushIfNotNil(setDistance(x, y+1, d+1, res))

		top = que.pop()

	}
	return res
}
