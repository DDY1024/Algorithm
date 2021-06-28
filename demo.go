package main

import "fmt"

func numBusesToDestination(routes [][]int, source int, target int) int {

	// 这种 case 就没啥意思了
	if source == target {
		return 0
	}

	n := len(routes)
	station := make(map[int][]int)
	stationMark := make(map[int]map[int]bool)
	vis := make(map[int]bool)
	dis := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < len(routes[i]); j++ {
			station[routes[i][j]] = append(station[routes[i][j]], i)
			if _, ok := stationMark[i]; !ok {
				stationMark[i] = make(map[int]bool)
			}
			stationMark[i][routes[i][j]] = true
		}
	}

	que := make([]int, 0, 100000)
	for _, u := range station[source] {
		dis[u] = 1
		vis[u] = true
		que = append(que, u)
	}

	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		if stationMark[u][target] {
			return dis[u]
		}

		for _, s := range routes[u] {
			for _, v := range station[s] {
				if !vis[v] {
					vis[v] = true
					dis[v] = dis[u] + 1
					que = append(que, v)
				}
			}
		}
	}
	return -1
}

// [[1,2,7],[3,6,7]]  1 --> 6
// [1,2,7]
// [3,6,7]
// [3,6,7]
// priority_queue + bfs --> 权值不相同时 --> 求解最短路径的方式

func main() {
	// fmt.Println("hello, world!")
	fmt.Println(numBusesToDestination([][]int{
		[]int{1, 2, 7},
		[]int{3, 6, 7},
	}, 1, 6))
}
