package main

import (
	"strings"
)

func main() {

}

func repeatedStringMatch(a string, b string) int {
	n, m := len(a), len(b)

	var construct = func(s string, c int) string {
		var res strings.Builder
		res.Grow(len(s) * c)
		for i := 1; i <= c; i++ {
			res.WriteString(s)
		}
		return res.String()
	}

	cc := m / n
	if m%n > 0 {
		cc++
	}

	if strings.Contains(construct(a, cc), b) {
		return cc
	}

	if strings.Contains(construct(a, cc+1), b) {
		return cc + 1
	}
	return -1
}
