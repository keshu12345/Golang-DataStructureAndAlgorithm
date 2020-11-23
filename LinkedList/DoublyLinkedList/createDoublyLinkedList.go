//githu.com/keshu12345/Golang-DataStructureAndAlgorithm
// AddNode at Front Time Complexity O(1)
//Add Node at the End Time Comlexity O(1)
//Add node at given position Time Complexity O(n)
//Remove Node from Front Time COmlexity O(1)
//Remove Node at End Time COmlexity O(1)
//All cases Space Comlexity O(1)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	temp   *Node
	length int
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

func (l *LinkedList) addAtFront(newNode *Node) {
	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
	l.length++

}

func (l *LinkedList) addAttheEnd(newNode *Node) {
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
	l.length++
}

func (l *LinkedList) addAnyPosition(newNode *Node, pos int) {
	var i int = 1
	l.temp = l.head
	for i < pos {
		i++
		l.temp = l.temp.next

	}
	newNode.next = l.temp.next
	l.temp.next.prev = newNode
	newNode.prev = l.temp
	l.temp.next = newNode
	l.length++

}
func (l *LinkedList) removeFromFront() {
	l.head = l.head.next
	l.head.prev.next = nil
	l.head.prev = nil
}

func (l *LinkedList) removeAtEnd() {
	l.tail = l.tail.prev
	l.tail.next.prev = nil
	l.tail.next = nil
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
func (l *LinkedList) PrintReverseDoublyLinkedList() {
	l.temp = l.tail
	fmt.Println()
	for l.temp != nil {
		fmt.Printf("%v ->", l.temp.data)
		l.temp = l.temp.prev
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pos, _ := strconv.Atoi(scanner.Text())
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
	linkedList.addAtFront(node5)
	linkedList.addAttheEnd(node6)
	linkedList.addAnyPosition(node7, pos)
	linkedList.removeFromFront()
	linkedList.removeAtEnd()
	linkedList.PrintDublyLinkedList()
	linkedList.PrintReverseDoublyLinkedList()
}
