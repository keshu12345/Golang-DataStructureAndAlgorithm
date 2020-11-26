package main

import (
	"fmt"
	"math"
)

// Node Structure

type Node struct {
	data  int
	left  *Node
	right *Node
}

//Construct Binary Tree
func (n *Node) Insert(key int) {

	if n.data < key {
		// move right
		if n.right == nil {
			n.right = &Node{data: key}
		} else {
			n.right.Insert(key)
		}

	} else if n.data > key {
		// move left

		if n.left == nil {
			n.left = &Node{data: key}
		} else {
			n.left.Insert(key)
		}
	}
}
func (n *Node) Search(key int) bool {

	if n == nil {
		return false
	}
	if n.data < key {

		// Move right
		return n.right.Search(key)
	} else if n.data > key {
		// move left
		return n.left.Search(key)
	}
	return true
}

func GetSize(root *Node) int {
	if root == nil {
		return 0
	} else {
		return 1 + GetSize(root.left) + GetSize(root.right)
	}
}

func GetHeight(root *Node) float64 {
	if root == nil {
		return 0
	} else {
		return math.Max(GetHeight(root.left), GetHeight(root.right)) + 1
	}
}

func main() {
	tree := &Node{data: 150}
	fmt.Println(tree)
	tree.Insert(100)
	tree.Insert(300)
	tree.Insert(500)
	tree.Insert(600)
	fmt.Println(tree.left)
	fmt.Println(tree.right)
	ok := tree.Search(300)
	fmt.Println(ok)
	ok = tree.Search(700)
	fmt.Println(ok)

	fmt.Println(GetSize(tree))
	fmt.Println(GetHeight(tree))

}
