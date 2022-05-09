package main

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type stack struct {
	data []*NestedInteger
}

func (s *stack) push(v *NestedInteger) {
	if v.IsInteger() {
		s.data = append(s.data, v)
	} else {
		l := v.GetList()
		if len(l) == 0 {
			return
		}
		for k := len(l) - 1; k >= 0; k-- {
			s.data = append(s.data, l[k])
		}
	}
}
func (s *stack) top() *NestedInteger {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}
func (s *stack) pop() *NestedInteger {
	if len(s.data) == 0 {
		return nil
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

type NestedIterator struct {
	s *stack
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	it := &NestedIterator{
		s: &stack{},
	}
	for jj := len(nestedList) - 1; jj >= 0; jj-- {
		it.s.push(nestedList[jj])
	}
	return it
}

func (this *NestedIterator) Next() int {
	ti := this.s.top()
	if ti == nil {
		panic(ti)
	}
	v := ti.GetInteger()
	this.s.pop()
	return v
}

func (this *NestedIterator) HasNext() bool {
	ti := this.s.top()
	if ti == nil {
		return false
	} else if ti.IsInteger() {
		return true
	}
	this.s.pop()
	nestedList := ti.GetList()
	for jj := len(nestedList) - 1; jj >= 0; jj-- {
		this.s.push(nestedList[jj])
	}
	return this.HasNext()

}
