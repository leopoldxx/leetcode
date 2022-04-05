package main

type matrix struct {
	data [][]int
	row  int
	col  int

	currentRow int
	currentCol int
}

func (m *matrix) init(mat [][]int) {
	m.data = mat
	m.row = len(mat)
	m.col = len(mat[0])
	m.currentRow = 0
	m.currentCol = 0
}
func (m *matrix) initWithD(r, c int) {
	m.row = r
	m.col = c
	m.currentRow = 0
	m.currentCol = 0
}
func (m *matrix) next() *int {

	if m.currentRow == m.row {
		return nil
	}

	res := &m.data[m.currentRow][m.currentCol]
	m.currentCol++
	if m.currentCol == m.col {
		m.currentRow++
		m.currentCol = 0
	}
	return res
}
func (m *matrix) append(v int) {
	if m.currentRow == m.row && m.currentCol == m.col {
		return
	}
	if m.data == nil {
		m.data = make([][]int, m.row)
	}
	if m.data[m.currentRow] == nil {
		m.data[m.currentRow] = make([]int, m.col)
	}

	m.data[m.currentRow][m.currentCol] = v
	m.currentCol++
	if m.currentCol == m.col {
		m.currentRow++
		m.currentCol = 0
	}
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rr := len(mat)
	cc := len(mat[0])

	if rr*cc != r*c {
		return mat
	}

	src := &matrix{}
	src.init(mat)

	dest := &matrix{}
	dest.initWithD(r, c)

	for v := src.next(); v != nil; v = src.next() {
		dest.append(*v)
	}
	return dest.data

}
