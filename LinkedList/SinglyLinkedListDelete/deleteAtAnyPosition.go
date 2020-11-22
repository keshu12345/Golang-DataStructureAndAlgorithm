package main

import (
	"fmt"
	"sync"
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

func (l *LinkedList) createAll(newNode *Node, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()

	if l.head == nil {
		l.head = newNode
		l.temp = newNode
		l.length++
	} else {
		l.temp.next = newNode
		l.temp = newNode
		l.length++
	}
	mutex.Unlock()
	wg.Done()
}

func (l *LinkedList) DeleteFirstElement(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	l.temp = l.head
	l.head = l.head.next
	l.temp.next = nil

	l.length--
	mutex.Unlock()
	wg.Done()
}

func (l *LinkedList) PrintLinkedList() {

	if l.length == 0 {
		fmt.Println("No Node available")

	}

	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}

}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	linkedList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	node5 := &Node{data: 50}
	wg.Add(6)
	go linkedList.createAll(node1, &wg, &mutex)
	go linkedList.createAll(node2, &wg, &mutex)
	go linkedList.createAll(node3, &wg, &mutex)
	go linkedList.createAll(node4, &wg, &mutex)
	go linkedList.createAll(node5, &wg, &mutex)
	//linkedList.PrintLinkedList()
	go linkedList.DeleteFirstElement(&wg, &mutex)
	linkedList.PrintLinkedList()
	wg.Wait()
	fmt.Println("End")

}
