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

func orangesRotting(grid [][]int) int {
	if grid == nil {
		return 0
	}
	que := &queue{}
	fresh, rotted := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				que.push([]int{i, j, 0})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	pushIfNotNil := func(c []int) {
		if c != nil {
			que.push(c)
		}
	}
	rot := func(i, j, m int, grid [][]int) []int {
		if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == 0 || grid[i][j] == 2 {
			return nil
		}
		rotted++
		grid[i][j] = 2
		return []int{i, j, m}
	}

	top := que.pop()
	last := top
	for top != nil {
		i, j, m := top[0], top[1], top[2]
		pushIfNotNil(rot(i-1, j, m+1, grid))
		pushIfNotNil(rot(i+1, j, m+1, grid))
		pushIfNotNil(rot(i, j-1, m+1, grid))
		pushIfNotNil(rot(i, j+1, m+1, grid))

		last = top
		top = que.pop()
	}
	if fresh-rotted != 0 {
		return -1
	}
	if last == nil {
		return 0
	}
	return last[2]
}
