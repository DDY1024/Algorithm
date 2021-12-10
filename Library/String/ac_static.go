package main

import "fmt"

const (
	CharNum = 26
	Maxn    = 500010
)

var (
	trie  [][]int // 从状态 u 通过字符 c 到达的状态
	fail  []int   // 失效指针
	count []int   // 字符串出现次数统计
	snum  int     // 状态数
)

func init() {
	snum = 0
	trie = make([][]int, Maxn)
	for i := 0; i < Maxn; i++ {
		trie[i] = make([]int, CharNum) // 默认索引 0 代表根节点的位置编号
	}
	fail = make([]int, Maxn)
	count = make([]int, Maxn)
}

func index(ch byte) int {
	return int(ch - 'a')
}

// insert trie, root index = 0
func insert(s string) {
	u, sl := 0, len(s)
	for i := 0; i < sl; i++ {
		v := index(s[i])
		if trie[u][v] == 0 {
			snum++
			trie[u][v] = snum
		}
		u = trie[u][v]
	}
	count[u]++ // 字符串个数
}

func calcFail() {
	que := make([]int, Maxn)
	front, rear := 0, 0
	for i := 0; i < CharNum; i++ {
		if trie[0][i] > 0 {
			que[rear] = trie[0][i]
			rear++
		}
	}

	for front < rear {
		u := que[front]
		front++
		for i := 0; i < CharNum; i++ {
			if trie[u][i] > 0 {
				fail[trie[u][i]] = trie[fail[u]][i]
				que[rear] = trie[u][i]
				rear++
			} else {
				trie[u][i] = trie[fail[u]][i]
			}
		}
	}
}

func matchAC(s string) int {
	u, rcnt := 0, 0
	for i := 0; i < len(s); i++ {
		u = trie[u][index(s[i])]
		for j := u; j > 0 && count[j] != -1; j = fail[j] {
			rcnt += count[j]
			count[j] = -1
		}
	}
	return rcnt
}

func main() {
	insert("a")
	insert("aa")
	insert("aa")
	insert("aaa")
	calcFail()
	fmt.Println(matchAC("aaa"))
}
