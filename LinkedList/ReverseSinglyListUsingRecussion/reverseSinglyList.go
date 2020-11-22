package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	temp   *Node
	length int
}

func (l *LinkedList) createList(newNode *Node) {
	if l.head == nil {
		l.head = newNode
		l.temp = newNode
		l.length++
	} else {
		l.temp.next = newNode
		l.temp = newNode
		l.length++
	}
}

func (l *LinkedList) Print() {
	l.temp = l.head
	if l.temp == nil {
		fmt.Println("No node ")
	}

	for l.temp != nil {
		fmt.Printf("%v ->", l.temp.data)

		l.temp = l.temp.next
	}

}
func reverseListUsingRecursion(head *Node) {
	if head == nil {
		return
	}
	reverseListUsingRecursion(head.next)
	fmt.Printf("%v ->", head.data)
}

func main() {
	linkedList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	linkedList.createList(node1)
	linkedList.createList(node2)
	linkedList.createList(node3)
	linkedList.createList(node4)
	reverseListUsingRecursion(linkedList.head)
	fmt.Println()
	linkedList.Print()
}
