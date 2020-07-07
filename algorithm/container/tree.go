package container

type BinaryTree struct {
	root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}
