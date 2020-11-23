package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head        *Node
	temp        *Node
	currentNode *Node
	nextNode    *Node
	prevNode    *Node
	length      int
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

	l.temp = l.head
	if l.head == nil {
		fmt.Println("No Node Available")
	}

	for l.temp != nil {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.next
	}
}

func (l *LinkedList) reverseIterative() {
	l.prevNode = nil
	l.nextNode = l.head
	l.currentNode = l.head

	for l.nextNode != nil {
		l.nextNode = l.nextNode.next
		l.currentNode.next = l.prevNode
		l.prevNode = l.currentNode
		l.currentNode = l.nextNode
	}
	l.head = l.prevNode

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
	linkedList.PrintList()
	linkedList.reverseIterative()
	fmt.Println()
	fmt.Print("Reverse the LinkedList is::::")
	linkedList.PrintList()
}
