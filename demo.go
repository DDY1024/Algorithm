package main

import (
	"fmt"
	"sort"
)

var ls = func(n int) int { return 1 << uint(n) }

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	var maxDP [30][30][30]int
	var minDP [30][30][30]int
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			for k := 0; k < 30; k++ {
				maxDP[i][j][k] = -1
				minDP[i][j][k] = -1
			}
		}
	}

	var dp func(n, fp, sp int) (int, int)
	dp = func(n, fp, sp int) (int, int) {
		if maxDP[n][fp][sp] != -1 {
			return minDP[n][fp][sp], maxDP[n][fp][sp]
		}

		if fp+sp == n-1 { // fp 和 sp 相撞了，状态结束停止计算
			minDP[n][fp][sp] = 1
			maxDP[n][fp][sp] = 1
			return minDP[n][fp][sp], maxDP[n][fp][sp]
		}

		arr, maxV, minV := make([]int, (n+1)/2), 0, 0xffff
		// 直接枚举所有可能的比赛结果
		for s := 0; s < ls(n/2); s++ {
			idx := 0
			for i := 0; i < n/2; i++ {
				if i == fp {
					arr[idx] = i
				} else if n-i-1 == fp {
					arr[idx] = n - i - 1
				} else if i == sp {
					arr[idx] = i
				} else if n-i-1 == sp {
					arr[idx] = n - i - 1
				} else if s&ls(i) == 0 {
					arr[idx] = i
				} else {
					arr[idx] = n - i - 1
				}
				idx++
			}
			if n&1 > 0 {
				arr[idx] = n / 2
				idx++
			}
			sort.Ints(arr)
			nf, ns := -1, -1
			for i := 0; i < (n+1)/2; i++ {
				if arr[i] == fp {
					nf = i
				}
				if arr[i] == sp {
					ns = i
				}
			}
			v1, v2 := dp((n+1)/2, nf, ns)
			maxV = maxInt(maxV, v2+1)
			minV = minInt(minV, v1+1)
		}
		maxDP[n][fp][sp] = maxV
		minDP[n][fp][sp] = minV
		return minDP[n][fp][sp], maxDP[n][fp][sp]
	}

	ans := make([]int, 2)
	ans[0], ans[1] = dp(n, firstPlayer-1, secondPlayer-1)
	return ans
}

func main() {
	fmt.Println(earliestAndLatest(11, 2, 4))
	fmt.Println(earliestAndLatest(5, 1, 5))
	fmt.Println(earliestAndLatest(4, 1, 2))
}

/*
const N = 30

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	var f [N][N][N]bool
	//f[k][i][j] 表示还剩k个人时firstP前有i个人，secondP后有j个人的情况是否可能出现
	f[n][firstPlayer-1][n-secondPlayer] = true

	for k := n; k > 1; k = (k + 1) / 2 {
		//枚举两端人数x，y
		for x := 0; x <= n; x++ {
			for y := 0; y <= n; y++ {
				if f[k][x][y] {
					mid2 := (k + 1) / 2
					mid := k - mid2
					if y >= mid {
						z := k - x - y - 2
						for i := 0; i <= x; i++ {
							for j := 0; j <= z; j++ {
								f[mid2][i][x-i+j+y-mid] = f[k][x][y]
							}
						}
					} else if x >= mid {
						z := k - x - y - 2
						for i := 0; i <= y; i++ {
							for j := 0; j <= z; j++ {
								f[mid2][y-i+j+x-mid][i] = f[k][x][y]
							}
						}
					} else if y > x {
						z := y - x - 1
						for i := 0; i <= x; i++ {
							for j := 0; j <= z; j++ {
								f[mid2][i][j+x-i] = f[k][x][y]
							}
						}
					} else if x > y {
						z := x - y - 1
						for i := 0; i <= y; i++ {
							for j := 0; j <= z; j++ {
								f[mid2][y-i+j][i] = f[k][x][y]
							}
						}
					}
				}
			}

		}
	}
	r1, r2 := 30, -1
	for k, t := n, 1; k > 1; k, t = (k+1)/2, t+1 {
		for i := 0; i < n; i++ {
			if f[k][i][i] {
				r1 = min(r1, t)
				r2 = max(r2, t)
			}
		}
	}
	return []int{r1, r2}
}

func max(r2 int, t int) int {
	if r2 > t {
		return r2
	}
	return t
}

func min(r1 int, t int) int {
	if r1 > t {
		return t
	}
	return r1
}

// 记忆化搜索 --> 反向递推求解
*/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			dp[i][j] = maxInt(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	if dp[0][n-1] >= 0 {
		return true
	}
	return false
}
