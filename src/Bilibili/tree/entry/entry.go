package main

import (
	"Bilibili/tree"

	"fmt"

	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSpares() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(100000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateNode(2)
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)

	myNode := myTreeNode{&root}
	myNode.postOrder()

	testSpares()

}
