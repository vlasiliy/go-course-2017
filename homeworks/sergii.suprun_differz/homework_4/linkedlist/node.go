package linkedlist

type node struct {
	value interface{}
	next  *node
}

func newNode(value interface{}, next *node) *node {
	return &node{value, next}
}

func (n *node) getValue() interface{} {
	return n.value
}

func (n *node) setValue(value interface{}) {
	n.value = value
}

func (n *node) getNext() *node {
	return n.next
}

func (n *node) setNext(next *node) {
	n.next = next
}
