package main

func wonderfulSubstrings(word string) int64 {
	n := len(word)
	stats, ans, mask := make([]int, 1024), 0, 0
	stats[0] = 1
	for i := 0; i < n; i++ {
		mask ^= 1 << uint(word[i]-'a')
		// fmt.Println(mask)
		for j := 0; j < 10; j++ {
			if mask&(1<<uint(j)) == 0 {
				ans += stats[mask|(1<<uint(j))]
			} else {
				ans += stats[mask&^(1<<uint(j))]
			}
		}
		stats[mask]++
	}
	return int64(ans)
}

func waysToBuildRooms(prevRoom []int) int {
	return 0
}

func main() {
	// wonderfulSubstrings("aba")
	// fmt.Println("hello, world!")
}
