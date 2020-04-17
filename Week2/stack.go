package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data int
	min  int
}

type MinStack struct {
	top *Node
}

func getNode(data, min int) *Node {
	return &Node{next: nil, prev: nil, data: data, min: min}
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{top: nil}
}

func (this *MinStack) Push(x int) {
	if this.top == nil {
		node := getNode(x, x)
		this.top = node
	} else {
		currentMin := this.top.min
		if this.top.min > x {
			currentMin = x
		}
		node := getNode(x, currentMin)
		node.prev = this.top
		this.top = node
	}
}

func (this *MinStack) Pop() {
	if this.top != nil && this.top.prev != nil {
		this.top = this.top.prev
	} else {
		this.top = nil
	}
}

func (this *MinStack) Top() int {
	if this.top == nil {
		return 0
	} else {
		return this.top.data
	}
}

func (this *MinStack) GetMin() int {
	return this.top.min
}

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Push(12)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	min := obj.GetMin()
}
