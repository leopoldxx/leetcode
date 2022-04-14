package main

type stack struct {
	data []int
}

func (s *stack) push(v int) {
	s.data = append(s.data, v)
}
func (s *stack) pop() *int {
	if len(s.data) == 0 {
		return nil
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &top
}
func (s *stack) top() *int {
	if len(s.data) == 0 {
		return nil
	}
	return &s.data[len(s.data)-1]
}
func (s *stack) len() int {
	return len(s.data)
}

type MyQueue struct {
	in  *stack
	out *stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  &stack{},
		out: &stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.in.push(x)
}

func (this *MyQueue) transfer() {
	if this.out.len() == 0 {
		top := this.in.pop()
		for top != nil {
			this.out.push(*top)
			top = this.in.pop()
		}
	}
}

func (this *MyQueue) Pop() int {
	this.transfer()
	top := this.out.pop()
	if top == nil {
		return 0
	}
	return *top
}

func (this *MyQueue) Peek() int {
	this.transfer()
	top := this.out.top()
	if top == nil {
		return 0
	}
	return *top
}

func (this *MyQueue) Empty() bool {
	return this.in.len() == 0 && this.out.len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
