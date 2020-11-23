package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	temp   *Node
	length string
}

func (l *LinkedList) createDublyLinkedList(newNode *Node) {
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}

}

func (l *LinkedList) PrintDublyLinkedList() {
	l.temp = l.head
	if l.head == nil {
		fmt.Println("No Node Available in LinkedList")
	}

	for l.temp != nil {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.next
	}

	fmt.Println()
	for l.temp != nil {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.prev
	}
}
func (l *LinkedList) PrintReverseDoublyLinkedList() {
	l.temp = l.tail
	fmt.Println()
	for l.temp != nil {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.prev
	}
}
func main() {
	linkedList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	linkedList.createDublyLinkedList(node1)
	linkedList.createDublyLinkedList(node2)
	linkedList.createDublyLinkedList(node3)
	linkedList.createDublyLinkedList(node4)
	linkedList.PrintDublyLinkedList()
	linkedList.PrintReverseDoublyLinkedList()
}
