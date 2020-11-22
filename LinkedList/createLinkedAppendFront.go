package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) frontAppend(newNode *Node) {

	if l.head == nil {
		l.head = newNode
		l.length++
	} else {
		newNode.next = l.head
		l.head = newNode
		l.length++
	}

}

func (l *LinkedList) Print() {
	if l.length == 0 {
		fmt.Println("NO Node")
	}

	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}

}

func main() {
	linkeList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	fmt.Println(linkeList)

	linkeList.frontAppend(node1)
	linkeList.frontAppend(node2)
	linkeList.frontAppend(node3)
	fmt.Println(linkeList)

	linkeList.Print()

}
