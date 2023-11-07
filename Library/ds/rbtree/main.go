package main

import (
	"fmt"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := rbt.NewWithIntComparator()

	// tree := rbt.NewWithStringComparator()

	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")
	tree.Put(6, "f")

	fmt.Println(tree)

	_ = tree.Values()
	_ = tree.Keys()

	tree.Remove(2)
	fmt.Println(tree)

	tree.Clear()
	tree.Empty()
	tree.Size()

	// 左边最小节点
	tree.Left()

	// 右边最大节点
	tree.Right()

	// <= x 的最大节点
	tree.Floor(1)

	// >= x 的最小节点
	tree.Ceiling(1)

	// lower_bound
	// upper_bound
}
