package main

import "fmt"

func main() {
	list := DoublyLinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	array := list.traverse()
	fmt.Println(array)
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (list *DoublyLinkedList) Insert(value int) {
	newNode := Node{value, nil, nil}
	if list.Head == nil {
		list.Head = &newNode
		list.Tail = &newNode
		return
	}

	list.Tail.next = &newNode
	newNode.prev = list.Tail
	list.Tail = list.Tail.next
}

func (list *DoublyLinkedList) Delete(value int) {

}

func (list *DoublyLinkedList) traverse() []int {
	array := []int{}
	node := list.Head
	for node != nil {
		array = append(array, node.value)
		node = node.next
	}
	return array
}
