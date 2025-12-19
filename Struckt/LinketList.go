package main

func main() {

}

type Node struct {
	Value any
	Next  *Node
}

type LinketList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinketList() *LinketList {
	return &LinketList{}
}

func (ll *LinketList) Append(v any) {
	newNode := &Node{Value: v}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}
