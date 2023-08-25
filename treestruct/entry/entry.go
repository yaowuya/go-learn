package main

import (
	"fmt"
	"go-learn/treestruct"
)

type myTreeNode struct {
	node *treestruct.Node
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

func main() {
	var root treestruct.Node
	root = treestruct.Node{Value: 3}
	root.Left = &treestruct.Node{}
	root.Right = &treestruct.Node{Value: 5}
	root.Right.Left = new(treestruct.Node)
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal: ")
	root.Traverse()
}
