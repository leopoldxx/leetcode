package main

type node struct {
	isWord bool
	next   map[byte]*node
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{
		root: &node{},
	}
}

func (this *Trie) Insert(word string) {
	ws := []byte(word)
	it := this.root
	for i := 0; i < len(ws); i++ {
		c := ws[i]
		if it.next == nil {
			it.next = map[byte]*node{}
		}
		if _, ok := it.next[c]; !ok {
			it.next[c] = &node{}
		}
		it = it.next[c]
	}
	it.isWord = true
}

func (this *Trie) Search(word string) bool {
	ws := []byte(word)
	it := this.root
	for i := 0; i < len(ws); i++ {
		c := ws[i]
		if _, ok := it.next[c]; !ok {
			return false
		}
		it = it.next[c]
	}
	return it.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	ws := []byte(prefix)
	it := this.root
	for i := 0; i < len(ws); i++ {
		c := ws[i]
		if _, ok := it.next[c]; !ok {
			return false
		}
		it = it.next[c]
	}
	return true
}
