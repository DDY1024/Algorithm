#### 一、题目链接
[最长重复子数组](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)

#### 二、题目大意
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
- 1 <= len(A), len(B) <= 1000
- 0 <= A[i], B[i] < 100

#### 三、解题思路
参考官方题解关于 `二分+Rabin-Karp` 解题方法：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zui-chang-zhong-fu-zi-shu-zu-by-leetcode-solution/

#### 四、复杂度分析

#### 五、代码
```go
func powMod(a, b, c int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = ret * a % c
		}
		a = a * a % c
		b >>= 1
	}
	return ret
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 利用 rabin-karp 滚动哈希的思想进行求解
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)

	// A[i]，B[i] <= 100
	mod := int(1e9 + 7) 
    // mod 要大一些，这样针对本题不需要判定 hash 冲突，实际情况下我们为了保证准确性仍然要检测 hash 冲突问题
	base := 113

	var check = func(l int) bool {
		mark := make(map[int]bool)
		hash := 0
		tmp := powMod(base, l-1, mod)
		for i := 0; i < l; i++ {
			hash = (hash*base + nums1[i]) % mod
		}
		mark[hash] = true
		for i := l; i < n; i++ {
			hash = (((hash-nums1[i-l]*tmp%mod)+mod)%mod*base + nums1[i]) % mod
			mark[hash] = true
		}

		hash = 0
		for i := 0; i < l; i++ {
			hash = (hash*base + nums2[i]) % mod
		}
		if mark[hash] {
			return true
		}

		for i := l; i < m; i++ {
			hash = (((hash-nums2[i-l]*tmp%mod)+mod)%mod*base + nums2[i]) % mod
			if mark[hash] {
				return true
			}
		}
		return false
	}

	l, r, ans := 1, minInt(n, m), 0
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}

```