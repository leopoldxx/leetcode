package main

type item struct {
	value int
	min   int
}
type MinStack struct {
	data []item
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.data = append(this.data, item{value: val, min: val})
		return
	}
	top := &this.data[len(this.data)-1]
	min := top.min
	if val < min {
		min = val
	}
	this.data = append(this.data, item{value: val, min: min})
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	top := &this.data[len(this.data)-1]
	return top.value
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return 0
	}
	top := &this.data[len(this.data)-1]
	return top.min
}
