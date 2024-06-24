package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}
type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	newNode := &Node{value: value}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	parent := t.root
	for {
		parent = current
		if value < current.value {
			current = current.left
			if current == nil {
				parent.left = newNode
				return
			}
		} else {
			current = current.right
			if current == nil {
				parent.right = newNode
				return
			}
		}

	}
}

func (t *Tree) Search(value int) *Node {
	if t.root == nil {
		return nil

	}
	current := t.root
	for {
		if value == current.value {
			return current
		} else if value < current.value {
			current = current.left

		} else {
			current = current.right
		}
		if current == nil {
			return nil
		}

	}

}

func inOrderTraversal(node *Node, result *[]int) {
	if node != nil {
		inOrderTraversal(node.left, result)
		*result = append(*result, node.value)
		inOrderTraversal(node.right, result)
	}
}

func (t *Tree) ToSortedArray() []int {
	var result []int
	inOrderTraversal(t.root, &result)
	return result
}

func main() {

	t := new(Tree)
	t.Insert(12)
	t.Insert(11)
	t.Insert(1)
	t.Insert(4)
	t.Insert(10)
	t.Insert(3)
	fmt.Println(t.Search(3))
	fmt.Println(t.ToSortedArray())

}
