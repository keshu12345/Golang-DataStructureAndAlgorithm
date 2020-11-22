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

func (l *LinkedList) PrintList() {
	if l.length == 0 {
		fmt.Println("No node available ")
	}

	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}
}

func main() {

	linkedList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	fmt.Println(linkedList)
	linkedList.createList(node1)
	linkedList.createList(node2)
	linkedList.createList(node3)
	linkedList.PrintList()
	fmt.Println()

	fmt.Println(linkedList)
}
