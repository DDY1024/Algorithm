package main

import "fmt"

type (
	State    = int
	CharType = int
)

const (
	StateInitial             State = iota // 0: 初始状态
	StateSign                             // 1: 符号位
	StateInteger                          // 2: 整数部分
	StatePointWithInteger                 // 3: 左侧有整数的小数点
	StatePointWithoutInteger              // 4: 左侧无整数的小数点
	StateFraction                         // 5: 小数部分
	StateExp                              // 6: 指数 e
	StateExpSign                          // 7: 指数符号位 e+1, e-1
	StateExpInteger                       // 8: 指数整数部分
	StateEnd                              // 9: 结束状态
)

const (
	CharInteger CharType = iota // 0: 整数
	CharExp                     // 1: 指数 e 或 E
	CharPoint                   // 2: 小数点
	CharSign                    // 3: 符号位 -/+
	CharIllegal                 // 4: 非法字符
)

// 定义 FSM 状态转换关系即每种状态下遇到某类字符时的状态转移
// 其实本质上状态转换已经对整数的表示形式进行了相对应的限制了
var transation = map[State]map[CharType]State{
	StateInitial: {
		CharInteger: StateInteger,
		CharPoint:   StatePointWithoutInteger,
		CharSign:    StateSign,
	},
	StateSign: {
		CharInteger: StateInteger,
		CharPoint:   StatePointWithoutInteger,
	},
	StateInteger: {
		CharInteger: StateInteger,
		CharExp:     StateExp,
		CharPoint:   StatePointWithInteger,
	},
	StatePointWithInteger: {
		CharInteger: StateFraction, // 进入小数部分了
		CharExp:     StateExp,
	},
	StatePointWithoutInteger: {
		CharInteger: StateFraction,
	},
	StateFraction: {
		CharInteger: StateFraction, // 小数部分
		CharExp:     StateExp,
	},
	StateExp: {
		CharInteger: StateExpInteger,
		CharSign:    StateExpSign,
	},
	StateExpSign: {
		CharInteger: StateExpInteger,
	},
	StateExpInteger: {
		CharInteger: StateExpInteger,
	},
}

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharInteger
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	default:
		return CharIllegal
	}
}

func isLegalEndState(state State) bool {
	return state == StateInteger || state == StatePointWithInteger || state == StateFraction || state == StateExpInteger
}

func isNumber(s string) bool {
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if typ == CharIllegal {
			return false
		}
		if _, ok := transation[state][typ]; !ok {
			return false
		}
		state = transation[state][typ]
	}
	return isLegalEndState(state)
}

func main() {
	fmt.Println("Hello, World!")
}
