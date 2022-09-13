package main

// 参考文章：https://zhuanlan.zhihu.com/p/58727559

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	dis := make([]int, n*m)
	for i := 1; i < n*m; i++ {
		dis[i] = 0x3f3f3f3f
	}

	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	que := make([]int, 0, 1000000)
	inq := make([]bool, n*m)
	que = append(que, 0)
	inq[0] = true
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		inq[u] = false
		x, y := u/m, u%m
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m {
				v := xx*m + yy
				w := maxInt(dis[u], absInt(heights[x][y]-heights[xx][yy]))
				if dis[v] > w {
					dis[v] = w
					if !inq[v] {
						inq[v] = true
						que = append(que, v)
					}
				}
			}
		}
	}
	return dis[n*m-1]
}

// SPFA也可以判负权环，我们可以用一个数组记录每个顶点进队的次数，当一个顶点进队超过n次时，就说明存在负权环
/*
void SPFA(int s)
{
    queue<int> Q;
    Q.push(s);
    while (!Q.empty())
    {
        int p = Q.front();
        Q.pop();
        inqueue[p] = 0;
        for (int e = head[p]; e != 0; e = edges[e].next)
        {
            int to = edges[e].to;
            if (dist[to] > dist[p] + edges[e].w)
            {
                dist[to] = dist[p] + edges[e].w;
                if (!inqueue[to])
                {
                    inqueue[to] = 1;
					// if ++cnt[to] > n {
					// 	// 负权环
					// }
                    Q.push(to);
                }
            }
        }
    }
}
*/
