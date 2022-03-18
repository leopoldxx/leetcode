package main

type stack struct {
	data []byte
}

func (s *stack) push(c byte) {
	s.data = append(s.data, c)
}
func (s *stack) len() int {
	return len(s.data)
}
func (s *stack) top() byte {
	return s.data[len(s.data)-1]
}
func (s *stack) pop() byte {
	c := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return c
}
func (s *stack) raw() []byte {
	return s.data
}

func removeDuplicateLetters(s string) string {
	count := map[byte]int{}

	for _, c := range s {
		count[byte(c)] = count[byte(c)] + 1
	}

	st := &stack{}
	visited := [26]bool{}
	for _, c := range s {
		bc := byte(c)
		index := bc - 'a'
		count[bc]--
		if visited[index] {
			continue
		}

		for {
			if st.len() == 0 {
				break
			}
			topc := st.top()
			if topc < bc || count[topc] == 0 {
				break
			}
			st.pop()
			visited[topc-'a'] = false

		}
		visited[index] = true
		st.push(bc)
	}

	return string(st.raw())
}
