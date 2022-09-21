package main

// 按位异或

const (
	MaxBits = 31
)

type TrieNode struct {
	child [2]*TrieNode
	pos   int
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func getBit(val, pos int) int {
	return (val >> uint(pos)) & 1
}

func Insert(root *TrieNode, val, pos int) {
	cur := root
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit] == nil {
			cur.child[bit] = NewTrieNode()
		}
		cur = cur.child[bit]
	}
	cur.pos = pos
}

func MaxXor(root *TrieNode, val int) int {
	cur := root
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit^1] != nil {
			cur = cur.child[bit^1]
			continue
		}
		if cur.child[bit] != nil {
			cur = cur.child[bit]
		}
	}
	return cur.pos
}

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	root := NewTrieNode()
	ans := make([]int, n)
	Insert(root, 0, n)
	xor := 0
	for i := n - 1; i >= 0; i-- {
		xor ^= nums[i]
		ans[i] = MaxXor(root, xor) - i
		Insert(root, xor, i)
	}
	return ans
}

// 按位 or
// or 操作只会发生 0->1、1->1 两种变化
// 因此，我们只需要维护某个 bit 变为 1 最近的坐标 i 即可
// 这样针对每个 i 开始的子数组，查看使得每一个 bit 位变为 1 的最近下标，然后在这些下标中寻找到最大值即可
func smallestSubarraysTwo(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	bitPos := make([]int, 31)
	for i := 0; i < 31; i++ {
		bitPos[i] = -1
	}
	for i := n - 1; i >= 0; i-- {
		pos := i
		for j := 0; j < 31; j++ {
			if nums[i]&(1<<uint(j)) > 0 {
				bitPos[j] = i // 更新最小下标
			}
			if bitPos[j] != -1 {
				pos = maxInt(pos, bitPos[j]) // 求解最大下标
			}
		}
		ans[i] = pos - i + 1
	}
	return ans
}
