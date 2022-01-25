package main

import "fmt"

func main() {
	bst := BST{10, nil, nil}
	bst.insert(5)
	bst.insert(2)
	bst.insert(1)
	bst.insert(15)
	bst.insert(13)
	bst.insert(21)

	array := []int{}
	bst.traversal(&array)

	fmt.Println(array)
}

type BST struct {
	value int
	left  *BST
	right *BST
}

func (bst *BST) insert(value int) {
	tree := bst
	for true {
		if value < tree.value {
			if tree.left == nil {
				tree.left = &BST{value, nil, nil}
				break
			} else {
				tree = tree.left
			}
		} else {
			if tree.right == nil {
				tree.right = &BST{value, nil, nil}
				break
			} else {
				tree = tree.right
			}
		}
	}
}

func (tree *BST) traversal(array *[]int) {
	if tree == nil {
		return
	}
	tree.left.traversal(array)
	*array = append(*array, tree.value)
	tree.right.traversal(array)
}
