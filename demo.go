package main

import (
	"fmt"
	"sort"
)

func getHint(secret string, guess string) string {
	n := len(secret)
	bA, bB, cA, cB := make([]byte, 0, n), make([]byte, 0, n), 0, 0
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			cA++
		} else {
			bA = append(bA, secret[i])
			bB = append(bB, guess[i])
		}
	}

	sort.Slice(bA, func(i, j int) bool {
		return bA[i] < bA[j]
	})
	sort.Slice(bB, func(i, j int) bool {
		return bB[i] < bB[j]
	})

	i, j, m := 0, 0, len(bA)
	for i < m && j < m {
		if bA[i] < bB[j] {
			i++
		} else if bA[i] == bB[j] {
			cB++
			i++
			j++
		} else {
			j++
		}
	}
	return fmt.Sprintf("%dA%dB", cA, cB)
}

func main() {
	fmt.Println("Hello, World!")
}
