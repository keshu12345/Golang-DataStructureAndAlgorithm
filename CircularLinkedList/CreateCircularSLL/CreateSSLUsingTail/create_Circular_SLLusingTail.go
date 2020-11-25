package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	tail   *Node
	temp   *Node
	length int
}

func (l *LinkedList) CreateSSLUsingTail(newNode *Node) {
	if l.tail == nil {
		l.tail = newNode
		l.tail.next = l.tail
	} else {
		newNode.next = l.tail.next
		l.tail.next = newNode
		l.tail = newNode

	}

}

func (l *LinkedList) PrintCSLL() {

	fmt.Printf("%v ->", l.tail.next.data)
	l.temp = l.tail.next.next
	for l.temp != l.tail.next {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.next
	}
}

func main() {
	linkedList := new(LinkedList)
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	linkedList.CreateSSLUsingTail(node1)
	linkedList.CreateSSLUsingTail(node2)
	linkedList.CreateSSLUsingTail(node3)
	linkedList.CreateSSLUsingTail(node4)
	linkedList.PrintCSLL()

}
