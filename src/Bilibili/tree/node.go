package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) SetValue(Value int) {
	node.Value = Value
}

func (node *Node) Traverse() {
	if node == nil { return  }
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

