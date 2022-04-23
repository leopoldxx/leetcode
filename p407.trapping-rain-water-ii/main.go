package main

type coordheight [3]int
type coord [2]int

type pqueue struct {
	data    []coordheight
	visited map[coord]bool
}

func (pq *pqueue) init() {
	pq.visited = make(map[coord]bool)
}
func (pq *pqueue) push(x coordheight) {
	c := coord{x[0], x[1]}
	if pq.visited[c] {
		return
	}
	pq.visited[c] = true

	pq.data = append(pq.data, x)
	pq.siftup(len(pq.data) - 1)
}

func (pq *pqueue) isVisited(x, y int) bool {
	return pq.visited[coord{x, y}]
}
func (pq *pqueue) parentIndex(idx int) int {
	return (idx - 1) / 2
}
func (pq *pqueue) leftIndex(idx int) int {
	return idx*2 + 1
}
func (pq *pqueue) rightIndex(idx int) int {
	return idx*2 + 2
}
func (pq *pqueue) in(idx int) bool {
	return idx < len(pq.data)
}

func (pq *pqueue) swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}
func (pq *pqueue) siftup(idx int) (changed bool) {
	for {
		if idx == 0 {
			return
		}
		parent := pq.parentIndex(idx)
		if pq.data[parent][2] <= pq.data[idx][2] {
			return
		}
		pq.swap(parent, idx)
		changed = true
		idx = parent
	}
}
func (pq *pqueue) siftdown(idx int) (changed bool) {
	for {
		left := pq.leftIndex(idx)
		if !pq.in(left) {
			return
		}
		candidate := left
		right := pq.rightIndex(idx)
		if pq.in(right) && pq.data[right][2] < pq.data[left][2] {
			candidate = right
		}
		if pq.data[candidate][2] >= pq.data[idx][2] {
			return
		}
		pq.swap(idx, candidate)
		idx = candidate
		changed = true
	}
}

func (pq *pqueue) top() *coordheight {
	if len(pq.data) == 0 {
		return nil
	}
	return &pq.data[0]
}

func (pq *pqueue) pop() *coordheight {
	if len(pq.data) == 0 {
		return nil
	}
	top := pq.data[0]
	pq.swap(0, len(pq.data)-1)
	pq.data = pq.data[:len(pq.data)-1]
	pq.siftdown(0)
	return &top
}

func surroundings(heightMap [][]int, c *coordheight) (res []coord) {
	rowCount := len(heightMap)
	columnCount := len(heightMap[0])

	inMap := func(x, y int) bool {
		return x >= 0 && x < rowCount-1 && y >= 0 && y < columnCount
	}
	x, y := c[0], c[1]
	if inMap(x-1, y) {
		res = append(res, coord{x - 1, y})
	}
	if inMap(x+1, y) {
		res = append(res, coord{x + 1, y})
	}
	if inMap(x, y-1) {
		res = append(res, coord{x, y - 1})
	}
	if inMap(x, y+1) {
		res = append(res, coord{x, y + 1})
	}
	return
}

func trapRainWater(heightMap [][]int) int {
	rowCount := len(heightMap)
	columnCount := len(heightMap[0])
	capacity := 0

	if rowCount <= 2 || columnCount <= 2 {
		return 0
	}
	pq := &pqueue{}
	pq.init()

	for i := 0; i < rowCount; i++ {
		pq.push(coordheight{i, 0, heightMap[i][0]})
		pq.push(coordheight{i, columnCount - 1, heightMap[i][columnCount-1]})
	}
	for i := 0; i < columnCount; i++ {
		pq.push(coordheight{0, i, heightMap[0][i]})
		pq.push(coordheight{rowCount - 1, i, heightMap[rowCount-1][i]})
	}

	for pq.top() != nil {
		lowlying := pq.pop()
		coords := surroundings(heightMap, lowlying)
		for _, c := range coords {
			if pq.isVisited(c[0], c[1]) {
				continue
			}
			if lowlying[2] > heightMap[c[0]][c[1]] {
				capacity += (lowlying[2] - heightMap[c[0]][c[1]])
				heightMap[c[0]][c[1]] = lowlying[2]
			}
			pq.push(coordheight{c[0], c[1], heightMap[c[0]][c[1]]})
		}
	}
	return capacity
}
