package main

import "Bilibili/tree"

type myTreeNode struct {
	node *tree.Node
}


func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil{
		return
	}
	
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateNode(2)
	root.Right = &tree.Node{Value:5, Left:nil, Right:nil}
	root.Right.Left = new(tree.Node)


	myNode := myTreeNode{&root}
	myNode.postOrder()
	
}
