package main

type generator struct {
	n      int
	offset int
	num    int
	maxNum int
	data   [][]int
}

func (g *generator) init(n int) {
	g.n = n
	g.data = make([][]int, n)
	for i := 0; i < n; i++ {
		g.data[i] = make([]int, n)
	}
	g.maxNum = g.n*g.n + 1
	g.num = 1
}

type next func() next

func (g *generator) right() next {
	maxj := g.n - g.offset
	indexi := g.offset
	for indexj := g.offset; indexj < maxj; indexj++ {
		g.data[indexi][indexj] = g.num
		g.num++
	}
	if g.num == g.maxNum {
		return nil
	}
	return g.down
}
func (g *generator) down() next {
	maxi := g.n - g.offset
	indexj := g.n - g.offset - 1
	for indexi := g.offset + 1; indexi < maxi; indexi++ {
		g.data[indexi][indexj] = g.num
		g.num++
	}
	if g.num == g.maxNum {
		return nil
	}
	return g.left

}
func (g *generator) left() next {
	minj := g.offset
	indexi := g.n - g.offset - 1
	for indexj := g.n - g.offset - 2; indexj >= minj; indexj-- {
		g.data[indexi][indexj] = g.num
		g.num++
	}
	if g.num == g.maxNum {
		return nil
	}
	return g.up
}
func (g *generator) up() next {
	mini := g.offset
	indexj := g.offset
	for indexi := g.n - g.offset - 2; indexi > mini; indexi-- {
		g.data[indexi][indexj] = g.num
		g.num++
	}
	if g.num == g.maxNum {
		return nil
	}
	g.offset++
	return g.right
}

func generateMatrix(n int) [][]int {
	g := &generator{}
	g.init(n)
	next := g.right()
	for next != nil {
		next = next()
	}
	return g.data

}
