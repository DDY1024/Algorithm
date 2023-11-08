## 定义
- 字符串 S 的循环同构即 S[i...n] + S[1...i-1]
- 字符串 S 的最小表示为与 S 循环同构的所有字符串中字典序最小的字符串

## 暴力解法
```go
	bs := []byte(s)
	i, j, k, n := 0, 1, 0, len(bs)
    // 如果 k >= n; 可以推导出整个字符串只由一类字符组成
	for i < n && j < n && k < n {
		if bs[(i+k)%n] == bs[(j+k)%n] {
			k++  
		} else {
			if bs[(i+k)%n] > bs[(j+k)%n] {
				i++
			} else {
				j++
			}
			k = 0
			if i == j {  // 存在 i 追上 j 的情况，需要 j++
				j++
			}
		}
	}
	return string(bs[i:]) + string(bs[:i])
```

## 优化解法
- 利用之前的匹配结果，避免一些无效判断
- 假设当前匹配到 S[i,...,i+k-1] = S[j,...,j+k-1]
	- 当 S[i+k] > S[j+k] 时，[i+1, i+k] 区间不需要进行判断，因为后续肯定存在一个 j 其结果更优，因此下一步 i = i + k + 1
	- 当 S[i+k] < S[j+k] 时，同理下一步 j = j + k + 1
```go
	bs := []byte(s)
	i, j, k, n := 0, 1, 0, len(bs)
	for i < n && j < n && k < n {
		if bs[(i+k)%n] == bs[(j+k)%n] {
			k++
		} else {
			// 跳跃式增进，避免无效判断
			if bs[(i+k)%n] > bs[(j+k)%n] {
				i = i + k + 1
			} else {
				j = j + k + 1
			}
			k = 0
			if i == j {
				j++
			}
		}
	}

	// 最终结果存在三种情况：
	// 1. i >= n，则 j 为最小字符串起始下标
	// 2. j >= n, 则 i 为最小字符串起始下标
	// 3. i < n && j < n，则 i 或 j 任意即可，通常选择索引下标更小的
	idx := minInt(i, j)
	return string(bs[idx:]) + string(bs[:idx])
```

## leetcode
- https://leetcode.cn/problems/orderly-queue/
	- 对于 k > 1 时，我们在任意次操作后，总会把字符串转化成原字符串字符重新排列后的最小字符串
	- 对于 k = 1，问题转化为字符串最小表示
```go
func orderlyQueue(s string, k int) string {
	bs := []byte(s)

	// k > 1 时，经过任意步操作后，必然构造最小的
	if k > 1 {
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	}

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// k = 1，每次只能移动一个，问题转化为最小字符串表示
	i, j, kk, n := 0, 1, 0, len(s)
	for i < n && j < n && kk < n {
		if s[(i+kk)%n] == s[(j+kk)%n] {
			kk++
		} else {
			if s[(i+kk)%n] > s[(j+kk)%n] {
				i = i + kk + 1
			} else {
				j = j + kk + 1
			}
			kk = 0
			if i == j {
				i++
			}
		}
	}
	idx := minInt(i, j)
	return s[idx:] + s[:idx]
}
```