// Reverse the DoublyLinkedList Time COmplexity O(n)

package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head        *Node
	tail        *Node
	temp        *Node
	length      int
	currentNode *Node
	nextNode    *Node
}

func (l *LinkedList) createDublyLinkedList(newNode *Node) {
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		l.length++
	}

}

func (l *LinkedList) reverseDLL() {
	if l.head == nil {
		fmt.Println("No Node Available")
	} else {
		l.currentNode = l.head
		l.nextNode = l.head
		for l.nextNode != nil {
			l.nextNode = l.nextNode.next
			l.currentNode.next = l.currentNode.prev
			l.currentNode.prev = l.nextNode
			l.currentNode = l.nextNode

		}

		l.currentNode = l.tail
		l.tail = l.head
		l.head = l.currentNode
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
	node5 := &Node{data: 50}
	node6 := &Node{data: 60}
	node7 := &Node{data: 70}
	linkedList.createDublyLinkedList(node1)
	linkedList.createDublyLinkedList(node2)
	linkedList.createDublyLinkedList(node3)
	linkedList.createDublyLinkedList(node4)
	linkedList.createDublyLinkedList(node5)
	linkedList.createDublyLinkedList(node6)
	linkedList.createDublyLinkedList(node7)
	linkedList.reverseDLL()
	linkedList.PrintDublyLinkedList()
}
