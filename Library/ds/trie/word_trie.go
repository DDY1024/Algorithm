package main

import "fmt"

type Node struct {
	childrens [26]*Node
	value     string
	count     int
}

// A ~ Z 或 a ~ z
func child(b byte) int {
	if b >= 'A' && b <= 'Z' {
		return int(b - 'A')
	}
	return int(b - 'a')
}

func insertNode(root *Node, s string) {
	for i := 0; i < len(s); i++ {
		idx := child(s[i])
		if root.childrens[idx] == nil {
			root.childrens[idx] = &Node{}
		}
		root = root.childrens[idx]
	}
	root.value = s
	root.count++
}

func findNode(root *Node, s string) bool {
	for i := 0; i < len(s); i++ {
		idx := child(s[i])
		if root.childrens[idx] == nil {
			return false
		}
		root = root.childrens[idx]
	}
	return root.count > 0 // 必须检查 count > 0，有可能是前缀字符串
}

func main() {
	root := &Node{}
	insertNode(root, "hello")
	fmt.Println(findNode(root, "hell"))
	fmt.Println(findNode(root, "hello"))
	fmt.Println(findNode(root, "h"))
	insertNode(root, "hxy")
	fmt.Println(findNode(root, "hc"))
	fmt.Println(findNode(root, "hx"))
	fmt.Println(findNode(root, "hxy"))
}
