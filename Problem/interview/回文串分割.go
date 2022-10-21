package main

func partition(s string) [][]string {
	n := len(s)
	isP := make([][]bool, n)
	for i := 0; i < n; i++ {
		isP[i] = make([]bool, n)
		isP[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if i+1 >= j-1 {
					isP[i][j] = true
				} else {
					isP[i][j] = isP[i+1][j-1]
				}
			}
		}
	}

	ret := make([][]string, 0)
	tmp := make([]string, n)

	var dfs func(pos, num int)
	dfs = func(pos, num int) {
		if pos >= n {
			ctmp := make([]string, 0, num)
			for i := 0; i < num; i++ {
				ctmp = append(ctmp, tmp[i])
			}
			ret = append(ret, ctmp)
			return
		}
		for i := pos; i < n; i++ {
			if isP[pos][i] {
				tmp[num] = s[pos : i+1]
				dfs(i+1, num+1)
			}
		}
	}
	dfs(0, 0)

	return ret
}
