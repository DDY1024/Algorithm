package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//
//
// 欠题
// TODO: 后缀数组
// https://leetcode-cn.com/problems/longest-common-subpath/
//
//
//
//
//

type Pair struct {
	tp   int // 0: 非括号 1: 左括号 2: 右括号
	freq map[string]int
}

type Node struct {
	elem string
	cnt  int
}

func countOfAtoms(formula string) string {
	n := len(formula)

	var mergePair = func(p1, p2 Pair) Pair {
		freq := make(map[string]int, len(p1.freq)+len(p2.freq))
		for k, v := range p1.freq {
			freq[k] += v
		}
		for k, v := range p2.freq {
			freq[k] += v
		}
		return Pair{freq: freq}
	}

	curElem, curCnt := []byte{}, 0
	stack := make([]Pair, 0, n)
	var pushStack = func() {
		if len(curElem) > 0 {
			if curCnt == 0 {
				curCnt = 1
			}
			stack = append(stack, Pair{freq: map[string]int{
				string(curElem): curCnt,
			}})
			curElem = []byte{}
			curCnt = 0
		}
	}

	for i := 0; i < n; i++ {
		if formula[i] >= 'a' && formula[i] <= 'z' {
			curElem = append(curElem, formula[i])
			continue
		}
		if formula[i] >= '0' && formula[i] <= '9' {
			curCnt = curCnt*10 + int(formula[i]-'0')
			continue
		}
		if formula[i] >= 'A' && formula[i] <= 'Z' {
			pushStack()
			curElem = append(curElem, formula[i])
			continue
		}
		if formula[i] == '(' {
			pushStack()
			stack = append(stack, Pair{tp: 1})
			continue
		}

		// ')'
		pushStack()
		mergeResult := Pair{}
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top.tp == 1 {
				break // 遇到 '(' break
			}
			mergeResult = mergePair(mergeResult, top)
		}
		// 计算 combo
		combo, idx := 0, i+1
		for ; idx < n; idx++ {
			if formula[idx] >= '0' && formula[idx] <= '9' {
				combo = combo*10 + int(formula[idx]-'0')
				continue
			}
			break
		}
		if combo > 0 {
			for k, v := range mergeResult.freq {
				mergeResult.freq[k] = v * combo
			}
		}
		stack = append(stack, mergeResult)
		i = idx - 1
	}
	pushStack()

	mergeResult := Pair{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		mergeResult = mergePair(mergeResult, top)
	}

	ndList := make([]Node, 0, len(mergeResult.freq))
	for k, v := range mergeResult.freq {
		ndList = append(ndList, Node{k, v})
	}
	sort.Slice(ndList, func(i, j int) bool {
		return ndList[i].elem <= ndList[j].elem
	})

	var result strings.Builder
	for _, nd := range ndList {
		result.WriteString(nd.elem)
		if nd.cnt > 1 {
			result.WriteString(strconv.FormatInt(int64(nd.cnt), 10))
		}
	}
	return result.String()
}

func main() {
	fmt.Println(countOfAtoms("H2O"))
	fmt.Println(countOfAtoms("Mg(OH)2"))
	fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
	fmt.Println(countOfAtoms("Be32"))
	fmt.Println("hello,world!")
}