package main

type stack struct {
	data []byte
}

func (s *stack) push(b byte) {
	s.data = append(s.data, b)
}
func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) pop() *byte {
	if len(s.data) == 0 {
		return nil
	}
	b := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &b
}

func isValid(s string) bool {
	bs := []byte(s)
	st := &stack{}
	for _, b := range bs {
		if b == '(' || b == '{' || b == '[' {
			st.push(b)
			continue
		}
		top := st.pop()
		if top == nil {
			return false
		}
		if (b == ')' && *top != '(') ||
			(b == ']' && *top != '[') ||
			(b == '}' && *top != '{') {
			return false
		}
	}
	return st.len() == 0
}
