package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) Find(k int) *Node {
	if n != nil {
		if k < n.data {
			return n.left.Find(k)
		} else if k > n.data {
			return n.right.Find(k)
		} else {
			return n
		}
	} else {
		return nil
	}
}

func (n *Node) Insert(k int) *Node {
	if k < n.data {
		if n.left != nil {
			return n.left.Insert(k)
		} else {
			n.left = new(Node)
			n.left.data = k
		}
	} else if k > n.data {
		if n.right != nil {
			return n.right.Insert(k)
		} else {
			n.right = new(Node)
			n.right.data = k
		}
	} else {
		return n
	}
	return nil
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.data)
	n.left.PreOrder()
	n.right.PreOrder()
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.left.InOrder()
	fmt.Printf("%d ", n.data)
	n.right.InOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.left.PostOrder()
	n.right.PostOrder()
	fmt.Printf("%d ", n.data)
}

func main() {
	Root := new(Node)
	Root.data = 100
	Root.Insert(100)
	Root.Insert(50)
	Root.Insert(150)
	Root.Insert(30)
	Root.Insert(70)
	Root.Insert(110)
	Root.Insert(170)

	//fmt.Println(Root)

	//fmt.Println(Root.Find(150))
	fmt.Println("前序：")
	Root.PreOrder()
	fmt.Println()
	fmt.Println("中序：")
	Root.InOrder()
	fmt.Println()
	fmt.Println("后序：")
	Root.PostOrder()
}
