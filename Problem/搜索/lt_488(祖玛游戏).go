package main

import (
	"sort"
)

// 经典搜索剪枝题目
// 题目链接: https://leetcode-cn.com/problems/zuma-game/
// 解题思路: https://leetcode-cn.com/problems/zuma-game/solution/zu-ma-you-xi-by-leetcode-solution-lrp4/

// func main() {
// 	fmt.Println(findMinStep("BRWGWYY", "YGBWY"))
// 	fmt.Println(findMinStep("RRWWRRBBRR", "WB"))
// }

// Tips: 需要很强的剪枝技巧以及处理字符串操作的优化技巧，一道比较赞的题目
// 特殊 case: "RRWWRRBBRR", "WB"
type Node struct {
	board string
	hand  string
	step  int
}

func findMinStep(board string, hand string) int {
	// 排序，方便后续处理
	handBS := []byte(hand)
	sort.Slice(handBS, func(i, j int) bool {
		return handBS[i] < handBS[j]
	})
	hand = string(handBS)

	// 利用 stack 思想来高效处理字符串的收缩问题
	// 此处很耗时整体搜索便会很耗时
	// 避免 go 中字符串的拷贝操作
	var process = func(s string) string {
		stk := make([]byte, 0, len(s))
		idx, isOk := 0, false
		for {
			if idx < len(s) {
				if len(stk) == 0 || stk[len(stk)-1] == s[idx] || isOk {
					stk = append(stk, s[idx])
					idx++
					isOk = false
					continue
				}
			}

			// 栈顶最长瘦身路径一定要及时处理
			slen := len(stk)
			if slen > 2 && stk[slen-1] == stk[slen-2] && stk[slen-2] == stk[slen-3] {
				pos := slen - 4
				for pos >= 0 && stk[pos] == stk[slen-1] {
					pos--
				}
				stk = stk[:pos+1]
			}
			// 表示进行过一次瘦身操作
			isOk = true

			// 无字符可追加且不再能进行瘦身，退出
			if idx >= len(s) && slen == len(stk) {
				break
			}
		}
		return string(stk)
	}

	var hash = func(board, hand string) string {
		return board + ":" + hand // 简单字符串拼接
	}

	que := make([]Node, 100000)
	vis := make(map[string]bool, 100000)
	front, rear := 0, 1
	que[front] = Node{board, hand, 0}
	for front < rear {
		cur := que[front]
		front++

		for j := 0; j < len(cur.hand); j++ {
			// 第 1 个剪枝条件: 当前选择的球的颜色和前一个球的颜色相同
			if j > 0 && cur.hand[j] == cur.hand[j-1] {
				continue
			}

			for i := 0; i <= len(cur.board); i++ {
				// 第 2 个剪枝条件: 只在连续相同颜色的球的开头位置插入新球
				if i > 0 && cur.board[i-1] == cur.hand[j] { // 在 i-1 位置的左边还是右边插入产生的效果都是相同的
					continue
				}

				// 第 3 个剪枝条件: 只在以下两种情况放置新球
				choose := false
				// 第 1 种情况 : 当前球颜色与后面的球的颜色相同（想一想不同，没有任何意义）
				if i < len(cur.board) && cur.board[i] == cur.hand[j] {
					choose = true
				}

				// 第 2 种情况 : 当前后颜色相同且与当前颜色不同时候放置球
				// 处理特殊 case: "RRWWRRBBRR", "WB"
				if i > 0 && i < len(cur.board) && cur.board[i-1] == cur.board[i] && cur.board[i] != cur.hand[j] {
					choose = true
				}

				if choose {
					// 插入第 i 个位置前面
					newBoard := process(cur.board[:i] + cur.hand[j:j+1] + cur.board[i:])
					newHand := cur.hand[:j] + cur.hand[j+1:]
					if len(newBoard) == 0 {
						return cur.step + 1
					}

					if !vis[hash(newBoard, newHand)] {
						que[rear] = Node{newBoard, newHand, cur.step + 1}
						rear++
						vis[hash(newBoard, newHand)] = true
					}
				}
			}
		}
	}
	return -1
}
