// Add the key  at front position time complexity O(1) and Space Complexity O(1)
//Add the key at the end position of Linkedlist time complexity O(n) and Space Complexity O(1)
// Add the at any position Time comlexity O(n) and Space comlexity O(1)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (l *LinkedList) create(newNode *Node) {

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

func (l *LinkedList) AddAtBeginning(newNode *Node) {
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *LinkedList) AddAtEnd(newNode *Node) {
	l.temp = l.head
	for l.temp.next != nil {
		l.temp = l.temp.next
	}
	l.temp.next = newNode
	l.length++
}

func (l *LinkedList) AddAtPos(newNode *Node, pos int) {
	i := 1
	l.temp = l.head
	if pos > l.length {
		fmt.Println("Invalid Position Entered")
	} else {
		for i < pos {
			l.temp = l.temp.next
			i++
		}

		newNode.next = l.temp.next
		l.temp.next = newNode
		l.length++
	}
}

func (l *LinkedList) AllPrint() {
	if l.length == 0 {
		fmt.Println("No Node")
	}

	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	pos, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Println(err)
	}

	linkedList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	node5 := &Node{data: 60}
	node6 := &Node{data: 80}
	linkedList.create(node1)
	linkedList.create(node2)
	linkedList.create(node3)
	linkedList.AddAtBeginning(node4)
	linkedList.AddAtEnd(node5)
	linkedList.AddAtPos(node6, pos)
	linkedList.AllPrint()
}
