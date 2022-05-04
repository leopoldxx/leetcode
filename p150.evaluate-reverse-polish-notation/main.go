package main

import "strconv"

type stack struct {
	tokens []int
}

func (s *stack) push(v int) {
	s.tokens = append(s.tokens, v)
}
func (s *stack) pop() *int {
	if len(s.tokens) == 0 {
		return nil
	}
	top := s.tokens[len(s.tokens)-1]
	s.tokens = s.tokens[:len(s.tokens)-1]
	return &top
}

func calc(left, right int, op string) int {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return 0
}

func evalRPN(tokens []string) int {
	s := &stack{}

	n := len(tokens)
	for i := 0; i < n; i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			right := s.pop()
			left := s.pop()
			s.push(calc(*left, *right, tokens[i]))
			continue
		}
		v, _ := strconv.Atoi(tokens[i])
		s.push(v)
	}
	return *s.pop()
}
