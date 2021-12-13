package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "wxy"
	for _, ch := range str {
		fmt.Println(ch)
		fmt.Println(reflect.TypeOf(ch).Kind()) // int32
		break
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	fmt.Println(reflect.TypeOf(str[0]).Kind()) // uint8
}

type TopVotedCandidate struct {
	times []int
	rank  []int
	n     int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	stats := make(map[int]int, n)
	maxC, maxP := 0, -1
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		stats[persons[i]]++
		cc := stats[persons[i]]
		if cc >= maxC {
			maxC = cc
			maxP = persons[i]
		}
		rank[i] = maxP
	}
	return TopVotedCandidate{
		times: times,
		rank:  rank,
		n:     n,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r, idx := 0, this.n-1, -1
	for l <= r {
		mid := l + (r-l)/2
		if this.times[mid] <= t {
			idx = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return this.rank[idx]
}

func toLowerCase(s string) string {

	var change = func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return byte(c - 'A' + 'a')
		}
		return c
	}

	result := make([]byte, 0, len(s))
	for i := range s {
		result = append(result, change(s[i]))
	}
	return string(result)
}
