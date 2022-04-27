package main

type unionFind struct {
	data []int
}

func (uf *unionFind) init(length int) {
	uf.data = make([]int, length)
	for i := 0; i < length; i++ {
		uf.data[i] = i
	}
}
func (uf *unionFind) union(i, j int) {
	if i > j {
		i, j = j, i
	}
	iroot := uf.find(i)
	jroot := uf.find(j)
	if iroot == jroot {
		return
	}

	uf.data[jroot] = iroot
}

func (uf *unionFind) find(i int) (root int) {
	for uf.data[i] != i {
		uf.data[i] = uf.data[uf.data[i]]
		i = uf.data[i]
	}
	return i
}

func (uf *unionFind) set() map[int][]int {
	res := make(map[int][]int)
	for i := 0; i < len(uf.data); i++ {
		root := uf.find(i)
		res[root] = append(res[root], i)
	}
	return res
}

func getByte(c []int) (res byte) {
	for i := 0; i < 26; i++ {
		if c[i] == 0 {
			continue
		}
		res = 'a' + byte(i)
		c[i]--
		break
	}
	return
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	bs := []byte(s)
	uf := &unionFind{}
	uf.init(len(bs))

	for i := 0; i < len(pairs); i++ {
		uf.union(pairs[i][0], pairs[i][1])
	}

	sets := uf.set()

	newset := make([][]int, len(bs))
	for _, set := range sets {
		count := make([]int, 26)
		for j := 0; j < len(set); j++ {
			idx := set[j]
			count[bs[idx]-'a']++
			newset[idx] = count
		}
	}
	for i := 0; i < len(bs); i++ {
		bs[i] = getByte(newset[i])
	}
	return string(bs)
}
