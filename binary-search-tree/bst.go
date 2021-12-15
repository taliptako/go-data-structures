package main

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	if k <= n.key {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		return n.right.Search(k)
	} else if n.key > k {
		return n.left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{key: 10}

	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(20)
	tree.Insert(22)
	tree.Insert(10)
	tree.Insert(4)

	tree.Search(15)
}
