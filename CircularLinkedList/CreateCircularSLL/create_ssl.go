package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	temp   *Node
	length int
}

func (l *LinkedList) createCircularLinkedList(newNode *Node) {

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		//l.tail.next = l.head
		l.length++
	} else {
		l.tail.next = newNode
		l.tail = newNode
		//l.tail.next = l.head  //
		l.length++
	}
	l.tail.next = l.head
}

func (l *LinkedList) PrintCircularLinkedList() {
	fmt.Printf("%v->", l.tail.next.data)
	l.temp = l.head.next
	for l.temp != l.head {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.next

	}
}

func main() {

	linkedLinst := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	node5 := &Node{data: 50}
	linkedLinst.createCircularLinkedList(node1)
	linkedLinst.createCircularLinkedList(node2)
	linkedLinst.createCircularLinkedList(node3)
	linkedLinst.createCircularLinkedList(node4)
	linkedLinst.createCircularLinkedList(node5)
	linkedLinst.PrintCircularLinkedList()
}
