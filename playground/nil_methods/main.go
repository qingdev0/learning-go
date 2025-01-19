package main

import "fmt"

// IntTree is a simple binary tree of integers
type IntTree struct {
	value       int      // value of the node
	left, right *IntTree // pointers to the left and right children
}

func (it *IntTree) Insert(v int) *IntTree {
	// Base case: if current node is nil, create and return a new node with value v
	// This happens when we've found the correct position to insert the new value
	// 1. Creates a new IntTree node
	// 2. Initializes its value field with v
	// 3. Returns pointer to the new node, which will be assigned to parent's left/right pointer
	//    This connects the new node to the tree structure
	if it == nil {
		return &IntTree{value: v}
	}
	if v < it.value {
		it.left = it.left.Insert(v)
	} else if v > it.value {
		it.right = it.right.Insert(v)
	}
	return it
}

// String implements the Stringer interface for pretty-printing the tree
func (it *IntTree) String() string {
	if it == nil {
		return "nil"
	}
	return fmt.Sprintf("(%d L:%v R:%v)", it.value, it.left, it.right)
}

func main() {
	tree := &IntTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	fmt.Println(tree.String())
}
