package main

type stack struct {
	data []interface{}
}

func (s *stack) push(elem interface{}) {
	s.data = append(s.data, elem)
}

func (s *stack) pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	ele := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ele
}

const (
	leftParenthesis     = byte('(')
	rightParenthesis    = byte(')')
	combinedParentheses = byte('|')
)

func scoreList(scores []int) int {
	if len(scores) == 1 {
		return scores[0] * 2
	}
	var result int
	for _, s := range scores {
		result += s
	}
	return result * 2
}

func scoreOfParentheses(s string) int {
	parenthesesStack := &stack{}
	scoreStack := &stack{}

	for _, c := range []byte(s) {
		if c == leftParenthesis {
			parenthesesStack.push(leftParenthesis)
			continue
		}
		// right parenthesis
		parenthesis := parenthesesStack.pop().(byte)
		// case 1: ()
		if parenthesis == leftParenthesis {
			parenthesesStack.push(combinedParentheses)
			scoreStack.push(int(1))
		} else {
			// case 2: (|) or (||)
			var tempList []int
			tempList = append(tempList, scoreStack.pop().(int))
			parenthesis = parenthesesStack.pop().(byte)
			for parenthesis != leftParenthesis {
				tempList = append(tempList, scoreStack.pop().(int))
				parenthesis = parenthesesStack.pop().(byte)
			}
			scoreStack.push(scoreList(tempList))
			parenthesesStack.push(combinedParentheses)

		}

	}

	var result int
	for s := scoreStack.pop(); s != nil; s = scoreStack.pop() {
		result += s.(int)
	}
	return result
}
