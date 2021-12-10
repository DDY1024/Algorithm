package main

// https://leetcode.com/contest/weekly-contest-221/problems/where-will-the-ball-fall/
func findBall(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	vis := make([][][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			vis[i][j] = make([]bool, 4)
		}
	}

	var dfs func(x, y, z int) int
	dfs = func(x, y, z int) int {
		vis[x][y][z] = true
		if x == n-1 && (z == 2 || z == 3) {
			return y
		}
		switch z {
		case 0:
			if y+1 < m {
				if grid[x][y+1] == 1 && !vis[x][y+1][2] {
					return dfs(x, y+1, 2)
				}
				if grid[x][y+1] == -1 && !vis[x][y+1][1] {
					return dfs(x, y+1, 1)
				}
			}
		case 1:
			if y-1 >= 0 {
				if grid[x][y-1] == 1 && !vis[x][y-1][0] {
					return dfs(x, y-1, 0)
				}
				if grid[x][y-1] == -1 && !vis[x][y-1][3] {
					return dfs(x, y-1, 3)
				}
			}
		case 2:
			if x+1 < n {
				if grid[x+1][y] == 1 && !vis[x+1][y][0] {
					return dfs(x+1, y, 0)
				}
				if grid[x+1][y] == -1 && !vis[x+1][y][1] {
					return dfs(x+1, y, 1)
				}
			}
		case 3:
			if x+1 < n {
				if grid[x+1][y] == 1 && !vis[x+1][y][0] {
					return dfs(x+1, y, 0)
				}
				if grid[x+1][y] == -1 && !vis[x+1][y][1] {
					return dfs(x+1, y, 1)
				}
			}
		}
		return -1
	}

	// 每个节点划分成四个状态并判断与相邻节点的连通性问题
	// 仔细观察可以发现连通通道的入口和出口是唯一的，理论上不需要严格的判重

	ret := make([]int, m)
	for i := 0; i < m; i++ {
		if grid[0][i] == 1 {
			ret[i] = dfs(0, i, 0)
		} else {
			ret[i] = dfs(0, i, 1)
		}
	}
	return ret
}

// 由题目中的规定可知，小球的移动轨迹其实是固定且唯一的。因为挡板的性质已经决定了。
// 因此对于每个下落入口处的小球，我们直接模拟其对应的下落轨迹即可。
/*
class Solution {
public:
    vector<int> findBall(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        vector<int> ans(n, -1);
        for(int i = 0; i < n; ++i) {//第i列球
            int now = i, flag = 1;//当前在第now列
            for(int j = 0; j < m; ++j) {//当前在第j行
                if(grid[j][now] == 1) {//向右滑
                    if(now + 1 == n || grid[j][now + 1] == -1) flag = 0;
                    ++ now;
                }else {
                    if(now == 0 || grid[j][now - 1] == 1) flag = 0;
                    -- now;
                }
                if(flag == 0) break;
            }
            if(flag) ans[i] = now;
        }  // 直接模拟掉落情况即可
        return ans;
    }
};
*/
