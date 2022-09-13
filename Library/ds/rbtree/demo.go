package main

import (
	"fmt"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

// gods: leetcode 标准库

func main() {
	tree := rbt.NewWithIntComparator() // empty (keys are of type int)
	// rbt.NewWithStringComparator()

	// (key, value) 结构
	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │           ┌── 6
	//	│       ┌── 5
	//	│   ┌── 4
	//	│   │   └── 3
	//	└── 2
	//		└── 1

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// 左边最小节点
	tree.Left() // gets the left-most (min) node

	// 右边最大节点
	tree.Right() // get the right-most (max) node

	// <= x 的最大节点
	tree.Floor(1) // get the floor node

	// >= x 的最小节点
	tree.Ceiling(1) // get the ceiling node

	// tree.Floor()   --> <= x 的最大节点
	// tree.Ceiling() --> >= x 的最小节点
}