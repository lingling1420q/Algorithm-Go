package main

import (
	"fmt"

	// "golang.org/x/tour/tree"
	tree "./tree"
)

func Reverse(root *tree.TNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left // 将node的左节点和右节点互换位置
	Reverse(root.Left)                            // 递归当前node的左子树
	Reverse(root.Right)                           // 递归当前node的右子树
}

func main() {
	root := tree.New(1)
	fmt.Println("Before reverse", root.String())
	Reverse(root)

	fmt.Println("After reverse", root.String())

}
