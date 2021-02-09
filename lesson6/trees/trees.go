package trees

import "fmt"

// TreeNode description
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// BinarySearchTree description
type BinarySearchTree struct {
	Root *TreeNode
}

// Insert description
func (bst *BinarySearchTree) Insert(item int) {
	newNode := &TreeNode{item, nil, nil}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		if err:= bst.Root.insert(newNode); err != nil {
			fmt.Println("Smth went wrong!")
		}
	}
}

func (n *TreeNode) insert(newNode *TreeNode) error {
	if n.Val > newNode.Val {
		if n.Left == nil {
			n.Left = newNode
			return nil
		}
		n.Left.insert(newNode)
	} else {
		if n.Right == nil {
			n.Right = newNode
			return nil
		}
		n.Right.insert(newNode)
	}
	return nil
}

// InOrderPrint description
func (n *TreeNode) InOrderPrint() {
	if n == nil {
		return
	}
	n.Left.InOrderPrint()
	fmt.Print(n.Val);fmt.Print(" ")
	n.Right.InOrderPrint()
}

// PreOrderPrint description
func (n *TreeNode) PreOrderPrint() {
	if n == nil {
		return
	}
	fmt.Print(n.Val); fmt.Print(" ")
	n.Left.PreOrderPrint()
	n.Right.PreOrderPrint()
}

// PostOrderPrint description
func (n *TreeNode) PostOrderPrint() {
	if n == nil {
		return
	}
	n.Left.PostOrderPrint()
	n.Right.PostOrderPrint()
	fmt.Print(n.Val); fmt.Print(" ")
}

// String description
func (bst *BinarySearchTree) String() {
	printer(bst.Root, 0)
}

func printer(n *TreeNode, lvl int) {
	if n != nil {
		str := ""
		for i:=0;i < lvl; i++ {
			str += "............."
		}
		str += "---[ "
		lvl++
		printer(n.Left, lvl)
		fmt.Printf(str + "%d\n",n.Val)
		printer(n.Right, lvl)
	}
}


// Search description
func (bst *BinarySearchTree) Search(item int) bool {
	return search(bst.Root, item)
}

func search(n *TreeNode, item int) bool {
	if n == nil {
		return false
	}
	if n.Val < item {
		return search(n.Left, item)
	}
	if n.Val > item {
		return search(n.Right, item)
	}
	return true
}
