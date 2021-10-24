package main

import "fmt"

// AC 自动机 golang 版本实现
// 参考实现: https://blog.csdn.net/li1615882553/article/details/80102530

// 1. 树结构实现 ac 自动机
const (
	CharNum = 26
)

type Node struct {
	isEnd bool    // 单次结尾
	cnt   int     // 出现次数
	word  string  // 单词
	child []*Node // 子节点列表
	fail  *Node   // fail 指针
}

func NewNode() *Node {
	return &Node{
		isEnd: false,
		word:  "",
		child: make([]*Node, CharNum),
		fail:  nil,
		cnt:   0, // 单词计数
	}
}

// trie 插入节点
func Insert(root *Node, s string) {
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if root.child[idx] == nil {
			root.child[idx] = NewNode()
		}
		root = root.child[idx]
	}
	root.isEnd = true
	root.word = s
	root.cnt++
}

// AC 自动机构建
// 按层次进行构建
// root.fail = nil
func BuildAC(root *Node) {
	que := make([]*Node, 10000)
	que[0] = root
	front, rear := 0, 1
	for front < rear {
		nd := que[front]
		front++
		for i := 0; i < CharNum; i++ {
			if nd.child[i] != nil {
				// 第一层节点的 fail 指针均为 root
				if nd == root {
					nd.child[i].fail = root
				} else {
					p := nd.fail
					for p != nil {
						if p.child[i] != nil {
							nd.child[i].fail = p.child[i]
							break
						}
						p = p.fail
					}
					if p == nil {
						nd.child[i].fail = root
					}
				}
				que[rear] = nd.child[i]
				rear++
			}
		}
	}
}

// 执行匹配过程
func MatchAC(s string, root *Node) {
	p, ans := root, 0
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		for p.child[idx] == nil && p != root {
			p = p.fail
		}
		p = p.child[idx]
		if p == nil {
			p = root
		}

		// 沿着 fail 指针寻找路径上可能存在所有 word
		// 如果该 fail 路径已经被计算过，则不需要再次进行重复计算
		tmp := p
		for tmp != root {
			if tmp.cnt >= 0 { // 判重操作，防止重复计算
				ans += tmp.cnt
				tmp.cnt = -1
			} else { // 已经被访问过了不再计数
				break
			}
			tmp = tmp.fail
		}
		// for tmp != root {
		// 	// 如果是计数的话，注意去重，每个模式串只会被计入一次
		// 	if tmp.isEnd {
		// 		fmt.Println("Match Pos:", i-len(tmp.word)+1)
		// 	}
		// 	tmp = tmp.fail
		// }
	}

	// 输出总共匹配的单词数
	fmt.Println("occur count:", ans)
}

// func Del(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	for i := 0; i < 26; i++ {
// 		Del(root.child[i])
// 	}
// 	root = nil
// }

// 2. 数组实现 ac 自动机

func main() {
	root := NewNode()
	Insert(root, "a")
	Insert(root, "aa")
	Insert(root, "aa")
	Insert(root, "aaa")
	BuildAC(root)
	MatchAC("aaa", root)
}
