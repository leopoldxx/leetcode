package main

type unionFind struct {
	data []int
}

func (uf *unionFind) init(count int) {
	uf.data = make([]int, count)
	for i := 0; i < count; i++ {
		uf.data[i] = i
	}
}
func (uf *unionFind) union(i, j int) {
	ir := uf.find(i)
	jr := uf.find(j)
	if ir == jr {
		return
	}
	if ir > jr {
		ir, jr = jr, ir
	}
	uf.data[jr] = ir
}
func (uf *unionFind) find(i int) int {
	for uf.data[i] != i {
		uf.data[i] = uf.data[uf.data[i]]
		i = uf.data[i]
	}
	return uf.data[i]
}
func (uf *unionFind) set() map[int][]int {
	res := map[int][]int{}
	for i := 0; i < len(uf.data); i++ {
		root := uf.find(i)
		res[root] = append(res[root], i)
	}
	return res
}

func findCircleNum(isConnected [][]int) int {
	count := len(isConnected)
	uf := &unionFind{}
	uf.init(count)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)

			}
		}
	}
	return len(uf.set())
}
