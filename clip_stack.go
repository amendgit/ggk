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
	toimpl()
	return
}
