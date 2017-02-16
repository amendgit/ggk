package ggk

import "container/list"

type ClipStack struct {
	deque     *list.List
	saveCount int
}

func NewClipStack() *ClipStack {
	var stack = &ClipStack{
		deque:     list.New(),
		saveCount: 0,
	}
	return stack
}

func (stack *ClipStack) Reset(otr *ClipStack) {
	// We used a placement new for each object in deque. so we're responsible
	// for calling the destructor on each of them as well.
	stack.deque = list.New()
	stack.saveCount = 0
	return
}
